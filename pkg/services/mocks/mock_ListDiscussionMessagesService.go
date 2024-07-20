// Code generated by mockery v2.43.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/in-rich/uservice-discussions/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockListDiscussionMessagesService is an autogenerated mock type for the ListDiscussionMessagesService type
type MockListDiscussionMessagesService struct {
	mock.Mock
}

type MockListDiscussionMessagesService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockListDiscussionMessagesService) EXPECT() *MockListDiscussionMessagesService_Expecter {
	return &MockListDiscussionMessagesService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, selector
func (_m *MockListDiscussionMessagesService) Exec(ctx context.Context, selector *models.ListDiscussionMessagesRequest) ([]*models.Message, error) {
	ret := _m.Called(ctx, selector)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 []*models.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.ListDiscussionMessagesRequest) ([]*models.Message, error)); ok {
		return rf(ctx, selector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.ListDiscussionMessagesRequest) []*models.Message); ok {
		r0 = rf(ctx, selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.ListDiscussionMessagesRequest) error); ok {
		r1 = rf(ctx, selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListDiscussionMessagesService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockListDiscussionMessagesService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - selector *models.ListDiscussionMessagesRequest
func (_e *MockListDiscussionMessagesService_Expecter) Exec(ctx interface{}, selector interface{}) *MockListDiscussionMessagesService_Exec_Call {
	return &MockListDiscussionMessagesService_Exec_Call{Call: _e.mock.On("Exec", ctx, selector)}
}

func (_c *MockListDiscussionMessagesService_Exec_Call) Run(run func(ctx context.Context, selector *models.ListDiscussionMessagesRequest)) *MockListDiscussionMessagesService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.ListDiscussionMessagesRequest))
	})
	return _c
}

func (_c *MockListDiscussionMessagesService_Exec_Call) Return(_a0 []*models.Message, _a1 error) *MockListDiscussionMessagesService_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListDiscussionMessagesService_Exec_Call) RunAndReturn(run func(context.Context, *models.ListDiscussionMessagesRequest) ([]*models.Message, error)) *MockListDiscussionMessagesService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockListDiscussionMessagesService creates a new instance of MockListDiscussionMessagesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockListDiscussionMessagesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockListDiscussionMessagesService {
	mock := &MockListDiscussionMessagesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
