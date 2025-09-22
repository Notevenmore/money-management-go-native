package incomeservices

import (
	"money-management-be/models"
	"money-management-be/repositories"
)

type IncomeServices struct {
	Repo *repositories.IncomeRepository
}

func NewIncomeServices(repo *repositories.IncomeRepository) *IncomeServices {
	return &IncomeServices{Repo: repo}
}

func (s *IncomeServices) GetAllIncomes() ([]models.Income, error) {
	incomes, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return incomes, nil
}

func (s *IncomeServices) CreateIncome(income models.Income) error {
	return s.Repo.Create(income)
}
