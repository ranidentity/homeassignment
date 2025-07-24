package services

import (
	"homeassignment/repositories"
	"log"
)

type TransactionService interface {
	SubmitTransaction(sourceAccountID int, destinationAccountID int, amount float64) error
}
type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(r repositories.TransactionRepository) TransactionService {
	return &transactionService{repo: r}
}
func (s *transactionService) SubmitTransaction(sourceAccountID int, destinationAccountID int, amount float64) error {
	err := s.repo.DoTransaction(sourceAccountID, destinationAccountID, amount)
	if err != nil {
		log.Printf("Transaction failed: %v", err)
	}
	return err
}
