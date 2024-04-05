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

package v1 // github.com/openshift-online/ocm-api-metamodel/tests/go/generated/addonsmgmt/v1

// AddonParameterOption represents the values of the 'addon_parameter_option' type.
//
// Representation of an addon parameter option.
type AddonParameterOption struct {
	bitmap_      uint32
	name         string
	rank         int
	requirements []*AddonRequirement
	value        string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *AddonParameterOption) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the addon parameter option.
func (o *AddonParameterOption) Name() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the addon parameter option.
func (o *AddonParameterOption) GetName() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.name
	}
	return
}

// Rank returns the value of the 'rank' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Rank of option to be used in cases where editable direction should be restricted.
func (o *AddonParameterOption) Rank() int {
	if o != nil && o.bitmap_&2 != 0 {
		return o.rank
	}
	return 0
}

// GetRank returns the value of the 'rank' attribute and
// a flag indicating if the attribute has a value.
//
// Rank of option to be used in cases where editable direction should be restricted.
func (o *AddonParameterOption) GetRank() (value int, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.rank
	}
	return
}

// Requirements returns the value of the 'requirements' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of addon requirements for this parameter option.
func (o *AddonParameterOption) Requirements() []*AddonRequirement {
	if o != nil && o.bitmap_&4 != 0 {
		return o.requirements
	}
	return nil
}

// GetRequirements returns the value of the 'requirements' attribute and
// a flag indicating if the attribute has a value.
//
// List of addon requirements for this parameter option.
func (o *AddonParameterOption) GetRequirements() (value []*AddonRequirement, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.requirements
	}
	return
}

// Value returns the value of the 'value' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Value of the addon parameter option.
func (o *AddonParameterOption) Value() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.value
	}
	return ""
}

// GetValue returns the value of the 'value' attribute and
// a flag indicating if the attribute has a value.
//
// Value of the addon parameter option.
func (o *AddonParameterOption) GetValue() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.value
	}
	return
}

// AddonParameterOptionListKind is the name of the type used to represent list of objects of
// type 'addon_parameter_option'.
const AddonParameterOptionListKind = "AddonParameterOptionList"

// AddonParameterOptionListLinkKind is the name of the type used to represent links to list
// of objects of type 'addon_parameter_option'.
const AddonParameterOptionListLinkKind = "AddonParameterOptionListLink"

// AddonParameterOptionNilKind is the name of the type used to nil lists of objects of
// type 'addon_parameter_option'.
const AddonParameterOptionListNilKind = "AddonParameterOptionListNil"

// AddonParameterOptionList is a list of values of the 'addon_parameter_option' type.
type AddonParameterOptionList struct {
	href  string
	link  bool
	items []*AddonParameterOption
}

// Len returns the length of the list.
func (l *AddonParameterOptionList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *AddonParameterOptionList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AddonParameterOptionList) Get(i int) *AddonParameterOption {
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
func (l *AddonParameterOptionList) Slice() []*AddonParameterOption {
	var slice []*AddonParameterOption
	if l == nil {
		slice = make([]*AddonParameterOption, 0)
	} else {
		slice = make([]*AddonParameterOption, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AddonParameterOptionList) Each(f func(item *AddonParameterOption) bool) {
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
func (l *AddonParameterOptionList) Range(f func(index int, item *AddonParameterOption) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
