// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/gossamer/devnet/cmd/scale-down-ecs-service/internal (interfaces: ECSAPI)

// Package main is a generated GoMock package.
package main

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	ecs "github.com/aws/aws-sdk-go/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockECSAPI is a mock of ECSAPI interface.
type MockECSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockECSAPIMockRecorder
}

// MockECSAPIMockRecorder is the mock recorder for MockECSAPI.
type MockECSAPIMockRecorder struct {
	mock *MockECSAPI
}

// NewMockECSAPI creates a new mock instance.
func NewMockECSAPI(ctrl *gomock.Controller) *MockECSAPI {
	mock := &MockECSAPI{ctrl: ctrl}
	mock.recorder = &MockECSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECSAPI) EXPECT() *MockECSAPIMockRecorder {
	return m.recorder
}

// DescribeServicesWithContext mocks base method.
func (m *MockECSAPI) DescribeServicesWithContext(arg0 context.Context, arg1 *ecs.DescribeServicesInput, arg2 ...request.Option) (*ecs.DescribeServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServicesWithContext", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServicesWithContext indicates an expected call of DescribeServicesWithContext.
func (mr *MockECSAPIMockRecorder) DescribeServicesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServicesWithContext", reflect.TypeOf((*MockECSAPI)(nil).DescribeServicesWithContext), varargs...)
}

// ListServicesWithContext mocks base method.
func (m *MockECSAPI) ListServicesWithContext(arg0 context.Context, arg1 *ecs.ListServicesInput, arg2 ...request.Option) (*ecs.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServicesWithContext", varargs...)
	ret0, _ := ret[0].(*ecs.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServicesWithContext indicates an expected call of ListServicesWithContext.
func (mr *MockECSAPIMockRecorder) ListServicesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServicesWithContext", reflect.TypeOf((*MockECSAPI)(nil).ListServicesWithContext), varargs...)
}

// UpdateServiceWithContext mocks base method.
func (m *MockECSAPI) UpdateServiceWithContext(arg0 context.Context, arg1 *ecs.UpdateServiceInput, arg2 ...request.Option) (*ecs.UpdateServiceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateServiceWithContext", varargs...)
	ret0, _ := ret[0].(*ecs.UpdateServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServiceWithContext indicates an expected call of UpdateServiceWithContext.
func (mr *MockECSAPIMockRecorder) UpdateServiceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceWithContext", reflect.TypeOf((*MockECSAPI)(nil).UpdateServiceWithContext), varargs...)
}