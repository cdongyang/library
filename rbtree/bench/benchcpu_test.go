package rbtree

import (
	"sync"
	"testing"
	"time"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}
var hmap = make(map[int]bool)
var smap = &sync.Map{}
var set = rbtree.NewSet(func(a, b interface{}) int {
	return a.(int) - b.(int)
})
var tmap = rbtree.NewMap(func(a, b interface{}) int {
	return a.(int) - b.(int)
})
var keys []int

func testBench(t *testing.T) {
	var rand = benchRand
	var fun = [](func(int)){
		benchSysHashMapInsert,
		benchSysHashMapFind,
		benchSysHashMapErase,
		benchSyncMapInsert,
		benchSyncMapFind,
		benchSyncMapErase,
		benchSetInsert,
		benchSetFind,
		benchSetErase,
		benchMapInsert,
		benchMapFind,
		benchMapErase,
	}
	for l := 100; l <= int(1e7); l *= 10 {
		hmap = make(map[int]bool)
		smap = &sync.Map{}
		set = rbtree.NewSet(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})
		tmap = rbtree.NewMap(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})
		keys = make([]int, l)
		for i := 0; i < len(keys); i++ {
			keys[i] = rand.Int()
		}
		for _, f := range fun {
			start := time.Now()
			f(len(keys))
			dur := time.Since(start)
			t.Log(l, "op", int(dur)/len(keys), "ns/op")
		}
	}
}

func benchSysHashMapInsert(n int) {
	for i := 0; i < n; i++ {
		hmap[keys[i]] = true
	}
}

func benchSysHashMapFind(n int) {
	for i := 0; i < n; i++ {
		_, _ = hmap[keys[i]]
	}
}

func benchSysHashMapErase(n int) {
	for i := 0; i < n; i++ {
		delete(hmap, keys[i])
	}
}

func benchSyncMapInsert(n int) {
	for i := 0; i < n; i++ {
		smap.Store(keys[i], true)
	}
}

func benchSyncMapFind(n int) {
	for i := 0; i < n; i++ {
		_, _ = smap.Load(keys[i])
	}
}

func benchSyncMapErase(n int) {
	for i := 0; i < n; i++ {
		smap.Delete(keys[i])
	}
}

func benchSetInsert(n int) {
	for i := 0; i < n; i++ {
		_, _ = set.Insert(keys[i])
	}
}

func benchSetFind(n int) {
	for i := 0; i < n; i++ {
		_ = set.Find(keys[i])
	}
}

func benchSetErase(n int) {
	for i := 0; i < n; i++ {
		_ = set.Erase(keys[i])
	}
}

func benchMapInsert(n int) {
	for i := 0; i < n; i++ {
		_, _ = tmap.Insert(rbtree.NewPair(keys[i], true))
	}
}

func benchMapFind(n int) {
	for i := 0; i < n; i++ {
		_ = tmap.Find(keys[i])
	}
}

func benchMapErase(n int) {
	for i := 0; i < n; i++ {
		_ = tmap.Erase(keys[i])
	}
}
