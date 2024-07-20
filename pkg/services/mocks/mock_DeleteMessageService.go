// Code generated by mockery v2.43.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/in-rich/uservice-discussions/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockDeleteMessageService is an autogenerated mock type for the DeleteMessageService type
type MockDeleteMessageService struct {
	mock.Mock
}

type MockDeleteMessageService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeleteMessageService) EXPECT() *MockDeleteMessageService_Expecter {
	return &MockDeleteMessageService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, selector
func (_m *MockDeleteMessageService) Exec(ctx context.Context, selector *models.DeleteMessageRequest) error {
	ret := _m.Called(ctx, selector)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.DeleteMessageRequest) error); ok {
		r0 = rf(ctx, selector)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDeleteMessageService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockDeleteMessageService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - selector *models.DeleteMessageRequest
func (_e *MockDeleteMessageService_Expecter) Exec(ctx interface{}, selector interface{}) *MockDeleteMessageService_Exec_Call {
	return &MockDeleteMessageService_Exec_Call{Call: _e.mock.On("Exec", ctx, selector)}
}

func (_c *MockDeleteMessageService_Exec_Call) Run(run func(ctx context.Context, selector *models.DeleteMessageRequest)) *MockDeleteMessageService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.DeleteMessageRequest))
	})
	return _c
}

func (_c *MockDeleteMessageService_Exec_Call) Return(_a0 error) *MockDeleteMessageService_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeleteMessageService_Exec_Call) RunAndReturn(run func(context.Context, *models.DeleteMessageRequest) error) *MockDeleteMessageService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeleteMessageService creates a new instance of MockDeleteMessageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeleteMessageService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeleteMessageService {
	mock := &MockDeleteMessageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}