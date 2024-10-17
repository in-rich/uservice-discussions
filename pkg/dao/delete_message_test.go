package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-discussions/pkg/dao"
	"github.com/in-rich/uservice-discussions/pkg/entities"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var deleteMessageFixtures = []*entities.Message{
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		AuthorID:         "author-id-1",
		TeamID:           "team-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		CreatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestDeleteMessage(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		messageID uuid.UUID

		expectErr error
	}{
		{
			name:      "DeleteMessage",
			messageID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		},
		{
			name:      "DeleteMessage/NotFound",
			messageID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		},
	}

	stx := BeginTX(db, deleteMessageFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewDeleteMessageRepository(tx)

			err := repo.DeleteMessage(context.Background(), tt.messageID)

			require.ErrorIs(t, err, tt.expectErr)
		})
	}
}
