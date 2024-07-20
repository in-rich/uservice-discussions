package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

var getDiscussionReadStatusFixtures = []*entities.ReadStatus{
	{
		UserID:              "user_id",
		TeamID:              "team_id",
		Target:              entities.TargetUser,
		PublicIdentifier:    "public_identifier",
		LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
