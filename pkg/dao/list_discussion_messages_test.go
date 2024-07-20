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

var listDiscussionMessagesFixtures = []*entities.Message{
	// Discussion about user public-identifier-1
	{
		ID:               uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		AuthorID:         "author-id-1",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:               uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		AuthorID:         "author-id-2",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-2",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:               uuid.MustParse("00000000-0000-0000-0000-000000000003"),
		AuthorID:         "author-id-3",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-3",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
	},

	// Discussion about user public-identifier-2
	{
		ID:               uuid.MustParse("00000000-0000-0000-0000-000000000004"),
		AuthorID:         "author-id-4",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-2",
		Target:           entities.TargetUser,
		Content:          "content-4",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
}

func TestListDiscussionMessages(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID           string
		publicIdentifier string
		target           entities.Target
		limit            int
		offset           int

		expect    []*entities.Message
		expectErr error
	}{
		{
			name:             "ListDiscussionMessagesRepository",
			teamID:           "team-id-1",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetUser,
			limit:            50,
			expect: []*entities.Message{
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					AuthorID:         "author-id-2",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-2",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					AuthorID:         "author-id-3",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-3",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					AuthorID:         "author-id-1",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-1",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:             "ListDiscussionMessagesRepository/Limit",
			teamID:           "team-id-1",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetUser,
			limit:            2,
			expect: []*entities.Message{
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					AuthorID:         "author-id-2",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-2",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					AuthorID:         "author-id-3",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-3",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:             "ListDiscussionMessagesRepository/Offset",
			teamID:           "team-id-1",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetUser,
			limit:            50,
			offset:           1,
			expect: []*entities.Message{
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					AuthorID:         "author-id-3",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-3",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					AuthorID:         "author-id-1",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-1",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:             "ListDiscussionMessagesRepository/NotFound",
			teamID:           "team-id-1",
			publicIdentifier: "public-identifier-3",
			target:           entities.TargetUser,
			limit:            50,
			expect:           []*entities.Message{},
		},
	}

	stx := BeginTX(db, listDiscussionMessagesFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewListDiscussionMessagesRepository(tx)

			messages, err := repo.ListDiscussionMessages(context.Background(), tt.teamID, tt.publicIdentifier, tt.target, tt.limit, tt.offset)

			require.ErrorIs(t, err, tt.expectErr)
			require.ElementsMatch(t, tt.expect, messages)
		})
	}
}
