// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

// Package mock_p2p is a generated GoMock package.
package mock_p2p

import (
	gomock "github.com/golang/mock/gomock"
	core "github.com/iost-official/prototype/core"
	p2p "github.com/iost-official/prototype/p2p"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockRouter) Init(base core.Network, port uint16) error {
	ret := m.ctrl.Call(m, "Init", base, port)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRouterMockRecorder) Init(base, port interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRouter)(nil).Init), base, port)
}

// FilteredChan mocks base method
func (m *MockRouter) FilteredChan(filter p2p.Filter) (chan core.Request, error) {
	ret := m.ctrl.Call(m, "FilteredChan", filter)
	ret0, _ := ret[0].(chan core.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilteredChan indicates an expected call of FilteredChan
func (mr *MockRouterMockRecorder) FilteredChan(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilteredChan", reflect.TypeOf((*MockRouter)(nil).FilteredChan), filter)
}

// Run mocks base method
func (m *MockRouter) Run() {
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockRouterMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRouter)(nil).Run))
}

// Stop mocks base method
func (m *MockRouter) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockRouterMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRouter)(nil).Stop))
}

// Send mocks base method
func (m *MockRouter) Send(req core.Request) {
	m.ctrl.Call(m, "Send", req)
}

// Send indicates an expected call of Send
func (mr *MockRouterMockRecorder) Send(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRouter)(nil).Send), req)
}

// Broadcast mocks base method
func (m *MockRouter) Broadcast(req core.Request) {
	m.ctrl.Call(m, "Broadcast", req)
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockRouterMockRecorder) Broadcast(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockRouter)(nil).Broadcast), req)
}

// Download mocks base method
func (m *MockRouter) Download(req core.Request) chan []byte {
	ret := m.ctrl.Call(m, "Download", req)
	ret0, _ := ret[0].(chan []byte)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockRouterMockRecorder) Download(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockRouter)(nil).Download), req)
}