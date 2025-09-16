package migrations

import (
	"fmt"
	"money-management-be/database"
)

func CreateIncomesTable() {
	query := `
	CREATE TABLE IF NOT EXISTS incomes (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		nominal INTEGER NOT NULL,
		date TIMESTAMP NOT NULL,
		category VARCHAR(100) NOT NULL
	);`

	_, err := database.DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration: incomes table created successfully")
}
