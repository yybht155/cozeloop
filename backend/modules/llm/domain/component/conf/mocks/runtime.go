// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/llm/domain/component/conf (interfaces: IConfigRuntime)
//
// Generated by this command:
//
//	mockgen -destination=mocks/runtime.go -package=mocks . IConfigRuntime
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIConfigRuntime is a mock of IConfigRuntime interface.
type MockIConfigRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockIConfigRuntimeMockRecorder
	isgomock struct{}
}

// MockIConfigRuntimeMockRecorder is the mock recorder for MockIConfigRuntime.
type MockIConfigRuntimeMockRecorder struct {
	mock *MockIConfigRuntime
}

// NewMockIConfigRuntime creates a new mock instance.
func NewMockIConfigRuntime(ctrl *gomock.Controller) *MockIConfigRuntime {
	mock := &MockIConfigRuntime{ctrl: ctrl}
	mock.recorder = &MockIConfigRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConfigRuntime) EXPECT() *MockIConfigRuntimeMockRecorder {
	return m.recorder
}

// NeedCvtURLToBase64 mocks base method.
func (m *MockIConfigRuntime) NeedCvtURLToBase64() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedCvtURLToBase64")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedCvtURLToBase64 indicates an expected call of NeedCvtURLToBase64.
func (mr *MockIConfigRuntimeMockRecorder) NeedCvtURLToBase64() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedCvtURLToBase64", reflect.TypeOf((*MockIConfigRuntime)(nil).NeedCvtURLToBase64))
}
