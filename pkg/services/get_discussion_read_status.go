package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type GetDiscussionReadStatusService interface {
	Exec(ctx context.Context, selector *models.GetDiscussionReadStatusRequest) (*models.DiscussionReadStatus, error)
}

type getDiscussionReadStatusServiceImpl struct {
	getDiscussionReadStatusRepository dao.GetDiscussionReadStatusRepository
}

func (s *getDiscussionReadStatusServiceImpl) Exec(ctx context.Context, selector *models.GetDiscussionReadStatusRequest) (*models.DiscussionReadStatus, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	readStatus, err := s.getDiscussionReadStatusRepository.GetDiscussionReadStatus(
		ctx, selector.TeamID, selector.UserID, entities.Target(selector.Target), selector.PublicIdentifier,
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

func NewGetDiscussionReadStatusService(getDiscussionReadStatusRepository dao.GetDiscussionReadStatusRepository) GetDiscussionReadStatusService {
	return &getDiscussionReadStatusServiceImpl{
		getDiscussionReadStatusRepository: getDiscussionReadStatusRepository,
	}
}
