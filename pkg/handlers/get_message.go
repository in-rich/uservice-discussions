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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetMessageHandler struct {
	discussions_pb.GetMessageServer
	service services.GetMessageService
	logger  monitor.GRPCLogger
}

func (h *GetMessageHandler) getMessage(ctx context.Context, in *discussions_pb.GetMessageRequest) (*discussions_pb.Message, error) {
	message, err := h.service.Exec(ctx, &models.GetMessageRequest{ID: in.GetMessageId()})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid message data: %v", err)
		}
		if errors.Is(err, dao.ErrMessageNotFound) {
			return nil, status.Error(codes.NotFound, "message not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to get message: %v", err)
	}

	return &discussions_pb.Message{
		Target:           message.Target,
		PublicIdentifier: message.PublicIdentifier,
		Content:          message.Content,
		CreatedAt:        timestamppb.New(*message.CreatedAt),
		AuthorId:         message.AuthorID,
		TeamId:           message.TeamID,
		MessageId:        message.ID,
	}, nil
}

func (h *GetMessageHandler) GetMessage(ctx context.Context, in *discussions_pb.GetMessageRequest) (*discussions_pb.Message, error) {
	res, err := h.getMessage(ctx, in)
	h.logger.Report(ctx, "GetMessage", err)
	return res, err
}

func NewGetMessageHandler(service services.GetMessageService, logger monitor.GRPCLogger) *GetMessageHandler {
	return &GetMessageHandler{
		service: service,
		logger:  logger,
	}
}
