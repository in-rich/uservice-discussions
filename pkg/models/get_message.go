package models

type GetMessageRequest struct {
	ID string `json:"id" validate:"required,max=255"`
}
