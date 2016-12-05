// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker"
)

type FakeDBContainerFactory struct {
	FindOrCreateBuildContainerStub        func(worker *dbng.Worker, build *dbng.Build, planID atc.PlanID, meta dbng.ContainerMetadata) (*dbng.CreatingContainer, error)
	findOrCreateBuildContainerMutex       sync.RWMutex
	findOrCreateBuildContainerArgsForCall []struct {
		worker *dbng.Worker
		build  *dbng.Build
		planID atc.PlanID
		meta   dbng.ContainerMetadata
	}
	findOrCreateBuildContainerReturns struct {
		result1 *dbng.CreatingContainer
		result2 error
	}
	CreateResourceGetContainerStub        func(worker *dbng.Worker, resourceCache *dbng.UsedResourceCache, stepName string) (*dbng.CreatingContainer, error)
	createResourceGetContainerMutex       sync.RWMutex
	createResourceGetContainerArgsForCall []struct {
		worker        *dbng.Worker
		resourceCache *dbng.UsedResourceCache
		stepName      string
	}
	createResourceGetContainerReturns struct {
		result1 *dbng.CreatingContainer
		result2 error
	}
	FindOrCreateResourceCheckContainerStub        func(worker *dbng.Worker, resourceConfig *dbng.UsedResourceConfig, stepName string) (*dbng.CreatingContainer, error)
	findOrCreateResourceCheckContainerMutex       sync.RWMutex
	findOrCreateResourceCheckContainerArgsForCall []struct {
		worker         *dbng.Worker
		resourceConfig *dbng.UsedResourceConfig
		stepName       string
	}
	findOrCreateResourceCheckContainerReturns struct {
		result1 *dbng.CreatingContainer
		result2 error
	}
	FindContainerStub        func(handle string) (*dbng.CreatedContainer, bool, error)
	findContainerMutex       sync.RWMutex
	findContainerArgsForCall []struct {
		handle string
	}
	findContainerReturns struct {
		result1 *dbng.CreatedContainer
		result2 bool
		result3 error
	}
	ContainerCreatedStub        func(*dbng.CreatingContainer) (*dbng.CreatedContainer, error)
	containerCreatedMutex       sync.RWMutex
	containerCreatedArgsForCall []struct {
		arg1 *dbng.CreatingContainer
	}
	containerCreatedReturns struct {
		result1 *dbng.CreatedContainer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBContainerFactory) FindOrCreateBuildContainer(worker *dbng.Worker, build *dbng.Build, planID atc.PlanID, meta dbng.ContainerMetadata) (*dbng.CreatingContainer, error) {
	fake.findOrCreateBuildContainerMutex.Lock()
	fake.findOrCreateBuildContainerArgsForCall = append(fake.findOrCreateBuildContainerArgsForCall, struct {
		worker *dbng.Worker
		build  *dbng.Build
		planID atc.PlanID
		meta   dbng.ContainerMetadata
	}{worker, build, planID, meta})
	fake.recordInvocation("FindOrCreateBuildContainer", []interface{}{worker, build, planID, meta})
	fake.findOrCreateBuildContainerMutex.Unlock()
	if fake.FindOrCreateBuildContainerStub != nil {
		return fake.FindOrCreateBuildContainerStub(worker, build, planID, meta)
	} else {
		return fake.findOrCreateBuildContainerReturns.result1, fake.findOrCreateBuildContainerReturns.result2
	}
}

func (fake *FakeDBContainerFactory) FindOrCreateBuildContainerCallCount() int {
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	return len(fake.findOrCreateBuildContainerArgsForCall)
}

func (fake *FakeDBContainerFactory) FindOrCreateBuildContainerArgsForCall(i int) (*dbng.Worker, *dbng.Build, atc.PlanID, dbng.ContainerMetadata) {
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	return fake.findOrCreateBuildContainerArgsForCall[i].worker, fake.findOrCreateBuildContainerArgsForCall[i].build, fake.findOrCreateBuildContainerArgsForCall[i].planID, fake.findOrCreateBuildContainerArgsForCall[i].meta
}

func (fake *FakeDBContainerFactory) FindOrCreateBuildContainerReturns(result1 *dbng.CreatingContainer, result2 error) {
	fake.FindOrCreateBuildContainerStub = nil
	fake.findOrCreateBuildContainerReturns = struct {
		result1 *dbng.CreatingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeDBContainerFactory) CreateResourceGetContainer(worker *dbng.Worker, resourceCache *dbng.UsedResourceCache, stepName string) (*dbng.CreatingContainer, error) {
	fake.createResourceGetContainerMutex.Lock()
	fake.createResourceGetContainerArgsForCall = append(fake.createResourceGetContainerArgsForCall, struct {
		worker        *dbng.Worker
		resourceCache *dbng.UsedResourceCache
		stepName      string
	}{worker, resourceCache, stepName})
	fake.recordInvocation("CreateResourceGetContainer", []interface{}{worker, resourceCache, stepName})
	fake.createResourceGetContainerMutex.Unlock()
	if fake.CreateResourceGetContainerStub != nil {
		return fake.CreateResourceGetContainerStub(worker, resourceCache, stepName)
	} else {
		return fake.createResourceGetContainerReturns.result1, fake.createResourceGetContainerReturns.result2
	}
}

func (fake *FakeDBContainerFactory) CreateResourceGetContainerCallCount() int {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return len(fake.createResourceGetContainerArgsForCall)
}

func (fake *FakeDBContainerFactory) CreateResourceGetContainerArgsForCall(i int) (*dbng.Worker, *dbng.UsedResourceCache, string) {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return fake.createResourceGetContainerArgsForCall[i].worker, fake.createResourceGetContainerArgsForCall[i].resourceCache, fake.createResourceGetContainerArgsForCall[i].stepName
}

func (fake *FakeDBContainerFactory) CreateResourceGetContainerReturns(result1 *dbng.CreatingContainer, result2 error) {
	fake.CreateResourceGetContainerStub = nil
	fake.createResourceGetContainerReturns = struct {
		result1 *dbng.CreatingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeDBContainerFactory) FindOrCreateResourceCheckContainer(worker *dbng.Worker, resourceConfig *dbng.UsedResourceConfig, stepName string) (*dbng.CreatingContainer, error) {
	fake.findOrCreateResourceCheckContainerMutex.Lock()
	fake.findOrCreateResourceCheckContainerArgsForCall = append(fake.findOrCreateResourceCheckContainerArgsForCall, struct {
		worker         *dbng.Worker
		resourceConfig *dbng.UsedResourceConfig
		stepName       string
	}{worker, resourceConfig, stepName})
	fake.recordInvocation("FindOrCreateResourceCheckContainer", []interface{}{worker, resourceConfig, stepName})
	fake.findOrCreateResourceCheckContainerMutex.Unlock()
	if fake.FindOrCreateResourceCheckContainerStub != nil {
		return fake.FindOrCreateResourceCheckContainerStub(worker, resourceConfig, stepName)
	} else {
		return fake.findOrCreateResourceCheckContainerReturns.result1, fake.findOrCreateResourceCheckContainerReturns.result2
	}
}

func (fake *FakeDBContainerFactory) FindOrCreateResourceCheckContainerCallCount() int {
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	return len(fake.findOrCreateResourceCheckContainerArgsForCall)
}

func (fake *FakeDBContainerFactory) FindOrCreateResourceCheckContainerArgsForCall(i int) (*dbng.Worker, *dbng.UsedResourceConfig, string) {
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	return fake.findOrCreateResourceCheckContainerArgsForCall[i].worker, fake.findOrCreateResourceCheckContainerArgsForCall[i].resourceConfig, fake.findOrCreateResourceCheckContainerArgsForCall[i].stepName
}

func (fake *FakeDBContainerFactory) FindOrCreateResourceCheckContainerReturns(result1 *dbng.CreatingContainer, result2 error) {
	fake.FindOrCreateResourceCheckContainerStub = nil
	fake.findOrCreateResourceCheckContainerReturns = struct {
		result1 *dbng.CreatingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeDBContainerFactory) FindContainer(handle string) (*dbng.CreatedContainer, bool, error) {
	fake.findContainerMutex.Lock()
	fake.findContainerArgsForCall = append(fake.findContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("FindContainer", []interface{}{handle})
	fake.findContainerMutex.Unlock()
	if fake.FindContainerStub != nil {
		return fake.FindContainerStub(handle)
	} else {
		return fake.findContainerReturns.result1, fake.findContainerReturns.result2, fake.findContainerReturns.result3
	}
}

func (fake *FakeDBContainerFactory) FindContainerCallCount() int {
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	return len(fake.findContainerArgsForCall)
}

func (fake *FakeDBContainerFactory) FindContainerArgsForCall(i int) string {
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	return fake.findContainerArgsForCall[i].handle
}

func (fake *FakeDBContainerFactory) FindContainerReturns(result1 *dbng.CreatedContainer, result2 bool, result3 error) {
	fake.FindContainerStub = nil
	fake.findContainerReturns = struct {
		result1 *dbng.CreatedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDBContainerFactory) ContainerCreated(arg1 *dbng.CreatingContainer) (*dbng.CreatedContainer, error) {
	fake.containerCreatedMutex.Lock()
	fake.containerCreatedArgsForCall = append(fake.containerCreatedArgsForCall, struct {
		arg1 *dbng.CreatingContainer
	}{arg1})
	fake.recordInvocation("ContainerCreated", []interface{}{arg1})
	fake.containerCreatedMutex.Unlock()
	if fake.ContainerCreatedStub != nil {
		return fake.ContainerCreatedStub(arg1)
	} else {
		return fake.containerCreatedReturns.result1, fake.containerCreatedReturns.result2
	}
}

func (fake *FakeDBContainerFactory) ContainerCreatedCallCount() int {
	fake.containerCreatedMutex.RLock()
	defer fake.containerCreatedMutex.RUnlock()
	return len(fake.containerCreatedArgsForCall)
}

func (fake *FakeDBContainerFactory) ContainerCreatedArgsForCall(i int) *dbng.CreatingContainer {
	fake.containerCreatedMutex.RLock()
	defer fake.containerCreatedMutex.RUnlock()
	return fake.containerCreatedArgsForCall[i].arg1
}

func (fake *FakeDBContainerFactory) ContainerCreatedReturns(result1 *dbng.CreatedContainer, result2 error) {
	fake.ContainerCreatedStub = nil
	fake.containerCreatedReturns = struct {
		result1 *dbng.CreatedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeDBContainerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	fake.findContainerMutex.RLock()
	defer fake.findContainerMutex.RUnlock()
	fake.containerCreatedMutex.RLock()
	defer fake.containerCreatedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDBContainerFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.DBContainerFactory = new(FakeDBContainerFactory)
