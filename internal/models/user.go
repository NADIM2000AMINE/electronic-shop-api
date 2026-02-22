package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleSuperAdmin Role = "SuperAdmin"
	RoleAdmin      Role = "Admin"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=2,max=255"`
	Email     string         `gorm:"type:varchar(255);not null;uniqueIndex" json:"email" validate:"required,email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Role      Role           `gorm:"type:varchar(50);not null" json:"role" validate:"required,oneof=SuperAdmin Admin"`
	ShopID    uint           `gorm:"not null;index" json:"shop_id" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Shop *Shop `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE" json:"-"`
}
