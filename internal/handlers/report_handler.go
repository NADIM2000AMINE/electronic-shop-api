package handlers

import (
	"electronic-shop-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetDashboard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantDB := c.Locals("db").(*gorm.DB)

		var totalSales, totalExpenses float64
		var totalProducts, totalTransactions int64

		// Calculer Total Sales
		tenantDB.Model(&models.Transaction{}).Where("type = ?", models.TransactionSale).Select("COALESCE(SUM(amount), 0)").Scan(&totalSales)

		// Calculer Total Expenses
		tenantDB.Model(&models.Transaction{}).Where("type = ?", models.TransactionExpense).Select("COALESCE(SUM(amount), 0)").Scan(&totalExpenses)

		netProfit := totalSales - totalExpenses

		// Compter total produits et transactions
		tenantDB.Model(&models.Product{}).Count(&totalProducts)
		tenantDB.Model(&models.Transaction{}).Count(&totalTransactions)

		// Produits en stock faible (ex: < 5)
		var lowStockProducts []fiber.Map
		tenantDB.Model(&models.Product{}).Where("stock < ?", 5).Select("id, name, stock, category").Find(&lowStockProducts)

		return c.JSON(fiber.Map{
			"total_sales":        totalSales,
			"total_expenses":     totalExpenses,
			"net_profit":         netProfit,
			"low_stock_products": lowStockProducts,
			"total_products":     totalProducts,
			"total_transactions": totalTransactions,
		})
	}
}
