package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/models"
	servicesmocks "github.com/in-rich/uservice-discussions/pkg/services/mocks"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGetAllMessages(t *testing.T) {
	testData := []struct {
		name string

		in *discussions_pb.GetAllMessagesRequest

		getAllResponse []*models.Message
		getAllErr      error

		expect     *discussions_pb.GetAllMessagesResponse
		expectCode codes.Code
	}{
		{
			name: "GetAllMessages",
			in: &discussions_pb.GetAllMessagesRequest{
				Limit:  50,
				Offset: 10,
			},
			getAllResponse: []*models.Message{
				{
					ID:               "id-1",
					PublicIdentifier: "public-identifier-1",
					AuthorID:         "author-id-1",
					Target:           "company",
					Content:          "content-1",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:               "id-2",
					PublicIdentifier: "public-identifier-2",
					AuthorID:         "author-id-1",
					Target:           "user",
					Content:          "content-2",
					CreatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
			expect: &discussions_pb.GetAllMessagesResponse{
				Messages: []*discussions_pb.Message{
					{
						MessageId:        "id-1",
						PublicIdentifier: "public-identifier-1",
						AuthorId:         "author-id-1",
						Target:           "company",
						Content:          "content-1",
						CreatedAt:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
					},
					{
						MessageId:        "id-2",
						PublicIdentifier: "public-identifier-2",
						AuthorId:         "author-id-1",
						Target:           "user",
						Content:          "content-2",
						CreatedAt:        timestamppb.New(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
		{
			name: "GetAllError",
			in: &discussions_pb.GetAllMessagesRequest{
				Limit:  50,
				Offset: 10,
			},
			getAllErr:  errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetAllMessagesService(t)

			service.
				On("Exec", context.TODO(), &models.GetAllMessages{
					Offset: tt.in.Offset,
					Limit:  tt.in.Limit,
				}).
				Return(tt.getAllResponse, tt.getAllErr)

			handler := handlers.NewGetAllMessagesHandler(service, monitor.NewDummyGRPCLogger())
			resp, err := handler.GetAllMessages(context.TODO(), tt.in)

			require.Equal(t, tt.expect, resp)
			RequireGRPCCodesEqual(t, err, tt.expectCode)

			service.AssertExpectations(t)
		})
	}
}
