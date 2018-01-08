package rbtree

import (
	"sort"
	"testing"

	"github.com/cdongyang/library/randint"
)

func (t colorType) String() string {
	if t == red {
		return "red"
	}
	return "black"
}

type node struct {
	RBTreeNode
	data int
}

func (n *node) Next() Iterator {
	return n.tree.Next(n)
}

func (n *node) Last() Iterator {
	return n.tree.Last(n)
}

func (n *node) GetData() interface{} {
	return n.data
}

func (n *node) GetKey() interface{} {
	return n.data
}

func (n *node) CopyData(src Iterator) {
	n.data = src.(*node).data
}

func (t *RBTree) leftmost() Iterator {
	var root = t.root()
	if t.sameIterator(root, t.End()) {
		return root
	}
	for !t.sameIterator(root.getChild(0), t.End()) {
		root = root.getChild(0)
	}
	return root
}

func (t *RBTree) rightmost() Iterator {
	var root = t.root()
	if t.sameIterator(root, t.End()) {
		return root
	}
	for !t.sameIterator(root.getChild(1), t.End()) {
		root = root.getChild(1)
	}
	return root
}

//check wheather the tree satisfy the rule of red-black tree
func (t *RBTree) Check() (int, int) {
	if !t.sameIterator(t.root(), t.End()) && t.root().getColor() != black {
		panic("not black root")
	}
	return t.check(t.root())
}
func (t *RBTree) check(root Iterator) (l int, size int) {
	if t.sameIterator(root, t.End()) {
		return 0, 0
	}
	//fmt.Printf("l:%p r:%p p:%p s:%p c:%s v:%d\n", root.getChild(0), root.getChild(1), root.getParent(), root, root.getColor().String(), root.GetData())
	for i := 0; i < 2; i++ {
		if !t.sameIterator(root.getChild(i), t.End()) {
			if root.getColor() == red && root.getChild(i).getColor() == red {
				panic("linked red node")
			} else if !t.sameIterator(root.getChild(i).getParent(), root) {
				panic("tree error")
			} else if i == 0 && t.compare(root.getChild(i).GetData(), root.GetData()) > 0 {
				panic("order error")
			} else if i == 1 && t.compare(root.getChild(i).GetData(), root.GetData()) < 0 {
				panic("order error")
			} else if t.compare(root.getChild(i).GetData(), root.GetData()) == 0 && t.unique { //unique set can't equal
				panic("order equal error")
			}
		}
	}
	var a, s1 = t.check(root.getChild(0))
	var b, s2 = t.check(root.getChild(1))
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
		compare = func(a, b interface{}) int {
			return a.(int) - b.(int)
		}
		newNode = func(data interface{}) Iterator {
			d, _ := data.(int)
			return &node{data: d}
		}
		deleteNode = func(node Iterator) {
		}
		sameNode = func(a Iterator, b Iterator) bool {
			return a.(*node) == b.(*node)
		}
	)
	tree := NewRBTree(newNode, deleteNode, compare, sameNode, true)
	var exists = make(map[int]bool, len(intSlice1K))
	if !tree.Empty() {
		t.Fatal("empty")
	}
	if tree.Begin() != tree.End() {
		t.Fatal("empty tree begin and end error")
	}
	//test RBTreer
	var _ RBTreer = tree
	var iter = tree.End()
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).data != intSlice1K[0] {
		if iter.(*node).data < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.(*node).data, intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).data < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.(*node).data, intSlice1K[0])
	}
	var maxVal, minVal = intSlice1K[0], intSlice1K[0]
	//test Insert method
	for _, val := range intSlice1K {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
		_, ok := tree.Insert(val)
		if ok == exists[val] {
			panic("insert error")
		}
		if tree.Begin().GetData() != minVal || tree.most(0) != tree.leftmost() {
			t.Fatal("leftmost error")
		}
		if tree.End().Last().GetData() != maxVal || tree.most(1) != tree.rightmost() {
			t.Fatal("rightmost error")
		}
		exists[val] = true
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.IntSlice(sortSlice).Sort()
	//unique sortslice
	if tree.unique {
		var uniqueN = 1
		for i := range sortSlice {
			if i > 0 && sortSlice[i] != sortSlice[i-1] {
				sortSlice[uniqueN] = sortSlice[i]
				uniqueN++
			}
		}
		sortSlice = sortSlice[:uniqueN]
	}
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).data != intSlice1K[0] {
		if iter.(*node).data < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.(*node).data, intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).data < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.(*node).data, intSlice1K[0])
	}
	//test Begin and EndNode method
	if tree.Begin().(*node).data != sortSlice[0] {
		t.Fatal("begin error", tree.Begin().(*node).data, sortSlice[0])
	}
	if tree.End().Last().(*node).data != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", tree.Begin().(*node).data, sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); !tree.sameIterator(it, tree.End()); it = it.Next() {
		if it.(*node).data != sortSlice[i] {
			t.Fatal("go through error", it.(*node).data, sortSlice[i])
		}
		i++
	}
	//test End and Last method
	i = len(sortSlice) - 1
	for it := tree.End().Last(); ; it = it.Last() {
		if it.(*node).data != sortSlice[i] {
			t.Fatal("go back tree error", it.(*node).data, sortSlice[i])
		}
		i--
		if tree.sameIterator(tree.Begin(), it) {
			break
		}
	}
	//test Find method
	iter = tree.Find(intSlice1K[0])
	if iter == tree.End() {
		t.Fatal("find error", intSlice1K[0])
	} else if iter.(*node).data != intSlice1K[0] {
		t.Fatal("find not equal", intSlice1K[0])
	}
	if tree.Find(max) != tree.End() {
		t.Fatal("find max error", max)
	}
	//test Erase method
	for _, val := range intSlice1K {
		/*if tree.most(0) != tree.leftmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.most(0), tree.leftmost(), tree.End(), tree.Size())
		}
		if tree.most(1) != tree.rightmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.most(1), tree.rightmost(), tree.End(), tree.Size())
		}*/
		num := tree.Erase(val)
		ok := num != 0
		if ok != exists[val] {
			t.Fatal("erase error")
		}
		delete(exists, val)
		_, size := tree.Check()
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
		_, ok := tree.Insert(val)
		if ok == exists[val] {
			t.Fatal("insert error")
		}
		exists[val] = true
		_, size := tree.Check()
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
func TestRBTree(t *testing.T) {
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}
	for i := 0; i < 100; i++ {
		testRBTree(t, rand.Int()%1000+1)
	}
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

//BenchmarkNewSetNode-4   	300000000	         4.90 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNewSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &SetNode{data: 1}
	}
}

//BenchmarkNewHeapSetNode-4   	20000000	        85.1 ns/op	      96 B/op	       1 allocs/op
func BenchmarkNewHeapSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		node := new(SetNode)
		node.data = 1
	}
}

//BenchmarkNewNode-4   	20000000	       125 ns/op	      96 B/op	       1 allocs/op
func BenchmarkNewNode(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := 0; i < b.N; i++ {
		_ = set.newNode(0)
	}
}

//BenchmarkNewPoiterNode-4   	20000000	       123 ns/op	      96 B/op	       1 allocs/op
func BenchmarkNewPoiterNode(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a = 0
	var c = &a
	for i := 0; i < b.N; i++ {
		_ = set.newNode(c)
	}
}

//BenchmarkCompare-4   	100000000	        14.6 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompare(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a, c = set.NewNode(0), set.NewNode(1)
	for i := 0; i < b.N; i++ {
		_ = set.compare(a.GetData(), c.GetData())
	}
}
