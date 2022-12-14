// Code generated by MockGen. DO NOT EDIT.
// Source: paste.repo.go

// Package mock_repo is a generated GoMock package.
package repo

import (
	models "carbon/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPasteRepo is a mock of PasteRepo interface.
type MockPasteRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPasteRepoMockRecorder
}

// MockPasteRepoMockRecorder is the mock recorder for MockPasteRepo.
type MockPasteRepoMockRecorder struct {
	mock *MockPasteRepo
}

// NewMockPasteRepo creates a new mock instance.
func NewMockPasteRepo(ctrl *gomock.Controller) *MockPasteRepo {
	mock := &MockPasteRepo{ctrl: ctrl}
	mock.recorder = &MockPasteRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasteRepo) EXPECT() *MockPasteRepoMockRecorder {
	return m.recorder
}

// CreatePaste mocks base method.
func (m *MockPasteRepo) CreatePaste(paste *models.Paste) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePaste", paste)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePaste indicates an expected call of CreatePaste.
func (mr *MockPasteRepoMockRecorder) CreatePaste(paste interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePaste", reflect.TypeOf((*MockPasteRepo)(nil).CreatePaste), paste)
}

// ViewPasteByUrl mocks base method.
func (m *MockPasteRepo) ViewPasteByUrl(url string) (*models.Paste, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewPasteByUrl", url)
	ret0, _ := ret[0].(*models.Paste)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewPasteByUrl indicates an expected call of ViewPasteByUrl.
func (mr *MockPasteRepoMockRecorder) ViewPasteByUrl(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewPasteByUrl", reflect.TypeOf((*MockPasteRepo)(nil).ViewPasteByUrl), url)
}
