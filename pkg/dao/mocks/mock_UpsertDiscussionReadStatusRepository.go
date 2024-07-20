// Code generated by mockery v2.43.2. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/in-rich/uservice-discussions/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUpsertDiscussionReadStatusRepository is an autogenerated mock type for the UpsertDiscussionReadStatusRepository type
type MockUpsertDiscussionReadStatusRepository struct {
	mock.Mock
}

type MockUpsertDiscussionReadStatusRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpsertDiscussionReadStatusRepository) EXPECT() *MockUpsertDiscussionReadStatusRepository_Expecter {
	return &MockUpsertDiscussionReadStatusRepository_Expecter{mock: &_m.Mock}
}

// UpsertDiscussionReadStatus provides a mock function with given fields: ctx, teamID, userID, target, publicIdentifier, messageID
func (_m *MockUpsertDiscussionReadStatusRepository) UpsertDiscussionReadStatus(ctx context.Context, teamID string, userID string, target entities.Target, publicIdentifier string, messageID uuid.UUID) (*entities.ReadStatus, error) {
	ret := _m.Called(ctx, teamID, userID, target, publicIdentifier, messageID)

	if len(ret) == 0 {
		panic("no return value specified for UpsertDiscussionReadStatus")
	}

	var r0 *entities.ReadStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, entities.Target, string, uuid.UUID) (*entities.ReadStatus, error)); ok {
		return rf(ctx, teamID, userID, target, publicIdentifier, messageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, entities.Target, string, uuid.UUID) *entities.ReadStatus); ok {
		r0 = rf(ctx, teamID, userID, target, publicIdentifier, messageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ReadStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, entities.Target, string, uuid.UUID) error); ok {
		r1 = rf(ctx, teamID, userID, target, publicIdentifier, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertDiscussionReadStatus'
type MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call struct {
	*mock.Call
}

// UpsertDiscussionReadStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - teamID string
//   - userID string
//   - target entities.Target
//   - publicIdentifier string
//   - messageID uuid.UUID
func (_e *MockUpsertDiscussionReadStatusRepository_Expecter) UpsertDiscussionReadStatus(ctx interface{}, teamID interface{}, userID interface{}, target interface{}, publicIdentifier interface{}, messageID interface{}) *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call {
	return &MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call{Call: _e.mock.On("UpsertDiscussionReadStatus", ctx, teamID, userID, target, publicIdentifier, messageID)}
}

func (_c *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call) Run(run func(ctx context.Context, teamID string, userID string, target entities.Target, publicIdentifier string, messageID uuid.UUID)) *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(entities.Target), args[4].(string), args[5].(uuid.UUID))
	})
	return _c
}

func (_c *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call) Return(_a0 *entities.ReadStatus, _a1 error) *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call) RunAndReturn(run func(context.Context, string, string, entities.Target, string, uuid.UUID) (*entities.ReadStatus, error)) *MockUpsertDiscussionReadStatusRepository_UpsertDiscussionReadStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpsertDiscussionReadStatusRepository creates a new instance of MockUpsertDiscussionReadStatusRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpsertDiscussionReadStatusRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpsertDiscussionReadStatusRepository {
	mock := &MockUpsertDiscussionReadStatusRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}