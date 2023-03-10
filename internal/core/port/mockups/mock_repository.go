// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aria3ppp/url-shortener-grpc/internal/core/port (interfaces: Repository)

// Package mockups is a generated GoMock package.
package mockups

import (
	reflect "reflect"

	domain "github.com/aria3ppp/url-shortener-grpc/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateLink mocks base method.
func (m *MockRepository) CreateLink(arg0 *domain.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLink", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLink indicates an expected call of CreateLink.
func (mr *MockRepositoryMockRecorder) CreateLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLink", reflect.TypeOf((*MockRepository)(nil).CreateLink), arg0)
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(arg0 *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), arg0)
}

// GetLink mocks base method.
func (m *MockRepository) GetLink(arg0 string) (*domain.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink", arg0)
	ret0, _ := ret[0].(*domain.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLink indicates an expected call of GetLink.
func (mr *MockRepositoryMockRecorder) GetLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*MockRepository)(nil).GetLink), arg0)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(arg0 string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), arg0)
}
