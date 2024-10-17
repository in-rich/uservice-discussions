package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-discussions/pkg/dao/mocks"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListDiscussionMessages(t *testing.T) {
	testData := []struct {
		name string

		in *models.ListDiscussionMessagesRequest

		shouldCallListDiscussionMessages bool
		listDiscussionMessagesResponse   []*entities.Message
		listDiscussionMessagesErr        error

		expect    []*models.Message
		expectErr error
	}{
		{
			name: "ListDiscussionMessages",
			in: &models.ListDiscussionMessagesRequest{
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           "user",
				Limit:            10,
				Offset:           0,
			},
			shouldCallListDiscussionMessages: true,
			listDiscussionMessagesResponse: []*entities.Message{
				{
					ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					AuthorID:         "author-id-2",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-2",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
					AuthorID:         "author-id-3",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-3",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					AuthorID:         "author-id-1",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           entities.TargetUser,
					Content:          "content-1",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
			expect: []*models.Message{
				{
					ID:               "00000000-0000-0000-0000-000000000002",
					AuthorID:         "author-id-2",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           "user",
					Content:          "content-2",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               "00000000-0000-0000-0000-000000000003",
					AuthorID:         "author-id-3",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           "user",
					Content:          "content-3",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               "00000000-0000-0000-0000-000000000001",
					AuthorID:         "author-id-1",
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           "user",
					Content:          "content-1",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "ListDiscussionMessagesError",
			in: &models.ListDiscussionMessagesRequest{
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           "user",
				Limit:            10,
				Offset:           0,
			},
			shouldCallListDiscussionMessages: true,
			listDiscussionMessagesErr:        FooErr,
			expectErr:                        FooErr,
		},
		{
			name: "InvalidData",
			in: &models.ListDiscussionMessagesRequest{
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           "invalid-target",
				Limit:            10,
				Offset:           0,
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			listDiscussionMessagesRepository := daomocks.NewMockListDiscussionMessagesRepository(t)

			if tt.shouldCallListDiscussionMessages {
				listDiscussionMessagesRepository.On(
					"ListDiscussionMessages", context.TODO(),
					tt.in.TeamID, tt.in.PublicIdentifier, entities.Target(tt.in.Target), tt.in.Limit, tt.in.Offset,
				).Return(tt.listDiscussionMessagesResponse, tt.listDiscussionMessagesErr)
			}

			service := services.NewListDiscussionMessagesService(listDiscussionMessagesRepository)

			result, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, result)

			listDiscussionMessagesRepository.AssertExpectations(t)
		})
	}
}
