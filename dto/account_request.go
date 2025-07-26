package dto

type CreateAccountRequest struct {
	AccountID int    `json:"account_id" validate:"required"`
	Balance   string `json:"initial_balance" validate:"required,gte=0,decimal2"`
}
