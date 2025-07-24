package repositories

import (
	"homeassignment/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetAllAccounts() ([]models.Account, error) {
	var account []models.Account
	err := r.db.
		Find(&account).
		Error
	return account, err
}

func (r *AccountRepository) GetAllAccountById(id int) (models.Account, error) {
	var account models.Account
	err := r.db.
		Where("id", id).
		Find(&account).
		Error
	return account, err
}
