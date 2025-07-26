package controllers

import (
	"homeassignment/common"
	"homeassignment/dto"
	"homeassignment/services"
	"net/http"
	"strconv"

	// Import the regexp package
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ITransactionHandler struct {
	Service   services.TransactionService
	Validator *validator.Validate
}

func NewTransactionHandler(svc services.TransactionService) *ITransactionHandler {
	return &ITransactionHandler{Service: svc}
}

func (h *ITransactionHandler) SubmitTransaction(c *gin.Context) {
	var req dto.SubmitTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_DATA_ERROR, "details": err.Error()})
		return
	}
	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.VALIDATION_ERROR, "details": common.FormatValidationError(err)})
		return
	}
	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.CONVERSION_ERROR, "details": err.Error()})
		return
	}
	err = h.Service.SubmitTransaction(req.SourceAccountID, req.DestinationAccountID, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": common.DATA_SUBMISSION_FAILURE, "details": err.Error()})
		return
	}
	//Send the successful JSON response using c.JSON()
	c.JSON(http.StatusOK, nil)
}
