package db

import "financetracker/models"

func SeedDefaults(dbConn *DB) {
	initialTypes := []string{"income", "expense"}
	for _, t := range initialTypes {
		dbConn.Gorm.FirstOrCreate(&models.TransactionType{}, models.TransactionType{Name: t})
	}

	initialCategories := []models.Category{
		{Name: "Salary", Type: "income"},
		{Name: "Freelance", Type: "income"},
		{Name: "Food", Type: "expense"},
		{Name: "Transport", Type: "expense"},
		{Name: "Investment", Type: "expense"},
		{Name: "Shopping", Type: "expense"},
	}

	for _, c := range initialCategories {
		dbConn.Gorm.FirstOrCreate(&models.Category{}, c)
	}
}
