// Copyright 2020 lack
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by main. Do NOT EDIT.

package sets

import (
	"reflect"
	"sort"
)

// sets.Float64 is a set of float64s, implemented via map[float64]struct{} for minimal memory consumption.
type Float64 map[float64]Empty

// NewFloat64 creates a Float64 from a list of values.
func NewFloat64(items ...float64) Float64 {
	ss := Float64{}
	ss.Insert(items...)
	return ss
}

// Float64KeySet creates a Float64 from a keys of a map[float64](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func Float64KeySet(theMap interface{}) Float64 {
	v := reflect.ValueOf(theMap)
	ret := Float64{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(float64))
	}
	return ret
}

// Insert adds items to the set.
func (s Float64) Insert(items ...float64) Float64 {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s Float64) Delete(items ...float64) Float64 {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s Float64) Has(item float64) bool {
	_, contained := s[item]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s Float64) HasAll(items ...float64) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s Float64) HasAny(items ...float64) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s Float64) Difference(s2 Float64) Float64 {
	result := NewFloat64()
	for key := range s {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Float64) Union(s2 Float64) Float64 {
	result := NewFloat64()
	for key := range s1 {
		result.Insert(key)
	}
	for key := range s2 {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Float64) Intersection(s2 Float64) Float64 {
	var walk, other Float64
	result := NewFloat64()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Float64) IsSuperset(s2 Float64) bool {
	for item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Float64) Equal(s2 Float64) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfFloat64 []float64

func (s sortableSliceOfFloat64) Len() int           { return len(s) }
func (s sortableSliceOfFloat64) Less(i, j int) bool { return lessFloat64(s[i], s[j]) }
func (s sortableSliceOfFloat64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted float64 slice.
func (s Float64) List() []float64 {
	res := make(sortableSliceOfFloat64, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []float64(res)
}

// UnsortedList returns the slice with contents in random order.
func (s Float64) UnsortedList() []float64 {
	res := make([]float64, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Returns a single element from the set.
func (s Float64) PopAny() (float64, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue float64
	return zeroValue, false
}

// Len returns the size of the set.
func (s Float64) Len() int {
	return len(s)
}

func lessFloat64(lhs, rhs float64) bool {
	return lhs < rhs
}