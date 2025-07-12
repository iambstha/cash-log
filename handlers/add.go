package handlers

import (
	"fmt"
	"log"
	"strconv"

	"financetracker/db"
	"financetracker/models"
	"financetracker/selectors"
	"financetracker/utils"
)

func InteractiveAdd(dbConn *db.DB) {

	tType, err := selectors.PromptSelectTransactionType(dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}

	category, err := selectors.PromptSelectCategoryByType(dbConn, tType)
	if err != nil {
		fmt.Println(err)
		return
	}

	amountStr := utils.PromptInput("Enter amount: ")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	description := utils.PromptInput("Enter description: ")

	t := models.Transaction{
		Amount:      amount,
		Category:    category,
		Description: description,
		Type:        tType,
	}

	result := dbConn.Gorm.Create(&t)
	if result.Error != nil {
		log.Fatal("Error inserting transaction:", result.Error)
	}
	fmt.Println("Transaction added successfully.")
}
