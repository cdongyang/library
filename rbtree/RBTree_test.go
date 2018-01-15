package rbtree

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/cdongyang/library/algorithm"
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
	n.data = src.GetKey().(int)
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
	//fmt.Printf("l:%p r:%p p:%p s:%p c:%s v:%d\n", root.getChild(0), root.getChild(1), root.getParent(), root, root.getColor().String(), root.GetKey())
	for i := 0; i < 2; i++ {
		if !t.sameIterator(root.getChild(i), t.End()) {
			if root.getColor() == red && root.getChild(i).getColor() == red {
				panic("linked red node")
			} else if !t.sameIterator(root.getChild(i).getParent(), root) {
				panic("tree error")
			} else if i == 0 && t.compare(root.getChild(i).GetKey(), root.GetKey()) > 0 {
				panic("order error")
			} else if i == 1 && t.compare(root.getChild(i).GetKey(), root.GetKey()) < 0 {
				panic("order error")
			} else if t.compare(root.getChild(i).GetKey(), root.GetKey()) == 0 && t.unique { //unique set can't equal
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

func testRBTree(t *testing.T, length int, unique bool) {
	// init
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
			return &node{data: data.(int)}
		}
		deleteNode = func(node Iterator) {
		}
		sameNode = func(a Iterator, b Iterator) bool {
			return a.(*node) == b.(*node)
		}
	)
	var tree = &RBTree{}
	tree = NewRBTreer(tree, &node{}, newNode, deleteNode, compare, sameNode, unique).(*RBTree)
	var count = make(map[int]int, len(intSlice1K))
	// test empty tree and empty tree Begin and End
	if !tree.Empty() {
		t.Fatal("empty")
	}
	if tree.Begin() != tree.End() {
		t.Fatal("empty tree begin and end error")
	}

	//test RBTreer
	var _ RBTreer = tree
	var iter = tree.End()
	var index int

	//test empty tree LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if !tree.SameIterator(iter, tree.End()) {
		t.Fatal("empty tree LowerBound error")
	}
	//test empty tree UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if !tree.SameIterator(iter, tree.End()) {
		t.Fatal("empty tree UpperBound error")
	}

	//test Insert,Begin,End method
	var maxVal, minVal = intSlice1K[0], intSlice1K[0]
	for _, val := range intSlice1K {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
		//fmt.Println("insert", val)
		_, ok := tree.Insert(val)
		if !tree.Unique() && !ok || tree.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if tree.Begin().GetKey() != minVal || tree.most(0) != tree.leftmost() {
			t.Fatal("leftmost error")
		}
		if tree.End().Last().GetKey() != maxVal || tree.most(1) != tree.rightmost() {
			t.Fatal("rightmost error")
		}
		if tree.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if tree.Count(val) != count[val] {
			t.Fatal("count error", tree.Count(val), count[val])
		}
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}

	// test Compare
	if tree.Compare(tree.Begin().GetKey(), minVal) != 0 {
		t.Fatal("Compare error")
	}
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.IntSlice(sortSlice).Sort()

	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	index = algorithm.LowerBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if tree.SameIterator(iter, tree.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
		//ok
	} else {
		t.Fatal("LowerBound error")
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	index = algorithm.UpperBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if tree.SameIterator(iter, tree.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
	}
	//test Begin and EndNode method
	if tree.Begin().GetKey() != sortSlice[0] {
		t.Fatal("begin error", tree.Begin().GetKey(), sortSlice[0])
	}
	if tree.End().Last().GetKey() != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", tree.Begin().GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); !tree.SameIterator(it, tree.End()); it = it.Next() {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go through error", it.GetKey(), sortSlice[i])
		}
		if tree.Unique() {
			for i+1 < len(sortSlice) && sortSlice[i+1] == sortSlice[i] {
				i++
			}
		}
		i++
	}
	//test End and Last method
	i = len(sortSlice) - 1
	for it := tree.End().Last(); ; it = it.Last() {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go back tree error", it.GetKey(), sortSlice[i])
		}
		if tree.Unique() {
			for i > 0 && sortSlice[i-1] == sortSlice[i] {
				i--
			}
		}
		i--
		if tree.SameIterator(tree.Begin(), it) {
			break
		}
	}
	//test Find method
	iter = tree.Find(intSlice1K[0])
	index = sort.SearchInts(sortSlice, intSlice1K[0])
	if tree.SameIterator(iter, tree.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == sortSlice[index] {
		//ok
	} else {
		t.Fatal("Find error")
	}
	if tree.Find(max) != tree.End() {
		t.Fatal("find max error", max)
	}
	//unique sortslice
	var uniqueN = 1
	for i := range sortSlice {
		if i > 0 && sortSlice[i] != sortSlice[i-1] {
			sortSlice[uniqueN] = sortSlice[i]
			uniqueN++
		}
	}
	sortSlice = sortSlice[:uniqueN]
	//test Erase method
	for _, val := range intSlice1K {
		//fmt.Println("erase", val)
		num := tree.Erase(val)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if tree.Count(val) != count[val] {
			t.Fatal("count error", tree.Count(val), count[val])
		}
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error")
		}
		if tree.most(0) != tree.leftmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.most(0), tree.leftmost(), tree.End(), tree.Size())
		}
		if tree.most(1) != tree.rightmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.most(1), tree.rightmost(), tree.End(), tree.Size())
		}
	}
	//test Empty method
	if !tree.Empty() {
		t.Fatal("empty error")
	}
	count = make(map[int]int)
	for _, val := range intSlice1K {
		_, ok := tree.Insert(val)
		if !tree.Unique() && !ok || tree.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if tree.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if tree.Count(val) != count[val] {
			t.Fatal("count error", tree.Count(val), count[val])
		}
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}
	//test EqualRange and EraseIteratorRange
	for _, val := range intSlice1K {
		//fmt.Println("erase", val)
		beg, end := tree.EqualRange(val)
		num := tree.EraseIteratorRange(beg, end)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if tree.Count(val) != count[val] {
			t.Fatal("count error", tree.Count(val), count[val])
		}
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error")
		}
		if tree.most(0) != tree.leftmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.most(0), tree.leftmost(), tree.End(), tree.Size())
		}
		if tree.most(1) != tree.rightmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.most(1), tree.rightmost(), tree.End(), tree.Size())
		}
	}
	if tree.Size() != 0 || !tree.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
	//test Clear method
	for _, val := range intSlice1K {
		_, _ = tree.Insert(val)
	}
	tree.Clear()
	if tree.Size() != 0 || !tree.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}
func TestRBTree(t *testing.T) {
	for i := 0; i < 500; i++ {
		testRBTree(t, i+1, i%2 == 0)
	}
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

func BenchmarkRBTreeInsert(t *testing.B) {
	var (
		compare = func(a, b interface{}) int {
			return a.(int) - b.(int)
		}
		newNode = func(data interface{}) Iterator {
			return &node{data: data.(int)}
		}
		deleteNode = func(node Iterator) {
		}
		sameNode = func(a Iterator, b Iterator) bool {
			return a.(*node) == b.(*node)
		}
	)
	var set = &RBTree{}
	set = NewRBTreer(set, &node{}, newNode, deleteNode, compare, sameNode, true).(*RBTree)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		_, _ = set.Insert(rand.Int())
	}
}

func BenchmarkRBTreeErase(t *testing.B) {
	var (
		compare = func(a, b interface{}) int {
			return a.(int) - b.(int)
		}
		newNode = func(data interface{}) Iterator {
			return &node{data: data.(int)}
		}
		deleteNode = func(node Iterator) {
		}
		sameNode = func(a Iterator, b Iterator) bool {
			return a.(*node) == b.(*node)
		}
	)
	var set = &RBTree{}
	set = NewRBTreer(set, &node{}, newNode, deleteNode, compare, sameNode, true).(*RBTree)
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = set.Erase(keys[i])
	}
}

func BenchmarkRBTreeFind(t *testing.B) {
	var (
		compare = func(a, b interface{}) int {
			return a.(int) - b.(int)
		}
		newNode = func(data interface{}) Iterator {
			return &node{data: data.(int)}
		}
		deleteNode = func(node Iterator) {
		}
		sameNode = func(a Iterator, b Iterator) bool {
			return a.(*node) == b.(*node)
		}
	)
	var set = &RBTree{}
	set = NewRBTreer(set, &node{}, newNode, deleteNode, compare, sameNode, true).(*RBTree)
	t.StopTimer()
	var keys = make([]int, t.N)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = set.Find(keys[i])
	}
}

//BenchmarkNewSetNode-4   	300000000	         4.90 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNewSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &SetNode{data: 1}
	}
}

//BenchmarkNewHeapSetNode-4   	20000000	        85.1 ns/op	      96 B/op	       1 allocs/op
func BenchmarkNewHeapSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var node *SetNode
		func() {
			node = new(SetNode)
		}()
		node.data = 1
	}
}

//BenchmarkNewNode-4   	20000000	       125 ns/op	      96 B/op	       1 allocs/op
func BenchmarkSetNewNode(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := 0; i < b.N; i++ {
		_ = set.newNode(0)
	}
}

//BenchmarkNewPoiterNode-4   	20000000	       123 ns/op	      96 B/op	       1 allocs/op
func BenchmarkSetNewPoiterNode(b *testing.B) {
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
	var a, c = set.newNode(0), set.newNode(1)
	for i := 0; i < b.N; i++ {
		_ = set.compare(a.GetKey(), c.GetKey())
	}
}

// the size of interface{} is 16B
// type eface struct {
// 	_type *_type
// 	data  unsafe.Pointer
// }
//
// type iface struct {
// 	tab  *itab
// 	data unsafe.Pointer
// }

func ExampleSetNodeSize() {
	fmt.Println(reflect.TypeOf(RBTreeNode{}).Size())
	fmt.Println(reflect.TypeOf(SetNode{}).Size())
	fmt.Println(reflect.TypeOf(MapNode{}).Size())
	//Output:
	//72
	//88
	//88
}
