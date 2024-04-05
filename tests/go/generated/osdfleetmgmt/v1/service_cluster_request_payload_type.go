/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-api-metamodel/tests/go/generated/osdfleetmgmt/v1

// ServiceClusterRequestPayload represents the values of the 'service_cluster_request_payload' type.
type ServiceClusterRequestPayload struct {
	bitmap_       uint32
	cloudProvider string
	labels        []*LabelRequestPayload
	region        string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ServiceClusterRequestPayload) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// CloudProvider returns the value of the 'cloud_provider' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *ServiceClusterRequestPayload) CloudProvider() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.cloudProvider
	}
	return ""
}

// GetCloudProvider returns the value of the 'cloud_provider' attribute and
// a flag indicating if the attribute has a value.
func (o *ServiceClusterRequestPayload) GetCloudProvider() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.cloudProvider
	}
	return
}

// Labels returns the value of the 'labels' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *ServiceClusterRequestPayload) Labels() []*LabelRequestPayload {
	if o != nil && o.bitmap_&2 != 0 {
		return o.labels
	}
	return nil
}

// GetLabels returns the value of the 'labels' attribute and
// a flag indicating if the attribute has a value.
func (o *ServiceClusterRequestPayload) GetLabels() (value []*LabelRequestPayload, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.labels
	}
	return
}

// Region returns the value of the 'region' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *ServiceClusterRequestPayload) Region() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.region
	}
	return ""
}

// GetRegion returns the value of the 'region' attribute and
// a flag indicating if the attribute has a value.
func (o *ServiceClusterRequestPayload) GetRegion() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.region
	}
	return
}

// ServiceClusterRequestPayloadListKind is the name of the type used to represent list of objects of
// type 'service_cluster_request_payload'.
const ServiceClusterRequestPayloadListKind = "ServiceClusterRequestPayloadList"

// ServiceClusterRequestPayloadListLinkKind is the name of the type used to represent links to list
// of objects of type 'service_cluster_request_payload'.
const ServiceClusterRequestPayloadListLinkKind = "ServiceClusterRequestPayloadListLink"

// ServiceClusterRequestPayloadNilKind is the name of the type used to nil lists of objects of
// type 'service_cluster_request_payload'.
const ServiceClusterRequestPayloadListNilKind = "ServiceClusterRequestPayloadListNil"

// ServiceClusterRequestPayloadList is a list of values of the 'service_cluster_request_payload' type.
type ServiceClusterRequestPayloadList struct {
	href  string
	link  bool
	items []*ServiceClusterRequestPayload
}

// Len returns the length of the list.
func (l *ServiceClusterRequestPayloadList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ServiceClusterRequestPayloadList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ServiceClusterRequestPayloadList) Get(i int) *ServiceClusterRequestPayload {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ServiceClusterRequestPayloadList) Slice() []*ServiceClusterRequestPayload {
	var slice []*ServiceClusterRequestPayload
	if l == nil {
		slice = make([]*ServiceClusterRequestPayload, 0)
	} else {
		slice = make([]*ServiceClusterRequestPayload, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ServiceClusterRequestPayloadList) Each(f func(item *ServiceClusterRequestPayload) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ServiceClusterRequestPayloadList) Range(f func(index int, item *ServiceClusterRequestPayload) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
