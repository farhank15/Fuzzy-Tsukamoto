// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/achievement/interface.go

// Package achievement is a generated GoMock package.
package achievement

import (
	context "context"
	achievement "go-tsukamoto/internal/app/dto/achievement"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAchievementService is a mock of AchievementService interface.
type MockAchievementService struct {
	ctrl     *gomock.Controller
	recorder *MockAchievementServiceMockRecorder
}

// MockAchievementServiceMockRecorder is the mock recorder for MockAchievementService.
type MockAchievementServiceMockRecorder struct {
	mock *MockAchievementService
}

// NewMockAchievementService creates a new mock instance.
func NewMockAchievementService(ctrl *gomock.Controller) *MockAchievementService {
	mock := &MockAchievementService{ctrl: ctrl}
	mock.recorder = &MockAchievementServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAchievementService) EXPECT() *MockAchievementServiceMockRecorder {
	return m.recorder
}

// CreateAchievement mocks base method.
func (m *MockAchievementService) CreateAchievement(ctx context.Context, req *achievement.CreateAchievementRequest) (*achievement.AchievementResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAchievement", ctx, req)
	ret0, _ := ret[0].(*achievement.AchievementResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAchievement indicates an expected call of CreateAchievement.
func (mr *MockAchievementServiceMockRecorder) CreateAchievement(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAchievement", reflect.TypeOf((*MockAchievementService)(nil).CreateAchievement), ctx, req)
}

// DeleteAchievement mocks base method.
func (m *MockAchievementService) DeleteAchievement(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAchievement", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAchievement indicates an expected call of DeleteAchievement.
func (mr *MockAchievementServiceMockRecorder) DeleteAchievement(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAchievement", reflect.TypeOf((*MockAchievementService)(nil).DeleteAchievement), ctx, id)
}

// GetAchievementByID mocks base method.
func (m *MockAchievementService) GetAchievementByID(ctx context.Context, id int) (*achievement.AchievementResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAchievementByID", ctx, id)
	ret0, _ := ret[0].(*achievement.AchievementResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAchievementByID indicates an expected call of GetAchievementByID.
func (mr *MockAchievementServiceMockRecorder) GetAchievementByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAchievementByID", reflect.TypeOf((*MockAchievementService)(nil).GetAchievementByID), ctx, id)
}

// GetAchievementsByUserID mocks base method.
func (m *MockAchievementService) GetAchievementsByUserID(ctx context.Context, userID int) ([]*achievement.AchievementResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAchievementsByUserID", ctx, userID)
	ret0, _ := ret[0].([]*achievement.AchievementResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAchievementsByUserID indicates an expected call of GetAchievementsByUserID.
func (mr *MockAchievementServiceMockRecorder) GetAchievementsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAchievementsByUserID", reflect.TypeOf((*MockAchievementService)(nil).GetAchievementsByUserID), ctx, userID)
}

// GetAllAchievements mocks base method.
func (m *MockAchievementService) GetAllAchievements(ctx context.Context) ([]*achievement.AchievementResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAchievements", ctx)
	ret0, _ := ret[0].([]*achievement.AchievementResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAchievements indicates an expected call of GetAllAchievements.
func (mr *MockAchievementServiceMockRecorder) GetAllAchievements(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAchievements", reflect.TypeOf((*MockAchievementService)(nil).GetAllAchievements), ctx)
}

// UpdateAchievement mocks base method.
func (m *MockAchievementService) UpdateAchievement(ctx context.Context, id int, req *achievement.UpdateAchievementRequest) (*achievement.AchievementResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAchievement", ctx, id, req)
	ret0, _ := ret[0].(*achievement.AchievementResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAchievement indicates an expected call of UpdateAchievement.
func (mr *MockAchievementServiceMockRecorder) UpdateAchievement(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAchievement", reflect.TypeOf((*MockAchievementService)(nil).UpdateAchievement), ctx, id, req)
}
