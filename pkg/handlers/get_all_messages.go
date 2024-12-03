package handlers

import (
	"context"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetAllMessagesHandler struct {
	discussions_pb.GetAllMessagesServer
	service services.GetAllMessagesService
	logger  monitor.GRPCLogger
}

func (h *GetAllMessagesHandler) getAllMessages(ctx context.Context, in *discussions_pb.GetAllMessagesRequest) (*discussions_pb.GetAllMessagesResponse, error) {
	messages, err := h.service.Exec(ctx, &models.GetAllMessages{Limit: in.GetLimit(), Offset: in.GetOffset()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all messages: %v", err)
	}

	res := &discussions_pb.GetAllMessagesResponse{
		Messages: make([]*discussions_pb.Message, len(messages)),
	}
	for i, message := range messages {
		res.Messages[i] = &discussions_pb.Message{
			MessageId:        message.ID,
			PublicIdentifier: message.PublicIdentifier,
			AuthorId:         message.AuthorID,
			Target:           message.Target,
			Content:          message.Content,
			TeamId:           message.TeamID,
			CreatedAt:        timestamppb.New(*message.CreatedAt),
		}
	}

	return res, nil
}

func (h *GetAllMessagesHandler) GetAllMessages(ctx context.Context, in *discussions_pb.GetAllMessagesRequest) (*discussions_pb.GetAllMessagesResponse, error) {
	res, err := h.getAllMessages(ctx, in)
	h.logger.Report(ctx, "GetAllMessages", err)
	return res, err
}

func NewGetAllMessagesHandler(service services.GetAllMessagesService, logger monitor.GRPCLogger) *GetAllMessagesHandler {
	return &GetAllMessagesHandler{
		service: service,
		logger:  logger,
	}
}
