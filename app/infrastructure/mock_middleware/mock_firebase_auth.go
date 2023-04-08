// Code generated by MockGen. DO NOT EDIT.
// Source: firebase_auth.go

// Package mock_middleware is a generated GoMock package.
package mock_middleware

import (
	context "context"
	reflect "reflect"

	auth "firebase.google.com/go/auth"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockFirebaseAuth is a mock of FirebaseAuth interface.
type MockFirebaseAuth struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseAuthMockRecorder
}

// MockFirebaseAuthMockRecorder is the mock recorder for MockFirebaseAuth.
type MockFirebaseAuthMockRecorder struct {
	mock *MockFirebaseAuth
}

// NewMockFirebaseAuth creates a new mock instance.
func NewMockFirebaseAuth(ctrl *gomock.Controller) *MockFirebaseAuth {
	mock := &MockFirebaseAuth{ctrl: ctrl}
	mock.recorder = &MockFirebaseAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirebaseAuth) EXPECT() *MockFirebaseAuthMockRecorder {
	return m.recorder
}

// Authentication mocks base method.
func (m *MockFirebaseAuth) Authentication(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Authentication", c)
}

// Authentication indicates an expected call of Authentication.
func (mr *MockFirebaseAuthMockRecorder) Authentication(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockFirebaseAuth)(nil).Authentication), c)
}

// MockFirebaseAuthClient is a mock of FirebaseAuthClient interface.
type MockFirebaseAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseAuthClientMockRecorder
}

// MockFirebaseAuthClientMockRecorder is the mock recorder for MockFirebaseAuthClient.
type MockFirebaseAuthClientMockRecorder struct {
	mock *MockFirebaseAuthClient
}

// NewMockFirebaseAuthClient creates a new mock instance.
func NewMockFirebaseAuthClient(ctrl *gomock.Controller) *MockFirebaseAuthClient {
	mock := &MockFirebaseAuthClient{ctrl: ctrl}
	mock.recorder = &MockFirebaseAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirebaseAuthClient) EXPECT() *MockFirebaseAuthClientMockRecorder {
	return m.recorder
}

// VerifyIDToken mocks base method.
func (m *MockFirebaseAuthClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", ctx, idToken)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockFirebaseAuthClientMockRecorder) VerifyIDToken(ctx, idToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockFirebaseAuthClient)(nil).VerifyIDToken), ctx, idToken)
}
