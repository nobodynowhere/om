// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type Executor struct {
	GetCommandHelpStub        func(string) ([]byte, error)
	getCommandHelpMutex       sync.RWMutex
	getCommandHelpArgsForCall []struct {
		arg1 string
	}
	getCommandHelpReturns struct {
		result1 []byte
		result2 error
	}
	getCommandHelpReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetCommandNamesAndDescriptionsStub        func() (map[string]string, error)
	getCommandNamesAndDescriptionsMutex       sync.RWMutex
	getCommandNamesAndDescriptionsArgsForCall []struct {
	}
	getCommandNamesAndDescriptionsReturns struct {
		result1 map[string]string
		result2 error
	}
	getCommandNamesAndDescriptionsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetDescriptionStub        func(string) (string, error)
	getDescriptionMutex       sync.RWMutex
	getDescriptionArgsForCall []struct {
		arg1 string
	}
	getDescriptionReturns struct {
		result1 string
		result2 error
	}
	getDescriptionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunOmCommandStub        func(...string) ([]byte, error)
	runOmCommandMutex       sync.RWMutex
	runOmCommandArgsForCall []struct {
		arg1 []string
	}
	runOmCommandReturns struct {
		result1 []byte
		result2 error
	}
	runOmCommandReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Executor) GetCommandHelp(arg1 string) ([]byte, error) {
	fake.getCommandHelpMutex.Lock()
	ret, specificReturn := fake.getCommandHelpReturnsOnCall[len(fake.getCommandHelpArgsForCall)]
	fake.getCommandHelpArgsForCall = append(fake.getCommandHelpArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetCommandHelp", []interface{}{arg1})
	fake.getCommandHelpMutex.Unlock()
	if fake.GetCommandHelpStub != nil {
		return fake.GetCommandHelpStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCommandHelpReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Executor) GetCommandHelpCallCount() int {
	fake.getCommandHelpMutex.RLock()
	defer fake.getCommandHelpMutex.RUnlock()
	return len(fake.getCommandHelpArgsForCall)
}

func (fake *Executor) GetCommandHelpCalls(stub func(string) ([]byte, error)) {
	fake.getCommandHelpMutex.Lock()
	defer fake.getCommandHelpMutex.Unlock()
	fake.GetCommandHelpStub = stub
}

func (fake *Executor) GetCommandHelpArgsForCall(i int) string {
	fake.getCommandHelpMutex.RLock()
	defer fake.getCommandHelpMutex.RUnlock()
	argsForCall := fake.getCommandHelpArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Executor) GetCommandHelpReturns(result1 []byte, result2 error) {
	fake.getCommandHelpMutex.Lock()
	defer fake.getCommandHelpMutex.Unlock()
	fake.GetCommandHelpStub = nil
	fake.getCommandHelpReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Executor) GetCommandHelpReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getCommandHelpMutex.Lock()
	defer fake.getCommandHelpMutex.Unlock()
	fake.GetCommandHelpStub = nil
	if fake.getCommandHelpReturnsOnCall == nil {
		fake.getCommandHelpReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getCommandHelpReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Executor) GetCommandNamesAndDescriptions() (map[string]string, error) {
	fake.getCommandNamesAndDescriptionsMutex.Lock()
	ret, specificReturn := fake.getCommandNamesAndDescriptionsReturnsOnCall[len(fake.getCommandNamesAndDescriptionsArgsForCall)]
	fake.getCommandNamesAndDescriptionsArgsForCall = append(fake.getCommandNamesAndDescriptionsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCommandNamesAndDescriptions", []interface{}{})
	fake.getCommandNamesAndDescriptionsMutex.Unlock()
	if fake.GetCommandNamesAndDescriptionsStub != nil {
		return fake.GetCommandNamesAndDescriptionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCommandNamesAndDescriptionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Executor) GetCommandNamesAndDescriptionsCallCount() int {
	fake.getCommandNamesAndDescriptionsMutex.RLock()
	defer fake.getCommandNamesAndDescriptionsMutex.RUnlock()
	return len(fake.getCommandNamesAndDescriptionsArgsForCall)
}

func (fake *Executor) GetCommandNamesAndDescriptionsCalls(stub func() (map[string]string, error)) {
	fake.getCommandNamesAndDescriptionsMutex.Lock()
	defer fake.getCommandNamesAndDescriptionsMutex.Unlock()
	fake.GetCommandNamesAndDescriptionsStub = stub
}

func (fake *Executor) GetCommandNamesAndDescriptionsReturns(result1 map[string]string, result2 error) {
	fake.getCommandNamesAndDescriptionsMutex.Lock()
	defer fake.getCommandNamesAndDescriptionsMutex.Unlock()
	fake.GetCommandNamesAndDescriptionsStub = nil
	fake.getCommandNamesAndDescriptionsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *Executor) GetCommandNamesAndDescriptionsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.getCommandNamesAndDescriptionsMutex.Lock()
	defer fake.getCommandNamesAndDescriptionsMutex.Unlock()
	fake.GetCommandNamesAndDescriptionsStub = nil
	if fake.getCommandNamesAndDescriptionsReturnsOnCall == nil {
		fake.getCommandNamesAndDescriptionsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.getCommandNamesAndDescriptionsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *Executor) GetDescription(arg1 string) (string, error) {
	fake.getDescriptionMutex.Lock()
	ret, specificReturn := fake.getDescriptionReturnsOnCall[len(fake.getDescriptionArgsForCall)]
	fake.getDescriptionArgsForCall = append(fake.getDescriptionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetDescription", []interface{}{arg1})
	fake.getDescriptionMutex.Unlock()
	if fake.GetDescriptionStub != nil {
		return fake.GetDescriptionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDescriptionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Executor) GetDescriptionCallCount() int {
	fake.getDescriptionMutex.RLock()
	defer fake.getDescriptionMutex.RUnlock()
	return len(fake.getDescriptionArgsForCall)
}

func (fake *Executor) GetDescriptionCalls(stub func(string) (string, error)) {
	fake.getDescriptionMutex.Lock()
	defer fake.getDescriptionMutex.Unlock()
	fake.GetDescriptionStub = stub
}

func (fake *Executor) GetDescriptionArgsForCall(i int) string {
	fake.getDescriptionMutex.RLock()
	defer fake.getDescriptionMutex.RUnlock()
	argsForCall := fake.getDescriptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Executor) GetDescriptionReturns(result1 string, result2 error) {
	fake.getDescriptionMutex.Lock()
	defer fake.getDescriptionMutex.Unlock()
	fake.GetDescriptionStub = nil
	fake.getDescriptionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Executor) GetDescriptionReturnsOnCall(i int, result1 string, result2 error) {
	fake.getDescriptionMutex.Lock()
	defer fake.getDescriptionMutex.Unlock()
	fake.GetDescriptionStub = nil
	if fake.getDescriptionReturnsOnCall == nil {
		fake.getDescriptionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getDescriptionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Executor) RunOmCommand(arg1 ...string) ([]byte, error) {
	fake.runOmCommandMutex.Lock()
	ret, specificReturn := fake.runOmCommandReturnsOnCall[len(fake.runOmCommandArgsForCall)]
	fake.runOmCommandArgsForCall = append(fake.runOmCommandArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("RunOmCommand", []interface{}{arg1})
	fake.runOmCommandMutex.Unlock()
	if fake.RunOmCommandStub != nil {
		return fake.RunOmCommandStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runOmCommandReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Executor) RunOmCommandCallCount() int {
	fake.runOmCommandMutex.RLock()
	defer fake.runOmCommandMutex.RUnlock()
	return len(fake.runOmCommandArgsForCall)
}

func (fake *Executor) RunOmCommandCalls(stub func(...string) ([]byte, error)) {
	fake.runOmCommandMutex.Lock()
	defer fake.runOmCommandMutex.Unlock()
	fake.RunOmCommandStub = stub
}

func (fake *Executor) RunOmCommandArgsForCall(i int) []string {
	fake.runOmCommandMutex.RLock()
	defer fake.runOmCommandMutex.RUnlock()
	argsForCall := fake.runOmCommandArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Executor) RunOmCommandReturns(result1 []byte, result2 error) {
	fake.runOmCommandMutex.Lock()
	defer fake.runOmCommandMutex.Unlock()
	fake.RunOmCommandStub = nil
	fake.runOmCommandReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Executor) RunOmCommandReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.runOmCommandMutex.Lock()
	defer fake.runOmCommandMutex.Unlock()
	fake.RunOmCommandStub = nil
	if fake.runOmCommandReturnsOnCall == nil {
		fake.runOmCommandReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.runOmCommandReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Executor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCommandHelpMutex.RLock()
	defer fake.getCommandHelpMutex.RUnlock()
	fake.getCommandNamesAndDescriptionsMutex.RLock()
	defer fake.getCommandNamesAndDescriptionsMutex.RUnlock()
	fake.getDescriptionMutex.RLock()
	defer fake.getDescriptionMutex.RUnlock()
	fake.runOmCommandMutex.RLock()
	defer fake.runOmCommandMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Executor) recordInvocation(key string, args []interface{}) {
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