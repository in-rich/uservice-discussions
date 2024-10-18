package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-discussions/pkg/dao/mocks"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateDiscussionReadStatus(t *testing.T) {
	testData := []struct {
		name string

		in *models.UpdateDiscussionReadStatusRequest

		shouldCallGetMessage bool
		getMessageResponse   *entities.Message
		getMessageErr        error

		shouldCallUpdateDiscussionReadStatus bool
		updateDiscussionReadStatusResponse   *entities.ReadStatus
		updateDiscussionReadStatusErr        error

		expect    *models.DiscussionReadStatus
		expectErr error
	}{
		{
			name: "UpdateDiscussionReadStatus",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageResponse: &entities.Message{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetCompany,
			},
			shouldCallUpdateDiscussionReadStatus: true,
			updateDiscussionReadStatusResponse: &entities.ReadStatus{
				TeamID:              "team-id-1",
				UserID:              "user-id-1",
				Target:              entities.TargetCompany,
				PublicIdentifier:    "public-identifier-1",
				LatestReadMessageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
			expect: &models.DiscussionReadStatus{
				TeamID:              "team-id-1",
				UserID:              "user-id-1",
				Target:              "company",
				PublicIdentifier:    "public-identifier-1",
				LatestReadMessageID: "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			name: "UpdateDiscussionReadStatusError",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageResponse: &entities.Message{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetCompany,
			},
			shouldCallUpdateDiscussionReadStatus: true,
			updateDiscussionReadStatusErr:        FooErr,
			expectErr:                            FooErr,
		},
		{
			name: "MessageInWrongTeam",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageResponse: &entities.Message{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				TeamID:           "team-id-2",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetCompany,
			},
			expectErr: services.ErrMessageInWrongDiscussion,
		},
		{
			name: "MessageInWrongPublicIdentifier",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageResponse: &entities.Message{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-2",
				Target:           entities.TargetCompany,
			},
			expectErr: services.ErrMessageInWrongDiscussion,
		},
		{
			name: "MessageInWrongTarget",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageResponse: &entities.Message{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
			},
			expectErr: services.ErrMessageInWrongDiscussion,
		},
		{
			name: "GetMessageErr",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageErr:        FooErr,
			expectErr:            FooErr,
		},
		{
			name: "InvalidData",
			in: &models.UpdateDiscussionReadStatusRequest{
				TeamID:           "team-id-1",
				UserID:           "user-id-1",
				Target:           "invalid-target",
				PublicIdentifier: "public-identifier-1",
				MessageID:        "invalid-message-id",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getMessageRepository := daomocks.NewMockGetMessageRepository(t)
			updateDiscussionReadStatusRepository := daomocks.NewMockUpsertDiscussionReadStatusRepository(t)

			if tt.shouldCallGetMessage {
				getMessageRepository.
					On("GetMessage", context.TODO(), mock.Anything).
					Return(tt.getMessageResponse, tt.getMessageErr)
			}

			if tt.shouldCallUpdateDiscussionReadStatus {
				updateDiscussionReadStatusRepository.
					On(
						"UpsertDiscussionReadStatus", context.TODO(),
						tt.in.TeamID, tt.in.UserID, entities.Target(tt.in.Target), tt.in.PublicIdentifier, mock.Anything,
					).
					Return(tt.updateDiscussionReadStatusResponse, tt.updateDiscussionReadStatusErr)
			}

			service := services.NewUpdateDiscussionReadStatusService(
				updateDiscussionReadStatusRepository,
				getMessageRepository,
			)

			resp, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, resp)

			getMessageRepository.AssertExpectations(t)
			updateDiscussionReadStatusRepository.AssertExpectations(t)
		})
	}
}
