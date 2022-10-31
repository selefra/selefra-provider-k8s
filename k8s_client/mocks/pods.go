package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockPodsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockPodsClientMockRecorder
}

type MockPodsClientMockRecorder struct {
	mock *MockPodsClient
}

func NewMockPodsClient(ctrl *gomock.Controller) *MockPodsClient {
	mock := &MockPodsClient{ctrl: ctrl}
	mock.recorder = &MockPodsClientMockRecorder{mock}
	return mock
}

func (m *MockPodsClient) EXPECT() *MockPodsClientMockRecorder {
	return m.recorder
}

func (m *MockPodsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPodsClient)(nil).List), arg0, arg1)
}
