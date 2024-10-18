package services_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	daomocks "github.com/in-rich/uservice-discussions/pkg/dao/mocks"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	testData := []struct {
		name string

		in *models.CreateMessageRequest

		shouldCallCreateMessage bool
		createMessageResponse   *entities.Message
		createMessageErr        error

		expect    *models.Message
		expectErr error
	}{
		{
			name: "CreateMessage",
			in: &models.CreateMessageRequest{
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				Content:          "content-1",
				PublicIdentifier: "public-identifier-1",
				Target:           "user",
			},
			shouldCallCreateMessage: true,
			createMessageResponse: &entities.Message{
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
			name: "CreateMessageError",
			in: &models.CreateMessageRequest{
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				Content:          "content-1",
				PublicIdentifier: "public-identifier-1",
				Target:           "user",
			},
			shouldCallCreateMessage: true,
			createMessageErr:        FooErr,
			expectErr:               FooErr,
		},
		{
			name: "InvalidData",
			in: &models.CreateMessageRequest{
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				Content:          "content-1",
				PublicIdentifier: "public-identifier-1",
				Target:           "invalid-target",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			createMessageRepository := daomocks.NewMockCreateMessageRepository(t)

			if tt.shouldCallCreateMessage {
				createMessageRepository.On("CreateMessage", context.TODO(), tt.in.AuthorID, tt.in.TeamID, &dao.CreateMessageData{
					Content:          tt.in.Content,
					PublicIdentifier: tt.in.PublicIdentifier,
					Target:           entities.Target(tt.in.Target),
				}).Return(tt.createMessageResponse, tt.createMessageErr)
			}

			service := services.NewCreateMessageService(createMessageRepository)

			resp, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, resp)

			createMessageRepository.AssertExpectations(t)
		})
	}
}
