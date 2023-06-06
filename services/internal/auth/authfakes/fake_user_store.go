// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"context"
	"sync"

	"github.com/JordanRad/chatbook/services/internal/auth"
)

type FakeUserStore struct {
	GetUserByEmailStub        func(context.Context, string) (*auth.User, error)
	getUserByEmailMutex       sync.RWMutex
	getUserByEmailArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getUserByEmailReturns struct {
		result1 *auth.User
		result2 error
	}
	getUserByEmailReturnsOnCall map[int]struct {
		result1 *auth.User
		result2 error
	}
	RegisterStub        func(context.Context, *auth.User) (*auth.User, error)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 context.Context
		arg2 *auth.User
	}
	registerReturns struct {
		result1 *auth.User
		result2 error
	}
	registerReturnsOnCall map[int]struct {
		result1 *auth.User
		result2 error
	}
	UpdateProfileNamesStub        func(context.Context, string, string, string) error
	updateProfileNamesMutex       sync.RWMutex
	updateProfileNamesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	updateProfileNamesReturns struct {
		result1 error
	}
	updateProfileNamesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserStore) GetUserByEmail(arg1 context.Context, arg2 string) (*auth.User, error) {
	fake.getUserByEmailMutex.Lock()
	ret, specificReturn := fake.getUserByEmailReturnsOnCall[len(fake.getUserByEmailArgsForCall)]
	fake.getUserByEmailArgsForCall = append(fake.getUserByEmailArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetUserByEmailStub
	fakeReturns := fake.getUserByEmailReturns
	fake.recordInvocation("GetUserByEmail", []interface{}{arg1, arg2})
	fake.getUserByEmailMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserStore) GetUserByEmailCallCount() int {
	fake.getUserByEmailMutex.RLock()
	defer fake.getUserByEmailMutex.RUnlock()
	return len(fake.getUserByEmailArgsForCall)
}

func (fake *FakeUserStore) GetUserByEmailCalls(stub func(context.Context, string) (*auth.User, error)) {
	fake.getUserByEmailMutex.Lock()
	defer fake.getUserByEmailMutex.Unlock()
	fake.GetUserByEmailStub = stub
}

func (fake *FakeUserStore) GetUserByEmailArgsForCall(i int) (context.Context, string) {
	fake.getUserByEmailMutex.RLock()
	defer fake.getUserByEmailMutex.RUnlock()
	argsForCall := fake.getUserByEmailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserStore) GetUserByEmailReturns(result1 *auth.User, result2 error) {
	fake.getUserByEmailMutex.Lock()
	defer fake.getUserByEmailMutex.Unlock()
	fake.GetUserByEmailStub = nil
	fake.getUserByEmailReturns = struct {
		result1 *auth.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserStore) GetUserByEmailReturnsOnCall(i int, result1 *auth.User, result2 error) {
	fake.getUserByEmailMutex.Lock()
	defer fake.getUserByEmailMutex.Unlock()
	fake.GetUserByEmailStub = nil
	if fake.getUserByEmailReturnsOnCall == nil {
		fake.getUserByEmailReturnsOnCall = make(map[int]struct {
			result1 *auth.User
			result2 error
		})
	}
	fake.getUserByEmailReturnsOnCall[i] = struct {
		result1 *auth.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserStore) Register(arg1 context.Context, arg2 *auth.User) (*auth.User, error) {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 context.Context
		arg2 *auth.User
	}{arg1, arg2})
	stub := fake.RegisterStub
	fakeReturns := fake.registerReturns
	fake.recordInvocation("Register", []interface{}{arg1, arg2})
	fake.registerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserStore) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeUserStore) RegisterCalls(stub func(context.Context, *auth.User) (*auth.User, error)) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = stub
}

func (fake *FakeUserStore) RegisterArgsForCall(i int) (context.Context, *auth.User) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	argsForCall := fake.registerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserStore) RegisterReturns(result1 *auth.User, result2 error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 *auth.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserStore) RegisterReturnsOnCall(i int, result1 *auth.User, result2 error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 *auth.User
			result2 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 *auth.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserStore) UpdateProfileNames(arg1 context.Context, arg2 string, arg3 string, arg4 string) error {
	fake.updateProfileNamesMutex.Lock()
	ret, specificReturn := fake.updateProfileNamesReturnsOnCall[len(fake.updateProfileNamesArgsForCall)]
	fake.updateProfileNamesArgsForCall = append(fake.updateProfileNamesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.UpdateProfileNamesStub
	fakeReturns := fake.updateProfileNamesReturns
	fake.recordInvocation("UpdateProfileNames", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateProfileNamesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserStore) UpdateProfileNamesCallCount() int {
	fake.updateProfileNamesMutex.RLock()
	defer fake.updateProfileNamesMutex.RUnlock()
	return len(fake.updateProfileNamesArgsForCall)
}

func (fake *FakeUserStore) UpdateProfileNamesCalls(stub func(context.Context, string, string, string) error) {
	fake.updateProfileNamesMutex.Lock()
	defer fake.updateProfileNamesMutex.Unlock()
	fake.UpdateProfileNamesStub = stub
}

func (fake *FakeUserStore) UpdateProfileNamesArgsForCall(i int) (context.Context, string, string, string) {
	fake.updateProfileNamesMutex.RLock()
	defer fake.updateProfileNamesMutex.RUnlock()
	argsForCall := fake.updateProfileNamesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUserStore) UpdateProfileNamesReturns(result1 error) {
	fake.updateProfileNamesMutex.Lock()
	defer fake.updateProfileNamesMutex.Unlock()
	fake.UpdateProfileNamesStub = nil
	fake.updateProfileNamesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserStore) UpdateProfileNamesReturnsOnCall(i int, result1 error) {
	fake.updateProfileNamesMutex.Lock()
	defer fake.updateProfileNamesMutex.Unlock()
	fake.UpdateProfileNamesStub = nil
	if fake.updateProfileNamesReturnsOnCall == nil {
		fake.updateProfileNamesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateProfileNamesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getUserByEmailMutex.RLock()
	defer fake.getUserByEmailMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.updateProfileNamesMutex.RLock()
	defer fake.updateProfileNamesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserStore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ auth.UserStore = new(FakeUserStore)
