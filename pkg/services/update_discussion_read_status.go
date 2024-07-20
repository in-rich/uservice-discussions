package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type UpdateDiscussionReadStatusService interface {
	Exec(ctx context.Context, selector *models.UpdateDiscussionReadStatusRequest) (*models.DiscussionReadStatus, error)
}

type updateDiscussionReadStatusServiceImpl struct {
	updateDiscussionReadStatusRepository dao.UpsertDiscussionReadStatusRepository
	getMessageRepository                 dao.GetMessageRepository
}

func (s *updateDiscussionReadStatusServiceImpl) Exec(ctx context.Context, selector *models.UpdateDiscussionReadStatusRequest) (*models.DiscussionReadStatus, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	messageID, err := uuid.Parse(selector.MessageID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	// Check the message exists, and is part of the targeted conversation.
	message, err := s.getMessageRepository.GetMessage(ctx, messageID)
	if err != nil {
		return nil, err
	}

	if message.TeamID != selector.TeamID ||
		message.PublicIdentifier != selector.PublicIdentifier ||
		message.Target != entities.Target(selector.Target) {
		return nil, ErrMessageInWrongDiscussion
	}

	// Upsert the read status.
	readStatus, err := s.updateDiscussionReadStatusRepository.UpsertDiscussionReadStatus(
		ctx, selector.TeamID, selector.UserID, entities.Target(selector.Target), selector.PublicIdentifier, messageID,
	)
	if err != nil {
		return nil, err
	}

	return &models.DiscussionReadStatus{
		Target:              string(readStatus.Target),
		PublicIdentifier:    readStatus.PublicIdentifier,
		TeamID:              readStatus.TeamID,
		LatestReadMessageID: readStatus.LatestReadMessageID.String(),
		UserID:              readStatus.UserID,
	}, nil
}

func NewUpdateDiscussionReadStatusService(updateDiscussionReadStatusRepository dao.UpsertDiscussionReadStatusRepository, getMessageRepository dao.GetMessageRepository) UpdateDiscussionReadStatusService {
	return &updateDiscussionReadStatusServiceImpl{
		updateDiscussionReadStatusRepository: updateDiscussionReadStatusRepository,
		getMessageRepository:                 getMessageRepository,
	}
}
