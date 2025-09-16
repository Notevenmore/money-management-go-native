package migrations

import (
	"fmt"
	"money-management-be/database"
)

func CreateAssetTable() {
	query := `
	CREATE TABLE IF NOT EXISTS assets (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		is_reusable BOOLEAN DEFAULT(TRUE)
	);`

	_, err := database.DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration: assets table created successfully")
}
