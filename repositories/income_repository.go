package repositories

import (
	"database/sql"
	"money-management-be/models"
)

type IncomeRepository struct {
	DB *sql.DB
}

func NewIncomeRepositories(db *sql.DB) *IncomeRepository {
	return &IncomeRepository{DB: db}
}

func (r *IncomeRepository) GetAll() ([]models.Income, error) {
	rows, err := r.DB.Query("SELECT id, name, nominal, date, category FROM incomes")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []models.Income
	for rows.Next() {
		var income models.Income
		if err := rows.Scan(&income.ID, &income.Name, &income.Nominal, &income.Date, &income.Category); err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}
	return incomes, nil
}

func (r *IncomeRepository) Create(income models.Income) error {
	_, err := r.DB.Exec("INSERT INTO incomes (name, nominal, date, category) VALUES ($1, $2, $3, $4)", income.Name, income.Nominal, income.Date, income.Category)
	return err
}
