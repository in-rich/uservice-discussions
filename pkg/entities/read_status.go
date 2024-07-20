package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReadStatus struct {
	bun.BaseModel `bun:"table:read_statuses"`

	Target              Target    `bun:"target,notnull"`
	PublicIdentifier    string    `bun:"public_identifier,notnull"`
	TeamID              string    `bun:"team_id,notnull"`
	UserID              string    `bun:"user_id,notnull"`
	LatestReadMessageID uuid.UUID `bun:"latest_read_message_id"`
}
