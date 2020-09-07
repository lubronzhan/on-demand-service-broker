// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/lubronzhan/on-demand-service-broker/purger"
)

type FakeDeregistrar struct {
	DeregisterStub        func(string) error
	deregisterMutex       sync.RWMutex
	deregisterArgsForCall []struct {
		arg1 string
	}
	deregisterReturns struct {
		result1 error
	}
	deregisterReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeregistrar) Deregister(arg1 string) error {
	fake.deregisterMutex.Lock()
	ret, specificReturn := fake.deregisterReturnsOnCall[len(fake.deregisterArgsForCall)]
	fake.deregisterArgsForCall = append(fake.deregisterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Deregister", []interface{}{arg1})
	fake.deregisterMutex.Unlock()
	if fake.DeregisterStub != nil {
		return fake.DeregisterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deregisterReturns
	return fakeReturns.result1
}

func (fake *FakeDeregistrar) DeregisterCallCount() int {
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	return len(fake.deregisterArgsForCall)
}

func (fake *FakeDeregistrar) DeregisterCalls(stub func(string) error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = stub
}

func (fake *FakeDeregistrar) DeregisterArgsForCall(i int) string {
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	argsForCall := fake.deregisterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeregistrar) DeregisterReturns(result1 error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = nil
	fake.deregisterReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeregistrar) DeregisterReturnsOnCall(i int, result1 error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = nil
	if fake.deregisterReturnsOnCall == nil {
		fake.deregisterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deregisterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeregistrar) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeregistrar) recordInvocation(key string, args []interface{}) {
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

var _ purger.Deregistrar = new(FakeDeregistrar)
