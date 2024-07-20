package services_test

import (
	"context"
	daomocks "github.com/in-rich/uservice-discussions/pkg/dao/mocks"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListDiscussionsByTeam(t *testing.T) {
	testData := []struct {
		name string

		in *models.ListDiscussionsByTeamRequest

		shouldCallListDiscussionsByTeam bool
		listDiscussionsByTeamResponse   []*entities.Discussion
		listDiscussionsByTeamErr        error

		expect    []*models.Discussion
		expectErr error
	}{
		{
			name: "ListDiscussionsByTeam",
			in: &models.ListDiscussionsByTeamRequest{
				TeamID: "team-id-1",
				Limit:  10,
				Offset: 0,
			},
			shouldCallListDiscussionsByTeam: true,
			listDiscussionsByTeamResponse: []*entities.Discussion{
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-3",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetCompany,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-2",
					Target:           entities.TargetUser,
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
			expect: []*models.Discussion{
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-3",
					Target:           "user",
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           "company",
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-2",
					Target:           "user",
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "ListDiscussionsByTeamError",
			in: &models.ListDiscussionsByTeamRequest{
				TeamID: "team-id-1",
				Limit:  10,
				Offset: 0,
			},
			shouldCallListDiscussionsByTeam: true,
			listDiscussionsByTeamErr:        FooErr,
			expectErr:                       FooErr,
		},
		{
			name: "InvalidData",
			in: &models.ListDiscussionsByTeamRequest{
				TeamID: "",
				Limit:  10,
				Offset: 0,
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			listDiscussionsByTeamRepository := daomocks.NewMockListDiscussionsByTeamRepository(t)

			if tt.shouldCallListDiscussionsByTeam {
				listDiscussionsByTeamRepository.On(
					"ListDiscussionsByTeam", context.TODO(), tt.in.TeamID, tt.in.Limit, tt.in.Offset,
				).Return(tt.listDiscussionsByTeamResponse, tt.listDiscussionsByTeamErr)
			}

			service := services.NewListDiscussionsByTeamService(listDiscussionsByTeamRepository)

			result, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, result)

			listDiscussionsByTeamRepository.AssertExpectations(t)
		})
	}
}
