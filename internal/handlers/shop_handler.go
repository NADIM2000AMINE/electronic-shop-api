package handlers

import (
	"electronic-shop-api/internal/models"
	"electronic-shop-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateWhatsAppRequest struct {
	WhatsAppNumber string `json:"whatsapp_number" validate:"required,e164"`
}

// UpdateWhatsAppNumber met à jour le numéro WhatsApp du shop (SuperAdmin only)
func UpdateWhatsAppNumber(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopID := c.Locals("shopID").(uint)
		// On récupère le rôle sous forme de string depuis le middleware
		roleStr := c.Locals("role").(string)
		role := models.Role(roleStr)

		// Vérifier que l'utilisateur est SuperAdmin
		if role != models.RoleSuperAdmin {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Only SuperAdmin can update WhatsApp number")
		}

		var req UpdateWhatsAppRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
		}

		// Mettre à jour le shop
		var shop models.Shop
		if err := db.First(&shop, shopID).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Shop not found")
		}

		shop.WhatsAppNumber = req.WhatsAppNumber
		if err := db.Save(&shop).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update shop")
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "WhatsApp number updated successfully", fiber.Map{
			"shop": shop,
		})
	}
}
