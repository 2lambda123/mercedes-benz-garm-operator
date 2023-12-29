// SPDX-License-Identifier: MIT
// Code generated by MockGen. DO NOT EDIT.
// Source: ../organization.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	organizations "github.com/cloudbase/garm/client/organizations"
	gomock "go.uber.org/mock/gomock"
)

// MockOrganizationClient is a mock of OrganizationClient interface.
type MockOrganizationClient struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationClientMockRecorder
}

// MockOrganizationClientMockRecorder is the mock recorder for MockOrganizationClient.
type MockOrganizationClientMockRecorder struct {
	mock *MockOrganizationClient
}

// NewMockOrganizationClient creates a new mock instance.
func NewMockOrganizationClient(ctrl *gomock.Controller) *MockOrganizationClient {
	mock := &MockOrganizationClient{ctrl: ctrl}
	mock.recorder = &MockOrganizationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationClient) EXPECT() *MockOrganizationClientMockRecorder {
	return m.recorder
}

// CreateOrganization mocks base method.
func (m *MockOrganizationClient) CreateOrganization(param *organizations.CreateOrgParams) (*organizations.CreateOrgOK, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", param)
	ret0, _ := ret[0].(*organizations.CreateOrgOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization.
func (mr *MockOrganizationClientMockRecorder) CreateOrganization(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockOrganizationClient)(nil).CreateOrganization), param)
}

// DeleteOrganization mocks base method.
func (m *MockOrganizationClient) DeleteOrganization(param *organizations.DeleteOrgParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganization", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganization indicates an expected call of DeleteOrganization.
func (mr *MockOrganizationClientMockRecorder) DeleteOrganization(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganization", reflect.TypeOf((*MockOrganizationClient)(nil).DeleteOrganization), param)
}

// GetOrganization mocks base method.
func (m *MockOrganizationClient) GetOrganization(param *organizations.GetOrgParams) (*organizations.GetOrgOK, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganization", param)
	ret0, _ := ret[0].(*organizations.GetOrgOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganization indicates an expected call of GetOrganization.
func (mr *MockOrganizationClientMockRecorder) GetOrganization(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganization", reflect.TypeOf((*MockOrganizationClient)(nil).GetOrganization), param)
}

// ListOrganizations mocks base method.
func (m *MockOrganizationClient) ListOrganizations(param *organizations.ListOrgsParams) (*organizations.ListOrgsOK, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrganizations", param)
	ret0, _ := ret[0].(*organizations.ListOrgsOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations.
func (mr *MockOrganizationClientMockRecorder) ListOrganizations(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockOrganizationClient)(nil).ListOrganizations), param)
}

// UpdateOrganization mocks base method.
func (m *MockOrganizationClient) UpdateOrganization(param *organizations.UpdateOrgParams) (*organizations.UpdateOrgOK, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrganization", param)
	ret0, _ := ret[0].(*organizations.UpdateOrgOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganization indicates an expected call of UpdateOrganization.
func (mr *MockOrganizationClientMockRecorder) UpdateOrganization(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganization", reflect.TypeOf((*MockOrganizationClient)(nil).UpdateOrganization), param)
}
