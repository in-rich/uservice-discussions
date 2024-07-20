package models

type ListDiscussionMessagesRequest struct {
	Target           string `json:"target" validate:"required,oneof=company user"`
	PublicIdentifier string `json:"publicIdentifier" validate:"required,max=255"`
	TeamID           string `json:"teamID" validate:"required,max=255"`
	Limit            int    `json:"limit" validate:"required,min=1,max=1000"`
	Offset           int    `json:"offset" validate:"min=0"`
}
