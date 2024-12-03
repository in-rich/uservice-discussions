package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type GetAllMessagesService interface {
	Exec(ctx context.Context, selector *models.GetAllMessages) ([]*models.Message, error)
}

type getAllMessagesServiceImpl struct {
	getAllMessagesRepository dao.GetAllMessagesRepository
}

func (s *getAllMessagesServiceImpl) Exec(ctx context.Context, selector *models.GetAllMessages) ([]*models.Message, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidMessageSelector, err)
	}

	messages, err := s.getAllMessagesRepository.GetAllMessages(ctx, selector.Limit, selector.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Message, len(messages))
	for i, message := range messages {
		result[i] = &models.Message{
			ID:               message.ID.String(),
			PublicIdentifier: message.PublicIdentifier,
			AuthorID:         message.AuthorID,
			Target:           string(message.Target),
			Content:          message.Content,
			CreatedAt:        message.CreatedAt,
		}
	}

	return result, nil
}

func NewGetAllMessagesService(getAllMessagesRepository dao.GetAllMessagesRepository) GetAllMessagesService {
	return &getAllMessagesServiceImpl{
		getAllMessagesRepository: getAllMessagesRepository,
	}
}
