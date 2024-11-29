package models

type DiscussionReadStatus struct {
	Target              string `json:"target"`
	PublicIdentifier    string `json:"publicIdentifier"`
	TeamID              string `json:"teamID"`
	LatestReadMessageID string `json:"latestReadMessageID"`
	UserID              string `json:"userID"`
	HasUnreadMessages   bool   `json:"hasUnreadMessages"`
}
