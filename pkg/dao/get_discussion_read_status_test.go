package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var getDiscussionReadStatusFixtures = []interface{}{
	&entities.ReadStatus{
		UserID:              "user_id",
		TeamID:              "team_id",
		Target:              entities.TargetUser,
		PublicIdentifier:    "public_identifier",
		LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
	},
	&entities.ReadStatus{
		UserID:              "user_id",
		TeamID:              "team_id_2",
		Target:              entities.TargetUser,
		PublicIdentifier:    "public_identifier",
		LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	},

	&entities.Message{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		TeamID:           "team_id",
		Target:           entities.TargetUser,
		PublicIdentifier: "public_identifier",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
	&entities.Message{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		TeamID:           "team_id_2",
		Target:           entities.TargetUser,
		PublicIdentifier: "public_identifier",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
	},
	&entities.Message{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		TeamID:           "team_id_2",
		Target:           entities.TargetUser,
		PublicIdentifier: "public_identifier",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
}

func TestGetDiscussionReadStatus(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID           string
		userID           string
		target           entities.Target
		publicIdentifier string

		expect    *entities.ReadStatus
		expectErr error
	}{
		{
			name:             "GetDiscussionReadStatus",
			teamID:           "team_id",
			userID:           "user_id",
			target:           entities.TargetUser,
			publicIdentifier: "public_identifier",
			expect: &entities.ReadStatus{
				UserID:              "user_id",
				TeamID:              "team_id",
				Target:              entities.TargetUser,
				PublicIdentifier:    "public_identifier",
				LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				HasUnreadMessages:   false,
			},
		},
		{
			name:             "NotFound",
			teamID:           "team_id",
			userID:           "user_id",
			target:           entities.TargetUser,
			publicIdentifier: "public_identifier_2",
			expectErr:        dao.ErrDiscussionReadStatusNotFound,
		},
		{
			name:             "HasUnreadMessages",
			teamID:           "team_id_2",
			userID:           "user_id",
			target:           entities.TargetUser,
			publicIdentifier: "public_identifier",
			expect: &entities.ReadStatus{
				UserID:              "user_id",
				TeamID:              "team_id_2",
				Target:              entities.TargetUser,
				PublicIdentifier:    "public_identifier",
				LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				HasUnreadMessages:   true,
			},
		},
	}

	stx := BeginTX(db, getDiscussionReadStatusFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewGetDiscussionReadStatusRepository(tx)

			readStatus, err := repo.GetDiscussionReadStatus(
				context.TODO(), tt.teamID, tt.userID, tt.target, tt.publicIdentifier,
			)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, readStatus)
		})
	}
}
