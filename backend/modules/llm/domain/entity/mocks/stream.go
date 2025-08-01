// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/llm/domain/entity (interfaces: IStreamReader)
//
// Generated by this command:
//
//	mockgen -destination=mocks/stream.go -package=mocks . IStreamReader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	entity "github.com/coze-dev/cozeloop/backend/modules/llm/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIStreamReader is a mock of IStreamReader interface.
type MockIStreamReader struct {
	ctrl     *gomock.Controller
	recorder *MockIStreamReaderMockRecorder
	isgomock struct{}
}

// MockIStreamReaderMockRecorder is the mock recorder for MockIStreamReader.
type MockIStreamReaderMockRecorder struct {
	mock *MockIStreamReader
}

// NewMockIStreamReader creates a new mock instance.
func NewMockIStreamReader(ctrl *gomock.Controller) *MockIStreamReader {
	mock := &MockIStreamReader{ctrl: ctrl}
	mock.recorder = &MockIStreamReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStreamReader) EXPECT() *MockIStreamReaderMockRecorder {
	return m.recorder
}

// Recv mocks base method.
func (m *MockIStreamReader) Recv() (*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockIStreamReaderMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockIStreamReader)(nil).Recv))
}
