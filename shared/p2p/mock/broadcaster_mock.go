// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/shared/p2p (interfaces: Broadcaster)

// Package mock_p2p is a generated GoMock package.
package mock_p2p

import (
	reflect "reflect"

	proto "github.com/gogo/protobuf/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockBroadcaster is a mock of Broadcaster interface
type MockBroadcaster struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcasterMockRecorder
}

// MockBroadcasterMockRecorder is the mock recorder for MockBroadcaster
type MockBroadcasterMockRecorder struct {
	mock *MockBroadcaster
}

// NewMockBroadcaster creates a new mock instance
func NewMockBroadcaster(ctrl *gomock.Controller) *MockBroadcaster {
	mock := &MockBroadcaster{ctrl: ctrl}
	mock.recorder = &MockBroadcasterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBroadcaster) EXPECT() *MockBroadcasterMockRecorder {
	return m.recorder
}

// Broadcast mocks base method
func (m *MockBroadcaster) Broadcast(arg0 proto.Message) {
	m.ctrl.Call(m, "Broadcast", arg0)
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockBroadcasterMockRecorder) Broadcast(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockBroadcaster)(nil).Broadcast), arg0)
}
