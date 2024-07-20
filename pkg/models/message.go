package models

import "time"

type Message struct {
	ID               string     `json:"id"`
	Target           string     `json:"target"`
	PublicIdentifier string     `json:"publicIdentifier"`
	AuthorID         string     `json:"authorID"`
	TeamID           string     `json:"teamID"`
	Content          string     `json:"content"`
	CreatedAt        *time.Time `json:"createdAt"`
}
