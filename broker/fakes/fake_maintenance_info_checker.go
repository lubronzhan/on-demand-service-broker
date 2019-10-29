// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/pivotal-cf/brokerapi/v7/domain"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
)

type FakeMaintenanceInfoChecker struct {
	CheckStub        func(string, *domain.MaintenanceInfo, []domain.Service, *log.Logger) error
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 string
		arg2 *domain.MaintenanceInfo
		arg3 []domain.Service
		arg4 *log.Logger
	}
	checkReturns struct {
		result1 error
	}
	checkReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMaintenanceInfoChecker) Check(arg1 string, arg2 *domain.MaintenanceInfo, arg3 []domain.Service, arg4 *log.Logger) error {
	var arg3Copy []domain.Service
	if arg3 != nil {
		arg3Copy = make([]domain.Service, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 string
		arg2 *domain.MaintenanceInfo
		arg3 []domain.Service
		arg4 *log.Logger
	}{arg1, arg2, arg3Copy, arg4})
	fake.recordInvocation("Check", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1
}

func (fake *FakeMaintenanceInfoChecker) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeMaintenanceInfoChecker) CheckCalls(stub func(string, *domain.MaintenanceInfo, []domain.Service, *log.Logger) error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeMaintenanceInfoChecker) CheckArgsForCall(i int) (string, *domain.MaintenanceInfo, []domain.Service, *log.Logger) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeMaintenanceInfoChecker) CheckReturns(result1 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMaintenanceInfoChecker) CheckReturnsOnCall(i int, result1 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMaintenanceInfoChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMaintenanceInfoChecker) recordInvocation(key string, args []interface{}) {
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

var _ broker.MaintenanceInfoChecker = new(FakeMaintenanceInfoChecker)
