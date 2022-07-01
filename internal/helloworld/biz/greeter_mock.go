// Code generated by MockGen. DO NOT EDIT.
// Source: greeter.go
// Package biz is a generated GoMock package.
package biz

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGreeterRepo is a mock of GreeterRepo interface.
type MockGreeterRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterRepoMockRecorder
}

// MockGreeterRepoMockRecorder is the mock recorder for MockGreeterRepo.
type MockGreeterRepoMockRecorder struct {
	mock *MockGreeterRepo
}

// NewMockGreeterRepo creates a new mock instance.
func NewMockGreeterRepo(ctrl *gomock.Controller) *MockGreeterRepo {
	mock := &MockGreeterRepo{ctrl: ctrl}
	mock.recorder = &MockGreeterRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeterRepo) EXPECT() *MockGreeterRepoMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockGreeterRepo) FindByID(arg0 context.Context, arg1 int64) (*Greeter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockGreeterRepoMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockGreeterRepo)(nil).FindByID), arg0, arg1)
}

// ListAll mocks base method.
func (m *MockGreeterRepo) ListAll(arg0 context.Context) ([]*Greeter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].([]*Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockGreeterRepoMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockGreeterRepo)(nil).ListAll), arg0)
}

// ListByHello mocks base method.
func (m *MockGreeterRepo) ListByHello(arg0 context.Context, arg1 string) ([]*Greeter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByHello", arg0, arg1)
	ret0, _ := ret[0].([]*Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByHello indicates an expected call of ListByHello.
func (mr *MockGreeterRepoMockRecorder) ListByHello(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByHello", reflect.TypeOf((*MockGreeterRepo)(nil).ListByHello), arg0, arg1)
}

// Save mocks base method.
func (m *MockGreeterRepo) Save(arg0 context.Context, arg1 *Greeter) (*Greeter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(*Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockGreeterRepoMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGreeterRepo)(nil).Save), arg0, arg1)
}

// Update mocks base method.
func (m *MockGreeterRepo) Update(arg0 context.Context, arg1 *Greeter) (*Greeter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockGreeterRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGreeterRepo)(nil).Update), arg0, arg1)
}
