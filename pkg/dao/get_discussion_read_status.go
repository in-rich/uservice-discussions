package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/uptrace/bun"
)

type GetDiscussionReadStatusRepository interface {
	GetDiscussionReadStatus(
		ctx context.Context, teamID string, userID string, target entities.Target, publicIdentifier string,
	) (*entities.ReadStatus, error)
}

type getDiscussionReadStatusRepositoryImpl struct {
	db bun.IDB
}

func (r *getDiscussionReadStatusRepositoryImpl) GetDiscussionReadStatus(
	ctx context.Context, teamID string, userID string, target entities.Target, publicIdentifier string,
) (*entities.ReadStatus, error) {
	readStatus := new(entities.ReadStatus)
	err := r.db.NewSelect().Model(readStatus).
		Where("team_id = ?", teamID).
		Where("target = ?", target).
		Where("public_identifier = ?", publicIdentifier).
		Where("user_id = ?", userID).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDiscussionReadStatusNotFound
		}

		return nil, err
	}

	return readStatus, nil
}

func NewGetDiscussionReadStatusRepository(db bun.IDB) GetDiscussionReadStatusRepository {
	return &getDiscussionReadStatusRepositoryImpl{
		db: db,
	}
}
