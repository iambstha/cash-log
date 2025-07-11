package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"financetracker/db"
	"financetracker/handlers"
	"financetracker/scheduler"

	"github.com/joho/godotenv"
)

var Version = "v1.0.0"

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
	versionFlag := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println("App Version:", Version)
		os.Exit(0)
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "production"
	}

	if env == "development" || env == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found or failed to load")
		}
	}

	if len(os.Args) < 2 {
		showUsage()
		return
	}

	dbConn := db.Connect()
	dbConn.AutoMigrate()
	db.SeedDefaults(dbConn)

	switch os.Args[1] {
	case "remind":
		scheduler.StartReminderScheduler()
		select {}
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
