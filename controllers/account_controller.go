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

// CreateAccount godoc
// @Summary      Create new account
// @Description  Create new account with initial balance. Account ID must be unique.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateAccountRequest true "Account creation data"
// @Success      201  "Account created successfully"
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse "Account already exists"
// @Router       /accounts [post]
func (h *IAccountHandler) CreateAccount(c *gin.Context) {
	var req dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.INVALID_DATA_ERROR, Details: err.Error()})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_DATA_ERROR, "details": err.Error()})
		return
	}
	if err := h.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.VALIDATION_ERROR, Details: common.FormatValidationError(err)})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.VALIDATION_ERROR, "details": common.FormatValidationError(err)})
		return
	}
	balance, err := strconv.ParseFloat(req.Balance, 64) //Convert to float64
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.CONVERSION_ERROR, Details: err.Error()})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.CONVERSION_ERROR, "details": err.Error()})
		return
	}
	err = h.Service.CreateAccount(req.AccountID, balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: common.ACCOUNT_CREATION_FAILURE, Details: err.Error()})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": common.ACCOUNT_CREATION_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, err)
}

// GetAccountById godoc
// @Summary      Get account by ID
// @Description  Returns a single account by ID. Account ID must be a positive integer.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account_id path int true "Account ID (positive integer)"
// @Success      200  {object}  dto.AccountDetail "Account details"
// @Failure      400  {object}  dto.ErrorResponse "Invalid account ID format"
// @Failure      500  {object}  dto.ErrorResponse "Account doesn't exists"
// @Router       /accounts/{account_id} [get]
func (h *IAccountHandler) GetAccountById(c *gin.Context) {
	accountIDStr := c.Param("account_id")
	if accountIDStr == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.INVALID_DATA_ERROR, Details: "Missing account id"})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_DATA_ERROR, "details": "Missing account id"})
		return
	}
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil || accountID <= 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: common.INVALID_DATA_ERROR, Details: "Account ID must be a positive integer"})
		// c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_DATA_ERROR, "details": "Account ID must be a positive integer"})
		return
	}
	account, err := h.Service.FindAccountById(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: common.DATA_RETRIEVE_FAILURE, Details: err.Error()})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": common.DATA_RETRIEVE_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *IAccountHandler) GetAccountInfo(c *gin.Context) {
	accounts, err := h.Service.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: common.DATA_RETRIEVE_FAILURE, Details: err.Error()})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": common.DATA_RETRIEVE_FAILURE, "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}
