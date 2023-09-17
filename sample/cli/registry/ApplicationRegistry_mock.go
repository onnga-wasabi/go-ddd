// Code generated by mockery. DO NOT EDIT.

package registry

import (
	application "github.com/onnga-wasabi/go-ddd/sample/application"
	mock "github.com/stretchr/testify/mock"
)

// MockApplicationRegistry is an autogenerated mock type for the ApplicationRegistry type
type MockApplicationRegistry struct {
	mock.Mock
}

type MockApplicationRegistry_Expecter struct {
	mock *mock.Mock
}

func (_m *MockApplicationRegistry) EXPECT() *MockApplicationRegistry_Expecter {
	return &MockApplicationRegistry_Expecter{mock: &_m.Mock}
}

// NewUserApplicationService provides a mock function with given fields:
func (_m *MockApplicationRegistry) NewUserApplicationService() application.UserApplicationService {
	ret := _m.Called()

	var r0 application.UserApplicationService
	if rf, ok := ret.Get(0).(func() application.UserApplicationService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.UserApplicationService)
		}
	}

	return r0
}

// MockApplicationRegistry_NewUserApplicationService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewUserApplicationService'
type MockApplicationRegistry_NewUserApplicationService_Call struct {
	*mock.Call
}

// NewUserApplicationService is a helper method to define mock.On call
func (_e *MockApplicationRegistry_Expecter) NewUserApplicationService() *MockApplicationRegistry_NewUserApplicationService_Call {
	return &MockApplicationRegistry_NewUserApplicationService_Call{Call: _e.mock.On("NewUserApplicationService")}
}

func (_c *MockApplicationRegistry_NewUserApplicationService_Call) Run(run func()) *MockApplicationRegistry_NewUserApplicationService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockApplicationRegistry_NewUserApplicationService_Call) Return(_a0 application.UserApplicationService) *MockApplicationRegistry_NewUserApplicationService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockApplicationRegistry_NewUserApplicationService_Call) RunAndReturn(run func() application.UserApplicationService) *MockApplicationRegistry_NewUserApplicationService_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockApplicationRegistry creates a new instance of MockApplicationRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockApplicationRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockApplicationRegistry {
	mock := &MockApplicationRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
