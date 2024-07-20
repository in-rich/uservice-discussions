package models

import "time"

type Discussion struct {
	Target           string     `json:"target"`
	PublicIdentifier string     `json:"publicIdentifier"`
	TeamID           string     `json:"teamID"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}
