package rbtree_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

func TestSet(t *testing.T) {
	type IntKey = rbtree.IntKey
	var intSlice = []IntKey{1, 2, 4, 5, 3, 4, 3, 5, 8, 3, 6, 3, 4, 2, 3, 5, 9}
	var set = rbtree.NewSet()
	var exists = make(map[IntKey]bool)
	for _, val := range intSlice {
		ok := set.Insert(val)
		if exists[val] == ok {
			t.Fatal("set insert error", val)
		}
		exists[val] = true
	}
	if set.Find(IntKey(7)) != set.End() {
		t.Fatal("find error:", set.Find(IntKey(7)).GetKey())
	}
	if set.Find(IntKey(8)) == set.End() {
		t.Fatal("find error:", 8)
	}
	if set.LowerBound(IntKey(7)).Compare(IntKey(8)) != 0 {
		t.Fatal("lowerBound error:", set.LowerBound(IntKey(7)).GetKey())
	}
	if set.LowerBound(IntKey(6)).Compare(IntKey(6)) != 0 {
		t.Fatal("lowerBound error:", set.LowerBound(IntKey(6)).GetKey())
	}
	if set.UpperBound(IntKey(5)).Compare(IntKey(6)) != 0 {
		t.Fatal("upperBound error:", set.LowerBound(IntKey(5)).GetKey())
	}
	if set.UpperBound(IntKey(6)).Compare(IntKey(8)) != 0 {
		t.Fatal("upperBound error:", set.UpperBound(IntKey(6)).GetKey())
	}
	if set.Begin().Compare(IntKey(1)) != 0 {
		t.Fatal("begin error:", set.Begin().GetKey())
	}
	if set.EndNode().Compare(IntKey(9)) != 0 {
		t.Fatal("endNode error:", set.EndNode().GetKey())
	}
	for i, val := range intSlice {
		ok := set.Erase(val)
		if exists[val] != ok {
			t.Fatal("set erase error", i, val)
		}
		delete(exists, val)
	}
	var _ rbtree.Seter = set
}

var mem runtime.MemStats
var outmem = false

func memStats() {
	if !outmem {
		return
	}
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}
