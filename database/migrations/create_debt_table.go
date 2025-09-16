package migrations

import (
	"fmt"
	"money-management-be/database"
)

func CreateDebtTable() {
	query := `
	CREATE TABLE IF NOT EXISTS debts (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		nominal INTEGER NOT NULL,
		deadline TIMESTAMP NOT NULL,
		is_finish BOOLEAN DEFAULT(FALSE)
	);`

	_, err := database.DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration: debts table created successfully")
}
