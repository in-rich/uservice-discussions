package entities

import (
	"github.com/uptrace/bun"
	"time"
)

type Discussion struct {
	bun.BaseModel `bun:"table:discussions"`

	Target           Target     `bun:"target,notnull"`
	PublicIdentifier string     `bun:"public_identifier,notnull"`
	TeamID           string     `bun:"team_id,notnull"`
	UpdatedAt        *time.Time `bun:"created_at"`
}
