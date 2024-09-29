package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	discussions_pb "github.com/in-rich/proto/proto-go/discussions"
	"github.com/in-rich/uservice-discussions/pkg/handlers"
	"github.com/in-rich/uservice-discussions/pkg/services"
	servicesmocks "github.com/in-rich/uservice-discussions/pkg/services/mocks"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestDeleteMessage(t *testing.T) {
	testData := []struct {
		name string

		in *discussions_pb.DeleteMessageRequest

		serviceErr error

		expectCode codes.Code
	}{
		{
			name: "DeleteMessage",
			in: &discussions_pb.DeleteMessageRequest{
				MessageId: "message-id-1",
			},
		},
		{
			name: "InvalidArgument",
			in: &discussions_pb.DeleteMessageRequest{
				MessageId: "message-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &discussions_pb.DeleteMessageRequest{
				MessageId: "message-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockDeleteMessageService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceErr)

			handler := handlers.NewDeleteMessageHandler(service, monitor.NewDummyGRPCLogger())

			_, err := handler.DeleteMessage(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)

			service.AssertExpectations(t)
		})
	}
}
