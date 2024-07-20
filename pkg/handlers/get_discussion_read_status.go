package handlers

import (
	"context"
	"errors"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetDiscussionReadStatusHandler struct {
	discussions_pb.GetDiscussionReadStatusServer
	service services.GetDiscussionReadStatusService
}

func (h *GetDiscussionReadStatusHandler) GetDiscussionReadStatus(ctx context.Context, in *discussions_pb.GetDiscussionReadStatusRequest) (*discussions_pb.DiscussionReadStatus, error) {
	readStatus, err := h.service.Exec(ctx, &models.GetDiscussionReadStatusRequest{
		Target:           in.GetTarget(),
		PublicIdentifier: in.GetPublicIdentifier(),
		TeamID:           in.GetTeamId(),
		UserID:           in.GetUserId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid discussion read status data: %v", err)
		}
		if errors.Is(err, dao.ErrDiscussionReadStatusNotFound) {
			return nil, status.Error(codes.NotFound, "discussion read status not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to get discussion read status: %v", err)
	}

	return &discussions_pb.DiscussionReadStatus{
		Target:              readStatus.Target,
		PublicIdentifier:    readStatus.PublicIdentifier,
		TeamId:              readStatus.TeamID,
		UserId:              readStatus.UserID,
		LatestReadMessageId: readStatus.LatestReadMessageID,
	}, nil
}

func NewGetDiscussionReadStatusHandler(service services.GetDiscussionReadStatusService) *GetDiscussionReadStatusHandler {
	return &GetDiscussionReadStatusHandler{
		service: service,
	}
}
