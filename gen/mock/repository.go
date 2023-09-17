// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/minguu42/simoom/pkg/domain/model"
	gomock "go.uber.org/mock/gomock"
)

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

// CreateProject mocks base method.
func (m *MockRepository) CreateProject(ctx context.Context, p model.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockRepositoryMockRecorder) CreateProject(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockRepository)(nil).CreateProject), ctx, p)
}

// CreateTask mocks base method.
func (m *MockRepository) CreateTask(ctx context.Context, t model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockRepositoryMockRecorder) CreateTask(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockRepository)(nil).CreateTask), ctx, t)
}

// DeleteProject mocks base method.
func (m *MockRepository) DeleteProject(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockRepositoryMockRecorder) DeleteProject(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockRepository)(nil).DeleteProject), ctx, id)
}

// DeleteTask mocks base method.
func (m *MockRepository) DeleteTask(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockRepositoryMockRecorder) DeleteTask(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockRepository)(nil).DeleteTask), ctx, id)
}

// GetProjectByID mocks base method.
func (m *MockRepository) GetProjectByID(ctx context.Context, id string) (model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectByID", ctx, id)
	ret0, _ := ret[0].(model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectByID indicates an expected call of GetProjectByID.
func (mr *MockRepositoryMockRecorder) GetProjectByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectByID", reflect.TypeOf((*MockRepository)(nil).GetProjectByID), ctx, id)
}

// GetTaskByID mocks base method.
func (m *MockRepository) GetTaskByID(ctx context.Context, id string) (model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, id)
	ret0, _ := ret[0].(model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockRepositoryMockRecorder) GetTaskByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockRepository)(nil).GetTaskByID), ctx, id)
}

// ListProjectsByUserID mocks base method.
func (m *MockRepository) ListProjectsByUserID(ctx context.Context, userID string, limit, offset uint) ([]model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsByUserID", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsByUserID indicates an expected call of ListProjectsByUserID.
func (mr *MockRepositoryMockRecorder) ListProjectsByUserID(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsByUserID", reflect.TypeOf((*MockRepository)(nil).ListProjectsByUserID), ctx, userID, limit, offset)
}

// ListTasksByProjectID mocks base method.
func (m *MockRepository) ListTasksByProjectID(ctx context.Context, projectID string, limit, offset uint) ([]model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasksByProjectID", ctx, projectID, limit, offset)
	ret0, _ := ret[0].([]model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasksByProjectID indicates an expected call of ListTasksByProjectID.
func (mr *MockRepositoryMockRecorder) ListTasksByProjectID(ctx, projectID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasksByProjectID", reflect.TypeOf((*MockRepository)(nil).ListTasksByProjectID), ctx, projectID, limit, offset)
}

// UpdateProject mocks base method.
func (m *MockRepository) UpdateProject(ctx context.Context, p model.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockRepositoryMockRecorder) UpdateProject(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockRepository)(nil).UpdateProject), ctx, p)
}

// UpdateTask mocks base method.
func (m *MockRepository) UpdateTask(ctx context.Context, t model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockRepositoryMockRecorder) UpdateTask(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockRepository)(nil).UpdateTask), ctx, t)
}
