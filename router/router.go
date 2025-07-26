package router

import (
	"homeassignment/controllers"
	"homeassignment/repositories"
	"homeassignment/services"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}
type Container struct {
	AccountHandler     controllers.IAccountHandler
	TransactionHandler controllers.ITransactionHandler
}

func RegisterRoutes(container *Container) *gin.Engine {
	r := gin.Default()
	r.GET("/accounts/:account_id", container.AccountHandler.GetAccountById)
	r.GET("/accounts", container.AccountHandler.GetAccountInfo)
	r.POST("/accounts", container.AccountHandler.CreateAccount)
	r.POST("/transactions", container.TransactionHandler.SubmitTransaction)
	return r

}

func Build(db *gorm.DB, validate *validator.Validate) *Container {
	// Repository
	accountRepo := repositories.NewAccountRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	// Service
	accountSvc := services.NewAccountService(accountRepo)
	transactionSvc := services.NewTransactionService(transactionRepo)

	//custom
	if err := validate.RegisterValidation("decimal2", func(fl validator.FieldLevel) bool {
		amountStr := fl.Field().String()
		match, _ := regexp.MatchString(`^\d+(\.\d{1,2})?$`, amountStr)
		return match
	}); err != nil {
		log.Fatalf("Error registering decimal2 validation: %v", err)
	}

	// Handler
	accountHandler := controllers.IAccountHandler{
		Service:   accountSvc,
		Validator: validate,
	}
	transactionHandler := controllers.ITransactionHandler{
		Service:   transactionSvc,
		Validator: validate,
	}
	return &Container{
		AccountHandler:     accountHandler,
		TransactionHandler: transactionHandler,
	}
}
