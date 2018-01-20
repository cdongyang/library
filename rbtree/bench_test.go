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

func BenchmarkSetInsert(b *testing.B) {
	b.Run("setInsert", runWith(benchmarkSetInsert, 1e5))
}

func BenchmarkSetErase(b *testing.B) {
	b.Run("setErase", runWith(benchmarkSetErase, 1e5))
}

func BenchmarkSetFind(b *testing.B) {
	b.Run("setFind", runWith(benchmarkSetFind, 1e5))
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
