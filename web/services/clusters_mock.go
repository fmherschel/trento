// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package services

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/trento-project/trento/web/models"
)

// MockClustersService is an autogenerated mock type for the ClustersService type
type MockClustersService struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: _a0, _a1
func (_m *MockClustersService) GetAll(_a0 *ClustersFilter, _a1 *Page) (models.ClusterList, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.ClusterList
	if rf, ok := ret.Get(0).(func(*ClustersFilter, *Page) models.ClusterList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.ClusterList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ClustersFilter, *Page) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllClusterTypes provides a mock function with given fields:
func (_m *MockClustersService) GetAllClusterTypes() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSIDs provides a mock function with given fields:
func (_m *MockClustersService) GetAllSIDs() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTags provides a mock function with given fields:
func (_m *MockClustersService) GetAllTags() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields:
func (_m *MockClustersService) GetCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
