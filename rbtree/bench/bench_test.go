package rbtree_test

import (
	"sync"
	"testing"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

func BenchmarkSetInsert(t *testing.B) {
	var set = rbtree.NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		_, _ = set.Insert(rand.Int())
	}
}

func BenchmarkSetErase(t *testing.B) {
	var set = rbtree.NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = set.Erase(keys[i])
	}
}

func BenchmarkSetFind(t *testing.B) {
	var set = rbtree.NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = set.Find(keys[i])
	}
}

func BenchmarkMapInsert(t *testing.B) {
	var mp = rbtree.NewMap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		_, _ = mp.Insert(rbtree.NewPair(rand.Int(), true))
	}
}

func BenchmarkMapErase(t *testing.B) {
	var mp = rbtree.NewMap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(rbtree.NewPair(keys[i], true))
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = mp.Erase(keys[i])
	}
}

func BenchmarkMapFind(t *testing.B) {
	var mp = rbtree.NewMap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		_, _ = mp.Insert(rbtree.NewPair(keys[i], true))
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = mp.Find(keys[i])
	}
}

func BenchmarkSysHashMapInsert(t *testing.B) {
	var mp = make(map[int]bool)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		mp[rand.Int()] = true
	}
}

func BenchmarkSysHashMapErase(t *testing.B) {
	var mp = make(map[int]bool)
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		delete(mp, keys[i])
	}
}

func BenchmarkSysHashMapFind(t *testing.B) {
	var mp = make(map[int]bool)
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_, _ = mp[keys[i]]
	}
}

func BenchmarkSyncMapInsert(t *testing.B) {
	var mp = &sync.Map{}
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		mp.Store(rand.Int(), true)
	}
}

func BenchmarkSyncMapErase(t *testing.B) {
	var mp = &sync.Map{}
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		mp.Store(rand.Int(), true)
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		mp.Delete(keys[i])
	}
}

func BenchmarkSyncMapFind(t *testing.B) {
	var mp = &sync.Map{}
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		mp.Store(rand.Int(), true)
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_, _ = mp.Load(keys[i])
	}
}
