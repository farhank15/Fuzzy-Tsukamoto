// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/course/interface.go

// Package course is a generated GoMock package.
package course

import (
	context "context"
	models "go-tsukamoto/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCourseRepositoryInterface is a mock of CourseRepositoryInterface interface.
type MockCourseRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCourseRepositoryInterfaceMockRecorder
}

// MockCourseRepositoryInterfaceMockRecorder is the mock recorder for MockCourseRepositoryInterface.
type MockCourseRepositoryInterfaceMockRecorder struct {
	mock *MockCourseRepositoryInterface
}

// NewMockCourseRepositoryInterface creates a new mock instance.
func NewMockCourseRepositoryInterface(ctrl *gomock.Controller) *MockCourseRepositoryInterface {
	mock := &MockCourseRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCourseRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourseRepositoryInterface) EXPECT() *MockCourseRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateCourse mocks base method.
func (m *MockCourseRepositoryInterface) CreateCourse(ctx context.Context, course *models.Course) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCourse", ctx, course)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCourse indicates an expected call of CreateCourse.
func (mr *MockCourseRepositoryInterfaceMockRecorder) CreateCourse(ctx, course interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCourse", reflect.TypeOf((*MockCourseRepositoryInterface)(nil).CreateCourse), ctx, course)
}

// DeleteCourse mocks base method.
func (m *MockCourseRepositoryInterface) DeleteCourse(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCourse", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCourse indicates an expected call of DeleteCourse.
func (mr *MockCourseRepositoryInterfaceMockRecorder) DeleteCourse(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCourse", reflect.TypeOf((*MockCourseRepositoryInterface)(nil).DeleteCourse), ctx, id)
}

// GetCourseByID mocks base method.
func (m *MockCourseRepositoryInterface) GetCourseByID(ctx context.Context, id int) (*models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseByID", ctx, id)
	ret0, _ := ret[0].(*models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseByID indicates an expected call of GetCourseByID.
func (mr *MockCourseRepositoryInterfaceMockRecorder) GetCourseByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseByID", reflect.TypeOf((*MockCourseRepositoryInterface)(nil).GetCourseByID), ctx, id)
}

// GetCourses mocks base method.
func (m *MockCourseRepositoryInterface) GetCourses(ctx context.Context) ([]*models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourses", ctx)
	ret0, _ := ret[0].([]*models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourses indicates an expected call of GetCourses.
func (mr *MockCourseRepositoryInterfaceMockRecorder) GetCourses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourses", reflect.TypeOf((*MockCourseRepositoryInterface)(nil).GetCourses), ctx)
}

// UpdateCourse mocks base method.
func (m *MockCourseRepositoryInterface) UpdateCourse(ctx context.Context, course *models.Course) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourse", ctx, course)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCourse indicates an expected call of UpdateCourse.
func (mr *MockCourseRepositoryInterfaceMockRecorder) UpdateCourse(ctx, course interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourse", reflect.TypeOf((*MockCourseRepositoryInterface)(nil).UpdateCourse), ctx, course)
}
