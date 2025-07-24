package repositories

import (
	"homeassignment/models"

	"gorm.io/gorm"
)

//	type AccountRepository struct {
//		db *gorm.DB
//	}
type AccountRepository interface {
	GetAllAccounts() ([]models.Account, error)
	GetAccountById(id int) (models.Account, error)
}
type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
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
		Where("id", id).
		Find(&account).
		Error
	return account, err
}
