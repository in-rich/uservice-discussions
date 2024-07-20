package handlers

import (
	"context"
	"errors"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListDiscussionsByTeamHandler struct {
	discussions_pb.ListDiscussionsByTeamServer
	service services.ListDiscussionsByTeamService
}

func (h *ListDiscussionsByTeamHandler) ListDiscussionsByTeam(ctx context.Context, in *discussions_pb.ListDiscussionsByTeamRequest) (*discussions_pb.ListDiscussionsByTeamResponse, error) {
	discussions, err := h.service.Exec(ctx, &models.ListDiscussionsByTeamRequest{
		TeamID: in.GetTeamId(),
		Limit:  int(in.GetLimit()),
		Offset: int(in.GetOffset()),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid selector: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to get team discussions: %v", err)
	}

	discussionsResp := make([]*discussions_pb.Discussion, len(discussions))
	for i, discussion := range discussions {
		discussionsResp[i] = &discussions_pb.Discussion{
			Target:           discussion.Target,
			PublicIdentifier: discussion.PublicIdentifier,
			TeamId:           discussion.TeamID,
			UpdatedAt:        timestamppb.New(*discussion.UpdatedAt),
		}
	}

	return &discussions_pb.ListDiscussionsByTeamResponse{
		Discussions: discussionsResp,
	}, nil
}

func NewListDiscussionsByTeamHandler(service services.ListDiscussionsByTeamService) *ListDiscussionsByTeamHandler {
	return &ListDiscussionsByTeamHandler{
		service: service,
	}
}
