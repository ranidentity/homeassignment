package repositories

import (
	"errors"
	"fmt"
	"homeassignment/models"

	"gorm.io/gorm"
)

//	type AccountRepository struct {
//		db *gorm.DB
//	}
type AccountRepository interface {
	GetAllAccounts() ([]models.Account, error)
	GetAccountById(id int) (models.Account, error)
	NewAccount(accountID int, balance float64) error
}
type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) NewAccount(accountID int, balance float64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingAccount models.Account
		err := tx.Where("account_id = ?", accountID).First(&existingAccount).Error
		if err == nil {
			return fmt.Errorf("%w: %d", errors.New("account with this ID already exists"), accountID)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// If the error is not ErrRecordNotFound, it's some other database error.
			return fmt.Errorf("failed to check for existing account: %w", err)
		}
		newAccount := models.Account{
			AccountID: int64(accountID),
			Balance:   balance,
		}
		if createErr := tx.Create(&newAccount).Error; createErr != nil {
			return fmt.Errorf("failed to create new account: %w", createErr)
		}
		return nil
	})
}

func (r *accountRepository) GetAllAccounts() ([]models.Account, error) {
	var account []models.Account
	err := r.db.
		Find(&account).
		Error
	return account, err
}

func (r *accountRepository) GetAccountById(id int) (models.Account, error) {
	var account models.Account
	err := r.db.
		Where("account_id", id).
		First(&account).
		Error
	return account, err
}
