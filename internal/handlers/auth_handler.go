package handlers

import (
	"electronic-shop-api/internal/models"
	"electronic-shop-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Role     models.Role `json:"role"`
	ShopID   uint        `json:"shop_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
		}

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
		}

		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
			Role:     req.Role,
			ShopID:   req.ShopID,
		}

		if err := db.Create(&user).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Could not create user")
		}

		return utils.SuccessResponse(c, fiber.StatusCreated, "User created successfully", fiber.Map{
			"user": user,
		})
	}
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
		}

		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
		}

		if !utils.CheckPassword(req.Password, user.Password) {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
		}

		token, err := utils.GenerateToken(user.ID, user.ShopID, user.Role)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Could not generate token")
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "Login successful", fiber.Map{
			"token": token,
			"user":  user,
		})
	}
}
