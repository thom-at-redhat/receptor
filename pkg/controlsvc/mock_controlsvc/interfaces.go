// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/controlsvc/interfaces.go

// Package mock_controlsvc is a generated GoMock package.
package mock_controlsvc

import (
	context "context"
	controlsvc "github.com/ansible/receptor/pkg/controlsvc"
	logger "github.com/ansible/receptor/pkg/logger"
	netceptor "github.com/ansible/receptor/pkg/netceptor"
	gomock "github.com/golang/mock/gomock"
	io "io"
	net "net"
	reflect "reflect"
)

// MockControlCommandType is a mock of ControlCommandType interface
type MockControlCommandType struct {
	ctrl     *gomock.Controller
	recorder *MockControlCommandTypeMockRecorder
}

// MockControlCommandTypeMockRecorder is the mock recorder for MockControlCommandType
type MockControlCommandTypeMockRecorder struct {
	mock *MockControlCommandType
}

// NewMockControlCommandType creates a new mock instance
func NewMockControlCommandType(ctrl *gomock.Controller) *MockControlCommandType {
	mock := &MockControlCommandType{ctrl: ctrl}
	mock.recorder = &MockControlCommandTypeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockControlCommandType) EXPECT() *MockControlCommandTypeMockRecorder {
	return m.recorder
}

// InitFromString mocks base method
func (m *MockControlCommandType) InitFromString(arg0 string) (controlsvc.ControlCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitFromString", arg0)
	ret0, _ := ret[0].(controlsvc.ControlCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitFromString indicates an expected call of InitFromString
func (mr *MockControlCommandTypeMockRecorder) InitFromString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitFromString", reflect.TypeOf((*MockControlCommandType)(nil).InitFromString), arg0)
}

// InitFromJSON mocks base method
func (m *MockControlCommandType) InitFromJSON(arg0 map[string]interface{}) (controlsvc.ControlCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitFromJSON", arg0)
	ret0, _ := ret[0].(controlsvc.ControlCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitFromJSON indicates an expected call of InitFromJSON
func (mr *MockControlCommandTypeMockRecorder) InitFromJSON(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitFromJSON", reflect.TypeOf((*MockControlCommandType)(nil).InitFromJSON), arg0)
}

// MockControlCommand is a mock of ControlCommand interface
type MockControlCommand struct {
	ctrl     *gomock.Controller
	recorder *MockControlCommandMockRecorder
}

// MockControlCommandMockRecorder is the mock recorder for MockControlCommand
type MockControlCommandMockRecorder struct {
	mock *MockControlCommand
}

// NewMockControlCommand creates a new mock instance
func NewMockControlCommand(ctrl *gomock.Controller) *MockControlCommand {
	mock := &MockControlCommand{ctrl: ctrl}
	mock.recorder = &MockControlCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockControlCommand) EXPECT() *MockControlCommandMockRecorder {
	return m.recorder
}

// ControlFunc mocks base method
func (m *MockControlCommand) ControlFunc(arg0 context.Context, arg1 *netceptor.Netceptor, arg2 controlsvc.ControlFuncOperations) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlFunc", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControlFunc indicates an expected call of ControlFunc
func (mr *MockControlCommandMockRecorder) ControlFunc(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlFunc", reflect.TypeOf((*MockControlCommand)(nil).ControlFunc), arg0, arg1, arg2)
}

// MockControlFuncOperations is a mock of ControlFuncOperations interface
type MockControlFuncOperations struct {
	ctrl     *gomock.Controller
	recorder *MockControlFuncOperationsMockRecorder
}

// MockControlFuncOperationsMockRecorder is the mock recorder for MockControlFuncOperations
type MockControlFuncOperationsMockRecorder struct {
	mock *MockControlFuncOperations
}

// NewMockControlFuncOperations creates a new mock instance
func NewMockControlFuncOperations(ctrl *gomock.Controller) *MockControlFuncOperations {
	mock := &MockControlFuncOperations{ctrl: ctrl}
	mock.recorder = &MockControlFuncOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockControlFuncOperations) EXPECT() *MockControlFuncOperationsMockRecorder {
	return m.recorder
}

// BridgeConn mocks base method
func (m *MockControlFuncOperations) BridgeConn(message string, bc io.ReadWriteCloser, bcName string, logger *logger.ReceptorLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BridgeConn", message, bc, bcName, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// BridgeConn indicates an expected call of BridgeConn
func (mr *MockControlFuncOperationsMockRecorder) BridgeConn(message, bc, bcName, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BridgeConn", reflect.TypeOf((*MockControlFuncOperations)(nil).BridgeConn), message, bc, bcName, logger)
}

// ReadFromConn mocks base method
func (m *MockControlFuncOperations) ReadFromConn(message string, out io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromConn", message, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFromConn indicates an expected call of ReadFromConn
func (mr *MockControlFuncOperationsMockRecorder) ReadFromConn(message, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromConn", reflect.TypeOf((*MockControlFuncOperations)(nil).ReadFromConn), message, out)
}

// WriteToConn mocks base method
func (m *MockControlFuncOperations) WriteToConn(message string, in chan []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToConn", message, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToConn indicates an expected call of WriteToConn
func (mr *MockControlFuncOperationsMockRecorder) WriteToConn(message, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToConn", reflect.TypeOf((*MockControlFuncOperations)(nil).WriteToConn), message, in)
}

// Close mocks base method
func (m *MockControlFuncOperations) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockControlFuncOperationsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockControlFuncOperations)(nil).Close))
}

// RemoteAddr mocks base method
func (m *MockControlFuncOperations) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockControlFuncOperationsMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockControlFuncOperations)(nil).RemoteAddr))
}