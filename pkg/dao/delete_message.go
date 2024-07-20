package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type DeleteMessageRepository interface {
	DeleteMessage(ctx context.Context, messageID uuid.UUID) error
}

type deleteMessageRepositoryImpl struct {
	db bun.IDB
}

func (r *deleteMessageRepositoryImpl) DeleteMessage(ctx context.Context, messageID uuid.UUID) error {
	_, err := r.db.NewDelete().
		Model((*entities.Message)(nil)).
		Where("id = ?", messageID).
		Exec(ctx)

	return err
}

func NewDeleteMessageRepository(db bun.IDB) DeleteMessageRepository {
	return &deleteMessageRepositoryImpl{
		db: db,
	}
}
