package algorithm_test

import "testing"
import "fmt"
import . "github.com/cdongyang/library/algorithm"

var intSlice = []int{1, 2, 3, 3, 3, 4, 4, 5}

func ExampleSliceLowerBound() {
	id := SliceLowerBound(intSlice, func(i int) bool {
		return intSlice[i] < 3
	})
	fmt.Println(id, intSlice[id])
	// Output: 2 3
}
func ExampleLowerBound() {
	id := LowerBound(SearchIntSlice{intSlice, 3})
	fmt.Println(id, intSlice[id])
	// Output: 2 3
}

func ExampleUpperBound() {
	id := UpperBound(SearchIntSlice{intSlice, 3})
	fmt.Println(id, intSlice[id])
	// Output: 5 4
}
func ExampleSliceUpperBound() {
	id := SliceUpperBound(intSlice, func(i int) bool {
		return intSlice[i] > 3
	})
	fmt.Println(id, intSlice[id])
	// Output: 5 4
}

func ExampleStructLowerBound() {
	id := SearchIntSlice{intSlice, 3}.LowerBound()
	fmt.Println(id, intSlice[id])
	id = intLowerBound(intSlice, len(intSlice), 3)
	fmt.Println(id, intSlice[id])
	id = SearchIntSlice{intSlice, 3}.UpperBound()
	fmt.Println(id, intSlice[id])
	id = intUpperBound(intSlice, len(intSlice), 3)
	fmt.Println(id, intSlice[id])
	// Output:
	//2 3
	//2 3
	//5 4
	//5 4
}

var intSlice1M = make([]int, 1<<20)
var intSlice1MKey int = 23456

func init() {
	for i := range intSlice1M {
		intSlice1M[i] = i
	}
}

//goos: linux
//goarch: amd64
//pkg: github.com/cdongyang/library/algorithm
//BenchmarkLowerBound-4                    5000000               364 ns/op              32 B/op          1 allocs/op
//BenchmarkSliceLowerBound-4              10000000               246 ns/op              32 B/op          1 allocs/op
//BenchmarkSearchStructLowerBound-4       10000000               119 ns/op               0 B/op          0 allocs/op
//BenchmarkIntLowerBound-4                30000000                56.7 ns/op             0 B/op          0 allocs/op
//BenchmarkUpperBound-4                    5000000               322 ns/op              32 B/op          1 allocs/op
//BenchmarkSliceUpperBound-4              10000000               248 ns/op              32 B/op          1 allocs/op
//BenchmarkSearchStructUpperBound-4       10000000               162 ns/op               0 B/op          0 allocs/op
//BenchmarkIntUpperBound-4                30000000                52.1 ns/op             0 B/op          0 allocs/op

func BenchmarkLowerBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		LowerBound(SearchIntSlice{intSlice1M, 3})
	}
}

func BenchmarkSliceLowerBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SliceLowerBound(intSlice1M, func(i int) bool {
			return intSlice1M[i] < intSlice1MKey
		})
	}
}

func BenchmarkSearchStructLowerBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SearchIntSlice{intSlice1M, intSlice1MKey}.LowerBound()
	}
}

func BenchmarkIntLowerBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		intLowerBound(intSlice1M, len(intSlice1M), intSlice1MKey)
	}
}

func intLowerBound(data []int, length int, key int) int {
	var l = 0
	var r = length - 1
	for l < r {
		m := (r + l) / 2
		if data[m] < key {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func BenchmarkUpperBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		UpperBound(SearchIntSlice{intSlice1M, 3})
	}
}

func BenchmarkSliceUpperBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SliceUpperBound(intSlice1M, func(i int) bool {
			return intSlice1M[i] > intSlice1MKey
		})
	}
}

func BenchmarkSearchStructUpperBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SearchIntSlice{intSlice1M, intSlice1MKey}.UpperBound()
	}
}

func BenchmarkIntUpperBound(t *testing.B) {
	for i := 0; i < t.N; i++ {
		intUpperBound(intSlice1M, len(intSlice1M), intSlice1MKey)
	}
}
func intUpperBound(data []int, length int, key int) int {
	var l = 0
	var r = length - 1
	for l < r {
		m := (r + l) / 2
		if data[m] > key {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
