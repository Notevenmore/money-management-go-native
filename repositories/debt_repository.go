package repositories

import (
	"database/sql"
	"money-management-be/models"
)

type DebtRepository struct {
	DB *sql.DB
}

func (r *DebtRepository) GetAll() ([]models.Debt, error) {
	rows, err := r.DB.Query("SELECT id, name, nominal, deadline, is_finish FROM debts")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var debts []models.Debt
	for rows.Next() {
		var debt models.Debt
		if err := rows.Scan(&debt.ID, &debt.Name, &debt.Nominal, &debt.Deadline, &debt.IsFinish); err != nil {
			return nil, err
		}
		debts = append(debts, debt)
	}
	return debts, nil
}

func (r *DebtRepository) Create(debt models.Debt) error {
	_, err := r.DB.Exec("INSERT INTO debts (name, nominal, deadline) VALUES ($1, $2, $3)", debt.Name, debt.Nominal, debt.Deadline)
	return err
}
