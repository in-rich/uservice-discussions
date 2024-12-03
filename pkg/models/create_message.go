package models

import "time"

type CreateMessageRequest struct {
	Target           string     `json:"target" validate:"required,oneof=company user"`
	PublicIdentifier string     `json:"publicIdentifier" validate:"required,max=255"`
	AuthorID         string     `json:"authorID" validate:"required,max=255"`
	TeamID           string     `json:"teamID" validate:"required,max=255"`
	Content          string     `json:"content" validate:"required,max=15000"`
	CreatedAt        *time.Time `json:"createdAt"`
}
