// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/lubronzhan/on-demand-service-broker/cf"
	"github.com/lubronzhan/on-demand-service-broker/service"
)

type FakeCFListerClient struct {
	GetServiceInstancesStub        func(cf.GetInstancesFilter, *log.Logger) ([]cf.Instance, error)
	getServiceInstancesMutex       sync.RWMutex
	getServiceInstancesArgsForCall []struct {
		arg1 cf.GetInstancesFilter
		arg2 *log.Logger
	}
	getServiceInstancesReturns struct {
		result1 []cf.Instance
		result2 error
	}
	getServiceInstancesReturnsOnCall map[int]struct {
		result1 []cf.Instance
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFListerClient) GetServiceInstances(arg1 cf.GetInstancesFilter, arg2 *log.Logger) ([]cf.Instance, error) {
	fake.getServiceInstancesMutex.Lock()
	ret, specificReturn := fake.getServiceInstancesReturnsOnCall[len(fake.getServiceInstancesArgsForCall)]
	fake.getServiceInstancesArgsForCall = append(fake.getServiceInstancesArgsForCall, struct {
		arg1 cf.GetInstancesFilter
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetServiceInstances", []interface{}{arg1, arg2})
	fake.getServiceInstancesMutex.Unlock()
	if fake.GetServiceInstancesStub != nil {
		return fake.GetServiceInstancesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceInstancesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFListerClient) GetServiceInstancesCallCount() int {
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	return len(fake.getServiceInstancesArgsForCall)
}

func (fake *FakeCFListerClient) GetServiceInstancesCalls(stub func(cf.GetInstancesFilter, *log.Logger) ([]cf.Instance, error)) {
	fake.getServiceInstancesMutex.Lock()
	defer fake.getServiceInstancesMutex.Unlock()
	fake.GetServiceInstancesStub = stub
}

func (fake *FakeCFListerClient) GetServiceInstancesArgsForCall(i int) (cf.GetInstancesFilter, *log.Logger) {
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	argsForCall := fake.getServiceInstancesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFListerClient) GetServiceInstancesReturns(result1 []cf.Instance, result2 error) {
	fake.getServiceInstancesMutex.Lock()
	defer fake.getServiceInstancesMutex.Unlock()
	fake.GetServiceInstancesStub = nil
	fake.getServiceInstancesReturns = struct {
		result1 []cf.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCFListerClient) GetServiceInstancesReturnsOnCall(i int, result1 []cf.Instance, result2 error) {
	fake.getServiceInstancesMutex.Lock()
	defer fake.getServiceInstancesMutex.Unlock()
	fake.GetServiceInstancesStub = nil
	if fake.getServiceInstancesReturnsOnCall == nil {
		fake.getServiceInstancesReturnsOnCall = make(map[int]struct {
			result1 []cf.Instance
			result2 error
		})
	}
	fake.getServiceInstancesReturnsOnCall[i] = struct {
		result1 []cf.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCFListerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFListerClient) recordInvocation(key string, args []interface{}) {
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

var _ service.CFListerClient = new(FakeCFListerClient)
