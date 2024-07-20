package services_test

import (
	"context"
	daomocks "github.com/in-rich/uservice-discussions/pkg/dao/mocks"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteMessage(t *testing.T) {
	testData := []struct {
		name string

		in *models.DeleteMessageRequest

		shouldCallDeleteMessage bool
		deleteMessageErr        error

		expectErr error
	}{
		{
			name: "DeleteMessage",
			in: &models.DeleteMessageRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallDeleteMessage: true,
		},
		{
			name: "DeleteMessageError",
			in: &models.DeleteMessageRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallDeleteMessage: true,
			deleteMessageErr:        FooErr,
			expectErr:               FooErr,
		},
		{
			name: "InvalidData",
			in: &models.DeleteMessageRequest{
				ID: "invalid-id",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			deleteMessageRepository := daomocks.NewMockDeleteMessageRepository(t)

			if tt.shouldCallDeleteMessage {
				deleteMessageRepository.On("DeleteMessage", context.TODO(), mock.Anything).Return(tt.deleteMessageErr)
			}

			service := services.NewDeleteMessageService(deleteMessageRepository)

			err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)

			deleteMessageRepository.AssertExpectations(t)
		})
	}
}
