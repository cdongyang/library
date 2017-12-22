package rbtree_test

import (
	"sort"
	"sync"
	"testing"

	"github.com/cdongyang/library/rbtree"
)

func BenchmarkSort1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
	}
	sort.IntSlice(keys).Sort()
	memStats()
}

func BenchmarkSetInsert1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		})
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkSetErase1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewCustomSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.Iterator {
			return rbtree.NewSetNode(elem)
		},
		func(iter rbtree.Iterator) {
		},
	)
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Erase(keys[i])
	}
	memStats()
}

func BenchmarkSetInsertAndErase1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var keys = make([]int, b.N)
	var insertn, erasen int
	var set = rbtree.NewCustomSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.Iterator {
			return rbtree.NewSetNode(elem)
		},
		func(iter rbtree.Iterator) {
		},
	)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			set.Insert(keys[insertn])
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			set.Erase(keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSetInsertAndEraseWithPool1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var nodePool = sync.Pool{New: func() interface{} {
		return &rbtree.SetNode{}
	}}
	var keys = make([]int, b.N)
	var insertn, erasen int
	var set = rbtree.NewCustomSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.Iterator {
			var iter = nodePool.Get().(*rbtree.SetNode)
			iter.SetKey(elem)
			return iter
		},
		func(iter rbtree.Iterator) {
			nodePool.Put(iter)
		},
	)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			set.Insert(keys[insertn])
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			set.Erase(keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSysHashMapInsert1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var mp = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		mp[rand.Int()] = true
	}
	memStats()
}

func BenchmarkSysHashMapErase1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		delete(mp, keys[i])
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndErase1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var keys = make([]int, b.N/10)
	var insertn, erasen int
	var mp = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			mp[keys[insertn]] = true
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			delete(mp, keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndEraseWithBuf1E7(b *testing.B) {
	b.N = 1e7
	var rand = benchRand
	var keys = make([]int, b.N/10)
	var insertn, erasen int
	var mp = make(map[int]bool, b.N)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			mp[keys[insertn]] = true
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			delete(mp, keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}
