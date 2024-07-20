package handlers_test

import (
	"context"
	"errors"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	servicesmocks "github.com/in-rich/uservice-discussions/pkg/services/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestGetDiscussionReadStatus(t *testing.T) {
	testData := []struct {
		name string

		in *discussions_pb.GetDiscussionReadStatusRequest

		serviceResp *models.DiscussionReadStatus
		serviceErr  error

		expect     *discussions_pb.DiscussionReadStatus
		expectCode codes.Code
	}{
		{
			name: "GetDiscussionReadStatus",
			in: &discussions_pb.GetDiscussionReadStatusRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				UserId:           "user-id-1",
			},
			serviceResp: &models.DiscussionReadStatus{
				Target:              "company",
				PublicIdentifier:    "public-identifier-1",
				TeamID:              "team-id-1",
				LatestReadMessageID: "message-id-1",
				UserID:              "user-id-1",
			},
			expect: &discussions_pb.DiscussionReadStatus{
				Target:              "company",
				PublicIdentifier:    "public-identifier-1",
				TeamId:              "team-id-1",
				UserId:              "user-id-1",
				LatestReadMessageId: "message-id-1",
			},
		},
		{
			name: "NotFound",
			in: &discussions_pb.GetDiscussionReadStatusRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				UserId:           "user-id-1",
			},
			serviceErr: dao.ErrDiscussionReadStatusNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "InvalidArgument",
			in: &discussions_pb.GetDiscussionReadStatusRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				UserId:           "user-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &discussions_pb.GetDiscussionReadStatusRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				UserId:           "user-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetDiscussionReadStatusService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewGetDiscussionReadStatusHandler(service)

			resp, err := handler.GetDiscussionReadStatus(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
