package utils

import (
	"financetracker/db"
	"financetracker/models"
)

func FetchCategoriesByType(dbConn *db.DB, tType string) []string {
	var cats []models.Category
	dbConn.Gorm.Where("type = ?", tType).Order("name asc").Find(&cats)

	var names []string
	for _, c := range cats {
		names = append(names, c.Name)
	}
	return names
}

func FetchAllCategories(dbConn *db.DB) []string {
	var cats []models.Category
	dbConn.Gorm.Order("name asc").Find(&cats)

	var names []string
	for _, c := range cats {
		names = append(names, c.Name)
	}
	return names
}

func FetchTypes(dbConn *db.DB) []string {
	var types []models.TransactionType
	dbConn.Gorm.Find(&types)
	var names []string
	for _, t := range types {
		names = append(names, t.Name)
	}
	return names
}
