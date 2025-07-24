package controllers

import (
	"homeassignment/dto"
	"homeassignment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAccountHandler struct {
	Service services.AccountService
}

func NewAccountHandler(svc services.AccountService) *IAccountHandler {
	return &IAccountHandler{Service: svc}
}

func (h *IAccountHandler) GetAccountInfo(c *gin.Context) {
	accounts, err := h.Service.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (h *IAccountHandler) GetAccountById(c *gin.Context) {
	var req dto.GetAccountByIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	account, err := h.Service.FindAccountById(req.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account", "details": err.Error()})
		return
	}
	//Send the successful JSON response using c.JSON()
	c.JSON(http.StatusOK, account)
}
