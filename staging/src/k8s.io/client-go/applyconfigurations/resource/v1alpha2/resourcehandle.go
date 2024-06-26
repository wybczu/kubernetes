/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

// ResourceHandleApplyConfiguration represents a declarative configuration of the ResourceHandle type for use
// with apply.
type ResourceHandleApplyConfiguration struct {
	DriverName     *string                                     `json:"driverName,omitempty"`
	Data           *string                                     `json:"data,omitempty"`
	StructuredData *StructuredResourceHandleApplyConfiguration `json:"structuredData,omitempty"`
}

// ResourceHandleApplyConfiguration constructs a declarative configuration of the ResourceHandle type for use with
// apply.
func ResourceHandle() *ResourceHandleApplyConfiguration {
	return &ResourceHandleApplyConfiguration{}
}

// WithDriverName sets the DriverName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DriverName field is set to the value of the last call.
func (b *ResourceHandleApplyConfiguration) WithDriverName(value string) *ResourceHandleApplyConfiguration {
	b.DriverName = &value
	return b
}

// WithData sets the Data field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Data field is set to the value of the last call.
func (b *ResourceHandleApplyConfiguration) WithData(value string) *ResourceHandleApplyConfiguration {
	b.Data = &value
	return b
}

// WithStructuredData sets the StructuredData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StructuredData field is set to the value of the last call.
func (b *ResourceHandleApplyConfiguration) WithStructuredData(value *StructuredResourceHandleApplyConfiguration) *ResourceHandleApplyConfiguration {
	b.StructuredData = value
	return b
}
