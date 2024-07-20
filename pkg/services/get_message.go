package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type GetMessageService interface {
	Exec(ctx context.Context, selector *models.GetMessageRequest) (*models.Message, error)
}

type getMessageServiceImpl struct {
	deleteMessageRepository dao.GetMessageRepository
}

func (s *getMessageServiceImpl) Exec(ctx context.Context, selector *models.GetMessageRequest) (*models.Message, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	messageID, err := uuid.Parse(selector.ID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	message, err := s.deleteMessageRepository.GetMessage(ctx, messageID)
	if err != nil {
		return nil, err
	}

	return &models.Message{
		ID:               message.ID.String(),
		Target:           string(message.Target),
		PublicIdentifier: message.PublicIdentifier,
		AuthorID:         message.AuthorID,
		TeamID:           message.TeamID,
		Content:          message.Content,
		CreatedAt:        message.CreatedAt,
	}, nil
}

func NewGetMessageService(deleteMessageRepository dao.GetMessageRepository) GetMessageService {
	return &getMessageServiceImpl{
		deleteMessageRepository: deleteMessageRepository,
	}
}
