package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	TransactionSale       TransactionType = "Sale"
	TransactionExpense    TransactionType = "Expense"
	TransactionWithdrawal TransactionType = "Withdrawal"
)

type Transaction struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Type      TransactionType `gorm:"type:varchar(50);not null" json:"type" validate:"required,oneof=Sale Expense Withdrawal"`
	ProductID *uint           `gorm:"index" json:"product_id,omitempty"`
	Quantity  int             `gorm:"default:0" json:"quantity" validate:"gte=0"`
	Amount    float64         `gorm:"type:decimal(10,2);not null" json:"amount" validate:"required,gt=0"`
	ShopID    uint            `gorm:"not null;index" json:"shop_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`

	// Relations
	Shop    *Shop    `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE" json:"-"`
	Product *Product `gorm:"foreignKey:ProductID;constraint:OnDelete:SET NULL" json:"-"`
}
