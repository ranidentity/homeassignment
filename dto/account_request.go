package dto

type CreateAccountRequest struct {
	AccountID int     `json:"account_id" validate:"required"`
	Balance   float64 `json:"initial_balance" validate:"required,gte=0"`
}
