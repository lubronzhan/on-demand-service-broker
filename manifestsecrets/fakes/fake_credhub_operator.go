// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
	"github.com/pivotal-cf/on-demand-service-broker/manifestsecrets"
)

type FakeCredhubOperator struct {
	BulkGetStub        func(map[string]boshdirector.Variable, *log.Logger) (map[string]string, error)
	bulkGetMutex       sync.RWMutex
	bulkGetArgsForCall []struct {
		arg1 map[string]boshdirector.Variable
		arg2 *log.Logger
	}
	bulkGetReturns struct {
		result1 map[string]string
		result2 error
	}
	bulkGetReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	FindNameLikeStub        func(name string, logger *log.Logger) ([]string, error)
	findNameLikeMutex       sync.RWMutex
	findNameLikeArgsForCall []struct {
		name   string
		logger *log.Logger
	}
	findNameLikeReturns struct {
		result1 []string
		result2 error
	}
	findNameLikeReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	BulkDeleteStub        func(paths []string, logger *log.Logger) error
	bulkDeleteMutex       sync.RWMutex
	bulkDeleteArgsForCall []struct {
		paths  []string
		logger *log.Logger
	}
	bulkDeleteReturns struct {
		result1 error
	}
	bulkDeleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredhubOperator) BulkGet(arg1 map[string]boshdirector.Variable, arg2 *log.Logger) (map[string]string, error) {
	fake.bulkGetMutex.Lock()
	ret, specificReturn := fake.bulkGetReturnsOnCall[len(fake.bulkGetArgsForCall)]
	fake.bulkGetArgsForCall = append(fake.bulkGetArgsForCall, struct {
		arg1 map[string]boshdirector.Variable
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("BulkGet", []interface{}{arg1, arg2})
	fake.bulkGetMutex.Unlock()
	if fake.BulkGetStub != nil {
		return fake.BulkGetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bulkGetReturns.result1, fake.bulkGetReturns.result2
}

func (fake *FakeCredhubOperator) BulkGetCallCount() int {
	fake.bulkGetMutex.RLock()
	defer fake.bulkGetMutex.RUnlock()
	return len(fake.bulkGetArgsForCall)
}

func (fake *FakeCredhubOperator) BulkGetArgsForCall(i int) (map[string]boshdirector.Variable, *log.Logger) {
	fake.bulkGetMutex.RLock()
	defer fake.bulkGetMutex.RUnlock()
	return fake.bulkGetArgsForCall[i].arg1, fake.bulkGetArgsForCall[i].arg2
}

func (fake *FakeCredhubOperator) BulkGetReturns(result1 map[string]string, result2 error) {
	fake.BulkGetStub = nil
	fake.bulkGetReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubOperator) BulkGetReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.BulkGetStub = nil
	if fake.bulkGetReturnsOnCall == nil {
		fake.bulkGetReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.bulkGetReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubOperator) FindNameLike(name string, logger *log.Logger) ([]string, error) {
	fake.findNameLikeMutex.Lock()
	ret, specificReturn := fake.findNameLikeReturnsOnCall[len(fake.findNameLikeArgsForCall)]
	fake.findNameLikeArgsForCall = append(fake.findNameLikeArgsForCall, struct {
		name   string
		logger *log.Logger
	}{name, logger})
	fake.recordInvocation("FindNameLike", []interface{}{name, logger})
	fake.findNameLikeMutex.Unlock()
	if fake.FindNameLikeStub != nil {
		return fake.FindNameLikeStub(name, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findNameLikeReturns.result1, fake.findNameLikeReturns.result2
}

func (fake *FakeCredhubOperator) FindNameLikeCallCount() int {
	fake.findNameLikeMutex.RLock()
	defer fake.findNameLikeMutex.RUnlock()
	return len(fake.findNameLikeArgsForCall)
}

func (fake *FakeCredhubOperator) FindNameLikeArgsForCall(i int) (string, *log.Logger) {
	fake.findNameLikeMutex.RLock()
	defer fake.findNameLikeMutex.RUnlock()
	return fake.findNameLikeArgsForCall[i].name, fake.findNameLikeArgsForCall[i].logger
}

func (fake *FakeCredhubOperator) FindNameLikeReturns(result1 []string, result2 error) {
	fake.FindNameLikeStub = nil
	fake.findNameLikeReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubOperator) FindNameLikeReturnsOnCall(i int, result1 []string, result2 error) {
	fake.FindNameLikeStub = nil
	if fake.findNameLikeReturnsOnCall == nil {
		fake.findNameLikeReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.findNameLikeReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubOperator) BulkDelete(paths []string, logger *log.Logger) error {
	var pathsCopy []string
	if paths != nil {
		pathsCopy = make([]string, len(paths))
		copy(pathsCopy, paths)
	}
	fake.bulkDeleteMutex.Lock()
	ret, specificReturn := fake.bulkDeleteReturnsOnCall[len(fake.bulkDeleteArgsForCall)]
	fake.bulkDeleteArgsForCall = append(fake.bulkDeleteArgsForCall, struct {
		paths  []string
		logger *log.Logger
	}{pathsCopy, logger})
	fake.recordInvocation("BulkDelete", []interface{}{pathsCopy, logger})
	fake.bulkDeleteMutex.Unlock()
	if fake.BulkDeleteStub != nil {
		return fake.BulkDeleteStub(paths, logger)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.bulkDeleteReturns.result1
}

func (fake *FakeCredhubOperator) BulkDeleteCallCount() int {
	fake.bulkDeleteMutex.RLock()
	defer fake.bulkDeleteMutex.RUnlock()
	return len(fake.bulkDeleteArgsForCall)
}

func (fake *FakeCredhubOperator) BulkDeleteArgsForCall(i int) ([]string, *log.Logger) {
	fake.bulkDeleteMutex.RLock()
	defer fake.bulkDeleteMutex.RUnlock()
	return fake.bulkDeleteArgsForCall[i].paths, fake.bulkDeleteArgsForCall[i].logger
}

func (fake *FakeCredhubOperator) BulkDeleteReturns(result1 error) {
	fake.BulkDeleteStub = nil
	fake.bulkDeleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredhubOperator) BulkDeleteReturnsOnCall(i int, result1 error) {
	fake.BulkDeleteStub = nil
	if fake.bulkDeleteReturnsOnCall == nil {
		fake.bulkDeleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bulkDeleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredhubOperator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bulkGetMutex.RLock()
	defer fake.bulkGetMutex.RUnlock()
	fake.findNameLikeMutex.RLock()
	defer fake.findNameLikeMutex.RUnlock()
	fake.bulkDeleteMutex.RLock()
	defer fake.bulkDeleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredhubOperator) recordInvocation(key string, args []interface{}) {
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

var _ manifestsecrets.CredhubOperator = new(FakeCredhubOperator)
