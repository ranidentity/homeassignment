package dto

type SubmitTransaction struct {
	SourceAccountID      int    `json:"source_account_id" validate:"required" `
	DestinationAccountID int    `json:"destination_account_id" validate:"required"`
	Amount               string `json:"amount" validate:"required,decimal2"`
}
