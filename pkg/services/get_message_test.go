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
	"time"
)

func TestGetMessage(t *testing.T) {
	testData := []struct {
		name string

		in *models.GetMessageRequest

		shouldCallGetMessage bool
		getMessageResponse   *entities.Message
		getMessageErr        error

		expect    *models.Message
		expectErr error
	}{
		{
			name: "GetMessage",
			in: &models.GetMessageRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageResponse: &entities.Message{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				Content:          "content-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
				CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Message{
				ID:               "00000000-0000-0000-0000-000000000001",
				Target:           "user",
				PublicIdentifier: "public-identifier-1",
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				Content:          "content-1",
				CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "GetMessageError",
			in: &models.GetMessageRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetMessage: true,
			getMessageErr:        FooErr,
			expectErr:            FooErr,
		},
		{
			name: "InvalidData",
			in: &models.GetMessageRequest{
				ID: "invalid-id",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getMessageRepository := daomocks.NewMockGetMessageRepository(t)

			if tt.shouldCallGetMessage {
				getMessageRepository.On("GetMessage", context.TODO(), mock.Anything).
					Return(tt.getMessageResponse, tt.getMessageErr)
			}

			service := services.NewGetMessageService(getMessageRepository)

			resp, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, resp)

			getMessageRepository.AssertExpectations(t)
		})
	}
}
