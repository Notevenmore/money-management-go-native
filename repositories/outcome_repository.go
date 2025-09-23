package repositories

import (
	"database/sql"
	"money-management-be/models"
)

type OutcomeRepository struct {
	DB *sql.DB
}

func NewOutcomeRepositories(db *sql.DB) *OutcomeRepository {
	return &OutcomeRepository{DB: db}
}

func (r *OutcomeRepository) GetAll() ([]models.Outcome, error) {
	rows, err := r.DB.Query("SELECT id, name, nominal, date, category FROM outcomes")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outcomes []models.Outcome
	for rows.Next() {
		var outcome models.Outcome
		if err := rows.Scan(&outcome.ID, &outcome.Name, &outcome.Nominal, &outcome.Date, &outcome.Category); err != nil {
			return nil, err
		}
		outcomes = append(outcomes, outcome)
	}
	return outcomes, nil
}

func (r *OutcomeRepository) Create(outcome models.Outcome) error {
	_, err := r.DB.Exec("INSERT INTO outcomes (name, nominal, date, category) VALUES ($1, $2, $3, $4)", outcome.Name, outcome.Nominal, outcome.Date, outcome.Category)
	return err
}
