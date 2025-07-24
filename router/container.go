package router

import (
	"homeassignment/controllers"
	"homeassignment/repositories"
	"homeassignment/services"

	"gorm.io/gorm"
)

type Container struct {
	AccountHandler     controllers.IAccountHandler
	TransactionHandler controllers.ITransactionHandler
}

func Build(db *gorm.DB) *Container {
	// Repository
	accountRepo := repositories.NewAccountRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	// Service
	accountSvc := services.NewAccountService(accountRepo)
	transactionSvc := services.NewTransactionService(transactionRepo)

	// Handler
	accountHandler := controllers.IAccountHandler{
		Service: accountSvc,
	}
	transactionHandler := controllers.ITransactionHandler{
		Service: transactionSvc,
	}
	return &Container{
		AccountHandler:     accountHandler,
		TransactionHandler: transactionHandler,
	}
}
