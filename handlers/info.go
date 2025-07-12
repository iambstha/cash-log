package handlers

import (
	"financetracker/db"
	"financetracker/models"
	"fmt"
)

func Info(dbConn *db.DB) {
	// Fetch and display all transaction types
	var types []models.TransactionType
	dbConn.Gorm.Order("name asc").Find(&types)

	fmt.Println("\n📂 Available Transaction Types:")
	for _, t := range types {
		fmt.Printf(" - %s\n", t.Name)
	}

	// Fetch and display all categories grouped by type
	var categories []models.Category
	dbConn.Gorm.Order("type asc, name asc").Find(&categories)

	typeToCats := make(map[string][]string)
	for _, cat := range categories {
		typeToCats[cat.Type] = append(typeToCats[cat.Type], cat.Name)
	}

	fmt.Println("\n🏷️ Categories by Type:")
	for tType, catList := range typeToCats {
		fmt.Printf(" [%s]:\n", tType)
		for _, cat := range catList {
			fmt.Printf("   - %s\n", cat)
		}
	}
}
