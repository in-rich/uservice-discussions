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

func TestGetAllMessages(t *testing.T) {
	testData := []struct {
		name string

		selector *models.GetAllMessages

		shouldCallGetAllMessages bool
		getAllMessagesResponse   []*entities.Message
		getAllMessagesError      error

		expect    []*models.Message
		expectErr error
	}{
		{
			name: "GetAllMessages",
			selector: &models.GetAllMessages{
				Limit: 10,
			},
			shouldCallGetAllMessages: true,
			getAllMessagesResponse: []*entities.Message{
				{
					ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					PublicIdentifier: "public-identifier",
					AuthorID:         "author-id",
					Target:           entities.Target("target"),
					Content:          "content",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
			expect: []*models.Message{
				{
					ID:               "00000000-0000-0000-0000-000000000001",
					PublicIdentifier: "public-identifier",
					AuthorID:         "author-id",
					Target:           "target",
					Content:          "content",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "GetAllMessagesError",
			selector: &models.GetAllMessages{
				Limit: 10,
			},
			shouldCallGetAllMessages: true,
			getAllMessagesError:      FooErr,
			expectErr:                FooErr,
		},
		{
			name: "InvalidSelector",
			selector: &models.GetAllMessages{
				Limit: -2,
			},
			expectErr: services.ErrInvalidMessageSelector,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getAllMessages := daomocks.NewMockGetAllMessagesRepository(t)

			if tt.shouldCallGetAllMessages {
				getAllMessages.
					On("GetAllMessages", context.TODO(), tt.selector.Limit, tt.selector.Offset).
					Return(tt.getAllMessagesResponse, tt.getAllMessagesError)
			}

			service := services.NewGetAllMessagesService(getAllMessages)

			result, err := service.Exec(context.TODO(), tt.selector)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, result)

			getAllMessages.AssertExpectations(t)
		})
	}
}
