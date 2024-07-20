package models

type DeleteMessageRequest struct {
	ID string `json:"id" validate:"required,max=255"`
}
