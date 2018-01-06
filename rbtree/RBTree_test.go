package rbtree

import (
	"fmt"
	"runtime"
	"sort"
	"testing"

	"github.com/cdongyang/library/randint"
)

type node = SetNode

//检查红黑树是否满足性质
func (t *RBTree) Check(root Iterator) (l int, size int) {
	if root == t.null {
		return 0, 0
	}
	if root.leftChild() != t.null {
		if root.getColor() == red && root.leftChild().getColor() == red {
			panic("left linked red node")
		} else if root.leftChild().getParent() != root {
			panic("left tree error")
		} else if root.leftChild().Compare(root.GetKey()) >= 0 { //unique set,can't equal
			panic("left order error")
		}
	}
	if root.rightChild() != t.null {
		if root.getColor() == red && root.rightChild().getColor() == red {
			panic("right linked red node")
		} else if root.rightChild().getParent() != root {
			panic("right tree error")
		} else if root.Compare(root.rightChild().GetKey()) >= 0 {
			panic("right order error")
		}
	}
	var a, s1 = t.Check(root.leftChild())
	var b, s2 = t.Check(root.rightChild())
	if a != b {
		panic("path length of black not equal")
	}
	if root.getColor() == black {
		return a + 1, s1 + s2 + 1
	}
	return a, s1 + s2 + 1
}

func testRBTree(t *testing.T, length int) {
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1000}
	var max = rand.Int()%length + 1
	var intSlice1K = make([]IntKey, length)
	for i := range intSlice1K {
		intSlice1K[i] = IntKey(rand.Int() % max)
	}
	tree := NewSet()
	var exists = make(map[IntKey]bool, len(intSlice1K))
	if !tree.Empty() {
		t.Fatal("empty")
	}
	//test Seter
	var _ Seter = tree
	var iter = tree.End()
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.GetKey() != intSlice1K[0] {
		if iter.Compare(intSlice1K[0]) < 0 {
			t.Fatal("LowerBound error", iter.GetKey(), intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.Compare(intSlice1K[0]) < 0 {
		t.Fatal("UpperBound error", iter.GetKey(), intSlice1K[0])
	}
	//test Insert method
	for _, val := range intSlice1K {
		ok := tree.Insert(val)
		if ok == exists[val] {
			panic("insert error")
		}
		exists[val] = true
		_, size := tree.Check(tree.root)
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}
	var sortSlice = make([]IntKey, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].Compare(sortSlice[j]) < 0
	})
	var uniqueN = 1
	for i := range sortSlice {
		if i > 0 && sortSlice[i] != sortSlice[i-1] {
			sortSlice[uniqueN] = sortSlice[i]
			uniqueN++
		}
	}
	sortSlice = sortSlice[:uniqueN]
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.Compare(intSlice1K[0]) != 0 {
		if iter.Compare(intSlice1K[0]) < 0 {
			t.Fatal("LowerBound error", iter.GetKey(), intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.Compare(intSlice1K[0]) < 0 {
		t.Fatal("UpperBound error", iter.GetKey(), intSlice1K[0])
	}
	//test Begin and EndNode method
	if tree.Begin().Compare(sortSlice[0]) != 0 {
		t.Fatal("begin error", tree.Begin().GetKey(), sortSlice[0])
	}
	if tree.EndNode().Compare(sortSlice[len(sortSlice)-1]) != 0 {
		t.Fatal("endNode error", tree.Begin().GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); it != tree.End(); it = it.Next(it) {
		if it.Compare(sortSlice[i]) != 0 {
			t.Fatal("go through error", it.GetKey(), sortSlice[i])
		}
		i++
	}
	//test EndNode and End and Last method
	i = len(sortSlice) - 1
	for it := tree.EndNode(); it != tree.End(); it = it.Last(it) {
		if it.Compare(sortSlice[i]) != 0 {
			t.Fatal("go back tree error", it.GetKey(), sortSlice[i])
		}
		i--
	}
	//test Find method
	iter = tree.Find(intSlice1K[0])
	if iter == tree.End() {
		t.Fatal("find error", intSlice1K[0])
	} else if iter.GetKey() != intSlice1K[0] {
		t.Fatal("find not equal", intSlice1K[0])
	}
	if tree.Find(IntKey(max)) != tree.End() {
		t.Fatal("find max error", max)
	}
	//test Erase method
	for _, val := range intSlice1K {
		ok := tree.Erase(val)
		if ok != exists[val] {
			t.Fatal("erase error")
		}
		delete(exists, val)
		_, size := tree.Check(tree.root)
		if size != tree.Size() {
			t.Fatal("size error")
		}
	}
	//test Empty method
	if !tree.Empty() {
		t.Fatal("empty error")
	}
	exists = make(map[IntKey]bool)
	for _, val := range intSlice1K {
		ok := tree.Insert(val)
		if ok == exists[val] {
			t.Fatal("insert error")
		}
		exists[val] = true
		_, size := tree.Check(tree.root)
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}
	//test Clear method
	tree.Clear()
	if tree.Size() != 0 || !tree.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}

func TestRBtree(t *testing.T) {
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}
	for i := 0; i < 100; i++ {
		testRBTree(t, rand.Int()%1000+1)
	}
}

var mem runtime.MemStats

func memStats() {
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

//BenchmarkNewSetNode-4   	300000000	         4.46 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNewSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewSetNode(IntKey(2))
	}
}

//BenchmarkNewElem-4   	10000000	       116 ns/op	      80 B/op	       1 allocs/o
func BenchmarkNewElem(b *testing.B) {
	var set = NewSet()
	for i := 0; i < b.N; i++ {
		_ = set.newElem(IntKey(0))
	}
}

//BenchmarkCompare-4   	100000000	        14.6 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompare(b *testing.B) {
	var a, c = IntKey(0), IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.Compare(c)
	}
}

//BenchmarkIteratorCompare-4   	50000000	        44.5 ns/op	       8 B/op	       1 allocs/op
func BenchmarkIteratorCompare(b *testing.B) {
	var a, c = NewSetNode(IntKey(0)), IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.Compare(c)
	}
}

//BenchmarkIteratorKeyCompare-4   	50000000	        41.4 ns/op	       8 B/op	       1 allocs/op
func BenchmarkIteratorKeyCompare(b *testing.B) {
	var a, c = NewSetNode(IntKey(0)), IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.GetKey().Compare(c)
	}
}

//BenchmarkIteratorKeyerCompare-4   	300000000	         4.69 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIteratorKeyerCompare(b *testing.B) {
	var a = NewSetNode(IntKey(0))
	var c Keyer = IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.Compare(c)
	}
}

//BenchmarkIteratorInterfaceKeyCompare-4   	200000000	         7.41 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIteratorInterfaceKeyCompare(b *testing.B) {
	var a Iterator = NewSetNode(IntKey(0))
	var c Keyer = IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.GetKey().Compare(c)
	}
}

//BenchmarkIteratorInterfaceKeyerCompare-4   	200000000	         8.42 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIteratorInterfaceKeyerCompare(b *testing.B) {
	var a Iterator = NewSetNode(IntKey(0))
	var c Keyer = IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.Compare(c)
	}
}

func BenchmarkInterfaceAssertTointerface(b *testing.B) {
	var a Keyer = IntKey(1)
	for i := 0; i < b.N; i++ {
		_ = a.(interface{})
	}
}
