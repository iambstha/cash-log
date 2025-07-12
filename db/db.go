package db

import (
	"financetracker/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
		log.Fatal("Invalid CASHLOG_DB_PORT:", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	var gormDB *gorm.DB
	for i := 0; i < 5; i++ {
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return &DB{Gorm: gormDB}
		}
		log.Printf("DB connection failed. Retrying in 5s (%d/5)...", i+1)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Failed to connect to database after 5 attempts:", err)
	return nil
}

func (db *DB) AutoMigrate() {
	err := db.Gorm.AutoMigrate(&models.Transaction{}, &models.Category{}, &models.TransactionType{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}
}
