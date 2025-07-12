package handlers

import (
	"fmt"
	"log"
	"strings"

	"financetracker/db"
	"financetracker/models"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func MonthlyReport(dbConn *db.DB, filter models.ReportFilter) {
	var results []models.Result

	baseQuery := `
        SELECT 
            EXTRACT(YEAR FROM created_at) AS year,
            EXTRACT(MONTH FROM created_at) AS month,
            SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END) AS income,
            SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) AS expense
        FROM transactions
    `

	whereClauses := []string{}
	params := []interface{}{}
	paramIdx := 1

	if filter.Year != 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("EXTRACT(YEAR FROM created_at) = $%d", paramIdx))
		params = append(params, filter.Year)
		paramIdx++
	}
	if filter.Month != 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("EXTRACT(MONTH FROM created_at) = $%d", paramIdx))
		params = append(params, filter.Month)
		paramIdx++
	}
	if filter.StartDate != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("created_at >= $%d", paramIdx))
		params = append(params, filter.StartDate)
		paramIdx++
	}
	if filter.EndDate != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("created_at <= $%d", paramIdx))
		params = append(params, filter.EndDate)
		paramIdx++
	}

	if len(whereClauses) > 0 {
		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	baseQuery += `
        GROUP BY year, month
        ORDER BY year DESC, month DESC
    `

	err := dbConn.Gorm.Raw(baseQuery, params...).Scan(&results).Error
	if err != nil {
		log.Fatal("Error running report query:", err)
	}

	fmt.Println("\n📊 Monthly Finance Summary")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("| %-10s | %12s | %12s | %12s |\n", "Year-Month", "Income", "Expense", "Balance")
	fmt.Println(strings.Repeat("-", 60))

	for _, r := range results {
		balance := r.Income - r.Expense
		fmt.Printf("| %-10s | %s%12.2f%s | %s%12.2f%s | %s%12.2f%s |\n",
			fmt.Sprintf("%d-%02d", r.Year, r.Month),
			Green, r.Income, Reset,
			Red, r.Expense, Reset,
			Yellow, balance, Reset)

		// fmt.Printf("| %-10s | %12.2f | %12.2f | %12.2f |\n",
		// 	fmt.Sprintf("%d-%02d", r.Year, r.Month), r.Income, r.Expense, balance)
	}
	fmt.Println(strings.Repeat("=", 60))
}
