package db

import (
	"financetracker/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Gorm *gorm.DB
}

func Connect() *DB {
	host := os.Getenv("CASHLOG_DB_HOST")
	user := os.Getenv("CASHLOG_DB_USER")
	password := os.Getenv("CASHLOG_DB_PASS")
	dbname := os.Getenv("CASHLOG_DB_NAME")
	portStr := os.Getenv("CASHLOG_DB_PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid DB_PORT:", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// uncomment this to see the full gorm log, only use for debugging
	// gormDB = gormDB.Debug()

	return &DB{Gorm: gormDB}
}

func (db *DB) AutoMigrate() {
	err := db.Gorm.AutoMigrate(&models.Transaction{}, &models.Category{}, &models.TransactionType{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}
}
