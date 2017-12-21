package rbtree_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/cdongyang/library/rbtree"
)

func TestSet(t *testing.T) {
	var intSlice = []int{1, 2, 4, 5, 3, 4, 3, 5, 8, 3, 6, 3, 4, 7, 2, 3, 5, 9}
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		})
	var exists = make(map[int]bool)
	for _, val := range intSlice {
		ok := set.Insert(val)
		if exists[val] == ok {
			t.Fatal("set insert error", val)
		}
		exists[val] = true
	}
	for _, val := range intSlice {
		ok := set.Erase(val)
		if exists[val] != ok {
			t.Fatal("set erase error", val)
		}
		delete(exists, val)
	}
	var _ rbtree.Seter = set
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
		})
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkSetInsertWithArr(b *testing.B) {
	var arr = make([]rbtree.SetNode, b.N)
	var num = 0
	var set = rbtree.NewCustomSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.Iterator {
			if num >= b.N {
				return rbtree.NewSetNode(elem)
			}
			arr[num].SetKey(elem)
			num++
			return &arr[num-1]
		},
		func(iter rbtree.Iterator) {
		},
	)
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}
