package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/models"
)

type ListDiscussionsByTeamService interface {
	Exec(ctx context.Context, selector *models.ListDiscussionsByTeamRequest) ([]*models.Discussion, error)
}

type listDiscussionsByTeamServiceImpl struct {
	listDiscussionsByTeamRepository dao.ListDiscussionsByTeamRepository
}

func (s *listDiscussionsByTeamServiceImpl) Exec(ctx context.Context, selector *models.ListDiscussionsByTeamRequest) ([]*models.Discussion, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	discussions, err := s.listDiscussionsByTeamRepository.ListDiscussionsByTeam(ctx, selector.TeamID, selector.Limit, selector.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Discussion, len(discussions))
	for i, discussion := range discussions {
		result[i] = &models.Discussion{
			TeamID:           discussion.TeamID,
			PublicIdentifier: discussion.PublicIdentifier,
			Target:           string(discussion.Target),
			UpdatedAt:        discussion.UpdatedAt,
		}
	}

	return result, nil
}

func NewListDiscussionsByTeamService(listDiscussionsByTeamRepository dao.ListDiscussionsByTeamRepository) ListDiscussionsByTeamService {
	return &listDiscussionsByTeamServiceImpl{
		listDiscussionsByTeamRepository: listDiscussionsByTeamRepository,
	}
}
