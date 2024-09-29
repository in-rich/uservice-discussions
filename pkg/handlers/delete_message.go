package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeleteMessageHandler struct {
	discussions_pb.DeleteMessageServer
	service services.DeleteMessageService
	logger  monitor.GRPCLogger
}

func (h *DeleteMessageHandler) deleteMessage(ctx context.Context, in *discussions_pb.DeleteMessageRequest) (*emptypb.Empty, error) {
	err := h.service.Exec(ctx, &models.DeleteMessageRequest{ID: in.GetMessageId()})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid message data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to create message: %v", err)
	}

	return new(emptypb.Empty), nil
}

func (h *DeleteMessageHandler) DeleteMessage(ctx context.Context, in *discussions_pb.DeleteMessageRequest) (*emptypb.Empty, error) {
	res, err := h.deleteMessage(ctx, in)
	h.logger.Report(ctx, "DeleteMessage", err)
	return res, err
}

func NewDeleteMessageHandler(service services.DeleteMessageService, logger monitor.GRPCLogger) *DeleteMessageHandler {
	return &DeleteMessageHandler{
		service: service,
		logger:  logger,
	}
}
