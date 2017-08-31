// Code generated by MockGen. DO NOT EDIT.
// Source: dir_reader.go

// Package os is a generated GoMock package.
package os

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDirReader is a mock of DirReader interface
type MockDirReader struct {
	ctrl     *gomock.Controller
	recorder *MockDirReaderMockRecorder
}

// MockDirReaderMockRecorder is the mock recorder for MockDirReader
type MockDirReaderMockRecorder struct {
	mock *MockDirReader
}

// NewMockDirReader creates a new mock instance
func NewMockDirReader(ctrl *gomock.Controller) *MockDirReader {
	mock := &MockDirReader{ctrl: ctrl}
	mock.recorder = &MockDirReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirReader) EXPECT() *MockDirReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockDirReader) Read() ([]string, error) {
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockDirReaderMockRecorder) Read() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDirReader)(nil).Read))
}

// Filter mocks base method
func (m *MockDirReader) Filter(arg0 DirEntryFilter) ([]string, error) {
	ret := m.ctrl.Call(m, "Filter", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockDirReaderMockRecorder) Filter(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockDirReader)(nil).Filter), arg0)
}
