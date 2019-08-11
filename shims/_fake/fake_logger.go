// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	log "github.com/rantav/go-logger"
)

type FakeLogger struct {
	DebugStub        func(msg ...interface{})
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		msg []interface{}
	}
	InfoStub        func(msg ...interface{})
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		msg []interface{}
	}
	WarnStub        func(msg ...interface{})
	warnMutex       sync.RWMutex
	warnArgsForCall []struct {
		msg []interface{}
	}
	ErrorStub        func(msg ...interface{})
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		msg []interface{}
	}
	DebugfStub        func(format string, args ...interface{})
	debugfMutex       sync.RWMutex
	debugfArgsForCall []struct {
		format string
		args   []interface{}
	}
	InfofStub        func(format string, args ...interface{})
	infofMutex       sync.RWMutex
	infofArgsForCall []struct {
		format string
		args   []interface{}
	}
	WarnfStub        func(format string, args ...interface{})
	warnfMutex       sync.RWMutex
	warnfArgsForCall []struct {
		format string
		args   []interface{}
	}
	ErrorfStub        func(format string, args ...interface{})
	errorfMutex       sync.RWMutex
	errorfArgsForCall []struct {
		format string
		args   []interface{}
	}
	WithFieldsStub        func(log.Fields) log.Logger
	withFieldsMutex       sync.RWMutex
	withFieldsArgsForCall []struct {
		arg1 log.Fields
	}
	withFieldsReturns struct {
		result1 log.Logger
	}
	withFieldsReturnsOnCall map[int]struct {
		result1 log.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Debug(msg ...interface{}) {
	fake.debugMutex.Lock()
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		msg []interface{}
	}{msg})
	fake.recordInvocation("Debug", []interface{}{msg})
	fake.debugMutex.Unlock()
	if fake.DebugStub != nil {
		fake.DebugStub(msg...)
	}
}

func (fake *FakeLogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *FakeLogger) DebugArgsForCall(i int) []interface{} {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return fake.debugArgsForCall[i].msg
}

func (fake *FakeLogger) Info(msg ...interface{}) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		msg []interface{}
	}{msg})
	fake.recordInvocation("Info", []interface{}{msg})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		fake.InfoStub(msg...)
	}
}

func (fake *FakeLogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeLogger) InfoArgsForCall(i int) []interface{} {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].msg
}

func (fake *FakeLogger) Warn(msg ...interface{}) {
	fake.warnMutex.Lock()
	fake.warnArgsForCall = append(fake.warnArgsForCall, struct {
		msg []interface{}
	}{msg})
	fake.recordInvocation("Warn", []interface{}{msg})
	fake.warnMutex.Unlock()
	if fake.WarnStub != nil {
		fake.WarnStub(msg...)
	}
}

func (fake *FakeLogger) WarnCallCount() int {
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	return len(fake.warnArgsForCall)
}

func (fake *FakeLogger) WarnArgsForCall(i int) []interface{} {
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	return fake.warnArgsForCall[i].msg
}

func (fake *FakeLogger) Error(msg ...interface{}) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		msg []interface{}
	}{msg})
	fake.recordInvocation("Error", []interface{}{msg})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		fake.ErrorStub(msg...)
	}
}

func (fake *FakeLogger) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeLogger) ErrorArgsForCall(i int) []interface{} {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return fake.errorArgsForCall[i].msg
}

func (fake *FakeLogger) Debugf(format string, args ...interface{}) {
	fake.debugfMutex.Lock()
	fake.debugfArgsForCall = append(fake.debugfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.recordInvocation("Debugf", []interface{}{format, args})
	fake.debugfMutex.Unlock()
	if fake.DebugfStub != nil {
		fake.DebugfStub(format, args...)
	}
}

func (fake *FakeLogger) DebugfCallCount() int {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return len(fake.debugfArgsForCall)
}

func (fake *FakeLogger) DebugfArgsForCall(i int) (string, []interface{}) {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return fake.debugfArgsForCall[i].format, fake.debugfArgsForCall[i].args
}

func (fake *FakeLogger) Infof(format string, args ...interface{}) {
	fake.infofMutex.Lock()
	fake.infofArgsForCall = append(fake.infofArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.recordInvocation("Infof", []interface{}{format, args})
	fake.infofMutex.Unlock()
	if fake.InfofStub != nil {
		fake.InfofStub(format, args...)
	}
}

func (fake *FakeLogger) InfofCallCount() int {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return len(fake.infofArgsForCall)
}

func (fake *FakeLogger) InfofArgsForCall(i int) (string, []interface{}) {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return fake.infofArgsForCall[i].format, fake.infofArgsForCall[i].args
}

func (fake *FakeLogger) Warnf(format string, args ...interface{}) {
	fake.warnfMutex.Lock()
	fake.warnfArgsForCall = append(fake.warnfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.recordInvocation("Warnf", []interface{}{format, args})
	fake.warnfMutex.Unlock()
	if fake.WarnfStub != nil {
		fake.WarnfStub(format, args...)
	}
}

func (fake *FakeLogger) WarnfCallCount() int {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	return len(fake.warnfArgsForCall)
}

func (fake *FakeLogger) WarnfArgsForCall(i int) (string, []interface{}) {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	return fake.warnfArgsForCall[i].format, fake.warnfArgsForCall[i].args
}

func (fake *FakeLogger) Errorf(format string, args ...interface{}) {
	fake.errorfMutex.Lock()
	fake.errorfArgsForCall = append(fake.errorfArgsForCall, struct {
		format string
		args   []interface{}
	}{format, args})
	fake.recordInvocation("Errorf", []interface{}{format, args})
	fake.errorfMutex.Unlock()
	if fake.ErrorfStub != nil {
		fake.ErrorfStub(format, args...)
	}
}

func (fake *FakeLogger) ErrorfCallCount() int {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return len(fake.errorfArgsForCall)
}

func (fake *FakeLogger) ErrorfArgsForCall(i int) (string, []interface{}) {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return fake.errorfArgsForCall[i].format, fake.errorfArgsForCall[i].args
}

func (fake *FakeLogger) WithFields(arg1 log.Fields) log.Logger {
	fake.withFieldsMutex.Lock()
	ret, specificReturn := fake.withFieldsReturnsOnCall[len(fake.withFieldsArgsForCall)]
	fake.withFieldsArgsForCall = append(fake.withFieldsArgsForCall, struct {
		arg1 log.Fields
	}{arg1})
	fake.recordInvocation("WithFields", []interface{}{arg1})
	fake.withFieldsMutex.Unlock()
	if fake.WithFieldsStub != nil {
		return fake.WithFieldsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.withFieldsReturns.result1
}

func (fake *FakeLogger) WithFieldsCallCount() int {
	fake.withFieldsMutex.RLock()
	defer fake.withFieldsMutex.RUnlock()
	return len(fake.withFieldsArgsForCall)
}

func (fake *FakeLogger) WithFieldsArgsForCall(i int) log.Fields {
	fake.withFieldsMutex.RLock()
	defer fake.withFieldsMutex.RUnlock()
	return fake.withFieldsArgsForCall[i].arg1
}

func (fake *FakeLogger) WithFieldsReturns(result1 log.Logger) {
	fake.WithFieldsStub = nil
	fake.withFieldsReturns = struct {
		result1 log.Logger
	}{result1}
}

func (fake *FakeLogger) WithFieldsReturnsOnCall(i int, result1 log.Logger) {
	fake.WithFieldsStub = nil
	if fake.withFieldsReturnsOnCall == nil {
		fake.withFieldsReturnsOnCall = make(map[int]struct {
			result1 log.Logger
		})
	}
	fake.withFieldsReturnsOnCall[i] = struct {
		result1 log.Logger
	}{result1}
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	fake.withFieldsMutex.RLock()
	defer fake.withFieldsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ log.Logger = new(FakeLogger)
