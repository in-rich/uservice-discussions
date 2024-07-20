package handlers_test

import (
	"context"
	"errors"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	servicesmocks "github.com/in-rich/uservice-discussions/pkg/services/mocks"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	testData := []struct {
		name string

		in *discussions_pb.CreateMessageRequest

		serviceResp *models.Message
		serviceErr  error

		expect     *discussions_pb.Message
		expectCode codes.Code
	}{
		{
			name: "CreateMessage",
			in: &discussions_pb.CreateMessageRequest{
				AuthorId:         "author-id-1",
				Content:          "content-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
			},
			serviceResp: &models.Message{
				ID:               "message-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				AuthorID:         "author-id-1",
				TeamID:           "team-id-1",
				Content:          "content-1",
				CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &discussions_pb.Message{
				MessageId:        "message-id-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				TeamId:           "team-id-1",
				Content:          "content-1",
				CreatedAt:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "InvalidArgument",
			in: &discussions_pb.CreateMessageRequest{
				AuthorId:         "author-id-1",
				Content:          "content-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &discussions_pb.CreateMessageRequest{
				AuthorId:         "author-id-1",
				Content:          "content-1",
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockCreateMessageService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewCreateMessageHandler(service)

			resp, err := handler.CreateMessage(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
