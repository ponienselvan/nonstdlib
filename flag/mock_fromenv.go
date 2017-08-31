// Code generated by MockGen. DO NOT EDIT.
// Source: fromenv.go

// Package flag is a generated GoMock package.
package flag

import (
	flag "flag"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFromEnv is a mock of FromEnv interface
type MockFromEnv struct {
	ctrl     *gomock.Controller
	recorder *MockFromEnvMockRecorder
}

// MockFromEnvMockRecorder is the mock recorder for MockFromEnv
type MockFromEnvMockRecorder struct {
	mock *MockFromEnv
}

// NewMockFromEnv creates a new mock instance
func NewMockFromEnv(ctrl *gomock.Controller) *MockFromEnv {
	mock := &MockFromEnv{ctrl: ctrl}
	mock.recorder = &MockFromEnvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFromEnv) EXPECT() *MockFromEnvMockRecorder {
	return m.recorder
}

// Prefix mocks base method
func (m *MockFromEnv) Prefix() string {
	ret := m.ctrl.Call(m, "Prefix")
	ret0, _ := ret[0].(string)
	return ret0
}

// Prefix indicates an expected call of Prefix
func (mr *MockFromEnvMockRecorder) Prefix() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefix", reflect.TypeOf((*MockFromEnv)(nil).Prefix))
}

// Fill mocks base method
func (m *MockFromEnv) Fill() {
	m.ctrl.Call(m, "Fill")
}

// Fill indicates an expected call of Fill
func (mr *MockFromEnvMockRecorder) Fill() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fill", reflect.TypeOf((*MockFromEnv)(nil).Fill))
}

// Filled mocks base method
func (m *MockFromEnv) Filled() map[string]string {
	ret := m.ctrl.Call(m, "Filled")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Filled indicates an expected call of Filled
func (mr *MockFromEnvMockRecorder) Filled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filled", reflect.TypeOf((*MockFromEnv)(nil).Filled))
}

// AllFlags mocks base method
func (m *MockFromEnv) AllFlags() []*flag.Flag {
	ret := m.ctrl.Call(m, "AllFlags")
	ret0, _ := ret[0].([]*flag.Flag)
	return ret0
}

// AllFlags indicates an expected call of AllFlags
func (mr *MockFromEnvMockRecorder) AllFlags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllFlags", reflect.TypeOf((*MockFromEnv)(nil).AllFlags))
}
