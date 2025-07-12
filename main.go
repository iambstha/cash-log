package main

import (
	"fmt"
	"log"
	"os"

	"financetracker/db"
	"financetracker/handlers"

	"github.com/joho/godotenv"
)

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add")
	fmt.Println("  view")
	fmt.Println("  balance")
	fmt.Println("  edit")
	fmt.Println("  delete")
	fmt.Println("  add-category")
	fmt.Println("  remove-category")
	fmt.Println("  add-type")
	fmt.Println("  remove-type")
	fmt.Println("  info")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	if len(os.Args) < 2 {
		showUsage()
		return
	}

	dbConn := db.Connect()
	dbConn.AutoMigrate()
	db.SeedDefaults(dbConn)

	switch os.Args[1] {
	case "add":
		handlers.InteractiveAdd(dbConn)
	case "view":
		handlers.InteractiveView(dbConn)
	case "balance":
		handlers.ViewBalance(dbConn)
	case "edit":
		handlers.InteractiveEdit(dbConn)
	case "delete":
		handlers.InteractiveDelete(dbConn)
	case "report":
		handlers.InteractiveReport(dbConn)
	case "add-category":
		handlers.AddCategory(dbConn)
	case "add-type":
		handlers.AddType(dbConn)
	case "remove-category":
		handlers.RemoveCategory(dbConn)
	case "remove-type":
		handlers.RemoveType(dbConn)
	case "info":
		handlers.Info(dbConn)

	default:
		showUsage()
	}
}
