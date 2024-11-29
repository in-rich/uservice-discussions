package dao

import (
	"github.com/uptrace/bun"
)

func mostRecentMessageForDiscussion(db bun.IDB, teamID string) *bun.SelectQuery {
	// List info from most recent message in each discussion.
	// We have to use a subquery, because we can't first order by created_at and then distinct on public_identifier.
	return db.NewSelect().
		// Only one message per discussion. Use the order clause to get the most recent message.
		DistinctOn("public_identifier, target").
		// Pre-filter by team_id.
		Where("team_id = ?", teamID).
		// We have to order first using the columns in the DISTINCT clause, then we order using created_at, so we
		// only get the most recent message.
		Order("public_identifier DESC", "target DESC", "created_at DESC")
}
