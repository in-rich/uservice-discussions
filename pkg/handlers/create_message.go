package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/models"
	"github.com/in-rich/uservice-discussions/pkg/services"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type CreateMessageHandler struct {
	discussions_pb.CreateMessageServer
	service services.CreateMessageService
	logger  monitor.GRPCLogger
}

func (h *CreateMessageHandler) createMessage(ctx context.Context, in *discussions_pb.CreateMessageRequest) (*discussions_pb.Message, error) {
	message, err := h.service.Exec(ctx, &models.CreateMessageRequest{
		Target:           in.GetTarget(),
		PublicIdentifier: in.GetPublicIdentifier(),
		AuthorID:         in.GetAuthorId(),
		TeamID:           in.GetTeamId(),
		Content:          strings.TrimSpace(in.GetContent()),
		CreatedAt: lo.TernaryF[*time.Time](
			in.GetUpdatedAt() == nil,
			func() *time.Time {
				return nil
			},
			func() *time.Time {
				return lo.ToPtr(in.GetUpdatedAt().AsTime())
			},
		),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid message data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to create message: %v", err)
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

func (h *CreateMessageHandler) CreateMessage(ctx context.Context, in *discussions_pb.CreateMessageRequest) (*discussions_pb.Message, error) {
	res, err := h.createMessage(ctx, in)
	h.logger.Report(ctx, "CreateMessage", err)
	return res, err
}

func NewCreateMessageHandler(service services.CreateMessageService, logger monitor.GRPCLogger) *CreateMessageHandler {
	return &CreateMessageHandler{
		service: service,
		logger:  logger,
	}
}
