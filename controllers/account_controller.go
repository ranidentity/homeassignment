package controllers

import (
	"homeassignment/common"
	"homeassignment/dto"
	"homeassignment/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IAccountHandler struct {
	Service   services.AccountService
	Validator *validator.Validate
}

func NewAccountHandler(svc services.AccountService) *IAccountHandler {
	return &IAccountHandler{Service: svc}
}

// Expected response is either an error or an empty response
func (h *IAccountHandler) CreateAccount(c *gin.Context) {
	var req dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_DATA_ERROR, "details": err.Error()})
		return
	}
	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.VALIDATION_ERROR, "details": err.Error()})
		return
	}

	err := h.Service.CreateAccount(req.AccountID, req.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.ACCOUNT_CREATION_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *IAccountHandler) GetAccountInfo(c *gin.Context) {
	accounts, err := h.Service.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.DATA_RETRIEVE_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (h *IAccountHandler) GetAccountById(c *gin.Context) {
	accountIDStr := c.Param("account_id")
	if accountIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_DATA_ERROR, "details": "Missing account id"})
		return
	}
	accountID, _ := strconv.Atoi(accountIDStr)
	account, err := h.Service.FindAccountById(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.DATA_RETRIEVE_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}
