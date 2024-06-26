// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/EgorMizerov/expansion_bot/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// RegistrationApplicationService is an autogenerated mock type for the RegistrationApplicationService type
type RegistrationApplicationService struct {
	mock.Mock
}

// ConfirmRegistrationApplication provides a mock function with given fields: ctx, application
func (_m *RegistrationApplicationService) ConfirmRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error {
	ret := _m.Called(ctx, application)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmRegistrationApplication")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.RegistrationApplication) error); ok {
		r0 = rf(ctx, application)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRegistrationApplication provides a mock function with given fields: ctx, applicationID
func (_m *RegistrationApplicationService) GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error) {
	ret := _m.Called(ctx, applicationID)

	if len(ret) == 0 {
		panic("no return value specified for GetRegistrationApplication")
	}

	var r0 *entity.RegistrationApplication
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.RegistrationApplicationID) (*entity.RegistrationApplication, error)); ok {
		return rf(ctx, applicationID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.RegistrationApplicationID) *entity.RegistrationApplication); ok {
		r0 = rf(ctx, applicationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.RegistrationApplication)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.RegistrationApplicationID) error); ok {
		r1 = rf(ctx, applicationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistrationApplications provides a mock function with given fields: ctx
func (_m *RegistrationApplicationService) GetRegistrationApplications(ctx context.Context) ([]*entity.RegistrationApplication, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRegistrationApplications")
	}

	var r0 []*entity.RegistrationApplication
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.RegistrationApplication, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.RegistrationApplication); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.RegistrationApplication)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveRegistrationApplication provides a mock function with given fields: ctx, application
func (_m *RegistrationApplicationService) SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error {
	ret := _m.Called(ctx, application)

	if len(ret) == 0 {
		panic("no return value specified for SaveRegistrationApplication")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.RegistrationApplication) error); ok {
		r0 = rf(ctx, application)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRegistrationApplicationService creates a new instance of RegistrationApplicationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegistrationApplicationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *RegistrationApplicationService {
	mock := &RegistrationApplicationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
