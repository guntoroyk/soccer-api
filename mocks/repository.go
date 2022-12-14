// Code generated by MockGen. DO NOT EDIT.
// Source: ../repository/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/guntoroyk/soccer-api/entity"
)

// MockTeamRepoItf is a mock of TeamRepoItf interface.
type MockTeamRepoItf struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRepoItfMockRecorder
}

// MockTeamRepoItfMockRecorder is the mock recorder for MockTeamRepoItf.
type MockTeamRepoItfMockRecorder struct {
	mock *MockTeamRepoItf
}

// NewMockTeamRepoItf creates a new mock instance.
func NewMockTeamRepoItf(ctrl *gomock.Controller) *MockTeamRepoItf {
	mock := &MockTeamRepoItf{ctrl: ctrl}
	mock.recorder = &MockTeamRepoItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamRepoItf) EXPECT() *MockTeamRepoItfMockRecorder {
	return m.recorder
}

// AddPlayerToTeam mocks base method.
func (m *MockTeamRepoItf) AddPlayerToTeam(teamID int, player *entity.Player) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlayerToTeam", teamID, player)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlayerToTeam indicates an expected call of AddPlayerToTeam.
func (mr *MockTeamRepoItfMockRecorder) AddPlayerToTeam(teamID, player interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlayerToTeam", reflect.TypeOf((*MockTeamRepoItf)(nil).AddPlayerToTeam), teamID, player)
}

// CreateTeam mocks base method.
func (m *MockTeamRepoItf) CreateTeam(team *entity.Team) (*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", team)
	ret0, _ := ret[0].(*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockTeamRepoItfMockRecorder) CreateTeam(team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockTeamRepoItf)(nil).CreateTeam), team)
}

// GetTeam mocks base method.
func (m *MockTeamRepoItf) GetTeam(id int) (*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeam", id)
	ret0, _ := ret[0].(*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeam indicates an expected call of GetTeam.
func (mr *MockTeamRepoItfMockRecorder) GetTeam(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeam", reflect.TypeOf((*MockTeamRepoItf)(nil).GetTeam), id)
}

// GetTeams mocks base method.
func (m *MockTeamRepoItf) GetTeams() ([]*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeams")
	ret0, _ := ret[0].([]*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeams indicates an expected call of GetTeams.
func (mr *MockTeamRepoItfMockRecorder) GetTeams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeams", reflect.TypeOf((*MockTeamRepoItf)(nil).GetTeams))
}

// MockPlayerRepoItf is a mock of PlayerRepoItf interface.
type MockPlayerRepoItf struct {
	ctrl     *gomock.Controller
	recorder *MockPlayerRepoItfMockRecorder
}

// MockPlayerRepoItfMockRecorder is the mock recorder for MockPlayerRepoItf.
type MockPlayerRepoItfMockRecorder struct {
	mock *MockPlayerRepoItf
}

// NewMockPlayerRepoItf creates a new mock instance.
func NewMockPlayerRepoItf(ctrl *gomock.Controller) *MockPlayerRepoItf {
	mock := &MockPlayerRepoItf{ctrl: ctrl}
	mock.recorder = &MockPlayerRepoItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayerRepoItf) EXPECT() *MockPlayerRepoItfMockRecorder {
	return m.recorder
}

// GetPlayer mocks base method.
func (m *MockPlayerRepoItf) GetPlayer(id int) (*entity.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayer", id)
	ret0, _ := ret[0].(*entity.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlayer indicates an expected call of GetPlayer.
func (mr *MockPlayerRepoItfMockRecorder) GetPlayer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayer", reflect.TypeOf((*MockPlayerRepoItf)(nil).GetPlayer), id)
}
