// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/interfaces/user.interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	entity "github.com/voltgizerz/rest-api-go-firestore/internal/app/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// DeleteUserDataByDocRefID mocks base method.
func (m *MockUserRepositoryInterface) DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserDataByDocRefID", ctx, docRefID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserDataByDocRefID indicates an expected call of DeleteUserDataByDocRefID.
func (mr *MockUserRepositoryInterfaceMockRecorder) DeleteUserDataByDocRefID(ctx, docRefID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserDataByDocRefID", reflect.TypeOf((*MockUserRepositoryInterface)(nil).DeleteUserDataByDocRefID), ctx, docRefID)
}

// GetAllUserData mocks base method.
func (m *MockUserRepositoryInterface) GetAllUserData(ctx context.Context) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserData", ctx)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserData indicates an expected call of GetAllUserData.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetAllUserData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserData", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetAllUserData), ctx)
}

// GetUserDataByDocRefID mocks base method.
func (m *MockUserRepositoryInterface) GetUserDataByDocRefID(ctx context.Context, docRefID string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataByDocRefID", ctx, docRefID)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDataByDocRefID indicates an expected call of GetUserDataByDocRefID.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUserDataByDocRefID(ctx, docRefID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataByDocRefID", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUserDataByDocRefID), ctx, docRefID)
}

// InsertUserData mocks base method.
func (m *MockUserRepositoryInterface) InsertUserData(ctx context.Context, data entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserData", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserData indicates an expected call of InsertUserData.
func (mr *MockUserRepositoryInterfaceMockRecorder) InsertUserData(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserData", reflect.TypeOf((*MockUserRepositoryInterface)(nil).InsertUserData), ctx, data)
}

// UpdateUserData mocks base method.
func (m *MockUserRepositoryInterface) UpdateUserData(ctx context.Context, docRefID string, data entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserData", ctx, docRefID, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserData indicates an expected call of UpdateUserData.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdateUserData(ctx, docRefID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserData", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdateUserData), ctx, docRefID, data)
}

// MockUserUsecaseInterface is a mock of UserUsecaseInterface interface.
type MockUserUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseInterfaceMockRecorder
}

// MockUserUsecaseInterfaceMockRecorder is the mock recorder for MockUserUsecaseInterface.
type MockUserUsecaseInterfaceMockRecorder struct {
	mock *MockUserUsecaseInterface
}

// NewMockUserUsecaseInterface creates a new mock instance.
func NewMockUserUsecaseInterface(ctrl *gomock.Controller) *MockUserUsecaseInterface {
	mock := &MockUserUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecaseInterface) EXPECT() *MockUserUsecaseInterfaceMockRecorder {
	return m.recorder
}

// DeleteUserDataByDocRefID mocks base method.
func (m *MockUserUsecaseInterface) DeleteUserDataByDocRefID(ctx context.Context, docRefID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserDataByDocRefID", ctx, docRefID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserDataByDocRefID indicates an expected call of DeleteUserDataByDocRefID.
func (mr *MockUserUsecaseInterfaceMockRecorder) DeleteUserDataByDocRefID(ctx, docRefID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserDataByDocRefID", reflect.TypeOf((*MockUserUsecaseInterface)(nil).DeleteUserDataByDocRefID), ctx, docRefID)
}

// GetAllUserData mocks base method.
func (m *MockUserUsecaseInterface) GetAllUserData(ctx context.Context) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserData", ctx)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserData indicates an expected call of GetAllUserData.
func (mr *MockUserUsecaseInterfaceMockRecorder) GetAllUserData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserData", reflect.TypeOf((*MockUserUsecaseInterface)(nil).GetAllUserData), ctx)
}

// GetUserDataByDocRefID mocks base method.
func (m *MockUserUsecaseInterface) GetUserDataByDocRefID(ctx context.Context, docRefID string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataByDocRefID", ctx, docRefID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDataByDocRefID indicates an expected call of GetUserDataByDocRefID.
func (mr *MockUserUsecaseInterfaceMockRecorder) GetUserDataByDocRefID(ctx, docRefID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataByDocRefID", reflect.TypeOf((*MockUserUsecaseInterface)(nil).GetUserDataByDocRefID), ctx, docRefID)
}

// InsertUserData mocks base method.
func (m *MockUserUsecaseInterface) InsertUserData(ctx context.Context, data entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserData", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserData indicates an expected call of InsertUserData.
func (mr *MockUserUsecaseInterfaceMockRecorder) InsertUserData(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserData", reflect.TypeOf((*MockUserUsecaseInterface)(nil).InsertUserData), ctx, data)
}

// UpdateUserDataByDocRefID mocks base method.
func (m *MockUserUsecaseInterface) UpdateUserDataByDocRefID(ctx context.Context, docRefID string, data entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDataByDocRefID", ctx, docRefID, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDataByDocRefID indicates an expected call of UpdateUserDataByDocRefID.
func (mr *MockUserUsecaseInterfaceMockRecorder) UpdateUserDataByDocRefID(ctx, docRefID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDataByDocRefID", reflect.TypeOf((*MockUserUsecaseInterface)(nil).UpdateUserDataByDocRefID), ctx, docRefID, data)
}

// MockAPIHandlerInterface is a mock of APIHandlerInterface interface.
type MockAPIHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAPIHandlerInterfaceMockRecorder
}

// MockAPIHandlerInterfaceMockRecorder is the mock recorder for MockAPIHandlerInterface.
type MockAPIHandlerInterfaceMockRecorder struct {
	mock *MockAPIHandlerInterface
}

// NewMockAPIHandlerInterface creates a new mock instance.
func NewMockAPIHandlerInterface(ctrl *gomock.Controller) *MockAPIHandlerInterface {
	mock := &MockAPIHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockAPIHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIHandlerInterface) EXPECT() *MockAPIHandlerInterfaceMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method.
func (m *MockAPIHandlerInterface) DeleteUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUser", c)
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockAPIHandlerInterfaceMockRecorder) DeleteUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockAPIHandlerInterface)(nil).DeleteUser), c)
}

// GetAllUsers mocks base method.
func (m *MockAPIHandlerInterface) GetAllUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllUsers", c)
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockAPIHandlerInterfaceMockRecorder) GetAllUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockAPIHandlerInterface)(nil).GetAllUsers), c)
}

// GetUserByID mocks base method.
func (m *MockAPIHandlerInterface) GetUserByID(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserByID", c)
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAPIHandlerInterfaceMockRecorder) GetUserByID(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAPIHandlerInterface)(nil).GetUserByID), c)
}

// InsertUser mocks base method.
func (m *MockAPIHandlerInterface) InsertUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertUser", c)
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockAPIHandlerInterfaceMockRecorder) InsertUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockAPIHandlerInterface)(nil).InsertUser), c)
}

// UpdateUser mocks base method.
func (m *MockAPIHandlerInterface) UpdateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUser", c)
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAPIHandlerInterfaceMockRecorder) UpdateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAPIHandlerInterface)(nil).UpdateUser), c)
}
