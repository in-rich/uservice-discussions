package dao

import (
	"context"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type GetAllMessagesRepository interface {
	GetAllMessages(ctx context.Context, limit int64, offset int64) ([]*entities.Message, error)
}

type getAllMessagesRepositoryImpl struct {
	db bun.IDB
}

func (r *getAllMessagesRepositoryImpl) GetAllMessages(ctx context.Context, limit int64, offset int64) ([]*entities.Message, error) {
	messages := make([]*entities.Message, 0)

	err := r.db.NewSelect().
		Model(&messages).
		Limit(int(limit)).
		Offset(int(offset)).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func NewGetAllMessagesRepository(db bun.IDB) GetAllMessagesRepository {
	return &getAllMessagesRepositoryImpl{
		db: db,
	}
}
