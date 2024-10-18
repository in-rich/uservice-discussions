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

var listDiscussionsByTeamFixtures = []*entities.Message{
	// Discussion 1 of team 1.
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		AuthorID:         "author-id-1",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		AuthorID:         "author-id-2",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-2",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},

	// Discussion 2 of team 1.
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		AuthorID:         "author-id-3",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-2",
		Target:           entities.TargetUser,
		Content:          "content-3",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
	},

	// Discussion 3 of team 1.
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000004")),
		AuthorID:         "author-id-4",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-3",
		Target:           entities.TargetUser,
		Content:          "content-4",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
	},

	// Discussion 1 of team 2.
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000005")),
		AuthorID:         "author-id-5",
		TeamID:           "team-id-2",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-5",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
}

func TestListDiscussionsByTeam(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID string
		limit  int
		offset int

		expect    []*entities.Discussion
		expectErr error
	}{
		{
			name:   "ListDiscussionsByTeam",
			teamID: "team-id-1",
			limit:  50,
			expect: []*entities.Discussion{
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-3",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-2",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "ListDiscussionsByTeam/Limit",
			teamID: "team-id-1",
			limit:  2,
			expect: []*entities.Discussion{
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-3",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "ListDiscussionsByTeam/Offset",
			teamID: "team-id-1",
			limit:  50,
			offset: 1,
			expect: []*entities.Discussion{
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-2",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "ListDiscussionsByTeam/Empty",
			teamID: "team-id-3",
			limit:  50,
			expect: []*entities.Discussion{},
		},
	}

	stx := BeginTX(db, listDiscussionsByTeamFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewListDiscussionsByTeamRepository(tx)

			discussions, err := repo.ListDiscussionsByTeam(context.Background(), tt.teamID, tt.limit, tt.offset)

			require.Equal(t, tt.expect, discussions)
			require.ErrorIs(t, err, tt.expectErr)
		})
	}
}
