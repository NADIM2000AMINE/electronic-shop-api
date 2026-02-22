package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TenantScope(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopID := c.Locals("shopID").(uint)
		c.Locals("db", db.Where("shop_id = ?", shopID))
		return c.Next()
	}
}
