package algorithm

import (
	"reflect"
)

type Interface interface {
	Len() int
	Compare(int) int
}

type SearchIntSlice struct {
	Data []int
	Key  int
}

//if use *SearchInt32Slice there will cause 1 allocs/op and the time use will become double or more

func (s SearchIntSlice) Compare(i int) int {
	return s.Data[i] - s.Key
}
func (s SearchIntSlice) Len() int {
	return len(s.Data)
}
func (s SearchIntSlice) LowerBound() int {
	return lowerBound(s.Compare, s.Len())
}
func (s SearchIntSlice) UpperBound() int {
	return upperBound(s.Compare, s.Len())
}

//LowerBound find the first index that data[index]>=key
//for example [1,2,3,3,3,4,4,5] find 3 return index 2 val 3
func LowerBound(data Interface) int {
	return lowerBound(data.Compare, data.Len())
}

//SliceLowerBound find the first index that data[index]>=key with compare function
//the compare func should like
//	greater x>0 	equal x==0   less x<0
//	func(i int) (x int) {
//		if slice[i] > key {
//			return 1
//		} else if slice[i] == key {
//			return 0
//		} else {
//			return -1
//		}
//	}
//
// The function panics if the provided interface is not a slice.
func SliceLowerBound(data interface{}, compare func(int) int) int {
	rv := reflect.ValueOf(data)
	return lowerBound(compare, rv.Len())
}

func lowerBound(compare func(int) int, length int) int {
	var l = 0
	var r = length - 1
	for l < r {
		m := (r + l) / 2
		if compare(m) < 0 {
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
	return upperBound(data.Compare, data.Len())
}

//SliceUpperBound find the first index that data[index]>key with compare function
//the compare func should like
//	greater x>0 	equal x==0   less x<0
//	func(i int) (x int) {
//		if slice[i] > key {
//			return 1
//		} else if slice[i] == key {
//			return 0
//		} else {
//			return -1
//		}
//	}
//
// The function panics if the provided interface is not a slice.
func SliceUpperBound(data interface{}, compare func(int) int) int {
	rv := reflect.ValueOf(data)
	return upperBound(compare, rv.Len())
}

func upperBound(compare func(int) int, length int) int {
	var l = 0
	var r = length - 1
	for l < r {
		m := (r + l) / 2
		if compare(m) > 0 {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
