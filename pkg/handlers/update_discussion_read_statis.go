package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateDiscussionReadStatusHandler struct {
	discussions_pb.UpdateDiscussionReadStatusServer
	service services.UpdateDiscussionReadStatusService
	logger  monitor.GRPCLogger
}

func (h *UpdateDiscussionReadStatusHandler) updateDiscussionReadStatus(ctx context.Context, in *discussions_pb.UpdateDiscussionReadStatusRequest) (*discussions_pb.DiscussionReadStatus, error) {
	readStatus, err := h.service.Exec(ctx, &models.UpdateDiscussionReadStatusRequest{
		Target:           in.GetTarget(),
		PublicIdentifier: in.GetPublicIdentifier(),
		TeamID:           in.GetTeamId(),
		UserID:           in.GetUserId(),
		MessageID:        in.GetMessageId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid selector: %v", err)
		}
		if errors.Is(err, dao.ErrMessageNotFound) {
			return nil, status.Error(codes.NotFound, "message not found")
		}
		if errors.Is(err, services.ErrMessageInWrongDiscussion) {
			return nil, status.Error(codes.NotFound, "message does not belong to the discussion")
		}

		return nil, status.Errorf(codes.Internal, "failed to update discussion read status: %v", err)
	}

	return &discussions_pb.DiscussionReadStatus{
		Target:              readStatus.Target,
		PublicIdentifier:    readStatus.PublicIdentifier,
		TeamId:              readStatus.TeamID,
		UserId:              readStatus.UserID,
		LatestReadMessageId: readStatus.LatestReadMessageID,
	}, nil
}

func (h *UpdateDiscussionReadStatusHandler) UpdateDiscussionReadStatus(ctx context.Context, in *discussions_pb.UpdateDiscussionReadStatusRequest) (*discussions_pb.DiscussionReadStatus, error) {
	res, err := h.updateDiscussionReadStatus(ctx, in)
	h.logger.Report(ctx, "UpdateDiscussionReadStatus", err)
	return res, err
}

func NewUpdateDiscussionReadStatusHandler(service services.UpdateDiscussionReadStatusService, logger monitor.GRPCLogger) *UpdateDiscussionReadStatusHandler {
	return &UpdateDiscussionReadStatusHandler{
		service: service,
		logger:  logger,
	}
}
