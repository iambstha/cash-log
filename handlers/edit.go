package handlers

import (
	"financetracker/db"
	"financetracker/models"
	"financetracker/utils"
	"fmt"
	"strconv"
	"time"
)

func InteractiveEdit(dbConn *db.DB) {
	idStr := utils.PromptInput("Enter transaction ID to edit: ")
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

	fmt.Println("Leave field empty to keep current value")

	amountStr := utils.PromptInput(fmt.Sprintf("Amount (%.2f): ", t.Amount))
	if amountStr != "" {
		newAmount, err := strconv.ParseFloat(amountStr, 64)
		if err == nil {
			t.Amount = newAmount
		}
	}

	tType := utils.PromptInput(fmt.Sprintf("Type (%s): ", t.Type))
	if tType != "" {
		t.Type = tType
	}

	category := utils.PromptInput(fmt.Sprintf("Category (%s): ", t.Category))
	if category != "" {
		t.Category = category
	}
	description := utils.PromptInput(fmt.Sprintf("Description (%s): ", t.Description))
	if description != "" {
		t.Description = description
	}

	t.UpdatedAt = time.Now()
	dbConn.Gorm.Save(&t)
	fmt.Println("Transaction updated.")
}
