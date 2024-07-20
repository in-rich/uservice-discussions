package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

var upsertDiscussionReadStatusFixtures = []*entities.ReadStatus{
	{
		UserID:              "user_id",
		TeamID:              "team_id",
		Target:              entities.TargetUser,
		PublicIdentifier:    "public_identifier",
		LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
	},
}

func TestUpsertDiscussionReadStatus(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID           string
		userID           string
		target           entities.Target
		publicIdentifier string
		messageID        uuid.UUID

		expect    *entities.ReadStatus
		expectErr error
	}{
		{
			name:             "CreateDiscussionReadStatus",
			teamID:           "team_id",
			userID:           "user_id",
			target:           entities.TargetUser,
			publicIdentifier: "public_identifier_2",
			messageID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			expect: &entities.ReadStatus{
				UserID:              "user_id",
				TeamID:              "team_id",
				Target:              entities.TargetUser,
				PublicIdentifier:    "public_identifier_2",
				LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name:             "UpdateDiscussionReadStatus",
			teamID:           "team_id",
			userID:           "user_id",
			target:           entities.TargetUser,
			publicIdentifier: "public_identifier",
			messageID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			expect: &entities.ReadStatus{
				UserID:              "user_id",
				TeamID:              "team_id",
				Target:              entities.TargetUser,
				PublicIdentifier:    "public_identifier",
				LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
	}

	stx := BeginTX(db, upsertDiscussionReadStatusFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewUpsertDiscussionReadStatusRepository(tx)

			readStatus, err := repo.UpsertDiscussionReadStatus(
				context.TODO(), tt.teamID, tt.userID, tt.target, tt.publicIdentifier, tt.messageID,
			)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, readStatus)
		})
	}
}
