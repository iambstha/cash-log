package handlers

import (
	"fmt"
	"log"
	"strconv"

	"financetracker/db"
	"financetracker/models"
	"financetracker/utils"

	"github.com/manifoldco/promptui"
)

func InteractiveAdd(dbConn *db.DB) {
	amountStr := utils.PromptInput("Enter amount: ")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	// fetch and pass types
	types := FetchTypes(dbConn)

	typePrompt := promptui.Select{
		Label: "Select transaction type",
		Items: types,
	}

	_, tType, err := typePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return
	}

	// fetch and pass categories
	categories := FetchCategoriesByType(dbConn, tType)
	if len(categories) == 0 {
		fmt.Printf("No categories found for type: %s\n", tType)
		return
	}

	fmt.Println("Choose a category:")
	for i, cat := range categories {
		fmt.Printf("[%d] %s\n", i+1, cat)
	}

	choiceStr := utils.PromptInput("Enter category number: ")
	choiceIdx, err := strconv.Atoi(choiceStr)
	if err != nil || choiceIdx < 1 || choiceIdx > len(categories) {
		fmt.Println("Invalid selection.")
		return
	}
	category := categories[choiceIdx-1]

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
