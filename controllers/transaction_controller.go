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

// SubmitTransaction godoc
// @Summary Submit a transaction between accounts
// @Description Transfers specified amount from source to destination account. Amount must be positive with up to 2 decimal places.
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body dto.SubmitTransaction true "Transaction details"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse "Invalid request data"
// @Failure 500 {object} dto.ErrorResponse "Insufficient funds"
// @Failure 500 {object} dto.ErrorResponse "Account not found"
// @Failure 400 {object} dto.ErrorResponse "Validation error"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /transactions [post]
func (h *ITransactionHandler) SubmitTransaction(c *gin.Context) {
	var req dto.SubmitTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.INVALID_DATA_ERROR, Details: err.Error()})
		return
	}
	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.VALIDATION_ERROR, Details: common.FormatValidationError(err)})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.VALIDATION_ERROR, "details": common.FormatValidationError(err)})
		return
	}
	if req.SourceAccountID == req.DestinationAccountID {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.VALIDATION_ERROR, Details: "Source and destination accounts cannot be the same"})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.VALIDATION_ERROR, "details": "Source and destination accounts cannot be the same"})
		return
	}
	amount, err := strconv.ParseFloat(req.Amount, 64) //Convert to float64
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.CONVERSION_ERROR, Details: err.Error()})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.CONVERSION_ERROR, "details": err.Error()})
		return
	}
	err = h.Service.SubmitTransaction(req.SourceAccountID, req.DestinationAccountID, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: common.DATA_SUBMISSION_FAILURE, Details: err.Error()})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": common.DATA_SUBMISSION_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
