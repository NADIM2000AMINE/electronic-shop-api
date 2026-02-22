package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"electronic-shop-api/internal/config"
	"electronic-shop-api/internal/database"
	"electronic-shop-api/internal/handlers"
	"electronic-shop-api/internal/middleware"
)

func main() {
	// Charger configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connexion DB
	db, err := database.Connect(cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Créer application Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// Middleware globaux
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Servir le frontend statique
	app.Static("/", "./public")

	// Routes publiques (sans authentification)
	public := app.Group("/api/public")
	public.Get("/:shopID/products", handlers.GetPublicProducts(db))
	public.Get("/products/:id/whatsapp", handlers.GetWhatsAppLink(db))

	// Routes d'authentification (publiques)
	auth := app.Group("/api/auth")
	auth.Post("/register", handlers.Register(db))
	auth.Post("/login", handlers.Login(db))

	// Routes protégées (JWT requis + Multi-tenant)
	api := app.Group("/api", middleware.JWTProtected(), middleware.TenantScope(db))

	// Products
	products := api.Group("/products")
	products.Get("/", handlers.GetProducts(db))
	products.Post("/", middleware.RequireRole("SuperAdmin", "Admin"), handlers.CreateProduct(db))
	products.Put("/:id", middleware.RequireRole("SuperAdmin", "Admin"), handlers.UpdateProduct(db))
	products.Delete("/:id", middleware.RequireRole("SuperAdmin", "Admin"), handlers.DeleteProduct(db))

	// Transactions
	transactions := api.Group("/transactions")
	transactions.Get("/", middleware.RequireRole("SuperAdmin", "Admin"), handlers.GetTransactions(db))
	transactions.Post("/", middleware.RequireRole("SuperAdmin", "Admin"), handlers.CreateTransaction(db))

	// Reports (SuperAdmin uniquement)
	reports := api.Group("/reports", middleware.RequireRole("SuperAdmin"))
	reports.Get("/dashboard", handlers.GetDashboard(db))

	// Shop (SuperAdmin uniquement)
	shop := api.Group("/shop", middleware.RequireRole("SuperAdmin"))
	shop.Put("/whatsapp", handlers.UpdateWhatsAppNumber(db))

	// Démarrer serveur
	log.Printf("Server starting on port %s...", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
