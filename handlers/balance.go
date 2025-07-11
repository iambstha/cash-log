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
	fmt.Printf("Income: %.2f\nExpenses: %.2f\nBalance: %.2f\n", incomeTotal, expenseTotal, balance)
}
