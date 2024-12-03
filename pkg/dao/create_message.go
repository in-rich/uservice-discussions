package dao

import (
	"context"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
	"time"
)

type CreateMessageData struct {
	PublicIdentifier string
	Target           entities.Target
	Content          string
	CreatedAt        *time.Time
}

type CreateMessageRepository interface {
	CreateMessage(ctx context.Context, authorID string, teamID string, in *CreateMessageData) (*entities.Message, error)
}

type createMessageRepositoryImpl struct {
	db bun.IDB
}

func (r *createMessageRepositoryImpl) CreateMessage(ctx context.Context, authorID string, teamID string, in *CreateMessageData) (*entities.Message, error) {
	message := &entities.Message{
		AuthorID:         authorID,
		TeamID:           teamID,
		PublicIdentifier: in.PublicIdentifier,
		Target:           in.Target,
		Content:          in.Content,
		CreatedAt:        in.CreatedAt,
	}

	_, err := r.db.NewInsert().Model(message).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func NewCreateMessageRepository(db bun.IDB) CreateMessageRepository {
	return &createMessageRepositoryImpl{
		db: db,
	}
}
