// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/interfaces/auth.interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthInterface is a mock of AuthInterface interface.
type MockAuthInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthInterfaceMockRecorder
}

// MockAuthInterfaceMockRecorder is the mock recorder for MockAuthInterface.
type MockAuthInterfaceMockRecorder struct {
	mock *MockAuthInterface
}

// NewMockAuthInterface creates a new mock instance.
func NewMockAuthInterface(ctrl *gomock.Controller) *MockAuthInterface {
	mock := &MockAuthInterface{ctrl: ctrl}
	mock.recorder = &MockAuthInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthInterface) EXPECT() *MockAuthInterfaceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuthInterface) GenerateToken(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GenerateToken", c)
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthInterfaceMockRecorder) GenerateToken(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthInterface)(nil).GenerateToken), c)
}
