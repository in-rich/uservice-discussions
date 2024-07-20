package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type UpsertDiscussionReadStatusRepository interface {
	UpsertDiscussionReadStatus(
		ctx context.Context, teamID string, userID string, target entities.Target, publicIdentifier string,
		messageID uuid.UUID,
	) (*entities.ReadStatus, error)
}

type upsertDiscussionReadStatusRepositoryImpl struct {
	db bun.IDB
}

func (r *upsertDiscussionReadStatusRepositoryImpl) UpsertDiscussionReadStatus(
	ctx context.Context, teamID string, userID string, target entities.Target, publicIdentifier string,
	messageID uuid.UUID,
) (*entities.ReadStatus, error) {
	readStatus := &entities.ReadStatus{
		TeamID:              teamID,
		UserID:              userID,
		Target:              target,
		PublicIdentifier:    publicIdentifier,
		LatestReadMessageID: messageID,
	}

	_, err := r.db.NewInsert().
		Model(readStatus).
		On("conflict (team_id, user_id, target, public_identifier) DO UPDATE").
		Set("latest_read_message_id = ?", messageID).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return readStatus, nil
}

func NewUpsertDiscussionReadStatusRepository(db bun.IDB) UpsertDiscussionReadStatusRepository {
	return &upsertDiscussionReadStatusRepositoryImpl{
		db: db,
	}
}
