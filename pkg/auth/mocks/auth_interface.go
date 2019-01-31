// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/auth (interfaces: ConfigSaver)

package auth_test

import (
	auth "github.com/jenkins-x/jx/pkg/auth"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockConfigSaver struct {
	fail func(message string, callerSkip ...int)
}

func NewMockConfigSaver() *MockConfigSaver {
	return &MockConfigSaver{fail: pegomock.GlobalFailHandler}
}

func (mock *MockConfigSaver) LoadConfig() (*auth.AuthConfig, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfigSaver().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("LoadConfig", params, []reflect.Type{reflect.TypeOf((**auth.AuthConfig)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *auth.AuthConfig
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*auth.AuthConfig)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConfigSaver) SaveConfig(_param0 *auth.AuthConfig) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfigSaver().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SaveConfig", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockConfigSaver) VerifyWasCalledOnce() *VerifierConfigSaver {
	return &VerifierConfigSaver{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockConfigSaver) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierConfigSaver {
	return &VerifierConfigSaver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockConfigSaver) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierConfigSaver {
	return &VerifierConfigSaver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockConfigSaver) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierConfigSaver {
	return &VerifierConfigSaver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierConfigSaver struct {
	mock                   *MockConfigSaver
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierConfigSaver) LoadConfig() *ConfigSaver_LoadConfig_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "LoadConfig", params, verifier.timeout)
	return &ConfigSaver_LoadConfig_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ConfigSaver_LoadConfig_OngoingVerification struct {
	mock              *MockConfigSaver
	methodInvocations []pegomock.MethodInvocation
}

func (c *ConfigSaver_LoadConfig_OngoingVerification) GetCapturedArguments() {
}

func (c *ConfigSaver_LoadConfig_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierConfigSaver) SaveConfig(_param0 *auth.AuthConfig) *ConfigSaver_SaveConfig_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SaveConfig", params, verifier.timeout)
	return &ConfigSaver_SaveConfig_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ConfigSaver_SaveConfig_OngoingVerification struct {
	mock              *MockConfigSaver
	methodInvocations []pegomock.MethodInvocation
}

func (c *ConfigSaver_SaveConfig_OngoingVerification) GetCapturedArguments() *auth.AuthConfig {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *ConfigSaver_SaveConfig_OngoingVerification) GetAllCapturedArguments() (_param0 []*auth.AuthConfig) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*auth.AuthConfig, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*auth.AuthConfig)
		}
	}
	return
}
