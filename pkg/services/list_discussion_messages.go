package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type ListDiscussionMessagesService interface {
	Exec(ctx context.Context, selector *models.ListDiscussionMessagesRequest) ([]*models.Message, error)
}

type listDiscussionMessagesServiceImpl struct {
	listDiscussionMessagesRepository dao.ListDiscussionMessagesRepository
}

func (s *listDiscussionMessagesServiceImpl) Exec(ctx context.Context, selector *models.ListDiscussionMessagesRequest) ([]*models.Message, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	messages, err := s.listDiscussionMessagesRepository.ListDiscussionMessages(
		ctx, selector.TeamID, selector.PublicIdentifier, entities.Target(selector.Target), selector.Limit, selector.Offset,
	)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Message, len(messages))
	for i, message := range messages {
		result[i] = &models.Message{
			ID:               message.ID.String(),
			Target:           string(message.Target),
			PublicIdentifier: message.PublicIdentifier,
			AuthorID:         message.AuthorID,
			TeamID:           message.TeamID,
			Content:          message.Content,
			CreatedAt:        message.CreatedAt,
		}
	}

	return result, nil
}

func NewListDiscussionMessagesService(listDiscussionMessagesRepository dao.ListDiscussionMessagesRepository) ListDiscussionMessagesService {
	return &listDiscussionMessagesServiceImpl{
		listDiscussionMessagesRepository: listDiscussionMessagesRepository,
	}
}
