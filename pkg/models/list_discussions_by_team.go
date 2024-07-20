package models

type ListDiscussionsByTeamRequest struct {
	TeamID string `json:"teamID" validate:"required,max=255"`
	Limit  int    `json:"limit" validate:"required,min=1,max=1000"`
	Offset int    `json:"offset" validate:"min=0"`
}
