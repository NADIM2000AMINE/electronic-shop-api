package models

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"type:varchar(255);not null;unique" json:"name" validate:"required,min=3,max=255"`
	Active         bool           `gorm:"default:true" json:"active"`
	WhatsAppNumber string         `gorm:"type:varchar(20)" json:"whatsapp_number" validate:"omitempty,e164"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Users        []User        `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE" json:"-"`
	Products     []Product     `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE" json:"-"`
	Transactions []Transaction `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE" json:"-"`
}
