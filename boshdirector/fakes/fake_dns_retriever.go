// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
)

type FakeDNSRetriever struct {
	LinkProviderIDStub        func(deploymentName, instanceGroupName, providerName string) (string, error)
	linkProviderIDMutex       sync.RWMutex
	linkProviderIDArgsForCall []struct {
		deploymentName    string
		instanceGroupName string
		providerName      string
	}
	linkProviderIDReturns struct {
		result1 string
		result2 error
	}
	linkProviderIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CreateLinkConsumerStub        func(providerID string) (string, error)
	createLinkConsumerMutex       sync.RWMutex
	createLinkConsumerArgsForCall []struct {
		providerID string
	}
	createLinkConsumerReturns struct {
		result1 string
		result2 error
	}
	createLinkConsumerReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetLinkAddressStub        func(consumerLinkID string) (string, error)
	getLinkAddressMutex       sync.RWMutex
	getLinkAddressArgsForCall []struct {
		consumerLinkID string
	}
	getLinkAddressReturns struct {
		result1 string
		result2 error
	}
	getLinkAddressReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DeleteLinkConsumerStub        func(consumerID string) error
	deleteLinkConsumerMutex       sync.RWMutex
	deleteLinkConsumerArgsForCall []struct {
		consumerID string
	}
	deleteLinkConsumerReturns struct {
		result1 error
	}
	deleteLinkConsumerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDNSRetriever) LinkProviderID(deploymentName string, instanceGroupName string, providerName string) (string, error) {
	fake.linkProviderIDMutex.Lock()
	ret, specificReturn := fake.linkProviderIDReturnsOnCall[len(fake.linkProviderIDArgsForCall)]
	fake.linkProviderIDArgsForCall = append(fake.linkProviderIDArgsForCall, struct {
		deploymentName    string
		instanceGroupName string
		providerName      string
	}{deploymentName, instanceGroupName, providerName})
	fake.recordInvocation("LinkProviderID", []interface{}{deploymentName, instanceGroupName, providerName})
	fake.linkProviderIDMutex.Unlock()
	if fake.LinkProviderIDStub != nil {
		return fake.LinkProviderIDStub(deploymentName, instanceGroupName, providerName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.linkProviderIDReturns.result1, fake.linkProviderIDReturns.result2
}

func (fake *FakeDNSRetriever) LinkProviderIDCallCount() int {
	fake.linkProviderIDMutex.RLock()
	defer fake.linkProviderIDMutex.RUnlock()
	return len(fake.linkProviderIDArgsForCall)
}

func (fake *FakeDNSRetriever) LinkProviderIDArgsForCall(i int) (string, string, string) {
	fake.linkProviderIDMutex.RLock()
	defer fake.linkProviderIDMutex.RUnlock()
	return fake.linkProviderIDArgsForCall[i].deploymentName, fake.linkProviderIDArgsForCall[i].instanceGroupName, fake.linkProviderIDArgsForCall[i].providerName
}

func (fake *FakeDNSRetriever) LinkProviderIDReturns(result1 string, result2 error) {
	fake.LinkProviderIDStub = nil
	fake.linkProviderIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSRetriever) LinkProviderIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.LinkProviderIDStub = nil
	if fake.linkProviderIDReturnsOnCall == nil {
		fake.linkProviderIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.linkProviderIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSRetriever) CreateLinkConsumer(providerID string) (string, error) {
	fake.createLinkConsumerMutex.Lock()
	ret, specificReturn := fake.createLinkConsumerReturnsOnCall[len(fake.createLinkConsumerArgsForCall)]
	fake.createLinkConsumerArgsForCall = append(fake.createLinkConsumerArgsForCall, struct {
		providerID string
	}{providerID})
	fake.recordInvocation("CreateLinkConsumer", []interface{}{providerID})
	fake.createLinkConsumerMutex.Unlock()
	if fake.CreateLinkConsumerStub != nil {
		return fake.CreateLinkConsumerStub(providerID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createLinkConsumerReturns.result1, fake.createLinkConsumerReturns.result2
}

func (fake *FakeDNSRetriever) CreateLinkConsumerCallCount() int {
	fake.createLinkConsumerMutex.RLock()
	defer fake.createLinkConsumerMutex.RUnlock()
	return len(fake.createLinkConsumerArgsForCall)
}

func (fake *FakeDNSRetriever) CreateLinkConsumerArgsForCall(i int) string {
	fake.createLinkConsumerMutex.RLock()
	defer fake.createLinkConsumerMutex.RUnlock()
	return fake.createLinkConsumerArgsForCall[i].providerID
}

func (fake *FakeDNSRetriever) CreateLinkConsumerReturns(result1 string, result2 error) {
	fake.CreateLinkConsumerStub = nil
	fake.createLinkConsumerReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSRetriever) CreateLinkConsumerReturnsOnCall(i int, result1 string, result2 error) {
	fake.CreateLinkConsumerStub = nil
	if fake.createLinkConsumerReturnsOnCall == nil {
		fake.createLinkConsumerReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createLinkConsumerReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSRetriever) GetLinkAddress(consumerLinkID string) (string, error) {
	fake.getLinkAddressMutex.Lock()
	ret, specificReturn := fake.getLinkAddressReturnsOnCall[len(fake.getLinkAddressArgsForCall)]
	fake.getLinkAddressArgsForCall = append(fake.getLinkAddressArgsForCall, struct {
		consumerLinkID string
	}{consumerLinkID})
	fake.recordInvocation("GetLinkAddress", []interface{}{consumerLinkID})
	fake.getLinkAddressMutex.Unlock()
	if fake.GetLinkAddressStub != nil {
		return fake.GetLinkAddressStub(consumerLinkID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getLinkAddressReturns.result1, fake.getLinkAddressReturns.result2
}

func (fake *FakeDNSRetriever) GetLinkAddressCallCount() int {
	fake.getLinkAddressMutex.RLock()
	defer fake.getLinkAddressMutex.RUnlock()
	return len(fake.getLinkAddressArgsForCall)
}

func (fake *FakeDNSRetriever) GetLinkAddressArgsForCall(i int) string {
	fake.getLinkAddressMutex.RLock()
	defer fake.getLinkAddressMutex.RUnlock()
	return fake.getLinkAddressArgsForCall[i].consumerLinkID
}

func (fake *FakeDNSRetriever) GetLinkAddressReturns(result1 string, result2 error) {
	fake.GetLinkAddressStub = nil
	fake.getLinkAddressReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSRetriever) GetLinkAddressReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetLinkAddressStub = nil
	if fake.getLinkAddressReturnsOnCall == nil {
		fake.getLinkAddressReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getLinkAddressReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDNSRetriever) DeleteLinkConsumer(consumerID string) error {
	fake.deleteLinkConsumerMutex.Lock()
	ret, specificReturn := fake.deleteLinkConsumerReturnsOnCall[len(fake.deleteLinkConsumerArgsForCall)]
	fake.deleteLinkConsumerArgsForCall = append(fake.deleteLinkConsumerArgsForCall, struct {
		consumerID string
	}{consumerID})
	fake.recordInvocation("DeleteLinkConsumer", []interface{}{consumerID})
	fake.deleteLinkConsumerMutex.Unlock()
	if fake.DeleteLinkConsumerStub != nil {
		return fake.DeleteLinkConsumerStub(consumerID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteLinkConsumerReturns.result1
}

func (fake *FakeDNSRetriever) DeleteLinkConsumerCallCount() int {
	fake.deleteLinkConsumerMutex.RLock()
	defer fake.deleteLinkConsumerMutex.RUnlock()
	return len(fake.deleteLinkConsumerArgsForCall)
}

func (fake *FakeDNSRetriever) DeleteLinkConsumerArgsForCall(i int) string {
	fake.deleteLinkConsumerMutex.RLock()
	defer fake.deleteLinkConsumerMutex.RUnlock()
	return fake.deleteLinkConsumerArgsForCall[i].consumerID
}

func (fake *FakeDNSRetriever) DeleteLinkConsumerReturns(result1 error) {
	fake.DeleteLinkConsumerStub = nil
	fake.deleteLinkConsumerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSRetriever) DeleteLinkConsumerReturnsOnCall(i int, result1 error) {
	fake.DeleteLinkConsumerStub = nil
	if fake.deleteLinkConsumerReturnsOnCall == nil {
		fake.deleteLinkConsumerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteLinkConsumerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSRetriever) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.linkProviderIDMutex.RLock()
	defer fake.linkProviderIDMutex.RUnlock()
	fake.createLinkConsumerMutex.RLock()
	defer fake.createLinkConsumerMutex.RUnlock()
	fake.getLinkAddressMutex.RLock()
	defer fake.getLinkAddressMutex.RUnlock()
	fake.deleteLinkConsumerMutex.RLock()
	defer fake.deleteLinkConsumerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDNSRetriever) recordInvocation(key string, args []interface{}) {
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

var _ boshdirector.DNSRetriever = new(FakeDNSRetriever)