package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
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

func TestListDiscussionMessages(t *testing.T) {
	testData := []struct {
		name string

		in *discussions_pb.ListDiscussionMessagesRequest

		serviceResp []*models.Message
		serviceErr  error

		expect     *discussions_pb.ListDiscussionMessagesResponse
		expectCode codes.Code
	}{
		{
			name: "ListDiscussionMessages",
			in: &discussions_pb.ListDiscussionMessagesRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				Limit:            10,
				Offset:           0,
			},
			serviceResp: []*models.Message{
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
			expect: &discussions_pb.ListDiscussionMessagesResponse{
				Messages: []*discussions_pb.Message{
					{
						MessageId:        "00000000-0000-0000-0000-000000000002",
						AuthorId:         "author-id-2",
						TeamId:           "team-id-1",
						PublicIdentifier: "public-identifier-1",
						Target:           "user",
						Content:          "content-2",
						CreatedAt:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
					},
					{
						MessageId:        "00000000-0000-0000-0000-000000000003",
						AuthorId:         "author-id-3",
						TeamId:           "team-id-1",
						PublicIdentifier: "public-identifier-1",
						Target:           "user",
						Content:          "content-3",
						CreatedAt:        timestamppb.New(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
					{
						MessageId:        "00000000-0000-0000-0000-000000000001",
						AuthorId:         "author-id-1",
						TeamId:           "team-id-1",
						PublicIdentifier: "public-identifier-1",
						Target:           "user",
						Content:          "content-1",
						CreatedAt:        timestamppb.New(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
		{
			name: "InvalidArgument",
			in: &discussions_pb.ListDiscussionMessagesRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				Limit:            10,
				Offset:           0,
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &discussions_pb.ListDiscussionMessagesRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				TeamId:           "team-id-1",
				Limit:            10,
				Offset:           0,
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockListDiscussionMessagesService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewListDiscussionMessagesHandler(service, monitor.NewDummyGRPCLogger())

			resp, err := handler.ListDiscussionMessages(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
