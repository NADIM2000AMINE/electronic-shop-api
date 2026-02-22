package main

import (
	"log"

	"electronic-shop-api/internal/config"
	"electronic-shop-api/internal/database"
	"electronic-shop-api/internal/models"
	"electronic-shop-api/internal/utils"
)

func main() {
	cfg, _ := config.LoadConfig()
	db, _ := database.Connect(cfg.GetDSN())

	// Créer 2 shops
	shop1 := models.Shop{
		Name:           "TechStore Paris",
		Active:         true,
		WhatsAppNumber: "+33612345678",
	}
	shop2 := models.Shop{
		Name:           "ElectroShop Lyon",
		Active:         true,
		WhatsAppNumber: "+33687654321",
	}
	db.Create(&shop1)
	db.Create(&shop2)

	// Créer utilisateurs Shop 1
	hashedPassword, _ := utils.HashPassword("password123")
	superAdmin := models.User{
		Name:     "Super Admin",
		Email:    "super@techstore.com",
		Password: hashedPassword,
		Role:     models.RoleSuperAdmin,
		ShopID:   shop1.ID,
	}
	admin := models.User{
		Name:     "Admin User",
		Email:    "admin@techstore.com",
		Password: hashedPassword,
		Role:     models.RoleAdmin,
		ShopID:   shop1.ID,
	}
	db.Create(&superAdmin)
	db.Create(&admin)

	// Créer produits Shop 1
	products := []models.Product{
		{
			Name:          "iPhone 15 Pro",
			Description:   "Le dernier iPhone avec puce A17 Pro.",
			Category:      "Smartphones",
			PurchasePrice: 900.0,
			SellingPrice:  1200.0,
			Stock:         10,
			ImageURL:      "https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=600&q=80",
			ShopID:        shop1.ID,
		},
		{
			Name:          "MacBook Pro M3",
			Description:   "Ordinateur portable surpuissant pour les pros.",
			Category:      "Ordinateurs",
			PurchasePrice: 1500.0,
			SellingPrice:  2000.0,
			Stock:         5,
			ImageURL:      "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=600&q=80",
			ShopID:        shop1.ID,
		},
		{
			Name:          "AirPods Pro 2",
			Description:   "Écouteurs sans fil avec réduction de bruit.",
			Category:      "Audio",
			PurchasePrice: 180.0,
			SellingPrice:  250.0,
			Stock:         2, // On laisse à 0 pour montrer la rupture de stock
			ImageURL:      "https://images.unsplash.com/photo-1606220588913-b3aacb4d2f46?w=600&q=80",
			ShopID:        shop1.ID,
		},
	}

	for _, p := range products {
		db.Create(&p)
	}

	log.Println("Seed data created successfully!")
	log.Println("Credentials:")
	log.Println(" SuperAdmin: super@techstore.com / password123")
	log.Println(" Admin: admin@techstore.com / password123")
}
