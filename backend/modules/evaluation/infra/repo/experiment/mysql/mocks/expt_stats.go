// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/modules/evaluation/infra/repo/experiment/mysql (interfaces: IExptStatsDAO)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/coze-dev/coze-loop/backend/modules/evaluation/domain/entity"
	model "github.com/coze-dev/coze-loop/backend/modules/evaluation/infra/repo/experiment/mysql/gorm_gen/model"
	gomock "go.uber.org/mock/gomock"
)

// MockIExptStatsDAO is a mock of IExptStatsDAO interface.
type MockIExptStatsDAO struct {
	ctrl     *gomock.Controller
	recorder *MockIExptStatsDAOMockRecorder
}

// MockIExptStatsDAOMockRecorder is the mock recorder for MockIExptStatsDAO.
type MockIExptStatsDAOMockRecorder struct {
	mock *MockIExptStatsDAO
}

// NewMockIExptStatsDAO creates a new mock instance.
func NewMockIExptStatsDAO(ctrl *gomock.Controller) *MockIExptStatsDAO {
	mock := &MockIExptStatsDAO{ctrl: ctrl}
	mock.recorder = &MockIExptStatsDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExptStatsDAO) EXPECT() *MockIExptStatsDAOMockRecorder {
	return m.recorder
}

// ArithOperateCount mocks base method.
func (m *MockIExptStatsDAO) ArithOperateCount(arg0 context.Context, arg1, arg2 int64, arg3 *entity.StatsCntArithOp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArithOperateCount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArithOperateCount indicates an expected call of ArithOperateCount.
func (mr *MockIExptStatsDAOMockRecorder) ArithOperateCount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArithOperateCount", reflect.TypeOf((*MockIExptStatsDAO)(nil).ArithOperateCount), arg0, arg1, arg2, arg3)
}

// Create mocks base method.
func (m *MockIExptStatsDAO) Create(arg0 context.Context, arg1 *model.ExptStats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIExptStatsDAOMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIExptStatsDAO)(nil).Create), arg0, arg1)
}

// Get mocks base method.
func (m *MockIExptStatsDAO) Get(arg0 context.Context, arg1, arg2 int64) (*model.ExptStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.ExptStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIExptStatsDAOMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIExptStatsDAO)(nil).Get), arg0, arg1, arg2)
}

// MGet mocks base method.
func (m *MockIExptStatsDAO) MGet(arg0 context.Context, arg1 []int64, arg2 int64) ([]*model.ExptStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MGet", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.ExptStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGet indicates an expected call of MGet.
func (mr *MockIExptStatsDAOMockRecorder) MGet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockIExptStatsDAO)(nil).MGet), arg0, arg1, arg2)
}

// Save mocks base method.
func (m *MockIExptStatsDAO) Save(arg0 context.Context, arg1 *model.ExptStats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIExptStatsDAOMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIExptStatsDAO)(nil).Save), arg0, arg1)
}

// UpdateByExptID mocks base method.
func (m *MockIExptStatsDAO) UpdateByExptID(arg0 context.Context, arg1, arg2 int64, arg3 *model.ExptStats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByExptID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByExptID indicates an expected call of UpdateByExptID.
func (mr *MockIExptStatsDAOMockRecorder) UpdateByExptID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByExptID", reflect.TypeOf((*MockIExptStatsDAO)(nil).UpdateByExptID), arg0, arg1, arg2, arg3)
}
