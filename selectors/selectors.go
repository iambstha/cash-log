package selectors

import (
	"fmt"
	"strconv"

	"financetracker/db"
	"financetracker/utils"

	"github.com/manifoldco/promptui"
)

func PromptSelectTransactionType(dbConn *db.DB) (string, error) {
	types := utils.FetchTypes(dbConn)

	typePrompt := promptui.Select{
		Label: "Select transaction type",
		Items: types,
	}

	_, selectedType, err := typePrompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return selectedType, nil
}

func PromptSelectCategoryByType(dbConn *db.DB, tType string) (string, error) {
	categories := utils.FetchCategoriesByType(dbConn, tType)
	if len(categories) == 0 {
		return "", fmt.Errorf("no categories found for type: %s", tType)
	}

	fmt.Println("Choose a category:")
	for i, cat := range categories {
		fmt.Printf("[%d] %s\n", i+1, cat)
	}

	choiceStr := utils.PromptInput("Enter category number: ")
	choiceIdx, err := strconv.Atoi(choiceStr)
	if err != nil || choiceIdx < 1 || choiceIdx > len(categories) {
		return "", fmt.Errorf("invalid selection")
	}

	return categories[choiceIdx-1], nil
}

func PromptSelectTypeAndCategory(dbConn *db.DB) (string, string, error) {
	tType, err := PromptSelectTransactionType(dbConn)
	if err != nil {
		return "", "", err
	}

	category, err := PromptSelectCategoryByType(dbConn, tType)
	if err != nil {
		return "", "", err
	}

	return tType, category, nil
}
