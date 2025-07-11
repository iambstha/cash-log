package handlers

import (
	"financetracker/db"
	"financetracker/models"
	"financetracker/utils"
	"fmt"
	"strconv"
	"strings"
)

func InteractiveDelete(dbConn *db.DB) {
	idStr := utils.PromptInput("Enter transaction ID to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	var t models.Transaction
	result := dbConn.Gorm.First(&t, id)
	if result.Error != nil {
		fmt.Println("Transaction not found")
		return
	}

	confirm := utils.PromptInput("Are you sure you want to delete this transaction? (yes/no): ")
	if strings.ToLower(confirm) == "yes" {
		dbConn.Gorm.Delete(&t)
		fmt.Println("Transaction deleted.")
	} else {
		fmt.Println("Deletion cancelled.")
	}
}
