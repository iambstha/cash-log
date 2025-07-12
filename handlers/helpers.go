package handlers

import (
	"financetracker/db"
	"financetracker/db/constants"
	"financetracker/models"
	"financetracker/utils"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/manifoldco/promptui"
)

func AddCategory(dbConn *db.DB) {
	name := utils.PromptInput("Enter new category name: ")
	types := FetchTypes(dbConn)

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

func RemoveType(dbConn *db.DB) {
	types := FetchTypes(dbConn)

	typePrompt := promptui.Select{
		Label: "Select transaction type",
		Items: types,
	}

	_, selectedType, err := typePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return
	}

	if isProtectedType(selectedType) {
		fmt.Println("⚠️  This transaction type is protected and cannot be deleted.")
		return
	}

	if err := blockIfTypeUsed(dbConn, selectedType); err != nil {
		fmt.Println(err)
		return
	}

	var tType models.TransactionType
	if err := dbConn.Gorm.Where("name = ?", selectedType).First(&tType).Error; err != nil {
		log.Fatal("Transaction type not found:", err)
	}

	result := dbConn.Gorm.Unscoped().Delete(&tType)
	if result.Error != nil {
		log.Fatal("Failed to delete transaction type:", result.Error)
	}

	fmt.Println("Transaction type deleted successfully.")
}

func RemoveCategory(dbConn *db.DB) {
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

	if isProtectedCategory(tType, category) {
		fmt.Println("⚠️  This category is protected and cannot be deleted.")
		return
	}

	if err := blockIfCategoryUsed(dbConn, category, tType); err != nil {
		fmt.Println(err)
		return
	}

	var cat models.Category
	if err := dbConn.Gorm.Where("name = ? AND type = ?", category, tType).First(&cat).Error; err != nil {
		log.Fatal("Category not found:", err)
	}

	result := dbConn.Gorm.Unscoped().Delete(&cat)
	if result.Error != nil {
		log.Fatal("Failed to delete category:", result.Error)
	}

	fmt.Println("Category deleted successfully.")
}

func blockIfCategoryUsed(dbConn *db.DB, category string, tType string) error {
	var count int64
	dbConn.Gorm.Model(&models.Transaction{}).
		Where("category = ? AND type = ?", category, tType).
		Count(&count)

	if count > 0 {
		return fmt.Errorf("⚠️  Cannot delete '%s' category. It is used in %d transaction(s)", category, count)
	}
	return nil
}

func blockIfTypeUsed(dbConn *db.DB, tType string) error {
	var count int64
	dbConn.Gorm.Model(&models.Transaction{}).
		Where("type = ?", tType).
		Count(&count)

	if count > 0 {
		return fmt.Errorf("⚠️  Cannot delete '%s' transaction type. It is used in %d transaction(s)", tType, count)
	}
	return nil
}

func isProtectedCategory(tType string, category string) bool {
	if protected, ok := constants.ProtectedCategories[tType]; ok {
		return slices.Contains(protected, category)
	}
	return false
}

func isProtectedType(tType string) bool {
	return slices.Contains(constants.ProtectedTransactionTypes, tType)
}
