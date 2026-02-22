package handlers

import (
	"electronic-shop-api/internal/models"
	"electronic-shop-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetTransactions(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantDB := c.Locals("db").(*gorm.DB)
		var transactions []models.Transaction

		if err := tenantDB.Find(&transactions).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Database error")
		}

		return c.JSON(fiber.Map{"transactions": transactions})
	}
}

func CreateTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// On utilise tenantDB pour s'assurer de l'isolation multi-tenant
		tenantDB := c.Locals("db").(*gorm.DB)
		shopID := c.Locals("shopID").(uint)

		var req models.Transaction
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request")
		}
		req.ShopID = shopID

		// Utilise tenantDB ici pour corriger l'erreur de variable non utilisée
		err := tenantDB.Transaction(func(tx *gorm.DB) error {
			if req.Type == models.TransactionSale && req.ProductID != nil {
				var product models.Product
				if err := tx.First(&product, *req.ProductID).Error; err != nil {
					return fiber.NewError(fiber.StatusNotFound, "Product not found")
				}
				if product.Stock < req.Quantity {
					return fiber.NewError(fiber.StatusBadRequest, "Insufficient stock")
				}
				product.Stock -= req.Quantity
				return tx.Save(&product).Error
			}
			return tx.Create(&req).Error
		})

		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
		return utils.SuccessResponse(c, fiber.StatusCreated, "Transaction created", req)
	}
}
