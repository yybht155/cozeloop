// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/component/rpc (interfaces: IUserProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"

	"go.uber.org/mock/gomock"

	"github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/entity"
)

// MockIUserProvider is a mock of IUserProvider interface.
type MockIUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIUserProviderMockRecorder
}

// MockIUserProviderMockRecorder is the mock recorder for MockIUserProvider.
type MockIUserProviderMockRecorder struct {
	mock *MockIUserProvider
}

// NewMockIUserProvider creates a new mock instance.
func NewMockIUserProvider(ctrl *gomock.Controller) *MockIUserProvider {
	mock := &MockIUserProvider{ctrl: ctrl}
	mock.recorder = &MockIUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserProvider) EXPECT() *MockIUserProviderMockRecorder {
	return m.recorder
}

// MGetUserInfo mocks base method.
func (m *MockIUserProvider) MGetUserInfo(arg0 context.Context, arg1 []string) ([]*entity.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MGetUserInfo", arg0, arg1)
	ret0, _ := ret[0].([]*entity.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGetUserInfo indicates an expected call of MGetUserInfo.
func (mr *MockIUserProviderMockRecorder) MGetUserInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGetUserInfo", reflect.TypeOf((*MockIUserProvider)(nil).MGetUserInfo), arg0, arg1)
}
