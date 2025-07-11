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
)

func InteractiveView(dbConn *db.DB) {
	tType := utils.PromptInput("Filter by type (income/expense/leave empty): ")
	category := utils.PromptInput("Filter by category (leave empty): ")

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

		if tType != "" {
			query = query.Where("type = ?", tType)
		}
		if category != "" {
			query = query.Where("category = ?", category)
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
		for _, t := range transactions {
			if !t.UpdatedAt.IsZero() {
				updatedAtStr = t.UpdatedAt.Format("2006-01-02 15:04:05")
			}
			fmt.Printf("%-4d %-10.2f %-10s %-15s %-30s %-25s %-25s\n",
				t.ID, t.Amount, t.Type, t.Category, t.Description, t.CreatedAt.Format("2006-01-02 15:04:05"), updatedAtStr)
		}

		input := utils.PromptInput("\nEnter 'n' for next page, 'q' to quit: ")
		if strings.ToLower(input) == "n" {
			page++
		} else {
			break
		}
	}
}
