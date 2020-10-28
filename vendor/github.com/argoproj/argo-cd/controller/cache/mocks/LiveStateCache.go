// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import kube "github.com/argoproj/argo-cd/util/kube"
import mock "github.com/stretchr/testify/mock"
import unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
import v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

// LiveStateCache is an autogenerated mock type for the LiveStateCache type
type LiveStateCache struct {
	mock.Mock
}

// GetManagedLiveObjs provides a mock function with given fields: a, targetObjs
func (_m *LiveStateCache) GetManagedLiveObjs(a *v1alpha1.Application, targetObjs []*unstructured.Unstructured) (map[kube.ResourceKey]*unstructured.Unstructured, error) {
	ret := _m.Called(a, targetObjs)

	var r0 map[kube.ResourceKey]*unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(*v1alpha1.Application, []*unstructured.Unstructured) map[kube.ResourceKey]*unstructured.Unstructured); ok {
		r0 = rf(a, targetObjs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[kube.ResourceKey]*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.Application, []*unstructured.Unstructured) error); ok {
		r1 = rf(a, targetObjs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Invalidate provides a mock function with given fields:
func (_m *LiveStateCache) Invalidate() {
	_m.Called()
}

// IsNamespaced provides a mock function with given fields: server, obj
func (_m *LiveStateCache) IsNamespaced(server string, obj *unstructured.Unstructured) (bool, error) {
	ret := _m.Called(server, obj)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *unstructured.Unstructured) bool); ok {
		r0 = rf(server, obj)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *unstructured.Unstructured) error); ok {
		r1 = rf(server, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IterateHierarchy provides a mock function with given fields: server, obj, action
func (_m *LiveStateCache) IterateHierarchy(server string, obj *unstructured.Unstructured, action func(v1alpha1.ResourceNode)) error {
	ret := _m.Called(server, obj, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *unstructured.Unstructured, func(v1alpha1.ResourceNode)) error); ok {
		r0 = rf(server, obj, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: ctx
func (_m *LiveStateCache) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
