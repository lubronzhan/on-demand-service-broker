// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/values"
	"code.cloudfoundry.org/credhub-cli/credhub/permissions"
	"github.com/pivotal-cf/on-demand-service-broker/credhub"
)

type FakeCredhubClient struct {
	GetByIdStub        func(id string) (credentials.Credential, error)
	getByIdMutex       sync.RWMutex
	getByIdArgsForCall []struct {
		id string
	}
	getByIdReturns struct {
		result1 credentials.Credential
		result2 error
	}
	getByIdReturnsOnCall map[int]struct {
		result1 credentials.Credential
		result2 error
	}
	GetLatestVersionStub        func(name string) (credentials.Credential, error)
	getLatestVersionMutex       sync.RWMutex
	getLatestVersionArgsForCall []struct {
		name string
	}
	getLatestVersionReturns struct {
		result1 credentials.Credential
		result2 error
	}
	getLatestVersionReturnsOnCall map[int]struct {
		result1 credentials.Credential
		result2 error
	}
	FindByPartialNameStub        func(partialName string) (credentials.FindResults, error)
	findByPartialNameMutex       sync.RWMutex
	findByPartialNameArgsForCall []struct {
		partialName string
	}
	findByPartialNameReturns struct {
		result1 credentials.FindResults
		result2 error
	}
	findByPartialNameReturnsOnCall map[int]struct {
		result1 credentials.FindResults
		result2 error
	}
	SetJSONStub        func(name string, value values.JSON) (credentials.JSON, error)
	setJSONMutex       sync.RWMutex
	setJSONArgsForCall []struct {
		name  string
		value values.JSON
	}
	setJSONReturns struct {
		result1 credentials.JSON
		result2 error
	}
	setJSONReturnsOnCall map[int]struct {
		result1 credentials.JSON
		result2 error
	}
	SetValueStub        func(name string, value values.Value) (credentials.Value, error)
	setValueMutex       sync.RWMutex
	setValueArgsForCall []struct {
		name  string
		value values.Value
	}
	setValueReturns struct {
		result1 credentials.Value
		result2 error
	}
	setValueReturnsOnCall map[int]struct {
		result1 credentials.Value
		result2 error
	}
	AddPermissionStub        func(credName string, actor string, ops []string) (*permissions.Permission, error)
	addPermissionMutex       sync.RWMutex
	addPermissionArgsForCall []struct {
		credName string
		actor    string
		ops      []string
	}
	addPermissionReturns struct {
		result1 *permissions.Permission
		result2 error
	}
	addPermissionReturnsOnCall map[int]struct {
		result1 *permissions.Permission
		result2 error
	}
	DeleteStub        func(name string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		name string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredhubClient) GetById(id string) (credentials.Credential, error) {
	fake.getByIdMutex.Lock()
	ret, specificReturn := fake.getByIdReturnsOnCall[len(fake.getByIdArgsForCall)]
	fake.getByIdArgsForCall = append(fake.getByIdArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("GetById", []interface{}{id})
	fake.getByIdMutex.Unlock()
	if fake.GetByIdStub != nil {
		return fake.GetByIdStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByIdReturns.result1, fake.getByIdReturns.result2
}

func (fake *FakeCredhubClient) GetByIdCallCount() int {
	fake.getByIdMutex.RLock()
	defer fake.getByIdMutex.RUnlock()
	return len(fake.getByIdArgsForCall)
}

func (fake *FakeCredhubClient) GetByIdArgsForCall(i int) string {
	fake.getByIdMutex.RLock()
	defer fake.getByIdMutex.RUnlock()
	return fake.getByIdArgsForCall[i].id
}

func (fake *FakeCredhubClient) GetByIdReturns(result1 credentials.Credential, result2 error) {
	fake.GetByIdStub = nil
	fake.getByIdReturns = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) GetByIdReturnsOnCall(i int, result1 credentials.Credential, result2 error) {
	fake.GetByIdStub = nil
	if fake.getByIdReturnsOnCall == nil {
		fake.getByIdReturnsOnCall = make(map[int]struct {
			result1 credentials.Credential
			result2 error
		})
	}
	fake.getByIdReturnsOnCall[i] = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) GetLatestVersion(name string) (credentials.Credential, error) {
	fake.getLatestVersionMutex.Lock()
	ret, specificReturn := fake.getLatestVersionReturnsOnCall[len(fake.getLatestVersionArgsForCall)]
	fake.getLatestVersionArgsForCall = append(fake.getLatestVersionArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetLatestVersion", []interface{}{name})
	fake.getLatestVersionMutex.Unlock()
	if fake.GetLatestVersionStub != nil {
		return fake.GetLatestVersionStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getLatestVersionReturns.result1, fake.getLatestVersionReturns.result2
}

func (fake *FakeCredhubClient) GetLatestVersionCallCount() int {
	fake.getLatestVersionMutex.RLock()
	defer fake.getLatestVersionMutex.RUnlock()
	return len(fake.getLatestVersionArgsForCall)
}

func (fake *FakeCredhubClient) GetLatestVersionArgsForCall(i int) string {
	fake.getLatestVersionMutex.RLock()
	defer fake.getLatestVersionMutex.RUnlock()
	return fake.getLatestVersionArgsForCall[i].name
}

func (fake *FakeCredhubClient) GetLatestVersionReturns(result1 credentials.Credential, result2 error) {
	fake.GetLatestVersionStub = nil
	fake.getLatestVersionReturns = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) GetLatestVersionReturnsOnCall(i int, result1 credentials.Credential, result2 error) {
	fake.GetLatestVersionStub = nil
	if fake.getLatestVersionReturnsOnCall == nil {
		fake.getLatestVersionReturnsOnCall = make(map[int]struct {
			result1 credentials.Credential
			result2 error
		})
	}
	fake.getLatestVersionReturnsOnCall[i] = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) FindByPartialName(partialName string) (credentials.FindResults, error) {
	fake.findByPartialNameMutex.Lock()
	ret, specificReturn := fake.findByPartialNameReturnsOnCall[len(fake.findByPartialNameArgsForCall)]
	fake.findByPartialNameArgsForCall = append(fake.findByPartialNameArgsForCall, struct {
		partialName string
	}{partialName})
	fake.recordInvocation("FindByPartialName", []interface{}{partialName})
	fake.findByPartialNameMutex.Unlock()
	if fake.FindByPartialNameStub != nil {
		return fake.FindByPartialNameStub(partialName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findByPartialNameReturns.result1, fake.findByPartialNameReturns.result2
}

func (fake *FakeCredhubClient) FindByPartialNameCallCount() int {
	fake.findByPartialNameMutex.RLock()
	defer fake.findByPartialNameMutex.RUnlock()
	return len(fake.findByPartialNameArgsForCall)
}

func (fake *FakeCredhubClient) FindByPartialNameArgsForCall(i int) string {
	fake.findByPartialNameMutex.RLock()
	defer fake.findByPartialNameMutex.RUnlock()
	return fake.findByPartialNameArgsForCall[i].partialName
}

func (fake *FakeCredhubClient) FindByPartialNameReturns(result1 credentials.FindResults, result2 error) {
	fake.FindByPartialNameStub = nil
	fake.findByPartialNameReturns = struct {
		result1 credentials.FindResults
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) FindByPartialNameReturnsOnCall(i int, result1 credentials.FindResults, result2 error) {
	fake.FindByPartialNameStub = nil
	if fake.findByPartialNameReturnsOnCall == nil {
		fake.findByPartialNameReturnsOnCall = make(map[int]struct {
			result1 credentials.FindResults
			result2 error
		})
	}
	fake.findByPartialNameReturnsOnCall[i] = struct {
		result1 credentials.FindResults
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetJSON(name string, value values.JSON) (credentials.JSON, error) {
	fake.setJSONMutex.Lock()
	ret, specificReturn := fake.setJSONReturnsOnCall[len(fake.setJSONArgsForCall)]
	fake.setJSONArgsForCall = append(fake.setJSONArgsForCall, struct {
		name  string
		value values.JSON
	}{name, value})
	fake.recordInvocation("SetJSON", []interface{}{name, value})
	fake.setJSONMutex.Unlock()
	if fake.SetJSONStub != nil {
		return fake.SetJSONStub(name, value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setJSONReturns.result1, fake.setJSONReturns.result2
}

func (fake *FakeCredhubClient) SetJSONCallCount() int {
	fake.setJSONMutex.RLock()
	defer fake.setJSONMutex.RUnlock()
	return len(fake.setJSONArgsForCall)
}

func (fake *FakeCredhubClient) SetJSONArgsForCall(i int) (string, values.JSON) {
	fake.setJSONMutex.RLock()
	defer fake.setJSONMutex.RUnlock()
	return fake.setJSONArgsForCall[i].name, fake.setJSONArgsForCall[i].value
}

func (fake *FakeCredhubClient) SetJSONReturns(result1 credentials.JSON, result2 error) {
	fake.SetJSONStub = nil
	fake.setJSONReturns = struct {
		result1 credentials.JSON
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetJSONReturnsOnCall(i int, result1 credentials.JSON, result2 error) {
	fake.SetJSONStub = nil
	if fake.setJSONReturnsOnCall == nil {
		fake.setJSONReturnsOnCall = make(map[int]struct {
			result1 credentials.JSON
			result2 error
		})
	}
	fake.setJSONReturnsOnCall[i] = struct {
		result1 credentials.JSON
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetValue(name string, value values.Value) (credentials.Value, error) {
	fake.setValueMutex.Lock()
	ret, specificReturn := fake.setValueReturnsOnCall[len(fake.setValueArgsForCall)]
	fake.setValueArgsForCall = append(fake.setValueArgsForCall, struct {
		name  string
		value values.Value
	}{name, value})
	fake.recordInvocation("SetValue", []interface{}{name, value})
	fake.setValueMutex.Unlock()
	if fake.SetValueStub != nil {
		return fake.SetValueStub(name, value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setValueReturns.result1, fake.setValueReturns.result2
}

func (fake *FakeCredhubClient) SetValueCallCount() int {
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	return len(fake.setValueArgsForCall)
}

func (fake *FakeCredhubClient) SetValueArgsForCall(i int) (string, values.Value) {
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	return fake.setValueArgsForCall[i].name, fake.setValueArgsForCall[i].value
}

func (fake *FakeCredhubClient) SetValueReturns(result1 credentials.Value, result2 error) {
	fake.SetValueStub = nil
	fake.setValueReturns = struct {
		result1 credentials.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetValueReturnsOnCall(i int, result1 credentials.Value, result2 error) {
	fake.SetValueStub = nil
	if fake.setValueReturnsOnCall == nil {
		fake.setValueReturnsOnCall = make(map[int]struct {
			result1 credentials.Value
			result2 error
		})
	}
	fake.setValueReturnsOnCall[i] = struct {
		result1 credentials.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) AddPermission(credName string, actor string, ops []string) (*permissions.Permission, error) {
	var opsCopy []string
	if ops != nil {
		opsCopy = make([]string, len(ops))
		copy(opsCopy, ops)
	}
	fake.addPermissionMutex.Lock()
	ret, specificReturn := fake.addPermissionReturnsOnCall[len(fake.addPermissionArgsForCall)]
	fake.addPermissionArgsForCall = append(fake.addPermissionArgsForCall, struct {
		credName string
		actor    string
		ops      []string
	}{credName, actor, opsCopy})
	fake.recordInvocation("AddPermission", []interface{}{credName, actor, opsCopy})
	fake.addPermissionMutex.Unlock()
	if fake.AddPermissionStub != nil {
		return fake.AddPermissionStub(credName, actor, ops)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.addPermissionReturns.result1, fake.addPermissionReturns.result2
}

func (fake *FakeCredhubClient) AddPermissionCallCount() int {
	fake.addPermissionMutex.RLock()
	defer fake.addPermissionMutex.RUnlock()
	return len(fake.addPermissionArgsForCall)
}

func (fake *FakeCredhubClient) AddPermissionArgsForCall(i int) (string, string, []string) {
	fake.addPermissionMutex.RLock()
	defer fake.addPermissionMutex.RUnlock()
	return fake.addPermissionArgsForCall[i].credName, fake.addPermissionArgsForCall[i].actor, fake.addPermissionArgsForCall[i].ops
}

func (fake *FakeCredhubClient) AddPermissionReturns(result1 *permissions.Permission, result2 error) {
	fake.AddPermissionStub = nil
	fake.addPermissionReturns = struct {
		result1 *permissions.Permission
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) AddPermissionReturnsOnCall(i int, result1 *permissions.Permission, result2 error) {
	fake.AddPermissionStub = nil
	if fake.addPermissionReturnsOnCall == nil {
		fake.addPermissionReturnsOnCall = make(map[int]struct {
			result1 *permissions.Permission
			result2 error
		})
	}
	fake.addPermissionReturnsOnCall[i] = struct {
		result1 *permissions.Permission
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) Delete(name string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Delete", []interface{}{name})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeCredhubClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCredhubClient) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].name
}

func (fake *FakeCredhubClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredhubClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredhubClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getByIdMutex.RLock()
	defer fake.getByIdMutex.RUnlock()
	fake.getLatestVersionMutex.RLock()
	defer fake.getLatestVersionMutex.RUnlock()
	fake.findByPartialNameMutex.RLock()
	defer fake.findByPartialNameMutex.RUnlock()
	fake.setJSONMutex.RLock()
	defer fake.setJSONMutex.RUnlock()
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	fake.addPermissionMutex.RLock()
	defer fake.addPermissionMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredhubClient) recordInvocation(key string, args []interface{}) {
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

var _ credhub.CredhubClient = new(FakeCredhubClient)
