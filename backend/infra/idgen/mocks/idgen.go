// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/infra/idgen (interfaces: IIDGenerator)
//
// Generated by this command:
//
//	mockgen -destination=mocks/idgen.go -package=mocks . IIDGenerator
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIIDGenerator is a mock of IIDGenerator interface.
type MockIIDGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIIDGeneratorMockRecorder
	isgomock struct{}
}

// MockIIDGeneratorMockRecorder is the mock recorder for MockIIDGenerator.
type MockIIDGeneratorMockRecorder struct {
	mock *MockIIDGenerator
}

// NewMockIIDGenerator creates a new mock instance.
func NewMockIIDGenerator(ctrl *gomock.Controller) *MockIIDGenerator {
	mock := &MockIIDGenerator{ctrl: ctrl}
	mock.recorder = &MockIIDGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIIDGenerator) EXPECT() *MockIIDGeneratorMockRecorder {
	return m.recorder
}

// GenID mocks base method.
func (m *MockIIDGenerator) GenID(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenID", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenID indicates an expected call of GenID.
func (mr *MockIIDGeneratorMockRecorder) GenID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenID", reflect.TypeOf((*MockIIDGenerator)(nil).GenID), ctx)
}

// GenMultiIDs mocks base method.
func (m *MockIIDGenerator) GenMultiIDs(ctx context.Context, counts int) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenMultiIDs", ctx, counts)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenMultiIDs indicates an expected call of GenMultiIDs.
func (mr *MockIIDGeneratorMockRecorder) GenMultiIDs(ctx, counts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenMultiIDs", reflect.TypeOf((*MockIIDGenerator)(nil).GenMultiIDs), ctx, counts)
}
