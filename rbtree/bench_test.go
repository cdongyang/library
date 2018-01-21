package rbtree_test

import (
	"runtime/debug"
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

// BenchmarkSet/setInsert/0-4               2000000               687 ns/op              72 B/op          2 allocs/op
// BenchmarkSet/setErase/0-4                5000000               330 ns/op               0 B/op          0 allocs/op
// BenchmarkSet/setFind/0-4                 5000000               292 ns/op               0 B/op          0 allocs/op
func BenchmarkSet(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 0))
	b.Run("setErase", runWith(benchmarkSetErase, 0))
	b.Run("setFind", runWith(benchmarkSetFind, 0))
}

func BenchmarkSet1E5(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e5))
	b.Run("setErase", runWith(benchmarkSetErase, 1e5))
	b.Run("setFind", runWith(benchmarkSetFind, 1e5))
}

func BenchmarkSet1E6(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e6))
	b.Run("setErase", runWith(benchmarkSetErase, 1e6))
	b.Run("setFind", runWith(benchmarkSetFind, 1e6))
}

// BenchmarkHashMap/hashMapInsert/0-4               3000000               345 ns/op              38 B/op          0 allocs/op
// BenchmarkHashMap/hashMapErase/0-4               10000000               166 ns/op               0 B/op          0 allocs/op
// BenchmarkHashMap/hashMapFind/0-4                20000000               144 ns/op               0 B/op          0 allocs/op
func BenchmarkHashMap(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 0))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 0))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 0))
}

func BenchmarkHashMap1E5(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 1e5))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 1e5))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 1e5))
}

func BenchmarkHashMap1E6(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 1e6))
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 1e6))
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 1e6))
}

// BenchmarkMap/mapInsert/0-4               2000000               876 ns/op             120 B/op          3 allocs/op
// BenchmarkMap/mapErase/0-4                5000000               345 ns/op               0 B/op          0 allocs/op
// BenchmarkMap/mapFind/0-4                 5000000               391 ns/op               0 B/op          0 allocs/op
func BenchmarkMap(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 0))
	b.Run("mapErase", runWith(benchmarkSetErase, 0))
	b.Run("mapFind", runWith(benchmarkMapFind, 0))
}

func BenchmarkMap1E5(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 1e5))
	b.Run("mapErase", runWith(benchmarkSetErase, 1e5))
	b.Run("mapFind", runWith(benchmarkMapFind, 1e5))
}

func BenchmarkMap1E6(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 1e6))
	b.Run("mapErase", runWith(benchmarkSetErase, 1e6))
	b.Run("mapFind", runWith(benchmarkMapFind, 1e6))
}

func BenchmarkSetInsert(b *testing.B) {
	debug.SetGCPercent(-1)
	b.Run("setInsert", runWith(benchmarkSetInsert, 0))
	debug.SetGCPercent(100)
}

func BenchmarkSetErase(b *testing.B) {
	debug.SetGCPercent(-1)
	b.Run("setErase", runWith(benchmarkSetErase, 0))
	debug.SetGCPercent(100)
}

func BenchmarkSetFind(b *testing.B) {
	debug.SetGCPercent(-1)
	b.Run("setFind", runWith(benchmarkSetFind, 0))
	debug.SetGCPercent(100)
}

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}

func benchmarkSetInsert(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(rbtree.CompareInt)
	var rand = benchRand
	b.ResetTimer()
	//rbtree.GetNodeCount = 0
	//rbtree.GetSetKeyPointerCount = 0
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = set.Insert(key)
	}
	//b.Log(b.N, rbtree.GetNodeCount, rbtree.GetSetKeyPointerCount)
}

func benchmarkSetErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(rbtree.CompareInt)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	//rbtree.GetNodeCount = 0
	//rbtree.GetSetKeyPointerCount = 0
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = set.Erase(key)
	}
	//b.Log(b.N, rbtree.GetNodeCount, rbtree.GetSetKeyPointerCount)
}

func benchmarkSetFind(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = rbtree.NewSet(rbtree.CompareInt)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	//rbtree.GetNodeCount = 0
	//rbtree.GetSetKeyPointerCount = 0
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_ = set.Find(key)
	}
	//b.Log(b.N, rbtree.GetNodeCount, rbtree.GetSetKeyPointerCount)
}

// 4 allocs/op on go 1.7
// 3 allocs/op on go 1.9.2
func benchmarkMapInsert(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = rbtree.NewMap(rbtree.CompareInt)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = mp.Insert(rbtree.Pair{Key: key, Value: true})
	}
}

func benchmarkMapErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = rbtree.NewMap(rbtree.CompareInt)
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
	if n != 0 {
		b.N = n
	}
	var mp = rbtree.NewMap(rbtree.CompareInt)
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
	if n != 0 {
		b.N = n
	}
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
	if n != 0 {
		b.N = n
	}
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
	if n != 0 {
		b.N = n
	}
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
