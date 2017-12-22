package rbtree_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

func TestSet(t *testing.T) {
	var intSlice = []int{1, 2, 4, 5, 3, 4, 3, 5, 8, 3, 6, 3, 4, 2, 3, 5, 9}
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
	if set.Find(7) != set.End() {
		t.Fatal("find error:", set.Find(7).GetKey())
	}
	if set.Find(8) == set.End() {
		t.Fatal("find error:", 8)
	}
	if set.LowerBound(7).GetKey() != 8 {
		t.Fatal("lowerBound error:", set.LowerBound(7).GetKey())
	}
	if set.LowerBound(6).GetKey() != 6 {
		t.Fatal("lowerBound error:", set.LowerBound(6).GetKey())
	}
	if set.UpperBound(5).GetKey() != 6 {
		t.Fatal("upperBound error:", set.LowerBound(5).GetKey())
	}
	if set.UpperBound(6).GetKey() != 8 {
		t.Fatal("upperBound error:", set.UpperBound(6).GetKey())
	}
	if set.Begin().GetKey() != 1 {
		t.Fatal("begin error:", set.Begin().GetKey())
	}
	if set.EndNode().GetKey() != 9 {
		t.Fatal("endNode error:", set.EndNode().GetKey())
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

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}
