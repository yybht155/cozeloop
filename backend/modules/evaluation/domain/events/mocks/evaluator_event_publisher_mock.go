// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/events (interfaces: EvaluatorEventPublisher)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/entity"
)

// MockEvaluatorEventPublisher is a mock of EvaluatorEventPublisher interface.
type MockEvaluatorEventPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockEvaluatorEventPublisherMockRecorder
}

// MockEvaluatorEventPublisherMockRecorder is the mock recorder for MockEvaluatorEventPublisher.
type MockEvaluatorEventPublisherMockRecorder struct {
	mock *MockEvaluatorEventPublisher
}

// NewMockEvaluatorEventPublisher creates a new mock instance.
func NewMockEvaluatorEventPublisher(ctrl *gomock.Controller) *MockEvaluatorEventPublisher {
	mock := &MockEvaluatorEventPublisher{ctrl: ctrl}
	mock.recorder = &MockEvaluatorEventPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvaluatorEventPublisher) EXPECT() *MockEvaluatorEventPublisherMockRecorder {
	return m.recorder
}

// PublishEvaluatorRecordCorrection mocks base method.
func (m *MockEvaluatorEventPublisher) PublishEvaluatorRecordCorrection(arg0 context.Context, arg1 *entity.EvaluatorRecordCorrectionEvent, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishEvaluatorRecordCorrection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishEvaluatorRecordCorrection indicates an expected call of PublishEvaluatorRecordCorrection.
func (mr *MockEvaluatorEventPublisherMockRecorder) PublishEvaluatorRecordCorrection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishEvaluatorRecordCorrection", reflect.TypeOf((*MockEvaluatorEventPublisher)(nil).PublishEvaluatorRecordCorrection), arg0, arg1, arg2)
}
