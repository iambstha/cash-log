package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"financetracker/db"
	"financetracker/models"
)

func InteractiveReport(dbConn *db.DB) {
	reader := bufio.NewReader(os.Stdin)

	readInt := func(prompt string) int {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			return 0
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid number, ignoring.")
			return 0
		}
		return num
	}

	year := readInt("Enter year (e.g. 2025) or leave empty: ")
	month := readInt("Enter month (1-12) or leave empty: ")

	fmt.Print("Enter start date (YYYY-MM-DD) or leave empty: ")
	startDate, _ := reader.ReadString('\n')
	startDate = strings.TrimSpace(startDate)

	fmt.Print("Enter end date (YYYY-MM-DD) or leave empty: ")
	endDate, _ := reader.ReadString('\n')
	endDate = strings.TrimSpace(endDate)

	filter := models.ReportFilter{
		Year:      year,
		Month:     month,
		StartDate: startDate,
		EndDate:   endDate,
	}

	MonthlyReport(dbConn, filter)
}
