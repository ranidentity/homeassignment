package repositories

import (
	"errors"
	"fmt"
	"homeassignment/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *AccountRepository) DoTransaction(sourceUserID int, destinationAccountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		// check if sender and receiver exist
		var source models.Account
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", sourceUserID).
			First(&source).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("source account not found: %w", err) // More specific error
			}
			return fmt.Errorf("failed to retrieve source account: %w", err)
		}
		if source.Balance < amount {
			return errors.New("insufficient balance")
		}
		var destination models.Account
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", destinationAccountID).
			First(&destination).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("destination account not found")
			}
			return err
		}

		newTransaction := models.Transaction{
			SourceAccountID:      int64(sourceUserID),
			DestinationAccountID: int64(destinationAccountID),
			Amount:               amount,
		}
		if err := tx.Create(&newTransaction).Error; err != nil {
			return err
		}

		// 4. Update balances
		if err := tx.Model(&models.Account{}).
			Where("id = ?", sourceUserID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Account{}).
			Where("id = ?", destinationAccountID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}
		return nil
	})
}
