package dao

import (
	"context"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type ListDiscussionMessagesRepository interface {
	ListDiscussionMessages(
		ctx context.Context, teamID string, publicIdentifier string, target entities.Target, limit int, offset int,
	) ([]*entities.Message, error)
}

type listDiscussionMessagesImpl struct {
	db bun.IDB
}

func (r *listDiscussionMessagesImpl) ListDiscussionMessages(
	ctx context.Context, teamID string, publicIdentifier string, target entities.Target, limit int, offset int,
) ([]*entities.Message, error) {
	messages := make([]*entities.Message, 0)

	err := r.db.NewSelect().
		Model(&messages).
		Where("team_id = ?", teamID).
		Where("public_identifier = ?", publicIdentifier).
		Where("target = ?", target).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(ctx)

	return messages, err
}

func NewListDiscussionMessagesRepository(db bun.IDB) ListDiscussionMessagesRepository {
	return &listDiscussionMessagesImpl{
		db: db,
	}
}
