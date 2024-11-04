package dao

import (
	"context"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type ListDiscussionsByTeamRepository interface {
	ListDiscussionsByTeam(
		ctx context.Context, teamID string, limit int, offset int,
	) ([]*entities.Discussion, error)
}

type listDiscussionsByTeamRepositoryImpl struct {
	db bun.IDB
}

func (r *listDiscussionsByTeamRepositoryImpl) ListDiscussionsByTeam(
	ctx context.Context, teamID string, limit int, offset int,
) ([]*entities.Discussion, error) {
	discussions := make([]*entities.Discussion, 0)

	// List info from most recent message in each discussion.
	// We have to use a subquery, because we can't first order by created_at and then distinct on public_identifier.
	mostRecentMessagePerDiscussion := r.db.NewSelect().
		Model((*entities.Message)(nil)).
		Column("public_identifier", "team_id", "target", "created_at").
		// Only one message per discussion. Use the order clause to get the most recent message.
		DistinctOn("public_identifier, target").
		// Pre-filter by team_id.
		Where("team_id = ?", teamID).
		// We have to order first using the columns in the DISTINCT clause, then we order using created_at, so we
		// only get the most recent message.
		Order("public_identifier DESC", "target DESC", "created_at DESC")

	query := r.db.NewSelect().
		Model(&discussions).
		With("discussions", mostRecentMessagePerDiscussion).
		// Select the most recently updated discussions.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset)

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return discussions, nil
}

func NewListDiscussionsByTeamRepository(db bun.IDB) ListDiscussionsByTeamRepository {
	return &listDiscussionsByTeamRepositoryImpl{
		db: db,
	}
}
