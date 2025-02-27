// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/fuzzy/interface.go

// Package fuzzy is a generated GoMock package.
package fuzzy

import (
	context "context"
	dto "go-tsukamoto/internal/app/dto/fuzzy"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFuzzyServiceInterface is a mock of FuzzyServiceInterface interface.
type MockFuzzyServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFuzzyServiceInterfaceMockRecorder
}

// MockFuzzyServiceInterfaceMockRecorder is the mock recorder for MockFuzzyServiceInterface.
type MockFuzzyServiceInterfaceMockRecorder struct {
	mock *MockFuzzyServiceInterface
}

// NewMockFuzzyServiceInterface creates a new mock instance.
func NewMockFuzzyServiceInterface(ctrl *gomock.Controller) *MockFuzzyServiceInterface {
	mock := &MockFuzzyServiceInterface{ctrl: ctrl}
	mock.recorder = &MockFuzzyServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFuzzyServiceInterface) EXPECT() *MockFuzzyServiceInterfaceMockRecorder {
	return m.recorder
}

// CalculateFuzzy mocks base method.
func (m *MockFuzzyServiceInterface) CalculateFuzzy(ctx context.Context, studentID int) (*dto.FuzzyResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateFuzzy", ctx, studentID)
	ret0, _ := ret[0].(*dto.FuzzyResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateFuzzy indicates an expected call of CalculateFuzzy.
func (mr *MockFuzzyServiceInterfaceMockRecorder) CalculateFuzzy(ctx, studentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateFuzzy", reflect.TypeOf((*MockFuzzyServiceInterface)(nil).CalculateFuzzy), ctx, studentID)
}
