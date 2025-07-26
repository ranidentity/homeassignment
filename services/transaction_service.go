package services

import (
	"homeassignment/repositories"
	"log"
	"sync"
)

type TransactionService interface {
	SubmitTransaction(sourceAccountID int, destinationAccountID int, amount float64) error
}
type transactionService struct {
	repo        repositories.TransactionRepository
	senderLocks map[int]*sync.Mutex
	mapLock     sync.Mutex
}

func NewTransactionService(r repositories.TransactionRepository) TransactionService {
	return &transactionService{
		repo:        r,
		senderLocks: make(map[int]*sync.Mutex),
	}
}
func (s *transactionService) SubmitTransaction(sourceAccountID int, destinationAccountID int, amount float64) error {
	// Get or create a mutex for the sender
	s.mapLock.Lock()
	if _, exists := s.senderLocks[sourceAccountID]; !exists {
		s.senderLocks[sourceAccountID] = &sync.Mutex{}
	}
	senderMutex := s.senderLocks[sourceAccountID]
	s.mapLock.Unlock()

	// Lock the sender-specific mutex to ensure only one transaction per sender
	senderMutex.Lock()
	defer senderMutex.Unlock()

	err := s.repo.DoTransaction(sourceAccountID, destinationAccountID, amount)
	if err != nil {
		log.Printf("Transaction failed for sender %d: %v", sourceAccountID, err)
	}
	return err
}
