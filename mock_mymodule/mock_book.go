// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/KPI-golang-5/Library/pkg/repositories (interfaces: BookRepository)

// Package mock_mymodule is a generated GoMock package.
package mock_mymodule

import (
	reflect "reflect"

	models "github.com/KPI-golang-5/Library/pkg/models"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBookRepository) Create(arg0 *models.Book) *models.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*models.Book)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockBookRepository) Delete(arg0 int64) models.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(models.Book)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookRepository)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockBookRepository) GetAll(arg0, arg1, arg2 string) []models.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Book)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBookRepositoryMockRecorder) GetAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBookRepository)(nil).GetAll), arg0, arg1, arg2)
}

// GetById mocks base method.
func (m *MockBookRepository) GetById(arg0 int64) (*models.Book, *gorm.DB) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(*gorm.DB)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBookRepositoryMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBookRepository)(nil).GetById), arg0)
}
