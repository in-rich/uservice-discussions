package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type GetMessageRepository interface {
	GetMessage(ctx context.Context, messageID uuid.UUID) (*entities.Message, error)
}

type getMessageRepositoryImpl struct {
	db bun.IDB
}

func (r *getMessageRepositoryImpl) GetMessage(ctx context.Context, messageID uuid.UUID) (*entities.Message, error) {
	message := new(entities.Message)
	err := r.db.NewSelect().Model(message).
		Where("id = ?", messageID).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrMessageNotFound
		}

		return nil, err
	}

	return message, nil
}

func NewGetMessageRepository(db bun.IDB) GetMessageRepository {
	return &getMessageRepositoryImpl{
		db: db,
	}
}
