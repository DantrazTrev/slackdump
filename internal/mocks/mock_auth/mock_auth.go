// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rusq/slackdump/v2/auth (interfaces: Provider)
//
// Generated by this command:
//
//	mockgen -destination ../internal/mocks/mock_auth/mock_auth.go github.com/rusq/slackdump/v2/auth Provider
//
// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	http "net/http"
	reflect "reflect"

	auth "github.com/rusq/slackdump/v2/auth"
	gomock "go.uber.org/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Cookies mocks base method.
func (m *MockProvider) Cookies() []*http.Cookie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookies")
	ret0, _ := ret[0].([]*http.Cookie)
	return ret0
}

// Cookies indicates an expected call of Cookies.
func (mr *MockProviderMockRecorder) Cookies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookies", reflect.TypeOf((*MockProvider)(nil).Cookies))
}

// HTTPClient mocks base method.
func (m *MockProvider) HTTPClient() (*http.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*http.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockProviderMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockProvider)(nil).HTTPClient))
}

// SlackToken mocks base method.
func (m *MockProvider) SlackToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlackToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// SlackToken indicates an expected call of SlackToken.
func (mr *MockProviderMockRecorder) SlackToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlackToken", reflect.TypeOf((*MockProvider)(nil).SlackToken))
}

// Test mocks base method.
func (m *MockProvider) Test(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Test", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Test indicates an expected call of Test.
func (mr *MockProviderMockRecorder) Test(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockProvider)(nil).Test), arg0)
}

// Type mocks base method.
func (m *MockProvider) Type() auth.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(auth.Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockProviderMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockProvider)(nil).Type))
}

// Validate mocks base method.
func (m *MockProvider) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockProviderMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockProvider)(nil).Validate))
}
