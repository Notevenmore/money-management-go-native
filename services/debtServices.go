package services

import (
	"database/sql"
	"strconv"

	"money-management-be/models"
	"money-management-be/repositories"
)

type DebtServices struct {
	Repo *repositories.DebtRepository
}

func NewDebtServices(repo *repositories.DebtRepository) *DebtServices {
	return &DebtServices{Repo: repo}
}

func (s *DebtServices) GetAllDebt() ([]models.Debt, error) {
	debts, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return debts, nil
}

func (s *DebtServices) CreateDebt(debt models.Debt) error {
	return s.Repo.Create(debt)
}

func (s *DebtServices) UpdateDebt(debt models.Debt, id string, db *sql.DB) error {

	if err := s.Repo.Update(debt, id); err != nil {
		return err
	}

	if debt.IsFinish {
		outcomeRepo := repositories.NewOutcomeRepositories(db)
		debtOutcomeRepo := repositories.NewDebtOutcomeRepositories(db)

		outcome := &models.Outcome{
			Name:     debt.Name,
			Nominal:  debt.Nominal,
			Date:     debt.Deadline,
			Category: "Tagihan",
		}

		if err := outcomeRepo.Create(outcome); err != nil {
			return err
		}

		idInt, err := strconv.Atoi(id)

		if err != nil {
			return err
		}

		debtOutcome := models.DebtOutcome{
			IDDebt:    idInt,
			IDOutcome: outcome.ID,
		}

		return debtOutcomeRepo.Create(debtOutcome)
	} else {
		debtOutcomeRepo := repositories.NewDebtOutcomeRepositories(db)

		return debtOutcomeRepo.Delete(id)
	}
}
