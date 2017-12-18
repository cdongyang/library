package algorithm

import (
	"reflect"
)

type Interface interface {
	Len() int
	Less(int) bool
	Greater(int) bool
}

type SearchIntSlice struct {
	Data []int
	Key  int
}

//if use *SearchInt32Slice there will cause 1 allocs/op and the time use will become double or more

func (s SearchIntSlice) Less(i int) bool {
	return s.Data[i] < s.Key
}
func (s SearchIntSlice) Greater(i int) bool {
	return s.Data[i] > s.Key
}
func (s SearchIntSlice) Len() int {
	return len(s.Data)
}
func (s SearchIntSlice) LowerBound() int {
	return lowerBound(s.Less, s.Len())
}
func (s SearchIntSlice) UpperBound() int {
	return upperBound(s.Greater, s.Len())
}

//LowerBound find the first index that data[index]>=key
//for example [1,2,3,3,3,4,4,5] find 3 return index 2 val 3
func LowerBound(data Interface) int {
	return lowerBound(data.Less, data.Len())
}

//SliceLowerBound find the first index that data[index]>=key with less function
//the less func should like
//
//	func(i int) bool {
//		return slice[i]<key
//	}
//
// The function panics if the provided interface is not a slice.
func SliceLowerBound(data interface{}, less func(int) bool) int {
	rv := reflect.ValueOf(data)
	return lowerBound(less, rv.Len())
}

func lowerBound(less func(int) bool, length int) int {
	var l = 0
	var r = length - 1
	for l < r {
		m := (r + l) / 2
		if less(m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

//UpperBound find the first index that data[index]>key
//for example [1,2,3,3,3,4,4,5] find 3 return index 5 val 4
func UpperBound(data Interface) int {
	return upperBound(data.Greater, data.Len())
}

//SliceUpperBound find the first index that data[index]>key with greater function
//the greater func should like
//
//	func(i int) bool {
//		return slice[i]>key
//	}
//
// The function panics if the provided interface is not a slice.
func SliceUpperBound(data interface{}, greater func(int) bool) int {
	rv := reflect.ValueOf(data)
	return upperBound(greater, rv.Len())
}

func upperBound(greater func(int) bool, length int) int {
	var l = 0
	var r = length - 1
	for l < r {
		m := (r + l) / 2
		if greater(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
