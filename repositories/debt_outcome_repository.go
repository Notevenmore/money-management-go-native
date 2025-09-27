package repositories

import (
	"database/sql"
	"money-management-be/models"
)

type DebtOutcomeRepository struct {
	DB *sql.DB
}

func NewDebtOutcomeRepositories(db *sql.DB) *DebtOutcomeRepository {
	return &DebtOutcomeRepository{DB: db}
}

func (r *DebtOutcomeRepository) Create(debtOutcome models.DebtOutcome) error {
	_, err := r.DB.Exec("INSERT INTO debt_outcome (id_debt, id_outcome) VALUES ($1, $2)", debtOutcome.IDDebt, debtOutcome.IDOutcome)
	return err
}

func (r *DebtOutcomeRepository) Delete(id_debt string) error {
	_, err := r.DB.Exec("DELETE from debt_outcome WHERE (id_debt = $1)", id_debt)

	return err
}
