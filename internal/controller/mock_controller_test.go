// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package controller is a generated GoMock package.
package controller

import (
	context "context"
	reflect "reflect"

	entity "github.com/asyauqi1511/test/internal/entity"
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockEmployeeResource is a mock of EmployeeResource interface.
type MockEmployeeResource struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeResourceMockRecorder
}

// MockEmployeeResourceMockRecorder is the mock recorder for MockEmployeeResource.
type MockEmployeeResourceMockRecorder struct {
	mock *MockEmployeeResource
}

// NewMockEmployeeResource creates a new mock instance.
func NewMockEmployeeResource(ctrl *gomock.Controller) *MockEmployeeResource {
	mock := &MockEmployeeResource{ctrl: ctrl}
	mock.recorder = &MockEmployeeResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeResource) EXPECT() *MockEmployeeResourceMockRecorder {
	return m.recorder
}

// DeleteEmployeeByID mocks base method.
func (m *MockEmployeeResource) DeleteEmployeeByID(ctx context.Context, tx *sqlx.Tx, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployeeByID", ctx, tx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployeeByID indicates an expected call of DeleteEmployeeByID.
func (mr *MockEmployeeResourceMockRecorder) DeleteEmployeeByID(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployeeByID", reflect.TypeOf((*MockEmployeeResource)(nil).DeleteEmployeeByID), ctx, tx, id)
}

// GetAllEmployees mocks base method.
func (m *MockEmployeeResource) GetAllEmployees(ctx context.Context) ([]entity.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmployees", ctx)
	ret0, _ := ret[0].([]entity.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEmployees indicates an expected call of GetAllEmployees.
func (mr *MockEmployeeResourceMockRecorder) GetAllEmployees(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmployees", reflect.TypeOf((*MockEmployeeResource)(nil).GetAllEmployees), ctx)
}

// GetEmployeeByID mocks base method.
func (m *MockEmployeeResource) GetEmployeeByID(ctx context.Context, id int64) (entity.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByID", ctx, id)
	ret0, _ := ret[0].(entity.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByID indicates an expected call of GetEmployeeByID.
func (mr *MockEmployeeResourceMockRecorder) GetEmployeeByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByID", reflect.TypeOf((*MockEmployeeResource)(nil).GetEmployeeByID), ctx, id)
}

// InsertEmployee mocks base method.
func (m *MockEmployeeResource) InsertEmployee(ctx context.Context, tx *sqlx.Tx, data entity.Employee) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEmployee", ctx, tx, data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertEmployee indicates an expected call of InsertEmployee.
func (mr *MockEmployeeResourceMockRecorder) InsertEmployee(ctx, tx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEmployee", reflect.TypeOf((*MockEmployeeResource)(nil).InsertEmployee), ctx, tx, data)
}

// UpdateEmployee mocks base method.
func (m *MockEmployeeResource) UpdateEmployee(ctx context.Context, tx *sqlx.Tx, data entity.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", ctx, tx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockEmployeeResourceMockRecorder) UpdateEmployee(ctx, tx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmployeeResource)(nil).UpdateEmployee), ctx, tx, data)
}
