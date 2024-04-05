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

package v1 // github.com/openshift-online/ocm-api-metamodel/tests/go/generated/servicemgmt/v1

// StatefulObject represents the values of the 'stateful_object' type.
type StatefulObject struct {
	bitmap_ uint32
	id      string
	href    string
	kind    string
	state   string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *StatefulObject) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// ID returns the value of the 'ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *StatefulObject) ID() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.id
	}
	return ""
}

// GetID returns the value of the 'ID' attribute and
// a flag indicating if the attribute has a value.
func (o *StatefulObject) GetID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.id
	}
	return
}

// Href returns the value of the 'href' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *StatefulObject) Href() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.href
	}
	return ""
}

// GetHref returns the value of the 'href' attribute and
// a flag indicating if the attribute has a value.
func (o *StatefulObject) GetHref() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.href
	}
	return
}

// Kind returns the value of the 'kind' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *StatefulObject) Kind() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.kind
	}
	return ""
}

// GetKind returns the value of the 'kind' attribute and
// a flag indicating if the attribute has a value.
func (o *StatefulObject) GetKind() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.kind
	}
	return
}

// State returns the value of the 'state' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *StatefulObject) State() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.state
	}
	return ""
}

// GetState returns the value of the 'state' attribute and
// a flag indicating if the attribute has a value.
func (o *StatefulObject) GetState() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.state
	}
	return
}

// StatefulObjectListKind is the name of the type used to represent list of objects of
// type 'stateful_object'.
const StatefulObjectListKind = "StatefulObjectList"

// StatefulObjectListLinkKind is the name of the type used to represent links to list
// of objects of type 'stateful_object'.
const StatefulObjectListLinkKind = "StatefulObjectListLink"

// StatefulObjectNilKind is the name of the type used to nil lists of objects of
// type 'stateful_object'.
const StatefulObjectListNilKind = "StatefulObjectListNil"

// StatefulObjectList is a list of values of the 'stateful_object' type.
type StatefulObjectList struct {
	href  string
	link  bool
	items []*StatefulObject
}

// Len returns the length of the list.
func (l *StatefulObjectList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *StatefulObjectList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *StatefulObjectList) Get(i int) *StatefulObject {
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
func (l *StatefulObjectList) Slice() []*StatefulObject {
	var slice []*StatefulObject
	if l == nil {
		slice = make([]*StatefulObject, 0)
	} else {
		slice = make([]*StatefulObject, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *StatefulObjectList) Each(f func(item *StatefulObject) bool) {
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
func (l *StatefulObjectList) Range(f func(index int, item *StatefulObject) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
