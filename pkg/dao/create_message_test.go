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

var createMessageFixtures = []*entities.Message{
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		AuthorID:         "author-id-1",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestCreateMessage(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testStart := time.Now()

	testData := []struct {
		name string

		authorID string
		teamID   string

		data *dao.CreateMessageData

		expect          *entities.Message
		expectTimestamp *time.Time
		expectErr       error
	}{
		{
			name:     "CreateMessage/NewDiscussion",
			authorID: "author-id-2",
			teamID:   "team-id-2",
			data: &dao.CreateMessageData{
				PublicIdentifier: "public-identifier-2",
				Target:           entities.TargetUser,
				Content:          "content-2",
			},
			expect: &entities.Message{
				AuthorID:         "author-id-2",
				TeamID:           "team-id-2",
				PublicIdentifier: "public-identifier-2",
				Target:           entities.TargetUser,
				Content:          "content-2",
			},
		},
		{
			name:     "CreateMessage/SameDiscussion",
			authorID: "author-id-1",
			teamID:   "team-id-1",
			data: &dao.CreateMessageData{
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
				Content:          "content-1",
			},
			expect: &entities.Message{
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
				Content:          "content-1",
			},
		},
		{
			name:     "CreateMessage/WithTimestamp",
			authorID: "author-id-2",
			teamID:   "team-id-2",
			data: &dao.CreateMessageData{
				PublicIdentifier: "public-identifier-2",
				Target:           entities.TargetUser,
				Content:          "content-2",
				CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectTimestamp: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			expect: &entities.Message{
				AuthorID:         "author-id-2",
				TeamID:           "team-id-2",
				PublicIdentifier: "public-identifier-2",
				Target:           entities.TargetUser,
				Content:          "content-2",
			},
		},
	}

	stx := BeginTX(db, createMessageFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewCreateMessageRepository(tx)

			message, err := repo.CreateMessage(context.Background(), tt.authorID, tt.teamID, tt.data)

			if message != nil {
				require.NotNil(t, message.ID)
				require.NotNil(t, message.CreatedAt)

				if tt.expectTimestamp != nil {
					require.Equal(t, tt.expectTimestamp, message.CreatedAt)
				} else {
					require.True(t, message.CreatedAt.After(testStart))
				}

				message.ID = nil
				message.CreatedAt = nil
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, message)
		})
	}
}
