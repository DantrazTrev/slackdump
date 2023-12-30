// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=../mocks/mock_app/mock_app.go Credentials,createOpener
//
// Package mock_app is a generated GoMock package.
package mock_app

import (
	context "context"
	io "io"
	reflect "reflect"

	auth "github.com/rusq/slackdump/v2/auth"
	browser "github.com/rusq/slackdump/v2/auth/browser"
	gomock "go.uber.org/mock/gomock"
)

// MockCredentials is a mock of Credentials interface.
type MockCredentials struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialsMockRecorder
}

// MockCredentialsMockRecorder is the mock recorder for MockCredentials.
type MockCredentialsMockRecorder struct {
	mock *MockCredentials
}

// NewMockCredentials creates a new mock instance.
func NewMockCredentials(ctrl *gomock.Controller) *MockCredentials {
	mock := &MockCredentials{ctrl: ctrl}
	mock.recorder = &MockCredentialsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentials) EXPECT() *MockCredentialsMockRecorder {
	return m.recorder
}

// AuthProvider mocks base method.
func (m *MockCredentials) AuthProvider(ctx context.Context, workspace string, browser browser.Browser, legacy bool) (auth.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthProvider", ctx, workspace, browser, legacy)
	ret0, _ := ret[0].(auth.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthProvider indicates an expected call of AuthProvider.
func (mr *MockCredentialsMockRecorder) AuthProvider(ctx, workspace, browser, legacy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthProvider", reflect.TypeOf((*MockCredentials)(nil).AuthProvider), ctx, workspace, browser, legacy)
}

// IsEmpty mocks base method.
func (m *MockCredentials) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty.
func (mr *MockCredentialsMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockCredentials)(nil).IsEmpty))
}

// MockcreateOpener is a mock of createOpener interface.
type MockcreateOpener struct {
	ctrl     *gomock.Controller
	recorder *MockcreateOpenerMockRecorder
}

// MockcreateOpenerMockRecorder is the mock recorder for MockcreateOpener.
type MockcreateOpenerMockRecorder struct {
	mock *MockcreateOpener
}

// NewMockcreateOpener creates a new mock instance.
func NewMockcreateOpener(ctrl *gomock.Controller) *MockcreateOpener {
	mock := &MockcreateOpener{ctrl: ctrl}
	mock.recorder = &MockcreateOpenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcreateOpener) EXPECT() *MockcreateOpenerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockcreateOpener) Create(arg0 string) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockcreateOpenerMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockcreateOpener)(nil).Create), arg0)
}

// Open mocks base method.
func (m *MockcreateOpener) Open(arg0 string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockcreateOpenerMockRecorder) Open(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockcreateOpener)(nil).Open), arg0)
}
