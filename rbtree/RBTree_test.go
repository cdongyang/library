package rbtree

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

type node struct {
	RBTreeNode
	key int
}

func (n *node) GetKey() interface{} {
	return n.key
}
func (n *node) SetKey(k interface{}) {
	n.key = k.(int)
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

func ExampleInterface() {
	var null Iterator
	var root *node
	var iter Iterator = root
	var NULL = root
	fmt.Println(iter, root)
	fmt.Println(iter == nil, root == nil, iter == root)
	fmt.Println(reflect.DeepEqual(iter, root), reflect.DeepEqual(iter, nil))
	fmt.Println(iter == null, root == null, null == nil)
	fmt.Println(iter == NULL, root == NULL, NULL == nil)
	fmt.Println(iter == (*node)(nil))
	switch iter.(type) {
	case *node:
		fmt.Println("*node")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("other")
	}
	//== judge type first,then judge value
	// Output:
	//<nil> <nil>
	//false true true
	//true false
	//false false true
	//true true true
	//true
	//*node
}

func testRBTree(t *testing.T, length int) {
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
	tree := NewRBTree((*node)(nil), compare, newElem, deleteElem)
	var exists = make(map[int]bool, len(intSlice1K))
	if !tree.Empty() {
		panic("empty")
	}
	var iter = tree.End()
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.GetKey() != intSlice1K[0] {
		if iter.GetKey().(int) < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.GetKey(), intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.GetKey().(int) < intSlice1K[0] {
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
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i] < sortSlice[j]
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
	} else if iter.GetKey() != intSlice1K[0] {
		if iter.GetKey().(int) < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.GetKey(), intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.GetKey().(int) < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.GetKey(), intSlice1K[0])
	}
	//test Begin and EndNode method
	if tree.Begin().GetKey() != sortSlice[0] {
		t.Fatal("begin error", tree.Begin().GetKey(), sortSlice[0])
	}
	if tree.EndNode().GetKey() != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", tree.Begin().GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); it != tree.End(); it = it.Next(it) {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go through error", it.GetKey(), sortSlice[i])
		}
		i++
	}
	//test EndNode and End and Last method
	i = len(sortSlice) - 1
	for it := tree.EndNode(); it != tree.End(); it = it.Last(it) {
		if it.GetKey() != sortSlice[i] {
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
	tree.Clear(&tree.root)
	if tree.Size() != 0 || !tree.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}

func TestRBtree(t *testing.T) {
	/*for i := 0; i < 100; i++ {
		testRBTree(t, rand.Int()%1000+1)
	}*/
}

var mem runtime.MemStats

func memStats() {
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}
func BenchmarkRBTree(t *testing.B) {
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
	tree := NewRBTree((*node)(nil), compare, newElem, deleteElem)
	for i := 0; i < t.N; i++ {
		tree.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkRBTreeWithPool(t *testing.B) {
	var (
		/*nodePool = &sync.Pool{New: func() interface{} {
			return &node{}
		}}*/
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeArr [1 << 20]node
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= (1 << 20) {
					return &node{key: e}
				}
				nodeArr[num].key = e
				num++
				return &nodeArr[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := NewRBTree((*node)(nil), compare, newElem, deleteElem)
	for i := 0; i < t.N; i++ {
		tree.Insert(rand.Int())
	}
	memStats()
}
