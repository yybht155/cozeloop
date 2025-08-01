// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/observability/domain/trace/service (interfaces: IngestionService)
//
// Generated by this command:
//
//	mockgen -destination=mocks/ingestion.go -package=mocks . IngestionService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIngestionService is a mock of IngestionService interface.
type MockIngestionService struct {
	ctrl     *gomock.Controller
	recorder *MockIngestionServiceMockRecorder
	isgomock struct{}
}

// MockIngestionServiceMockRecorder is the mock recorder for MockIngestionService.
type MockIngestionServiceMockRecorder struct {
	mock *MockIngestionService
}

// NewMockIngestionService creates a new mock instance.
func NewMockIngestionService(ctrl *gomock.Controller) *MockIngestionService {
	mock := &MockIngestionService{ctrl: ctrl}
	mock.recorder = &MockIngestionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngestionService) EXPECT() *MockIngestionServiceMockRecorder {
	return m.recorder
}

// RunAsync mocks base method.
func (m *MockIngestionService) RunAsync(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunAsync", ctx)
}

// RunAsync indicates an expected call of RunAsync.
func (mr *MockIngestionServiceMockRecorder) RunAsync(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAsync", reflect.TypeOf((*MockIngestionService)(nil).RunAsync), ctx)
}

// RunSync mocks base method.
func (m *MockIngestionService) RunSync(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunSync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSync indicates an expected call of RunSync.
func (mr *MockIngestionServiceMockRecorder) RunSync(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSync", reflect.TypeOf((*MockIngestionService)(nil).RunSync), ctx)
}
