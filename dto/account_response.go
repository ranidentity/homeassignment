package dto

type AccountDetail struct {
	AccountID int    `json:"account_id" validate:"required"`
	Balance   string `json:"balance" validate:"required"`
}
