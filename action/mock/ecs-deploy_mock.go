// Code generated by MockGen. DO NOT EDIT.
// Source: ./action/ecs-deploy.go

// Package mock_action is a generated GoMock package.
package mock_action

import (
	ecs "github.com/adroll/ecs-ship/ecs"
	ecs0 "github.com/aws/aws-sdk-go/service/ecs"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockECSDeployClient is a mock of ECSDeployClient interface
type MockECSDeployClient struct {
	ctrl     *gomock.Controller
	recorder *MockECSDeployClientMockRecorder
}

// MockECSDeployClientMockRecorder is the mock recorder for MockECSDeployClient
type MockECSDeployClientMockRecorder struct {
	mock *MockECSDeployClient
}

// NewMockECSDeployClient creates a new mock instance
func NewMockECSDeployClient(ctrl *gomock.Controller) *MockECSDeployClient {
	mock := &MockECSDeployClient{ctrl: ctrl}
	mock.recorder = &MockECSDeployClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECSDeployClient) EXPECT() *MockECSDeployClientMockRecorder {
	return m.recorder
}

// GetService mocks base method
func (m *MockECSDeployClient) GetService(clusterName, serviceName string) (*ecs0.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", clusterName, serviceName)
	ret0, _ := ret[0].(*ecs0.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockECSDeployClientMockRecorder) GetService(clusterName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockECSDeployClient)(nil).GetService), clusterName, serviceName)
}

// LooksGood mocks base method
func (m *MockECSDeployClient) LooksGood(service *ecs0.Service) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LooksGood", service)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LooksGood indicates an expected call of LooksGood
func (mr *MockECSDeployClientMockRecorder) LooksGood(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LooksGood", reflect.TypeOf((*MockECSDeployClient)(nil).LooksGood), service)
}

// CopyTaskDefinition mocks base method
func (m *MockECSDeployClient) CopyTaskDefinition(service *ecs0.Service) (*ecs0.RegisterTaskDefinitionInput, *ecs0.TaskDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyTaskDefinition", service)
	ret0, _ := ret[0].(*ecs0.RegisterTaskDefinitionInput)
	ret1, _ := ret[1].(*ecs0.TaskDefinition)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CopyTaskDefinition indicates an expected call of CopyTaskDefinition
func (mr *MockECSDeployClientMockRecorder) CopyTaskDefinition(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyTaskDefinition", reflect.TypeOf((*MockECSDeployClient)(nil).CopyTaskDefinition), service)
}

// WaitUntilGood mocks base method
func (m *MockECSDeployClient) WaitUntilGood(service *ecs0.Service, timeout *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilGood", service, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilGood indicates an expected call of WaitUntilGood
func (mr *MockECSDeployClientMockRecorder) WaitUntilGood(service, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilGood", reflect.TypeOf((*MockECSDeployClient)(nil).WaitUntilGood), service, timeout)
}

// RegisterTaskDefinition mocks base method
func (m *MockECSDeployClient) RegisterTaskDefinition(input *ecs0.RegisterTaskDefinitionInput) (*ecs0.TaskDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTaskDefinition", input)
	ret0, _ := ret[0].(*ecs0.TaskDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTaskDefinition indicates an expected call of RegisterTaskDefinition
func (mr *MockECSDeployClientMockRecorder) RegisterTaskDefinition(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTaskDefinition", reflect.TypeOf((*MockECSDeployClient)(nil).RegisterTaskDefinition), input)
}

// UpdateTaskDefinition mocks base method
func (m *MockECSDeployClient) UpdateTaskDefinition(service *ecs0.Service, task *ecs0.TaskDefinition) (*ecs0.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskDefinition", service, task)
	ret0, _ := ret[0].(*ecs0.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskDefinition indicates an expected call of UpdateTaskDefinition
func (mr *MockECSDeployClientMockRecorder) UpdateTaskDefinition(service, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskDefinition", reflect.TypeOf((*MockECSDeployClient)(nil).UpdateTaskDefinition), service, task)
}

// MockECSDeployTaskConfig is a mock of ECSDeployTaskConfig interface
type MockECSDeployTaskConfig struct {
	ctrl     *gomock.Controller
	recorder *MockECSDeployTaskConfigMockRecorder
}

// MockECSDeployTaskConfigMockRecorder is the mock recorder for MockECSDeployTaskConfig
type MockECSDeployTaskConfigMockRecorder struct {
	mock *MockECSDeployTaskConfig
}

// NewMockECSDeployTaskConfig creates a new mock instance
func NewMockECSDeployTaskConfig(ctrl *gomock.Controller) *MockECSDeployTaskConfig {
	mock := &MockECSDeployTaskConfig{ctrl: ctrl}
	mock.recorder = &MockECSDeployTaskConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECSDeployTaskConfig) EXPECT() *MockECSDeployTaskConfigMockRecorder {
	return m.recorder
}

// ApplyTo mocks base method
func (m *MockECSDeployTaskConfig) ApplyTo(input *ecs0.RegisterTaskDefinitionInput) (*ecs0.RegisterTaskDefinitionInput, *ecs.TaskConfigDiff) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTo", input)
	ret0, _ := ret[0].(*ecs0.RegisterTaskDefinitionInput)
	ret1, _ := ret[1].(*ecs.TaskConfigDiff)
	return ret0, ret1
}

// ApplyTo indicates an expected call of ApplyTo
func (mr *MockECSDeployTaskConfigMockRecorder) ApplyTo(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTo", reflect.TypeOf((*MockECSDeployTaskConfig)(nil).ApplyTo), input)
}
