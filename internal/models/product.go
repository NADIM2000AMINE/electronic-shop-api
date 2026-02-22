package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=2,max=255"`
	Description   string         `gorm:"type:text" json:"description" validate:"required"`
	Category      string         `gorm:"type:varchar(100)" json:"category" validate:"required"`
	PurchasePrice float64        `gorm:"type:decimal(10,2);not null" json:"purchase_price,omitempty" validate:"required,gt=0"`
	SellingPrice  float64        `gorm:"type:decimal(10,2);not null" json:"selling_price" validate:"required,gt=0"`
	Stock         int            `gorm:"default:0" json:"stock" validate:"gte=0"`
	ImageURL      string         `gorm:"type:varchar(500)" json:"image_url" validate:"omitempty,url"`
	ShopID        uint           `gorm:"not null;index" json:"shop_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Shop         *Shop         `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE" json:"-"`
	Transactions []Transaction `gorm:"foreignKey:ProductID" json:"-"`
}

// PublicProduct représente un produit pour l'affichage public (sans purchase_price)
type PublicProduct struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	SellingPrice float64 `json:"selling_price"`
	Stock        int     `json:"stock"`
	ImageURL     string  `json:"image_url"`
}
