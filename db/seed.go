package db

import (
	"financetracker/db/constants"
	"financetracker/models"
)

func SeedDefaults(dbConn *DB) {
	// Seed protected types
	for _, t := range constants.ProtectedTransactionTypes {
		dbConn.Gorm.FirstOrCreate(&models.TransactionType{}, models.TransactionType{Name: t})
	}

	// Convert protected categories map to []models.Category
	var initialCategories []models.Category
	for tType, names := range constants.ProtectedCategories {
		for _, name := range names {
			initialCategories = append(initialCategories, models.Category{Name: name, Type: tType})
		}
	}

	// Seed categories
	for _, c := range initialCategories {
		dbConn.Gorm.FirstOrCreate(&models.Category{}, c)
	}
}
