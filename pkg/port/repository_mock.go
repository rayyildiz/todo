// Code generated by MockGen. DO NOT EDIT.
// Source: go.rayyildiz.dev/todo/pkg/port (interfaces: RepositorySaver,RepositoryReader,Repository)

// Package port is a generated GoMock package.
package port

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "go.rayyildiz.dev/todo/pkg/domain"
)

// MockRepositorySaver is a mock of RepositorySaver interface.
type MockRepositorySaver struct {
	ctrl     *gomock.Controller
	recorder *MockRepositorySaverMockRecorder
}

// MockRepositorySaverMockRecorder is the mock recorder for MockRepositorySaver.
type MockRepositorySaverMockRecorder struct {
	mock *MockRepositorySaver
}

// NewMockRepositorySaver creates a new mock instance.
func NewMockRepositorySaver(ctrl *gomock.Controller) *MockRepositorySaver {
	mock := &MockRepositorySaver{ctrl: ctrl}
	mock.recorder = &MockRepositorySaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositorySaver) EXPECT() *MockRepositorySaverMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepositorySaver) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositorySaverMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepositorySaver)(nil).Delete), arg0, arg1)
}

// Store mocks base method.
func (m *MockRepositorySaver) Store(arg0 context.Context, arg1 string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockRepositorySaverMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRepositorySaver)(nil).Store), arg0, arg1)
}

// Toggle mocks base method.
func (m *MockRepositorySaver) Toggle(arg0 context.Context, arg1 string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Toggle", arg0, arg1)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Toggle indicates an expected call of Toggle.
func (mr *MockRepositorySaverMockRecorder) Toggle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Toggle", reflect.TypeOf((*MockRepositorySaver)(nil).Toggle), arg0, arg1)
}

// MockRepositoryReader is a mock of RepositoryReader interface.
type MockRepositoryReader struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryReaderMockRecorder
}

// MockRepositoryReaderMockRecorder is the mock recorder for MockRepositoryReader.
type MockRepositoryReaderMockRecorder struct {
	mock *MockRepositoryReader
}

// NewMockRepositoryReader creates a new mock instance.
func NewMockRepositoryReader(ctrl *gomock.Controller) *MockRepositoryReader {
	mock := &MockRepositoryReader{ctrl: ctrl}
	mock.recorder = &MockRepositoryReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryReader) EXPECT() *MockRepositoryReaderMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockRepositoryReader) FindAll(arg0 context.Context) ([]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryReaderMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepositoryReader)(nil).FindAll), arg0)
}

// FindById mocks base method.
func (m *MockRepositoryReader) FindById(arg0 context.Context, arg1 string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockRepositoryReaderMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepositoryReader)(nil).FindById), arg0, arg1)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(arg0 context.Context) ([]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), arg0)
}

// FindById mocks base method.
func (m *MockRepository) FindById(arg0 context.Context, arg1 string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockRepositoryMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), arg0, arg1)
}

// Store mocks base method.
func (m *MockRepository) Store(arg0 context.Context, arg1 string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockRepositoryMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRepository)(nil).Store), arg0, arg1)
}

// Toggle mocks base method.
func (m *MockRepository) Toggle(arg0 context.Context, arg1 string) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Toggle", arg0, arg1)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Toggle indicates an expected call of Toggle.
func (mr *MockRepositoryMockRecorder) Toggle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Toggle", reflect.TypeOf((*MockRepository)(nil).Toggle), arg0, arg1)
}