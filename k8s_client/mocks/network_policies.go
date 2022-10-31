package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/networking/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockNetworkPoliciesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockNetworkPoliciesClientMockRecorder
}

type MockNetworkPoliciesClientMockRecorder struct {
	mock *MockNetworkPoliciesClient
}

func NewMockNetworkPoliciesClient(ctrl *gomock.Controller) *MockNetworkPoliciesClient {
	mock := &MockNetworkPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockNetworkPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkPoliciesClient) EXPECT() *MockNetworkPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkPoliciesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NetworkPolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.NetworkPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPoliciesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworkPoliciesClient)(nil).List), arg0, arg1)
}
