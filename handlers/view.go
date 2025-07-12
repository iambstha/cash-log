package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"financetracker/db"
	"financetracker/models"
	"financetracker/utils"

	"github.com/manifoldco/promptui"
)

func InteractiveView(dbConn *db.DB) {

	types := append([]string{"All"}, FetchTypes(dbConn)...)

	typePrompt := promptui.Select{
		Label: "Select transaction type",
		Items: types,
	}

	_, selectedType, err := typePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
		return
	}

	var selectedCategory string

	if selectedType != "All" {
		categories := FetchCategoriesByType(dbConn, selectedType)
		if len(categories) > 0 {
			categories = append([]string{"All"}, categories...)

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
			selectedCategory = categories[choiceIdx-1]
		} else {
			fmt.Println("No categories found for this type.")
			selectedCategory = "All"
		}
	} else {
		selectedCategory = "All"
	}

	searchField := utils.PromptInput("Search field (category/description/amount/date/leave empty): ")
	searchQuery := ""
	if searchField != "" {
		searchQuery = utils.PromptInput("Enter search query: ")
	}

	sortField := utils.PromptInput("Sort by (amount/category/created_at): ")
	if sortField == "" {
		sortField = "created_at"
	}
	sortOrder := utils.PromptInput("Sort order (asc/desc): ")
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	pageSizeStr := utils.PromptInput("Enter page size (default 5): ")
	pageSize := 5
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pageSize = ps
		} else {
			fmt.Println("Invalid page size, using default 5.")
		}
	}

	page := 1
	for {
		var transactions []models.Transaction
		query := dbConn.Gorm

		if selectedType != "All" {
			query = query.Where("type = ?", selectedType)
		}
		if selectedCategory != "All" {
			query = query.Where("category = ?", selectedCategory)
		}

		if searchField != "" && searchQuery != "" {
			switch searchField {
			case "category", "description":
				query = query.Where(fmt.Sprintf("%s ILIKE ?", searchField), "%"+searchQuery+"%")
			case "amount":
				amount, err := strconv.ParseFloat(searchQuery, 64)
				if err == nil {
					query = query.Where("amount = ?", amount)
				} else {
					fmt.Println("Invalid amount for search, ignoring.")
				}
			case "date":
				date, err := time.Parse("2006-01-02", searchQuery)
				if err == nil {
					nextDate := date.AddDate(0, 0, 1)
					query = query.Where("created_at >= ? AND created_at < ?", date, nextDate)
				} else {
					fmt.Println("Invalid date format, ignoring.")
				}
			default:
				fmt.Println("Unknown search field, ignoring.")
			}
		}

		orderClause := fmt.Sprintf("%s %s", sortField, sortOrder)

		offset := (page - 1) * pageSize
		result := query.Order(orderClause).Limit(pageSize).Offset(offset).Find(&transactions)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		if len(transactions) == 0 {
			fmt.Println("No more transactions.")
			break
		}

		fmt.Printf("\n--- Transaction History (Page %d) ---\n", page)
		fmt.Printf("%-4s %-10s %-10s %-15s %-30s %-25s %-25s\n", "ID", "Amount", "Type", "Category", "Description", "Created At", "UpdatedAt")
		fmt.Println(strings.Repeat("-", 125))
		updatedAtStr := ""
		var incomeTotal, expenseTotal float64

		for _, t := range transactions {
			if !t.UpdatedAt.IsZero() {
				updatedAtStr = t.UpdatedAt.Format("2006-01-02 15:04:05")
			}
			fmt.Printf("%-4d %-10.2f %-10s %-15s %-30s %-25s %-25s\n",
				t.ID, t.Amount, t.Type, t.Category, t.Description, t.CreatedAt.Format("2006-01-02 15:04:05"), updatedAtStr)
			switch t.Type {
			case "income":
				incomeTotal += t.Amount
			case "expense":
				expenseTotal += t.Amount
			}

		}

		fmt.Println("\n--- Summary ---")

		hasIncome := incomeTotal > 0
		hasExpense := expenseTotal > 0

		switch {
		case hasIncome && !hasExpense:
			fmt.Printf("%-15s: %.2f\n", "Total Income", incomeTotal)
		case !hasIncome && hasExpense:
			fmt.Printf("%-15s: %.2f\n", "Total Expense", expenseTotal)
		default:
			balance := incomeTotal - expenseTotal
			fmt.Printf("%-15s: %.2f\n", "Total Income", incomeTotal)
			fmt.Printf("%-15s: %.2f\n", "Total Expense", expenseTotal)
			fmt.Printf("%-15s: %.2f\n", "Balance", balance)
		}

		input := utils.PromptInput("\nEnter 'n' for next page, 'q' to quit: ")
		if strings.ToLower(input) == "n" {
			page++
		} else {
			break
		}
	}
}
