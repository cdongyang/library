package rbtree_test

import (
	"strconv"
	"testing"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

func BenchmarkAll(b *testing.B) {
	var ns = []int{1e3, 1e4, 1e5, 1e6, 1e7}
	b.Run("setInsert", runWith(benchmarkSetInsert, ns...))
	b.Run("setErase", runWith(benchmarkSetErase, ns...))
	b.Run("setFind", runWith(benchmarkSetFind, ns...))
	b.Run("mapInsert", runWith(benchmarkMapInsert, ns...))
	b.Run("mapErase", runWith(benchmarkSetErase, ns...))
	b.Run("mapFind", runWith(benchmarkMapFind, ns...))
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, ns...))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, ns...))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, ns...))
}

func BenchmarkSet(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e5))
	b.Run("setErase", runWith(benchmarkSetErase, 1e5))
	b.Run("setFind", runWith(benchmarkSetFind, 1e5))
}

func BenchmarkHashMap(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 1e5))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 1e5))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 1e5))
}

func BenchmarkMap(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 1e5))
	b.Run("mapErase", runWith(benchmarkSetErase, 1e5))
	b.Run("mapFind", runWith(benchmarkMapFind, 1e5))
}

func BenchmarkSetInsert(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e5))
}

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}

func benchmarkSetInsert(b *testing.B, n int) {
	b.N = n
	var set = rbtree.NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = set.Insert(key)
	}
}

func benchmarkSetErase(b *testing.B, n int) {
	b.N = n
	var set = rbtree.NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = set.Erase(key)
	}
}

func benchmarkSetFind(b *testing.B, n int) {
	b.N = n
	var set = rbtree.NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = set.Find(key)
	}
}

func benchmarkMapInsert(b *testing.B, n int) {
	b.N = n
	var mp = rbtree.NewMap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = mp.Insert(rbtree.Pair{Key: key, Value: true})
	}
}

func benchmarkMapErase(b *testing.B, n int) {
	b.N = n
	var mp = rbtree.NewMap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(rbtree.Pair{Key: keys[i], Value: true})
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = mp.Erase(key)
	}
}

func benchmarkMapFind(b *testing.B, n int) {
	b.N = n
	var mp = rbtree.NewMap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(rbtree.Pair{Key: keys[i], Value: true})
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = mp.Find(key)
	}
}

func benchmarkHashMapInsert(b *testing.B, n int) {
	b.N = n
	var mp = make(map[int]bool)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		mp[key] = true
	}
}

func benchmarkHashMapErase(b *testing.B, n int) {
	b.N = n
	var mp = make(map[int]bool)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		delete(mp, key)
	}
}

func benchmarkHashMapFind(b *testing.B, n int) {
	b.N = n
	var mp = make(map[int]bool)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_, _ = mp[key]
	}
}

// BenchmarkAll/setInsert/1000-4               1000              1041 ns/op             104 B/op          2 allocs/op
// BenchmarkAll/setInsert/10000-4             10000              1084 ns/op             104 B/op          2 allocs/op
// BenchmarkAll/setInsert/100000-4           100000              1312 ns/op             104 B/op          2 allocs/op
// BenchmarkAll/setInsert/1000000-4         1000000              1357 ns/op             104 B/op          2 allocs/op
// BenchmarkAll/setInsert/10000000-4               10000000              1563 ns/op             104 B/op          2 allocs/op
// BenchmarkAll/setErase/1000-4                        1000               627 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setErase/10000-4                      10000               709 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setErase/100000-4                    100000               695 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setErase/1000000-4                  1000000              1332 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setErase/10000000-4                10000000              1087 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setFind/1000-4                         1000               288 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setFind/10000-4                       10000               382 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setFind/100000-4                     100000               448 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setFind/1000000-4                   1000000              1190 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/setFind/10000000-4                 10000000               753 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapInsert/1000-4                       1000              1088 ns/op             152 B/op          3 allocs/op
// BenchmarkAll/mapInsert/10000-4                     10000              1244 ns/op             152 B/op          3 allocs/op
// BenchmarkAll/mapInsert/100000-4                   100000              1546 ns/op             152 B/op          3 allocs/op
// BenchmarkAll/mapInsert/1000000-4                 1000000              1378 ns/op             152 B/op          3 allocs/op
// BenchmarkAll/mapInsert/10000000-4               10000000              1509 ns/op             152 B/op          3 allocs/op
// BenchmarkAll/mapErase/1000-4                        1000               634 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapErase/10000-4                      10000               647 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapErase/100000-4                    100000               885 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapErase/1000000-4                  1000000              1022 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapErase/10000000-4                10000000              1181 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapFind/1000-4                         1000               465 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapFind/10000-4                       10000               506 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapFind/100000-4                     100000               631 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapFind/1000000-4                   1000000              1085 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/mapFind/10000000-4                 10000000              1751 ns/op               8 B/op          1 allocs/op
// BenchmarkAll/hashMapInsert/1000-4                   1000               345 ns/op              53 B/op          0 allocs/op
// BenchmarkAll/hashMapInsert/10000-4                 10000               265 ns/op              43 B/op          0 allocs/op
// BenchmarkAll/hashMapInsert/100000-4               100000               261 ns/op              36 B/op          0 allocs/op
// BenchmarkAll/hashMapInsert/1000000-4             1000000               352 ns/op              54 B/op          0 allocs/op
// BenchmarkAll/hashMapInsert/10000000-4           10000000               332 ns/op              44 B/op          0 allocs/op
// BenchmarkAll/hashMapErase/1000-4                    1000                57.4 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapErase/10000-4                  10000                57.9 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapErase/100000-4                100000                80.0 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapErase/1000000-4              1000000               148 ns/op               0 B/op          0 allocs/op
// BenchmarkAll/hashMapErase/10000000-4            10000000               173 ns/op               0 B/op          0 allocs/op
// BenchmarkAll/hashMapFind/1000-4                     1000                43.9 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapFind/10000-4                   10000                60.4 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapFind/100000-4                 100000                77.9 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapFind/1000000-4               1000000                95.2 ns/op             0 B/op          0 allocs/op
// BenchmarkAll/hashMapFind/10000000-4             10000000               119 ns/op               0 B/op          0 allocs/op
// BenchmarkSetInsert/setInsert/100000-4             100000              1494 ns/op             104 B/op          2 allocs/op
// BenchmarkSetErase/setErase/100000-4               100000               774 ns/op               8 B/op          1 allocs/op
// BenchmarkSetFind/setFind/100000-4                 100000               469 ns/op               8 B/op          1 allocs/op
