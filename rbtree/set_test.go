package rbtree_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"

	"github.com/cdongyang/library/rbtree"
)

func TestSet(t *testing.T) {
	var intSlice = []int{1, 2, 4, 3, 6, 3, 4, 7, 2, 3, 5, 9}
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.SetIterator {
			return rbtree.NewSetNode(elem)
		},
		func(iter rbtree.SetIterator) {
		},
	)
	for _, val := range intSlice {
		ok := set.Insert(val)
		fmt.Println("insert", val, ok)
	}
	for _, val := range intSlice {
		ok := set.Erase(val)
		fmt.Println("erase", val, ok)
	}
}

var mem runtime.MemStats

func memStats() {
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

func BenchmarkSetInsert(b *testing.B) {
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.SetIterator {
			return rbtree.NewSetNode(elem)
		},
		func(iter rbtree.SetIterator) {
		},
	)
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkSetInsertWithPool(b *testing.B) {
	var nodePool = &sync.Pool{New: func() interface{} {
		return &rbtree.SetNode{}
	}}
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.SetIterator {
			var iter = nodePool.Get().(*rbtree.SetNode)
			iter.SetKey(elem)
			return iter
		},
		func(iter rbtree.SetIterator) {
			nodePool.Put(iter)
		},
	)
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkSetInsertWithArr(b *testing.B) {
	var arr [1 << 20]rbtree.SetNode
	var num = 0
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.SetIterator {
			if num >= (1 << 20) {
				return rbtree.NewSetNode(elem)
			}
			arr[num].SetKey(elem)
			num++
			return &arr[num-1]
		},
		func(iter rbtree.SetIterator) {
		},
	)
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}
