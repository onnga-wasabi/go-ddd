// Code generated by mockery. DO NOT EDIT.

package data_model

import (
	context "context"
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// MockDB is an autogenerated mock type for the DB type
type MockDB struct {
	mock.Mock
}

type MockDB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDB) EXPECT() *MockDB_Expecter {
	return &MockDB_Expecter{mock: &_m.Mock}
}

// ExecContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDB) ExecContext(_a0 context.Context, _a1 string, _a2 ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type MockDB_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...interface{}
func (_e *MockDB_Expecter) ExecContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *MockDB_ExecContext_Call {
	return &MockDB_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *MockDB_ExecContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...interface{})) *MockDB_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDB_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *MockDB_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *MockDB_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDB) QueryContext(_a0 context.Context, _a1 string, _a2 ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type MockDB_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...interface{}
func (_e *MockDB_Expecter) QueryContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *MockDB_QueryContext_Call {
	return &MockDB_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *MockDB_QueryContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...interface{})) *MockDB_QueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDB_QueryContext_Call) Return(_a0 *sql.Rows, _a1 error) *MockDB_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_QueryContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Rows, error)) *MockDB_QueryContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRowContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDB) QueryRowContext(_a0 context.Context, _a1 string, _a2 ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// MockDB_QueryRowContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowContext'
type MockDB_QueryRowContext_Call struct {
	*mock.Call
}

// QueryRowContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...interface{}
func (_e *MockDB_Expecter) QueryRowContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *MockDB_QueryRowContext_Call {
	return &MockDB_QueryRowContext_Call{Call: _e.mock.On("QueryRowContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *MockDB_QueryRowContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...interface{})) *MockDB_QueryRowContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDB_QueryRowContext_Call) Return(_a0 *sql.Row) *MockDB_QueryRowContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDB_QueryRowContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *sql.Row) *MockDB_QueryRowContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDB creates a new instance of MockDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDB {
	mock := &MockDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
