package services

import (
	"homeassignment/models"
	"homeassignment/repositories"
)

type AccountService interface {
	GetAllAccounts() ([]models.Account, error)
	FindAccountById(id int) (models.Account, error)
}
type accountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(r repositories.AccountRepository) AccountService {
	return &accountService{repo: r}
}

func (s *accountService) GetAllAccounts() ([]models.Account, error) {
	// Add business logic here if needed
	return s.repo.GetAllAccounts()
}
func (s *accountService) FindAccountById(id int) (models.Account, error) {
	// Add business logic here if needed
	return s.repo.GetAccountById(id)
}
