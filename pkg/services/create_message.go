package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type CreateMessageService interface {
	Exec(ctx context.Context, data *models.CreateMessageRequest) (*models.Message, error)
}

type createMessageServiceImpl struct {
	createMessageRepository dao.CreateMessageRepository
}

func (s *createMessageServiceImpl) Exec(ctx context.Context, data *models.CreateMessageRequest) (*models.Message, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	message, err := s.createMessageRepository.CreateMessage(ctx, data.AuthorID, data.TeamID, &dao.CreateMessageData{
		Content:          data.Content,
		PublicIdentifier: data.PublicIdentifier,
		Target:           entities.Target(data.Target),
		CreatedAt:        data.CreatedAt,
	})
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

func NewCreateMessageService(createMessageRepository dao.CreateMessageRepository) CreateMessageService {
	return &createMessageServiceImpl{
		createMessageRepository: createMessageRepository,
	}
}
