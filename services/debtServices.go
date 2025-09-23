package services

import (
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
