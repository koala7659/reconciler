// Code generated by mockery (devel). DO NOT EDIT.

package connectivityproxymocks

import (
	service "github.com/kyma-incubator/reconciler/pkg/reconciler/service"
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"
)

// Commands is an autogenerated mock type for the Commands type
type Commands struct {
	mock.Mock
}

// CopyResources provides a mock function with given fields: _a0
func (_m *Commands) CopyResources(_a0 *service.ActionContext) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*service.ActionContext) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Install provides a mock function with given fields: _a0
func (_m *Commands) Install(_a0 *service.ActionContext) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*service.ActionContext) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PopulateConfigs provides a mock function with given fields: _a0, _a1
func (_m *Commands) PopulateConfigs(_a0 *service.ActionContext, _a1 *v1.Secret) {
	_m.Called(_a0, _a1)
}

// Remove provides a mock function with given fields: _a0
func (_m *Commands) Remove(_a0 *service.ActionContext) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*service.ActionContext) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
