// Code generated by mockery v1.0.0
package pet

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockAPI is an autogenerated mock type for the API type
type MockAPI struct {
	mock.Mock
}

// PetCreate provides a mock function with given fields: ctx, params
func (_m *MockAPI) PetCreate(ctx context.Context, params *PetCreateParams) (*PetCreateCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *PetCreateCreated
	if rf, ok := ret.Get(0).(func(context.Context, *PetCreateParams) *PetCreateCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PetCreateCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PetCreateParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PetDelete provides a mock function with given fields: ctx, params
func (_m *MockAPI) PetDelete(ctx context.Context, params *PetDeleteParams) (*PetDeleteNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *PetDeleteNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *PetDeleteParams) *PetDeleteNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PetDeleteNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PetDeleteParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PetGet provides a mock function with given fields: ctx, params
func (_m *MockAPI) PetGet(ctx context.Context, params *PetGetParams) (*PetGetOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *PetGetOK
	if rf, ok := ret.Get(0).(func(context.Context, *PetGetParams) *PetGetOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PetGetOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PetGetParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PetList provides a mock function with given fields: ctx, params
func (_m *MockAPI) PetList(ctx context.Context, params *PetListParams) (*PetListOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *PetListOK
	if rf, ok := ret.Get(0).(func(context.Context, *PetListParams) *PetListOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PetListOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PetListParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PetUpdate provides a mock function with given fields: ctx, params
func (_m *MockAPI) PetUpdate(ctx context.Context, params *PetUpdateParams) (*PetUpdateCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *PetUpdateCreated
	if rf, ok := ret.Get(0).(func(context.Context, *PetUpdateParams) *PetUpdateCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PetUpdateCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PetUpdateParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
