package rbtree

import (
	"fmt"
	"runtime"
	"sort"
	"testing"

	"github.com/cdongyang/library/randint"
)

type node struct {
	RBTreeNode
	key int
}

func (n *node) Copy(des, src Iterator) {
	des.(*node).key = src.(*node).key
}
func (t *RBTree) Check(root Iterator) (l int, size int) {
	if root == t.null {
		return 0, 0
	}
	if root.leftChild() != t.null {
		if root.getColor() == red && root.leftChild().getColor() == red {
			panic("left linked red node")
		} else if root.leftChild().getParent() != root {
			panic("left tree error")
		} else if t.compare(root.leftChild(), root) >= 0 { //unique set,can't equal
			panic("left order error")
		}
	}
	if root.rightChild() != t.null {
		if root.getColor() == red && root.rightChild().getColor() == red {
			panic("right linked red node")
		} else if root.rightChild().getParent() != root {
			panic("right tree error")
		} else if t.compare(root, root.rightChild()) >= 0 {
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
	var intSlice1K = make([]int, length)
	for i := range intSlice1K {
		intSlice1K[i] = rand.Int() % max
	}
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var exists = make(map[int]bool, len(intSlice1K))
	if !tree.Empty() {
		t.Fatal("empty")
	}
	//test RBTreer
	var _ RBTreer = tree
	var iter = tree.End()
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).key != intSlice1K[0] {
		if iter.(*node).key < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.(*node).key, intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).key < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.(*node).key, intSlice1K[0])
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
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.IntSlice(sortSlice).Sort()
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
	} else if iter.(*node).key != intSlice1K[0] {
		if iter.(*node).key < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.(*node).key, intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).key < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.(*node).key, intSlice1K[0])
	}
	//test Begin and EndNode method
	if tree.Begin().(*node).key != sortSlice[0] {
		t.Fatal("begin error", tree.Begin().(*node).key, sortSlice[0])
	}
	if tree.EndNode().(*node).key != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", tree.Begin().(*node).key, sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); it != tree.End(); it = it.Next(it) {
		if it.(*node).key != sortSlice[i] {
			t.Fatal("go through error", it.(*node).key, sortSlice[i])
		}
		i++
	}
	//test EndNode and End and Last method
	i = len(sortSlice) - 1
	for it := tree.EndNode(); it != tree.End(); it = it.Last(it) {
		if it.(*node).key != sortSlice[i] {
			t.Fatal("go back tree error", it.(*node).key, sortSlice[i])
		}
		i--
	}
	//test Find method
	iter = tree.Find(intSlice1K[0])
	if iter == tree.End() {
		t.Fatal("find error", intSlice1K[0])
	} else if iter.(*node).key != intSlice1K[0] {
		t.Fatal("find not equal", intSlice1K[0])
	}
	if tree.Find(max) != tree.End() {
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
	exists = make(map[int]bool)
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

/*
func BenchmarkRBTreeInsert1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	for i := 0; i < b.N; i++ {
		tree.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkRBTreeInsertWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeBuf = make([]node, b.N)
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= b.N {
					return &node{key: e}
				}
				nodeBuf[num].key = e
				num++
				return &nodeBuf[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	for i := 0; i < b.N; i++ {
		tree.Insert(rand.Int())
	}
	memStats()
}
func BenchmarkRBTreeErase1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	b.StopTimer()
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}
func BenchmarkRBTreeEraseWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	b.StopTimer()
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeBuf = make([]node, b.N)
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= b.N {
					return &node{key: e}
				}
				nodeBuf[num].key = e
				num++
				return &nodeBuf[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}

func BenchmarkRBTreeInsertAndErase1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}
func BenchmarkRBTreeInsertAndEraseWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeBuf = make([]node, b.N)
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= b.N {
					return &node{key: e}
				}
				nodeBuf[num].key = e
				num++
				return &nodeBuf[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}

*/
