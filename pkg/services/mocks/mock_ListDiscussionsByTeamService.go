// Code generated by mockery v2.49.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/in-rich/uservice-discussions/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockListDiscussionsByTeamService is an autogenerated mock type for the ListDiscussionsByTeamService type
type MockListDiscussionsByTeamService struct {
	mock.Mock
}

type MockListDiscussionsByTeamService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockListDiscussionsByTeamService) EXPECT() *MockListDiscussionsByTeamService_Expecter {
	return &MockListDiscussionsByTeamService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, selector
func (_m *MockListDiscussionsByTeamService) Exec(ctx context.Context, selector *models.ListDiscussionsByTeamRequest) ([]*models.Discussion, error) {
	ret := _m.Called(ctx, selector)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 []*models.Discussion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.ListDiscussionsByTeamRequest) ([]*models.Discussion, error)); ok {
		return rf(ctx, selector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.ListDiscussionsByTeamRequest) []*models.Discussion); ok {
		r0 = rf(ctx, selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Discussion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.ListDiscussionsByTeamRequest) error); ok {
		r1 = rf(ctx, selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListDiscussionsByTeamService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockListDiscussionsByTeamService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - selector *models.ListDiscussionsByTeamRequest
func (_e *MockListDiscussionsByTeamService_Expecter) Exec(ctx interface{}, selector interface{}) *MockListDiscussionsByTeamService_Exec_Call {
	return &MockListDiscussionsByTeamService_Exec_Call{Call: _e.mock.On("Exec", ctx, selector)}
}

func (_c *MockListDiscussionsByTeamService_Exec_Call) Run(run func(ctx context.Context, selector *models.ListDiscussionsByTeamRequest)) *MockListDiscussionsByTeamService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.ListDiscussionsByTeamRequest))
	})
	return _c
}

func (_c *MockListDiscussionsByTeamService_Exec_Call) Return(_a0 []*models.Discussion, _a1 error) *MockListDiscussionsByTeamService_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListDiscussionsByTeamService_Exec_Call) RunAndReturn(run func(context.Context, *models.ListDiscussionsByTeamRequest) ([]*models.Discussion, error)) *MockListDiscussionsByTeamService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockListDiscussionsByTeamService creates a new instance of MockListDiscussionsByTeamService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockListDiscussionsByTeamService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockListDiscussionsByTeamService {
	mock := &MockListDiscussionsByTeamService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
