package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type DeleteMessageService interface {
	Exec(ctx context.Context, selector *models.DeleteMessageRequest) error
}

type deleteMessageServiceImpl struct {
	deleteMessageRepository dao.DeleteMessageRepository
}

func (s *deleteMessageServiceImpl) Exec(ctx context.Context, selector *models.DeleteMessageRequest) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return errors.Join(ErrInvalidData, err)
	}

	messageID, err := uuid.Parse(selector.ID)
	if err != nil {
		return errors.Join(ErrInvalidData, err)
	}

	err = s.deleteMessageRepository.DeleteMessage(ctx, messageID)
	if err != nil {
		return err
	}

	return nil
}

func NewDeleteMessageService(deleteMessageRepository dao.DeleteMessageRepository) DeleteMessageService {
	return &deleteMessageServiceImpl{
		deleteMessageRepository: deleteMessageRepository,
	}
}
