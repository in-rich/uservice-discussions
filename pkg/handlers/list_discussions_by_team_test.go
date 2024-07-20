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

func TestListDiscussionsByTeam(t *testing.T) {
	testData := []struct {
		name string

		in *discussions_pb.ListDiscussionsByTeamRequest

		serviceResp []*models.Discussion
		serviceErr  error

		expect     *discussions_pb.ListDiscussionsByTeamResponse
		expectCode codes.Code
	}{
		{
			name: "ListDiscussionsByTeam",
			in: &discussions_pb.ListDiscussionsByTeamRequest{
				TeamId: "team-id-1",
				Limit:  10,
				Offset: 0,
			},
			serviceResp: []*models.Discussion{
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-3",
					Target:           "user",
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-1",
					Target:           "company",
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					TeamID:           "team-id-1",
					PublicIdentifier: "public-identifier-2",
					Target:           "user",
					UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
			expect: &discussions_pb.ListDiscussionsByTeamResponse{
				Discussions: []*discussions_pb.Discussion{
					{
						TeamId:           "team-id-1",
						PublicIdentifier: "public-identifier-3",
						Target:           "user",
						UpdatedAt:        timestamppb.New(time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)),
					},
					{
						TeamId:           "team-id-1",
						PublicIdentifier: "public-identifier-1",
						Target:           "company",
						UpdatedAt:        timestamppb.New(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
					},
					{
						TeamId:           "team-id-1",
						PublicIdentifier: "public-identifier-2",
						Target:           "user",
						UpdatedAt:        timestamppb.New(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
		{
			name: "InvalidArgument",
			in: &discussions_pb.ListDiscussionsByTeamRequest{
				TeamId: "team-id-1",
				Limit:  10,
				Offset: 0,
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &discussions_pb.ListDiscussionsByTeamRequest{
				TeamId: "team-id-1",
				Limit:  10,
				Offset: 0,
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockListDiscussionsByTeamService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewListDiscussionsByTeamHandler(service)

			resp, err := handler.ListDiscussionsByTeam(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
