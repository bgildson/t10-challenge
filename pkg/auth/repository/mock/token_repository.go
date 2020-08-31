// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/auth/repository/token_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "github.com/bgildson/t10-challenge/pkg/auth/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTokenRepository is a mock of TokenRepository interface
type MockTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRepositoryMockRecorder
}

// MockTokenRepositoryMockRecorder is the mock recorder for MockTokenRepository
type MockTokenRepositoryMockRecorder struct {
	mock *MockTokenRepository
}

// NewMockTokenRepository creates a new mock instance
func NewMockTokenRepository(ctrl *gomock.Controller) *MockTokenRepository {
	mock := &MockTokenRepository{ctrl: ctrl}
	mock.recorder = &MockTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenRepository) EXPECT() *MockTokenRepositoryMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockTokenRepository) Generate(arg0 *entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate
func (mr *MockTokenRepositoryMockRecorder) Generate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTokenRepository)(nil).Generate), arg0)
}

// Validate mocks base method
func (m *MockTokenRepository) Validate(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate
func (mr *MockTokenRepositoryMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenRepository)(nil).Validate), arg0)
}