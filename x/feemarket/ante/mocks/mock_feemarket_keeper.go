// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	feemarkettypes "cosmossdk.io/x/feemarket/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// FeeMarketKeeper is an autogenerated mock type for the FeeMarketKeeper type
type FeeMarketKeeper struct {
	mock.Mock
}

// GetMinGasPrice provides a mock function with given fields: ctx, denom
func (_m *FeeMarketKeeper) GetMinGasPrice(ctx types.Context, denom string) (types.DecCoin, error) {
	ret := _m.Called(ctx, denom)

	if len(ret) == 0 {
		panic("no return value specified for GetMinGasPrice")
	}

	var r0 types.DecCoin
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, string) (types.DecCoin, error)); ok {
		return rf(ctx, denom)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) types.DecCoin); ok {
		r0 = rf(ctx, denom)
	} else {
		r0 = ret.Get(0).(types.DecCoin)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) error); ok {
		r1 = rf(ctx, denom)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParams provides a mock function with given fields: ctx
func (_m *FeeMarketKeeper) GetParams(ctx types.Context) (feemarkettypes.Params, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetParams")
	}

	var r0 feemarkettypes.Params
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (feemarkettypes.Params, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) feemarkettypes.Params); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(feemarkettypes.Params)
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: ctx
func (_m *FeeMarketKeeper) GetState(ctx types.Context) (feemarkettypes.State, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 feemarkettypes.State
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (feemarkettypes.State, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) feemarkettypes.State); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(feemarkettypes.State)
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveToDenom provides a mock function with given fields: ctx, coin, denom
func (_m *FeeMarketKeeper) ResolveToDenom(ctx types.Context, coin types.DecCoin, denom string) (types.DecCoin, error) {
	ret := _m.Called(ctx, coin, denom)

	if len(ret) == 0 {
		panic("no return value specified for ResolveToDenom")
	}

	var r0 types.DecCoin
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.DecCoin, string) (types.DecCoin, error)); ok {
		return rf(ctx, coin, denom)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.DecCoin, string) types.DecCoin); ok {
		r0 = rf(ctx, coin, denom)
	} else {
		r0 = ret.Get(0).(types.DecCoin)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.DecCoin, string) error); ok {
		r1 = rf(ctx, coin, denom)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetParams provides a mock function with given fields: ctx, params
func (_m *FeeMarketKeeper) SetParams(ctx types.Context, params feemarkettypes.Params) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for SetParams")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, feemarkettypes.Params) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetState provides a mock function with given fields: ctx, state
func (_m *FeeMarketKeeper) SetState(ctx types.Context, state feemarkettypes.State) error {
	ret := _m.Called(ctx, state)

	if len(ret) == 0 {
		panic("no return value specified for SetState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, feemarkettypes.State) error); ok {
		r0 = rf(ctx, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFeeMarketKeeper creates a new instance of FeeMarketKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeeMarketKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *FeeMarketKeeper {
	mock := &FeeMarketKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
