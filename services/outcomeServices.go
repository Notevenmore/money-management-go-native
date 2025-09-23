package services

import (
	"money-management-be/models"
	"money-management-be/repositories"
)

type OutcomeServices struct {
	Repo *repositories.OutcomeRepository
}

func NewOutcomeServices(repo *repositories.OutcomeRepository) *OutcomeServices {
	return &OutcomeServices{Repo: repo}
}

func (s *OutcomeServices) GetAllOutcomes() ([]models.Outcome, error) {
	outcomes, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return outcomes, nil
}

func (s *OutcomeServices) CreateOutcome(outcome models.Outcome) error {
	return s.Repo.Create(outcome)
}
