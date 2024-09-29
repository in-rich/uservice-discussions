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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListDiscussionMessagesHandler struct {
	discussions_pb.ListDiscussionMessagesServer
	service services.ListDiscussionMessagesService
	logger  monitor.GRPCLogger
}

func (h *ListDiscussionMessagesHandler) listDiscussionMessages(ctx context.Context, in *discussions_pb.ListDiscussionMessagesRequest) (*discussions_pb.ListDiscussionMessagesResponse, error) {
	messages, err := h.service.Exec(ctx, &models.ListDiscussionMessagesRequest{
		Target:           in.GetTarget(),
		PublicIdentifier: in.GetPublicIdentifier(),
		TeamID:           in.GetTeamId(),
		Limit:            int(in.GetLimit()),
		Offset:           int(in.GetOffset()),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid message data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to list discussion messages: %v", err)
	}

	response := &discussions_pb.ListDiscussionMessagesResponse{
		Messages: make([]*discussions_pb.Message, len(messages)),
	}
	for i, message := range messages {
		response.Messages[i] = &discussions_pb.Message{
			Target:           message.Target,
			PublicIdentifier: message.PublicIdentifier,
			Content:          message.Content,
			CreatedAt:        timestamppb.New(*message.CreatedAt),
			AuthorId:         message.AuthorID,
			TeamId:           message.TeamID,
			MessageId:        message.ID,
		}
	}

	return response, nil
}

func (h *ListDiscussionMessagesHandler) ListDiscussionMessages(ctx context.Context, in *discussions_pb.ListDiscussionMessagesRequest) (*discussions_pb.ListDiscussionMessagesResponse, error) {
	res, err := h.listDiscussionMessages(ctx, in)
	h.logger.Report(ctx, "ListDiscussionMessages", err)
	return res, err
}

func NewListDiscussionMessagesHandler(service services.ListDiscussionMessagesService, logger monitor.GRPCLogger) *ListDiscussionMessagesHandler {
	return &ListDiscussionMessagesHandler{
		service: service,
		logger:  logger,
	}
}
