package handlers

import (
	"financetracker/db"
	"financetracker/models"
	"financetracker/utils"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func AddCategory(dbConn *db.DB) {
	name := utils.PromptInput("Enter new category name: ")
	types := fetchTypes(dbConn)

	typePrompt := promptui.Select{
		Label: "Select transaction type",
		Items: types,
	}

	_, tType, err := typePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
		return
	}

	category := models.Category{Name: name, Type: tType}
	result := dbConn.Gorm.FirstOrCreate(&category, models.Category{Name: name, Type: tType})
	if result.Error != nil {
		log.Fatal("Failed to add category:", result.Error)
	}
	fmt.Println("Category added successfully.")
}

func AddType(dbConn *db.DB) {
	name := utils.PromptInput("Enter new transaction type: ")
	tType := models.TransactionType{Name: name}
	result := dbConn.Gorm.FirstOrCreate(&tType, models.TransactionType{Name: name})
	if result.Error != nil {
		log.Fatal("Failed to add type:", result.Error)
	}
	fmt.Println("Transaction type added successfully.")
}
