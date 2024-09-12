// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// Withdrawal is an autogenerated mock type for the Withdrawal type
type Withdrawal[T interface{}] struct {
	mock.Mock
}

type Withdrawal_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *Withdrawal[T]) EXPECT() *Withdrawal_Expecter[T] {
	return &Withdrawal_Expecter[T]{mock: &_m.Mock}
}

// New provides a mock function with given fields: index, validator, address, amount
func (_m *Withdrawal[T]) New(index math.U64, validator math.U64, address common.ExecutionAddress, amount math.U64) T {
	ret := _m.Called(index, validator, address, amount)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 T
	if rf, ok := ret.Get(0).(func(math.U64, math.U64, common.ExecutionAddress, math.U64) T); ok {
		r0 = rf(index, validator, address, amount)
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// Withdrawal_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type Withdrawal_New_Call[T interface{}] struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - index math.U64
//   - validator math.U64
//   - address common.ExecutionAddress
//   - amount math.U64
func (_e *Withdrawal_Expecter[T]) New(index interface{}, validator interface{}, address interface{}, amount interface{}) *Withdrawal_New_Call[T] {
	return &Withdrawal_New_Call[T]{Call: _e.mock.On("New", index, validator, address, amount)}
}

func (_c *Withdrawal_New_Call[T]) Run(run func(index math.U64, validator math.U64, address common.ExecutionAddress, amount math.U64)) *Withdrawal_New_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(math.U64), args[2].(common.ExecutionAddress), args[3].(math.U64))
	})
	return _c
}

func (_c *Withdrawal_New_Call[T]) Return(_a0 T) *Withdrawal_New_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Withdrawal_New_Call[T]) RunAndReturn(run func(math.U64, math.U64, common.ExecutionAddress, math.U64) T) *Withdrawal_New_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewWithdrawal creates a new instance of Withdrawal. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWithdrawal[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Withdrawal[T] {
	mock := &Withdrawal[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
