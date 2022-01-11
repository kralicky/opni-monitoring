// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/rbac/rbac.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// ListTenantsForUser mocks base method.
func (m *MockProvider) ListTenantsForUser(ctx context.Context, userID string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTenantsForUser", ctx, userID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTenantsForUser indicates an expected call of ListTenantsForUser.
func (mr *MockProviderMockRecorder) ListTenantsForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTenantsForUser", reflect.TypeOf((*MockProvider)(nil).ListTenantsForUser), ctx, userID)
}
