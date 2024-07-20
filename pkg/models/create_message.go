package models

type CreateMessageRequest struct {
	Target           string `json:"target" validate:"required,oneof=company user"`
	PublicIdentifier string `json:"publicIdentifier" validate:"required,max=255"`
	AuthorID         string `json:"authorID" validate:"required,max=255"`
	TeamID           string `json:"teamID" validate:"required,max=255"`
	Content          string `json:"content" validate:"max=15000"`
}
