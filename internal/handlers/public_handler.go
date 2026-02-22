package handlers

import (
	"electronic-shop-api/internal/models"
	"electronic-shop-api/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetPublicProducts retourne les produits publics d'un shop (sans auth)
func GetPublicProducts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopID, err := strconv.Atoi(c.Params("shopID"))
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid shop ID")
		}

		// Vérifier que le shop existe et est actif
		var shop models.Shop
		if err := db.First(&shop, shopID).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Shop not found")
		}

		if !shop.Active {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Shop is not active")
		}

		// Récupérer les produits du shop (stock > 0 uniquement)
		var products []models.Product
		if err := db.Where("shop_id = ? AND stock > ?", shopID, 0).Find(&products).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Database error")
		}

		// Convertir en PublicProduct (sans purchase_price)
		publicProducts := make([]models.PublicProduct, len(products))
		for i, p := range products {
			publicProducts[i] = models.PublicProduct{
				ID:           p.ID,
				Name:         p.Name,
				Description:  p.Description,
				Category:     p.Category,
				SellingPrice: p.SellingPrice,
				Stock:        p.Stock,
				ImageURL:     p.ImageURL,
			}
		}

		return c.JSON(fiber.Map{
			"shop": fiber.Map{
				"id":              shop.ID,
				"name":            shop.Name,
				"whatsapp_number": shop.WhatsAppNumber,
			},
			"products": publicProducts,
		})
	}
}

// GetWhatsAppLink génère un lien WhatsApp pour un produit
func GetWhatsAppLink(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
		}

		var product models.Product
		if err := db.Preload("Shop").First(&product, productID).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
		}

		if product.Shop.WhatsAppNumber == "" {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "WhatsApp number not configured")
		}

		whatsappURL := utils.GenerateWhatsAppLink(
			product.Shop.WhatsAppNumber,
			product.Name,
			product.SellingPrice,
		)

		return c.JSON(fiber.Map{
			"whatsapp_url": whatsappURL,
		})
	}
}
