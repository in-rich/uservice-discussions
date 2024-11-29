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

	mostRecentMessagePerDiscussion := mostRecentMessageForDiscussion(r.db, teamID).
		Column("public_identifier", "team_id", "target", "created_at").
		Model((*entities.Message)(nil))

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
