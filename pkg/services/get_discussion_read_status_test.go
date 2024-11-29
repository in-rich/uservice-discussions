package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-discussions/pkg/dao/mocks"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDiscussionReadStatus(t *testing.T) {
	testData := []struct {
		name string

		in *models.GetDiscussionReadStatusRequest

		shouldCallGetDiscussionReadStatus bool
		getDiscussionReadStatusResponse   *entities.ReadStatus
		getDiscussionReadStatusErr        error

		expect    *models.DiscussionReadStatus
		expectErr error
	}{
		{
			name: "GetDiscussionReadStatus",
			in: &models.GetDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
			},
			shouldCallGetDiscussionReadStatus: true,
			getDiscussionReadStatusResponse: &entities.ReadStatus{
				TeamID:              "team-id-1",
				UserID:              "user-id-1",
				Target:              entities.TargetCompany,
				PublicIdentifier:    "public-identifier-1",
				LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				HasUnreadMessages:   true,
			},
			expect: &models.DiscussionReadStatus{
				TeamID:              "team-id-1",
				UserID:              "user-id-1",
				Target:              "company",
				PublicIdentifier:    "public-identifier-1",
				LatestReadMessageID: "00000000-0000-0000-0000-000000000001",
				HasUnreadMessages:   true,
			},
		},
		{
			name: "GetDiscussionReadStatusError",
			in: &models.GetDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
			},
			shouldCallGetDiscussionReadStatus: true,
			getDiscussionReadStatusErr:        FooErr,
			expectErr:                         FooErr,
		},
		{
			name: "InvalidData",
			in: &models.GetDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "invalid-target",
				PublicIdentifier: "public-identifier-1",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getDiscussionReadStatusRepository := daomocks.NewMockGetDiscussionReadStatusRepository(t)

			if tt.shouldCallGetDiscussionReadStatus {
				getDiscussionReadStatusRepository.
					On(
						"GetDiscussionReadStatus", context.TODO(),
						tt.in.TeamID, tt.in.UserID, entities.Target(tt.in.Target), tt.in.PublicIdentifier,
					).Return(tt.getDiscussionReadStatusResponse, tt.getDiscussionReadStatusErr)
			}

			service := services.NewGetDiscussionReadStatusService(getDiscussionReadStatusRepository)

			resp, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, resp)

			getDiscussionReadStatusRepository.AssertExpectations(t)
		})
	}
}
