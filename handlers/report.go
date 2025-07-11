package handlers

import (
	"fmt"
	"log"
	"strings"

	"financetracker/db"
	"financetracker/models"
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

	fmt.Println("Monthly Finance Summary")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Printf("%-10s %-10s %-10s %-10s\n", "Year-Month", "Income", "Expense", "Balance")

	for _, r := range results {
		balance := r.Income - r.Expense
		fmt.Printf("%4d-%02d    %10.2f %10.2f %10.2f\n", r.Year, r.Month, r.Income, r.Expense, balance)
	}
}
