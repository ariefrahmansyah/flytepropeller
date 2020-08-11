// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// ImmutableExecutionContext is an autogenerated mock type for the ImmutableExecutionContext type
type ImmutableExecutionContext struct {
	mock.Mock
}

type ImmutableExecutionContext_GetAnnotations struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetAnnotations) Return(_a0 map[string]string) *ImmutableExecutionContext_GetAnnotations {
	return &ImmutableExecutionContext_GetAnnotations{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetAnnotations() *ImmutableExecutionContext_GetAnnotations {
	c := _m.On("GetAnnotations")
	return &ImmutableExecutionContext_GetAnnotations{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetAnnotationsMatch(matchers ...interface{}) *ImmutableExecutionContext_GetAnnotations {
	c := _m.On("GetAnnotations", matchers...)
	return &ImmutableExecutionContext_GetAnnotations{Call: c}
}

// GetAnnotations provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type ImmutableExecutionContext_GetCreationTimestamp struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetCreationTimestamp) Return(_a0 v1.Time) *ImmutableExecutionContext_GetCreationTimestamp {
	return &ImmutableExecutionContext_GetCreationTimestamp{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetCreationTimestamp() *ImmutableExecutionContext_GetCreationTimestamp {
	c := _m.On("GetCreationTimestamp")
	return &ImmutableExecutionContext_GetCreationTimestamp{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetCreationTimestampMatch(matchers ...interface{}) *ImmutableExecutionContext_GetCreationTimestamp {
	c := _m.On("GetCreationTimestamp", matchers...)
	return &ImmutableExecutionContext_GetCreationTimestamp{Call: c}
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetCreationTimestamp() v1.Time {
	ret := _m.Called()

	var r0 v1.Time
	if rf, ok := ret.Get(0).(func() v1.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.Time)
	}

	return r0
}

type ImmutableExecutionContext_GetEventVersion struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetEventVersion) Return(_a0 v1alpha1.EventVersion) *ImmutableExecutionContext_GetEventVersion {
	return &ImmutableExecutionContext_GetEventVersion{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetEventVersion() *ImmutableExecutionContext_GetEventVersion {
	c := _m.On("GetEventVersion")
	return &ImmutableExecutionContext_GetEventVersion{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetEventVersionMatch(matchers ...interface{}) *ImmutableExecutionContext_GetEventVersion {
	c := _m.On("GetEventVersion", matchers...)
	return &ImmutableExecutionContext_GetEventVersion{Call: c}
}

// GetEventVersion provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetEventVersion() v1alpha1.EventVersion {
	ret := _m.Called()

	var r0 v1alpha1.EventVersion
	if rf, ok := ret.Get(0).(func() v1alpha1.EventVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.EventVersion)
	}

	return r0
}

type ImmutableExecutionContext_GetExecutionID struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetExecutionID) Return(_a0 v1alpha1.WorkflowExecutionIdentifier) *ImmutableExecutionContext_GetExecutionID {
	return &ImmutableExecutionContext_GetExecutionID{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetExecutionID() *ImmutableExecutionContext_GetExecutionID {
	c := _m.On("GetExecutionID")
	return &ImmutableExecutionContext_GetExecutionID{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetExecutionIDMatch(matchers ...interface{}) *ImmutableExecutionContext_GetExecutionID {
	c := _m.On("GetExecutionID", matchers...)
	return &ImmutableExecutionContext_GetExecutionID{Call: c}
}

// GetExecutionID provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetExecutionID() v1alpha1.WorkflowExecutionIdentifier {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowExecutionIdentifier
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowExecutionIdentifier); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowExecutionIdentifier)
	}

	return r0
}

type ImmutableExecutionContext_GetID struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetID) Return(_a0 string) *ImmutableExecutionContext_GetID {
	return &ImmutableExecutionContext_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetID() *ImmutableExecutionContext_GetID {
	c := _m.On("GetID")
	return &ImmutableExecutionContext_GetID{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetIDMatch(matchers ...interface{}) *ImmutableExecutionContext_GetID {
	c := _m.On("GetID", matchers...)
	return &ImmutableExecutionContext_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ImmutableExecutionContext_GetK8sWorkflowID struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetK8sWorkflowID) Return(_a0 types.NamespacedName) *ImmutableExecutionContext_GetK8sWorkflowID {
	return &ImmutableExecutionContext_GetK8sWorkflowID{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetK8sWorkflowID() *ImmutableExecutionContext_GetK8sWorkflowID {
	c := _m.On("GetK8sWorkflowID")
	return &ImmutableExecutionContext_GetK8sWorkflowID{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetK8sWorkflowIDMatch(matchers ...interface{}) *ImmutableExecutionContext_GetK8sWorkflowID {
	c := _m.On("GetK8sWorkflowID", matchers...)
	return &ImmutableExecutionContext_GetK8sWorkflowID{Call: c}
}

// GetK8sWorkflowID provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetK8sWorkflowID() types.NamespacedName {
	ret := _m.Called()

	var r0 types.NamespacedName
	if rf, ok := ret.Get(0).(func() types.NamespacedName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NamespacedName)
	}

	return r0
}

type ImmutableExecutionContext_GetLabels struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetLabels) Return(_a0 map[string]string) *ImmutableExecutionContext_GetLabels {
	return &ImmutableExecutionContext_GetLabels{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetLabels() *ImmutableExecutionContext_GetLabels {
	c := _m.On("GetLabels")
	return &ImmutableExecutionContext_GetLabels{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetLabelsMatch(matchers ...interface{}) *ImmutableExecutionContext_GetLabels {
	c := _m.On("GetLabels", matchers...)
	return &ImmutableExecutionContext_GetLabels{Call: c}
}

// GetLabels provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type ImmutableExecutionContext_GetName struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetName) Return(_a0 string) *ImmutableExecutionContext_GetName {
	return &ImmutableExecutionContext_GetName{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetName() *ImmutableExecutionContext_GetName {
	c := _m.On("GetName")
	return &ImmutableExecutionContext_GetName{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetNameMatch(matchers ...interface{}) *ImmutableExecutionContext_GetName {
	c := _m.On("GetName", matchers...)
	return &ImmutableExecutionContext_GetName{Call: c}
}

// GetName provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ImmutableExecutionContext_GetNamespace struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetNamespace) Return(_a0 string) *ImmutableExecutionContext_GetNamespace {
	return &ImmutableExecutionContext_GetNamespace{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetNamespace() *ImmutableExecutionContext_GetNamespace {
	c := _m.On("GetNamespace")
	return &ImmutableExecutionContext_GetNamespace{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetNamespaceMatch(matchers ...interface{}) *ImmutableExecutionContext_GetNamespace {
	c := _m.On("GetNamespace", matchers...)
	return &ImmutableExecutionContext_GetNamespace{Call: c}
}

// GetNamespace provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ImmutableExecutionContext_GetOnFailurePolicy struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetOnFailurePolicy) Return(_a0 v1alpha1.WorkflowOnFailurePolicy) *ImmutableExecutionContext_GetOnFailurePolicy {
	return &ImmutableExecutionContext_GetOnFailurePolicy{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetOnFailurePolicy() *ImmutableExecutionContext_GetOnFailurePolicy {
	c := _m.On("GetOnFailurePolicy")
	return &ImmutableExecutionContext_GetOnFailurePolicy{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetOnFailurePolicyMatch(matchers ...interface{}) *ImmutableExecutionContext_GetOnFailurePolicy {
	c := _m.On("GetOnFailurePolicy", matchers...)
	return &ImmutableExecutionContext_GetOnFailurePolicy{Call: c}
}

// GetOnFailurePolicy provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetOnFailurePolicy() v1alpha1.WorkflowOnFailurePolicy {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowOnFailurePolicy
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowOnFailurePolicy); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowOnFailurePolicy)
	}

	return r0
}

type ImmutableExecutionContext_GetOwnerReference struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetOwnerReference) Return(_a0 v1.OwnerReference) *ImmutableExecutionContext_GetOwnerReference {
	return &ImmutableExecutionContext_GetOwnerReference{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetOwnerReference() *ImmutableExecutionContext_GetOwnerReference {
	c := _m.On("GetOwnerReference")
	return &ImmutableExecutionContext_GetOwnerReference{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetOwnerReferenceMatch(matchers ...interface{}) *ImmutableExecutionContext_GetOwnerReference {
	c := _m.On("GetOwnerReference", matchers...)
	return &ImmutableExecutionContext_GetOwnerReference{Call: c}
}

// GetOwnerReference provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetOwnerReference() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}

type ImmutableExecutionContext_GetServiceAccountName struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_GetServiceAccountName) Return(_a0 string) *ImmutableExecutionContext_GetServiceAccountName {
	return &ImmutableExecutionContext_GetServiceAccountName{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnGetServiceAccountName() *ImmutableExecutionContext_GetServiceAccountName {
	c := _m.On("GetServiceAccountName")
	return &ImmutableExecutionContext_GetServiceAccountName{Call: c}
}

func (_m *ImmutableExecutionContext) OnGetServiceAccountNameMatch(matchers ...interface{}) *ImmutableExecutionContext_GetServiceAccountName {
	c := _m.On("GetServiceAccountName", matchers...)
	return &ImmutableExecutionContext_GetServiceAccountName{Call: c}
}

// GetServiceAccountName provides a mock function with given fields:
func (_m *ImmutableExecutionContext) GetServiceAccountName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ImmutableExecutionContext_IsInterruptible struct {
	*mock.Call
}

func (_m ImmutableExecutionContext_IsInterruptible) Return(_a0 bool) *ImmutableExecutionContext_IsInterruptible {
	return &ImmutableExecutionContext_IsInterruptible{Call: _m.Call.Return(_a0)}
}

func (_m *ImmutableExecutionContext) OnIsInterruptible() *ImmutableExecutionContext_IsInterruptible {
	c := _m.On("IsInterruptible")
	return &ImmutableExecutionContext_IsInterruptible{Call: c}
}

func (_m *ImmutableExecutionContext) OnIsInterruptibleMatch(matchers ...interface{}) *ImmutableExecutionContext_IsInterruptible {
	c := _m.On("IsInterruptible", matchers...)
	return &ImmutableExecutionContext_IsInterruptible{Call: c}
}

// IsInterruptible provides a mock function with given fields:
func (_m *ImmutableExecutionContext) IsInterruptible() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
