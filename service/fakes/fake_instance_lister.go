// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/lubronzhan/on-demand-service-broker/service"
)

type FakeInstanceLister struct {
	InstancesStub        func(map[string]string) ([]service.Instance, error)
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct {
		arg1 map[string]string
	}
	instancesReturns struct {
		result1 []service.Instance
		result2 error
	}
	instancesReturnsOnCall map[int]struct {
		result1 []service.Instance
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstanceLister) Instances(arg1 map[string]string) ([]service.Instance, error) {
	fake.instancesMutex.Lock()
	ret, specificReturn := fake.instancesReturnsOnCall[len(fake.instancesArgsForCall)]
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	fake.recordInvocation("Instances", []interface{}{arg1})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instancesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstanceLister) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeInstanceLister) InstancesCalls(stub func(map[string]string) ([]service.Instance, error)) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = stub
}

func (fake *FakeInstanceLister) InstancesArgsForCall(i int) map[string]string {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	argsForCall := fake.instancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstanceLister) InstancesReturns(result1 []service.Instance, result2 error) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceLister) InstancesReturnsOnCall(i int, result1 []service.Instance, result2 error) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []service.Instance
			result2 error
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeInstanceLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstanceLister) recordInvocation(key string, args []interface{}) {
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

var _ service.InstanceLister = new(FakeInstanceLister)
