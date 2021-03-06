// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"encoding/json"
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ConfigureDirectorService struct {
	UpdateStagedDirectorAvailabilityZonesStub        func(api.AvailabilityZoneInput) error
	updateStagedDirectorAvailabilityZonesMutex       sync.RWMutex
	updateStagedDirectorAvailabilityZonesArgsForCall []struct {
		arg1 api.AvailabilityZoneInput
	}
	updateStagedDirectorAvailabilityZonesReturns struct {
		result1 error
	}
	updateStagedDirectorAvailabilityZonesReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStagedDirectorNetworksStub        func(json.RawMessage) error
	updateStagedDirectorNetworksMutex       sync.RWMutex
	updateStagedDirectorNetworksArgsForCall []struct {
		arg1 json.RawMessage
	}
	updateStagedDirectorNetworksReturns struct {
		result1 error
	}
	updateStagedDirectorNetworksReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStagedDirectorNetworkAndAZStub        func(api.NetworkAndAZConfiguration) error
	updateStagedDirectorNetworkAndAZMutex       sync.RWMutex
	updateStagedDirectorNetworkAndAZArgsForCall []struct {
		arg1 api.NetworkAndAZConfiguration
	}
	updateStagedDirectorNetworkAndAZReturns struct {
		result1 error
	}
	updateStagedDirectorNetworkAndAZReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStagedDirectorPropertiesStub        func(api.DirectorProperties) error
	updateStagedDirectorPropertiesMutex       sync.RWMutex
	updateStagedDirectorPropertiesArgsForCall []struct {
		arg1 api.DirectorProperties
	}
	updateStagedDirectorPropertiesReturns struct {
		result1 error
	}
	updateStagedDirectorPropertiesReturnsOnCall map[int]struct {
		result1 error
	}
	ListStagedProductJobsStub        func(string) (map[string]string, error)
	listStagedProductJobsMutex       sync.RWMutex
	listStagedProductJobsArgsForCall []struct {
		arg1 string
	}
	listStagedProductJobsReturns struct {
		result1 map[string]string
		result2 error
	}
	listStagedProductJobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetStagedProductJobResourceConfigStub        func(string, string) (api.JobProperties, error)
	getStagedProductJobResourceConfigMutex       sync.RWMutex
	getStagedProductJobResourceConfigArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getStagedProductJobResourceConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	UpdateStagedProductJobResourceConfigStub        func(string, string, api.JobProperties) error
	updateStagedProductJobResourceConfigMutex       sync.RWMutex
	updateStagedProductJobResourceConfigArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 api.JobProperties
	}
	updateStagedProductJobResourceConfigReturns struct {
		result1 error
	}
	updateStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 error
	}
	GetStagedProductByNameStub        func(name string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		name string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	GetStagedProductManifestStub        func(guid string) (manifest string, err error)
	getStagedProductManifestMutex       sync.RWMutex
	getStagedProductManifestArgsForCall []struct {
		guid string
	}
	getStagedProductManifestReturns struct {
		result1 string
		result2 error
	}
	getStagedProductManifestReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorAvailabilityZones(arg1 api.AvailabilityZoneInput) error {
	fake.updateStagedDirectorAvailabilityZonesMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorAvailabilityZonesReturnsOnCall[len(fake.updateStagedDirectorAvailabilityZonesArgsForCall)]
	fake.updateStagedDirectorAvailabilityZonesArgsForCall = append(fake.updateStagedDirectorAvailabilityZonesArgsForCall, struct {
		arg1 api.AvailabilityZoneInput
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorAvailabilityZones", []interface{}{arg1})
	fake.updateStagedDirectorAvailabilityZonesMutex.Unlock()
	if fake.UpdateStagedDirectorAvailabilityZonesStub != nil {
		return fake.UpdateStagedDirectorAvailabilityZonesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorAvailabilityZonesReturns.result1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorAvailabilityZonesCallCount() int {
	fake.updateStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.updateStagedDirectorAvailabilityZonesMutex.RUnlock()
	return len(fake.updateStagedDirectorAvailabilityZonesArgsForCall)
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorAvailabilityZonesArgsForCall(i int) api.AvailabilityZoneInput {
	fake.updateStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.updateStagedDirectorAvailabilityZonesMutex.RUnlock()
	return fake.updateStagedDirectorAvailabilityZonesArgsForCall[i].arg1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorAvailabilityZonesReturns(result1 error) {
	fake.UpdateStagedDirectorAvailabilityZonesStub = nil
	fake.updateStagedDirectorAvailabilityZonesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorAvailabilityZonesReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorAvailabilityZonesStub = nil
	if fake.updateStagedDirectorAvailabilityZonesReturnsOnCall == nil {
		fake.updateStagedDirectorAvailabilityZonesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorAvailabilityZonesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworks(arg1 json.RawMessage) error {
	fake.updateStagedDirectorNetworksMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorNetworksReturnsOnCall[len(fake.updateStagedDirectorNetworksArgsForCall)]
	fake.updateStagedDirectorNetworksArgsForCall = append(fake.updateStagedDirectorNetworksArgsForCall, struct {
		arg1 json.RawMessage
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorNetworks", []interface{}{arg1})
	fake.updateStagedDirectorNetworksMutex.Unlock()
	if fake.UpdateStagedDirectorNetworksStub != nil {
		return fake.UpdateStagedDirectorNetworksStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorNetworksReturns.result1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworksCallCount() int {
	fake.updateStagedDirectorNetworksMutex.RLock()
	defer fake.updateStagedDirectorNetworksMutex.RUnlock()
	return len(fake.updateStagedDirectorNetworksArgsForCall)
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworksArgsForCall(i int) json.RawMessage {
	fake.updateStagedDirectorNetworksMutex.RLock()
	defer fake.updateStagedDirectorNetworksMutex.RUnlock()
	return fake.updateStagedDirectorNetworksArgsForCall[i].arg1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworksReturns(result1 error) {
	fake.UpdateStagedDirectorNetworksStub = nil
	fake.updateStagedDirectorNetworksReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworksReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorNetworksStub = nil
	if fake.updateStagedDirectorNetworksReturnsOnCall == nil {
		fake.updateStagedDirectorNetworksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorNetworksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworkAndAZ(arg1 api.NetworkAndAZConfiguration) error {
	fake.updateStagedDirectorNetworkAndAZMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorNetworkAndAZReturnsOnCall[len(fake.updateStagedDirectorNetworkAndAZArgsForCall)]
	fake.updateStagedDirectorNetworkAndAZArgsForCall = append(fake.updateStagedDirectorNetworkAndAZArgsForCall, struct {
		arg1 api.NetworkAndAZConfiguration
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorNetworkAndAZ", []interface{}{arg1})
	fake.updateStagedDirectorNetworkAndAZMutex.Unlock()
	if fake.UpdateStagedDirectorNetworkAndAZStub != nil {
		return fake.UpdateStagedDirectorNetworkAndAZStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorNetworkAndAZReturns.result1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworkAndAZCallCount() int {
	fake.updateStagedDirectorNetworkAndAZMutex.RLock()
	defer fake.updateStagedDirectorNetworkAndAZMutex.RUnlock()
	return len(fake.updateStagedDirectorNetworkAndAZArgsForCall)
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworkAndAZArgsForCall(i int) api.NetworkAndAZConfiguration {
	fake.updateStagedDirectorNetworkAndAZMutex.RLock()
	defer fake.updateStagedDirectorNetworkAndAZMutex.RUnlock()
	return fake.updateStagedDirectorNetworkAndAZArgsForCall[i].arg1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworkAndAZReturns(result1 error) {
	fake.UpdateStagedDirectorNetworkAndAZStub = nil
	fake.updateStagedDirectorNetworkAndAZReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorNetworkAndAZReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorNetworkAndAZStub = nil
	if fake.updateStagedDirectorNetworkAndAZReturnsOnCall == nil {
		fake.updateStagedDirectorNetworkAndAZReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorNetworkAndAZReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorProperties(arg1 api.DirectorProperties) error {
	fake.updateStagedDirectorPropertiesMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorPropertiesReturnsOnCall[len(fake.updateStagedDirectorPropertiesArgsForCall)]
	fake.updateStagedDirectorPropertiesArgsForCall = append(fake.updateStagedDirectorPropertiesArgsForCall, struct {
		arg1 api.DirectorProperties
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorProperties", []interface{}{arg1})
	fake.updateStagedDirectorPropertiesMutex.Unlock()
	if fake.UpdateStagedDirectorPropertiesStub != nil {
		return fake.UpdateStagedDirectorPropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorPropertiesReturns.result1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorPropertiesCallCount() int {
	fake.updateStagedDirectorPropertiesMutex.RLock()
	defer fake.updateStagedDirectorPropertiesMutex.RUnlock()
	return len(fake.updateStagedDirectorPropertiesArgsForCall)
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorPropertiesArgsForCall(i int) api.DirectorProperties {
	fake.updateStagedDirectorPropertiesMutex.RLock()
	defer fake.updateStagedDirectorPropertiesMutex.RUnlock()
	return fake.updateStagedDirectorPropertiesArgsForCall[i].arg1
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorPropertiesReturns(result1 error) {
	fake.UpdateStagedDirectorPropertiesStub = nil
	fake.updateStagedDirectorPropertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedDirectorPropertiesReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorPropertiesStub = nil
	if fake.updateStagedDirectorPropertiesReturnsOnCall == nil {
		fake.updateStagedDirectorPropertiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorPropertiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) ListStagedProductJobs(arg1 string) (map[string]string, error) {
	fake.listStagedProductJobsMutex.Lock()
	ret, specificReturn := fake.listStagedProductJobsReturnsOnCall[len(fake.listStagedProductJobsArgsForCall)]
	fake.listStagedProductJobsArgsForCall = append(fake.listStagedProductJobsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListStagedProductJobs", []interface{}{arg1})
	fake.listStagedProductJobsMutex.Unlock()
	if fake.ListStagedProductJobsStub != nil {
		return fake.ListStagedProductJobsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedProductJobsReturns.result1, fake.listStagedProductJobsReturns.result2
}

func (fake *ConfigureDirectorService) ListStagedProductJobsCallCount() int {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return len(fake.listStagedProductJobsArgsForCall)
}

func (fake *ConfigureDirectorService) ListStagedProductJobsArgsForCall(i int) string {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return fake.listStagedProductJobsArgsForCall[i].arg1
}

func (fake *ConfigureDirectorService) ListStagedProductJobsReturns(result1 map[string]string, result2 error) {
	fake.ListStagedProductJobsStub = nil
	fake.listStagedProductJobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) ListStagedProductJobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ListStagedProductJobsStub = nil
	if fake.listStagedProductJobsReturnsOnCall == nil {
		fake.listStagedProductJobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.listStagedProductJobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) GetStagedProductJobResourceConfig(arg1 string, arg2 string) (api.JobProperties, error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.getStagedProductJobResourceConfigReturnsOnCall[len(fake.getStagedProductJobResourceConfigArgsForCall)]
	fake.getStagedProductJobResourceConfigArgsForCall = append(fake.getStagedProductJobResourceConfigArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetStagedProductJobResourceConfig", []interface{}{arg1, arg2})
	fake.getStagedProductJobResourceConfigMutex.Unlock()
	if fake.GetStagedProductJobResourceConfigStub != nil {
		return fake.GetStagedProductJobResourceConfigStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductJobResourceConfigReturns.result1, fake.getStagedProductJobResourceConfigReturns.result2
}

func (fake *ConfigureDirectorService) GetStagedProductJobResourceConfigCallCount() int {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.getStagedProductJobResourceConfigArgsForCall)
}

func (fake *ConfigureDirectorService) GetStagedProductJobResourceConfigArgsForCall(i int) (string, string) {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return fake.getStagedProductJobResourceConfigArgsForCall[i].arg1, fake.getStagedProductJobResourceConfigArgsForCall[i].arg2
}

func (fake *ConfigureDirectorService) GetStagedProductJobResourceConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetStagedProductJobResourceConfigStub = nil
	fake.getStagedProductJobResourceConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) GetStagedProductJobResourceConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.GetStagedProductJobResourceConfigStub = nil
	if fake.getStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.getStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) UpdateStagedProductJobResourceConfig(arg1 string, arg2 string, arg3 api.JobProperties) error {
	fake.updateStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.updateStagedProductJobResourceConfigReturnsOnCall[len(fake.updateStagedProductJobResourceConfigArgsForCall)]
	fake.updateStagedProductJobResourceConfigArgsForCall = append(fake.updateStagedProductJobResourceConfigArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 api.JobProperties
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateStagedProductJobResourceConfig", []interface{}{arg1, arg2, arg3})
	fake.updateStagedProductJobResourceConfigMutex.Unlock()
	if fake.UpdateStagedProductJobResourceConfigStub != nil {
		return fake.UpdateStagedProductJobResourceConfigStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedProductJobResourceConfigReturns.result1
}

func (fake *ConfigureDirectorService) UpdateStagedProductJobResourceConfigCallCount() int {
	fake.updateStagedProductJobResourceConfigMutex.RLock()
	defer fake.updateStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.updateStagedProductJobResourceConfigArgsForCall)
}

func (fake *ConfigureDirectorService) UpdateStagedProductJobResourceConfigArgsForCall(i int) (string, string, api.JobProperties) {
	fake.updateStagedProductJobResourceConfigMutex.RLock()
	defer fake.updateStagedProductJobResourceConfigMutex.RUnlock()
	return fake.updateStagedProductJobResourceConfigArgsForCall[i].arg1, fake.updateStagedProductJobResourceConfigArgsForCall[i].arg2, fake.updateStagedProductJobResourceConfigArgsForCall[i].arg3
}

func (fake *ConfigureDirectorService) UpdateStagedProductJobResourceConfigReturns(result1 error) {
	fake.UpdateStagedProductJobResourceConfigStub = nil
	fake.updateStagedProductJobResourceConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) UpdateStagedProductJobResourceConfigReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedProductJobResourceConfigStub = nil
	if fake.updateStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.updateStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureDirectorService) GetStagedProductByName(name string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetStagedProductByName", []interface{}{name})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductByNameReturns.result1, fake.getStagedProductByNameReturns.result2
}

func (fake *ConfigureDirectorService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *ConfigureDirectorService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return fake.getStagedProductByNameArgsForCall[i].name
}

func (fake *ConfigureDirectorService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) GetStagedProductManifest(guid string) (manifest string, err error) {
	fake.getStagedProductManifestMutex.Lock()
	ret, specificReturn := fake.getStagedProductManifestReturnsOnCall[len(fake.getStagedProductManifestArgsForCall)]
	fake.getStagedProductManifestArgsForCall = append(fake.getStagedProductManifestArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("GetStagedProductManifest", []interface{}{guid})
	fake.getStagedProductManifestMutex.Unlock()
	if fake.GetStagedProductManifestStub != nil {
		return fake.GetStagedProductManifestStub(guid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductManifestReturns.result1, fake.getStagedProductManifestReturns.result2
}

func (fake *ConfigureDirectorService) GetStagedProductManifestCallCount() int {
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	return len(fake.getStagedProductManifestArgsForCall)
}

func (fake *ConfigureDirectorService) GetStagedProductManifestArgsForCall(i int) string {
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	return fake.getStagedProductManifestArgsForCall[i].guid
}

func (fake *ConfigureDirectorService) GetStagedProductManifestReturns(result1 string, result2 error) {
	fake.GetStagedProductManifestStub = nil
	fake.getStagedProductManifestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) GetStagedProductManifestReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetStagedProductManifestStub = nil
	if fake.getStagedProductManifestReturnsOnCall == nil {
		fake.getStagedProductManifestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStagedProductManifestReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ConfigureDirectorService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.updateStagedDirectorAvailabilityZonesMutex.RUnlock()
	fake.updateStagedDirectorNetworksMutex.RLock()
	defer fake.updateStagedDirectorNetworksMutex.RUnlock()
	fake.updateStagedDirectorNetworkAndAZMutex.RLock()
	defer fake.updateStagedDirectorNetworkAndAZMutex.RUnlock()
	fake.updateStagedDirectorPropertiesMutex.RLock()
	defer fake.updateStagedDirectorPropertiesMutex.RUnlock()
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	fake.updateStagedProductJobResourceConfigMutex.RLock()
	defer fake.updateStagedProductJobResourceConfigMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigureDirectorService) recordInvocation(key string, args []interface{}) {
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
