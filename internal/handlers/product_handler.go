package handlers

import (
	"electronic-shop-api/internal/models"
	"electronic-shop-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Le middleware TenantScope a déjà filtré le db sur le bon shop
		tenantDB := c.Locals("db").(*gorm.DB)
		roleStr := c.Locals("role").(string)

		var products []models.Product
		if err := tenantDB.Find(&products).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Database error")
		}

		// Si c'est un Admin, on masque le prix d'achat
		if roleStr == string(models.RoleAdmin) {
			for i := range products {
				products[i].PurchasePrice = 0 // Sera masqué grâce au 'omitempty' dans le json
			}
		}

		return c.JSON(fiber.Map{"products": products})
	}
}

func CreateProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantDB := c.Locals("db").(*gorm.DB)
		shopID := c.Locals("shopID").(uint)

		var product models.Product
		if err := c.BodyParser(&product); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
		}

		product.ShopID = shopID

		if err := tenantDB.Create(&product).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create product")
		}

		return utils.SuccessResponse(c, fiber.StatusCreated, "Product created successfully", fiber.Map{
			"product": product,
		})
	}
}

func UpdateProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantDB := c.Locals("db").(*gorm.DB)
		id := c.Params("id")

		var product models.Product
		if err := tenantDB.First(&product, id).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
		}

		if err := c.BodyParser(&product); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
		}

		tenantDB.Save(&product)
		return utils.SuccessResponse(c, fiber.StatusOK, "Product updated", product)
	}
}

func DeleteProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantDB := c.Locals("db").(*gorm.DB)
		id := c.Params("id")

		if err := tenantDB.Delete(&models.Product{}, id).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete product")
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "Product deleted", nil)
	}
}
