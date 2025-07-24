package controllers

import (
	"homeassignment/dto"
	"homeassignment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITransactionHandler struct {
	Service services.TransactionService
}

func NewTransactionHandler(svc services.TransactionService) *ITransactionHandler {
	return &ITransactionHandler{Service: svc}
}

func (h *ITransactionHandler) SubmitTransaction(c *gin.Context) {
	var req dto.SubmitTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	err := h.Service.SubmitTransaction(req.SourceAccountID, req.DestinationAccountID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account", "details": err.Error()})
		return
	}
	//Send the successful JSON response using c.JSON()
	c.JSON(http.StatusOK, nil)
}
