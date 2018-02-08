package rbtree_test

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"testing"

	"github.com/cdongyang/library/randint"
)

var mem runtime.MemStats
var outmem = true

func memStats() {
	if !outmem {
		return
	}
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

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
	b.Run("setInsert", runWith(benchmarkSetInsert, 0))
	memStats()
	b.Run("setErase", runWith(benchmarkSetErase, 0))
	memStats()
	b.Run("setFind", runWith(benchmarkSetFind, 0))
	memStats()
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

func BenchmarkHashMap(b *testing.B) {
	b.Run("hashMapInsert", runWith(benchmarkHashMapInsert, 0))
	memStats()
	b.Run("hashMapErase", runWith(benchmarkHashMapErase, 0))
	memStats()
	b.Run("hashMapFind", runWith(benchmarkHashMapFind, 0))
	memStats()
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

func BenchmarkMap(b *testing.B) {
	b.Run("mapInsert", runWith(benchmarkMapInsert, 0))
	memStats()
	b.Run("mapErase", runWith(benchmarkSetErase, 0))
	memStats()
	b.Run("mapFind", runWith(benchmarkMapFind, 0))
	memStats()
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
	var set = NewSet(true)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = set.Insert(key)
	}
}

func benchmarkSetErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var set = NewSet(true)
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
	if n != 0 {
		b.N = n
	}
	var set = NewSet(true)
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
	if n != 0 {
		b.N = n
	}
	var mp = NewMap(true)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		_, _ = mp.Insert(key, key*2)
	}
}

func benchmarkMapErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = NewMap(true)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(keys[i], keys[i]*2)
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
	var mp = NewMap(true)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(keys[i], keys[i]*2)
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
	var mp = make(map[int]int)
	var rand = benchRand
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = rand.Int()
		mp[key] = key * 2
	}
}

func benchmarkHashMapErase(b *testing.B, n int) {
	if n != 0 {
		b.N = n
	}
	var mp = make(map[int]int)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = keys[i] * 2
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
	var mp = make(map[int]int)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = keys[i] * 2
	}
	b.ResetTimer()
	var key int
	for i := 0; i < b.N; i++ {
		key = keys[i]
		_, _ = mp[key]
	}
}
