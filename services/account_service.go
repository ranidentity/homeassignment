package services

import (
	"homeassignment/dto"
	"homeassignment/models"
	"homeassignment/repositories"
	"strconv"
)

type AccountService interface {
	GetAllAccounts() ([]models.Account, error)
	FindAccountById(id int) (dto.AccountDetail, error)
	CreateAccount(accountID int, balance float64) error
}
type accountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(r repositories.AccountRepository) AccountService {
	return &accountService{repo: r}
}

func (s *accountService) CreateAccount(accountID int, balance float64) error {
	return s.repo.NewAccount(accountID, balance)
}

func (s *accountService) FindAccountById(id int) (dto.AccountDetail, error) {
	var (
		result dto.AccountDetail
		err    error
	)
	account, err := s.repo.GetAccountById(id)
	if err == nil {
		result = dto.AccountDetail{
			AccountID: int(account.AccountID),
			Balance:   strconv.FormatFloat(account.Balance, 'f', 2, 64),
		}
	}
	return result, err
}

func (s *accountService) GetAllAccounts() ([]models.Account, error) {
	return s.repo.GetAllAccounts()
}
