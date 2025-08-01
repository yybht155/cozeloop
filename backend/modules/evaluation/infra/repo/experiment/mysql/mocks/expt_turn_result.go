// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/cozeloop/backend/modules/evaluation/infra/repo/experiment/mysql (interfaces: ExptTurnResultDAO)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "github.com/coze-dev/cozeloop/backend/infra/db"
	entity "github.com/coze-dev/cozeloop/backend/modules/evaluation/domain/entity"
	model "github.com/coze-dev/cozeloop/backend/modules/evaluation/infra/repo/experiment/mysql/gorm_gen/model"
	gomock "go.uber.org/mock/gomock"
)

// MockExptTurnResultDAO is a mock of ExptTurnResultDAO interface.
type MockExptTurnResultDAO struct {
	ctrl     *gomock.Controller
	recorder *MockExptTurnResultDAOMockRecorder
}

// MockExptTurnResultDAOMockRecorder is the mock recorder for MockExptTurnResultDAO.
type MockExptTurnResultDAOMockRecorder struct {
	mock *MockExptTurnResultDAO
}

// NewMockExptTurnResultDAO creates a new mock instance.
func NewMockExptTurnResultDAO(ctrl *gomock.Controller) *MockExptTurnResultDAO {
	mock := &MockExptTurnResultDAO{ctrl: ctrl}
	mock.recorder = &MockExptTurnResultDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExptTurnResultDAO) EXPECT() *MockExptTurnResultDAOMockRecorder {
	return m.recorder
}

// BatchCreateNX mocks base method.
func (m *MockExptTurnResultDAO) BatchCreateNX(arg0 context.Context, arg1 []*model.ExptTurnResult, arg2 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCreateNX", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreateNX indicates an expected call of BatchCreateNX.
func (mr *MockExptTurnResultDAOMockRecorder) BatchCreateNX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateNX", reflect.TypeOf((*MockExptTurnResultDAO)(nil).BatchCreateNX), varargs...)
}

// BatchCreateNXRunLog mocks base method.
func (m *MockExptTurnResultDAO) BatchCreateNXRunLog(arg0 context.Context, arg1 []*model.ExptTurnResultRunLog, arg2 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCreateNXRunLog", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreateNXRunLog indicates an expected call of BatchCreateNXRunLog.
func (mr *MockExptTurnResultDAOMockRecorder) BatchCreateNXRunLog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateNXRunLog", reflect.TypeOf((*MockExptTurnResultDAO)(nil).BatchCreateNXRunLog), varargs...)
}

// BatchGet mocks base method.
func (m *MockExptTurnResultDAO) BatchGet(arg0 context.Context, arg1, arg2 int64, arg3 []int64, arg4 ...db.Option) ([]*model.ExptTurnResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGet", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGet indicates an expected call of BatchGet.
func (mr *MockExptTurnResultDAOMockRecorder) BatchGet(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGet", reflect.TypeOf((*MockExptTurnResultDAO)(nil).BatchGet), varargs...)
}

// CreateTurnEvaluatorRefs mocks base method.
func (m *MockExptTurnResultDAO) CreateTurnEvaluatorRefs(arg0 context.Context, arg1 []*model.ExptTurnEvaluatorResultRef, arg2 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTurnEvaluatorRefs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTurnEvaluatorRefs indicates an expected call of CreateTurnEvaluatorRefs.
func (mr *MockExptTurnResultDAOMockRecorder) CreateTurnEvaluatorRefs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTurnEvaluatorRefs", reflect.TypeOf((*MockExptTurnResultDAO)(nil).CreateTurnEvaluatorRefs), varargs...)
}

// GetItemTurnResults mocks base method.
func (m *MockExptTurnResultDAO) GetItemTurnResults(arg0 context.Context, arg1, arg2, arg3 int64, arg4 ...db.Option) ([]*model.ExptTurnResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItemTurnResults", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemTurnResults indicates an expected call of GetItemTurnResults.
func (mr *MockExptTurnResultDAOMockRecorder) GetItemTurnResults(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemTurnResults", reflect.TypeOf((*MockExptTurnResultDAO)(nil).GetItemTurnResults), varargs...)
}

// GetItemTurnRunLogs mocks base method.
func (m *MockExptTurnResultDAO) GetItemTurnRunLogs(arg0 context.Context, arg1, arg2, arg3, arg4 int64, arg5 ...db.Option) ([]*model.ExptTurnResultRunLog, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItemTurnRunLogs", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResultRunLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemTurnRunLogs indicates an expected call of GetItemTurnRunLogs.
func (mr *MockExptTurnResultDAOMockRecorder) GetItemTurnRunLogs(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemTurnRunLogs", reflect.TypeOf((*MockExptTurnResultDAO)(nil).GetItemTurnRunLogs), varargs...)
}

// ListTurnResult mocks base method.
func (m *MockExptTurnResultDAO) ListTurnResult(arg0 context.Context, arg1, arg2 int64, arg3 *entity.ExptTurnResultFilter, arg4 entity.Page, arg5 bool, arg6 ...db.Option) ([]*model.ExptTurnResult, int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4, arg5}
	for _, a := range arg6 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTurnResult", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResult)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTurnResult indicates an expected call of ListTurnResult.
func (mr *MockExptTurnResultDAOMockRecorder) ListTurnResult(arg0, arg1, arg2, arg3, arg4, arg5 interface{}, arg6 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5}, arg6...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTurnResult", reflect.TypeOf((*MockExptTurnResultDAO)(nil).ListTurnResult), varargs...)
}

// MGetItemTurnRunLogs mocks base method.
func (m *MockExptTurnResultDAO) MGetItemTurnRunLogs(arg0 context.Context, arg1, arg2 int64, arg3 []int64, arg4 int64, arg5 ...db.Option) ([]*model.ExptTurnResultRunLog, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGetItemTurnRunLogs", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResultRunLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGetItemTurnRunLogs indicates an expected call of MGetItemTurnRunLogs.
func (mr *MockExptTurnResultDAOMockRecorder) MGetItemTurnRunLogs(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGetItemTurnRunLogs", reflect.TypeOf((*MockExptTurnResultDAO)(nil).MGetItemTurnRunLogs), varargs...)
}

// SaveTurnResults mocks base method.
func (m *MockExptTurnResultDAO) SaveTurnResults(arg0 context.Context, arg1 []*model.ExptTurnResult, arg2 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveTurnResults", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTurnResults indicates an expected call of SaveTurnResults.
func (mr *MockExptTurnResultDAOMockRecorder) SaveTurnResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTurnResults", reflect.TypeOf((*MockExptTurnResultDAO)(nil).SaveTurnResults), varargs...)
}

// SaveTurnRunLogs mocks base method.
func (m *MockExptTurnResultDAO) SaveTurnRunLogs(arg0 context.Context, arg1 []*model.ExptTurnResultRunLog, arg2 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveTurnRunLogs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTurnRunLogs indicates an expected call of SaveTurnRunLogs.
func (mr *MockExptTurnResultDAOMockRecorder) SaveTurnRunLogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTurnRunLogs", reflect.TypeOf((*MockExptTurnResultDAO)(nil).SaveTurnRunLogs), varargs...)
}

// ScanTurnResults mocks base method.
func (m *MockExptTurnResultDAO) ScanTurnResults(arg0 context.Context, arg1 int64, arg2 []int32, arg3, arg4, arg5 int64, arg6 ...db.Option) ([]*model.ExptTurnResult, int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4, arg5}
	for _, a := range arg6 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanTurnResults", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResult)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ScanTurnResults indicates an expected call of ScanTurnResults.
func (mr *MockExptTurnResultDAOMockRecorder) ScanTurnResults(arg0, arg1, arg2, arg3, arg4, arg5 interface{}, arg6 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5}, arg6...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanTurnResults", reflect.TypeOf((*MockExptTurnResultDAO)(nil).ScanTurnResults), varargs...)
}

// ScanTurnRunLogs mocks base method.
func (m *MockExptTurnResultDAO) ScanTurnRunLogs(arg0 context.Context, arg1, arg2, arg3, arg4 int64, arg5 ...db.Option) ([]*model.ExptTurnResultRunLog, int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanTurnRunLogs", varargs...)
	ret0, _ := ret[0].([]*model.ExptTurnResultRunLog)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ScanTurnRunLogs indicates an expected call of ScanTurnRunLogs.
func (mr *MockExptTurnResultDAOMockRecorder) ScanTurnRunLogs(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanTurnRunLogs", reflect.TypeOf((*MockExptTurnResultDAO)(nil).ScanTurnRunLogs), varargs...)
}

// UpdateTurnResults mocks base method.
func (m *MockExptTurnResultDAO) UpdateTurnResults(arg0 context.Context, arg1 int64, arg2 []*entity.ItemTurnID, arg3 int64, arg4 map[string]interface{}, arg5 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTurnResults", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTurnResults indicates an expected call of UpdateTurnResults.
func (mr *MockExptTurnResultDAOMockRecorder) UpdateTurnResults(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTurnResults", reflect.TypeOf((*MockExptTurnResultDAO)(nil).UpdateTurnResults), varargs...)
}

// UpdateTurnResultsWithItemIDs mocks base method.
func (m *MockExptTurnResultDAO) UpdateTurnResultsWithItemIDs(arg0 context.Context, arg1 int64, arg2 []int64, arg3 int64, arg4 map[string]interface{}, arg5 ...db.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTurnResultsWithItemIDs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTurnResultsWithItemIDs indicates an expected call of UpdateTurnResultsWithItemIDs.
func (mr *MockExptTurnResultDAOMockRecorder) UpdateTurnResultsWithItemIDs(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTurnResultsWithItemIDs", reflect.TypeOf((*MockExptTurnResultDAO)(nil).UpdateTurnResultsWithItemIDs), varargs...)
}
