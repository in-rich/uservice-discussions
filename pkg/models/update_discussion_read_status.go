package models

type UpdateDiscussionReadStatusRequest struct {
	Target           string `json:"target" validate:"required,oneof=company user"`
	PublicIdentifier string `json:"publicIdentifier" validate:"required,max=255"`
	TeamID           string `json:"teamID" validate:"required,max=255"`
	UserID           string `json:"userID" validate:"required,max=255"`
	MessageID        string `json:"messageID" validate:"required,max=255"`
}
