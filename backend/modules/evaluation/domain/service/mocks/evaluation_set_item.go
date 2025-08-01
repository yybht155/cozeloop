// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/service (interfaces: EvaluationSetItemService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockEvaluationSetItemService is a mock of EvaluationSetItemService interface.
type MockEvaluationSetItemService struct {
	ctrl     *gomock.Controller
	recorder *MockEvaluationSetItemServiceMockRecorder
}

// MockEvaluationSetItemServiceMockRecorder is the mock recorder for MockEvaluationSetItemService.
type MockEvaluationSetItemServiceMockRecorder struct {
	mock *MockEvaluationSetItemService
}

// NewMockEvaluationSetItemService creates a new mock instance.
func NewMockEvaluationSetItemService(ctrl *gomock.Controller) *MockEvaluationSetItemService {
	mock := &MockEvaluationSetItemService{ctrl: ctrl}
	mock.recorder = &MockEvaluationSetItemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvaluationSetItemService) EXPECT() *MockEvaluationSetItemServiceMockRecorder {
	return m.recorder
}

// BatchCreateEvaluationSetItems mocks base method.
func (m *MockEvaluationSetItemService) BatchCreateEvaluationSetItems(arg0 context.Context, arg1 *entity.BatchCreateEvaluationSetItemsParam) (map[int64]int64, []*entity.ItemErrorGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateEvaluationSetItems", arg0, arg1)
	ret0, _ := ret[0].(map[int64]int64)
	ret1, _ := ret[1].([]*entity.ItemErrorGroup)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BatchCreateEvaluationSetItems indicates an expected call of BatchCreateEvaluationSetItems.
func (mr *MockEvaluationSetItemServiceMockRecorder) BatchCreateEvaluationSetItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateEvaluationSetItems", reflect.TypeOf((*MockEvaluationSetItemService)(nil).BatchCreateEvaluationSetItems), arg0, arg1)
}

// BatchDeleteEvaluationSetItems mocks base method.
func (m *MockEvaluationSetItemService) BatchDeleteEvaluationSetItems(arg0 context.Context, arg1, arg2 int64, arg3 []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteEvaluationSetItems", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchDeleteEvaluationSetItems indicates an expected call of BatchDeleteEvaluationSetItems.
func (mr *MockEvaluationSetItemServiceMockRecorder) BatchDeleteEvaluationSetItems(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteEvaluationSetItems", reflect.TypeOf((*MockEvaluationSetItemService)(nil).BatchDeleteEvaluationSetItems), arg0, arg1, arg2, arg3)
}

// BatchGetEvaluationSetItems mocks base method.
func (m *MockEvaluationSetItemService) BatchGetEvaluationSetItems(arg0 context.Context, arg1 *entity.BatchGetEvaluationSetItemsParam) ([]*entity.EvaluationSetItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetEvaluationSetItems", arg0, arg1)
	ret0, _ := ret[0].([]*entity.EvaluationSetItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetEvaluationSetItems indicates an expected call of BatchGetEvaluationSetItems.
func (mr *MockEvaluationSetItemServiceMockRecorder) BatchGetEvaluationSetItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetEvaluationSetItems", reflect.TypeOf((*MockEvaluationSetItemService)(nil).BatchGetEvaluationSetItems), arg0, arg1)
}

// ClearEvaluationSetDraftItem mocks base method.
func (m *MockEvaluationSetItemService) ClearEvaluationSetDraftItem(arg0 context.Context, arg1, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearEvaluationSetDraftItem", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearEvaluationSetDraftItem indicates an expected call of ClearEvaluationSetDraftItem.
func (mr *MockEvaluationSetItemServiceMockRecorder) ClearEvaluationSetDraftItem(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearEvaluationSetDraftItem", reflect.TypeOf((*MockEvaluationSetItemService)(nil).ClearEvaluationSetDraftItem), arg0, arg1, arg2)
}

// ListEvaluationSetItems mocks base method.
func (m *MockEvaluationSetItemService) ListEvaluationSetItems(arg0 context.Context, arg1 *entity.ListEvaluationSetItemsParam) ([]*entity.EvaluationSetItem, *int64, *string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEvaluationSetItems", arg0, arg1)
	ret0, _ := ret[0].([]*entity.EvaluationSetItem)
	ret1, _ := ret[1].(*int64)
	ret2, _ := ret[2].(*string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListEvaluationSetItems indicates an expected call of ListEvaluationSetItems.
func (mr *MockEvaluationSetItemServiceMockRecorder) ListEvaluationSetItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEvaluationSetItems", reflect.TypeOf((*MockEvaluationSetItemService)(nil).ListEvaluationSetItems), arg0, arg1)
}

// UpdateEvaluationSetItem mocks base method.
func (m *MockEvaluationSetItemService) UpdateEvaluationSetItem(arg0 context.Context, arg1, arg2, arg3 int64, arg4 []*entity.Turn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvaluationSetItem", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEvaluationSetItem indicates an expected call of UpdateEvaluationSetItem.
func (mr *MockEvaluationSetItemServiceMockRecorder) UpdateEvaluationSetItem(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvaluationSetItem", reflect.TypeOf((*MockEvaluationSetItemService)(nil).UpdateEvaluationSetItem), arg0, arg1, arg2, arg3, arg4)
}
