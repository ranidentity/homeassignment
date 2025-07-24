package services

import (
	"homeassignment/models"
	"homeassignment/repositories"
)

type AccountService struct {
	repo repositories.AccountRepository
}

func (s *AccountService) GetAllAccounts() ([]models.Account, error) {
	// Add business logic here if needed
	return s.repo.GetAllAccounts()
}
func (s *AccountService) FindAccountById(id int) (models.Account, error) {
	// Add business logic here if needed
	return s.repo.GetAllAccountById(id)
}
