package repositories

import (
	"database/sql"
	"money-management-be/models"
)

type DebtRepository struct {
	DB *sql.DB
}

func NewDebtRepositories(db *sql.DB) *DebtRepository {
	return &DebtRepository{DB: db}
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

func (r *DebtRepository) Update(debt models.Debt, id string) error {
	_, err := r.DB.Exec("UPDATE debts SET is_finish = $1 WHERE (id = $2)", debt.IsFinish, id)

	return err
}
