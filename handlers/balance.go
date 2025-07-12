package handlers

import (
	"fmt"

	"financetracker/db"
	"financetracker/models"
)

func ViewBalance(dbConn *db.DB) {
	var incomeTotal, expenseTotal float64
	dbConn.Gorm.Model(models.Transaction{}).Where("type = ?", "income").Select("SUM(amount)").Scan(&incomeTotal)
	dbConn.Gorm.Model(models.Transaction{}).Where("type = ?", "expense").Select("SUM(amount)").Scan(&expenseTotal)
	balance := incomeTotal - expenseTotal

	fmt.Println()
	fmt.Println("┌──────────────────────────────┐")
	fmt.Println("│        💰 Balance Sheet       │")
	fmt.Println("├──────────────────────────────┤")
	fmt.Printf("│ %-20s %8.2f │\n", "Total Income:", incomeTotal)
	fmt.Printf("│ %-20s %8.2f │\n", "Total Expenses:", expenseTotal)
	fmt.Println("├──────────────────────────────┤")
	fmt.Printf("│ %-20s %8.2f │\n", "Net Balance:", balance)
	fmt.Println("└──────────────────────────────┘")
}
