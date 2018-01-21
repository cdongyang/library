package rbtree

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"unsafe"

	"github.com/cdongyang/library/algorithm"
	"github.com/cdongyang/library/randint"
)

func (t colorType) String() string {
	if t == red {
		return "red"
	}
	return "black"
}

var offsetNode uintptr

func init() {
	var header = &node{}
	offsetNode = uintptr(unsafe.Pointer(&header.Node)) - uintptr(unsafe.Pointer(header))
	//fmt.Println("offsetNode:", offsetNode)
}

type node struct {
	data int
	Node
}

type IntSet struct {
	Tree
}

func NewTree(unique bool) *IntSet {
	var (
		newNode = func(data interface{}) Iterator {
			//return &node{data: data.(int)}
			return &node{data: data.(int)}
		}
		deleteNode = func(node Iterator) {
		}
	)
	var tree = &Tree{}
	var header = &node{}
	//fmt.Printf("header: %p header: %+v\n", header, header)
	return &IntSet{*NewTreer(tree, header, offsetNode, newNode, deleteNode, CompareInt,
		func(p unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(&(*node)(p).data)
		},
		unique).(*Tree)}
}

func NewTreeWithPool(unique bool) *IntSet {
	var (
		pool = &sync.Pool{New: func() interface{} {
			return &node{}
		}}
		newNode = func(data interface{}) Iterator {
			var n = pool.Get().(*node)
			n.data = data.(int)
			return n
		}
		deleteNode = func(node Iterator) {
			pool.Put(node)
		}
	)
	var tree = &Tree{}
	var header = &node{}
	//fmt.Printf("header: %p header: %+v\n", header, header)
	return &IntSet{*NewTreer(tree, header, offsetNode, newNode, deleteNode, CompareInt,
		func(p unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(&(*node)(p).data)
		},
		unique).(*Tree)}
}

func (s *IntSet) Insert(data interface{}) (iter Iterator, ok bool) {
	var dataEface = interface2eface(data)
	var dataCopy = eface{dataEface._type, noescape(dataEface.pointer)}
	return s.Tree.Insert(*(*interface{})(noescape(unsafe.Pointer(&dataCopy))))
}

func ExampleNodeOffset() {
	var it = &node{}
	var base = uintptr(unsafe.Pointer(it))
	fmt.Println(uintptr(unsafe.Pointer(&it.Node)) - base)
	fmt.Println(uintptr(unsafe.Pointer(&it.child[0])) - base)
	fmt.Println(uintptr(unsafe.Pointer(&it.child[1])) - base)
	fmt.Println(uintptr(unsafe.Pointer(&it.parent)) - base)
	fmt.Println(uintptr(unsafe.Pointer(&it.tree)) - base)
	fmt.Println(uintptr(unsafe.Pointer(&it.color)) - base)
	// Output:
	//8
	//8
	//16
	//24
	//32
	//40
}

func (n *node) Next() Iterator {
	return n.tree.Next(n)
}

func (n *node) Last() Iterator {
	return n.tree.Last(n)
}

var intType = interface2type(1)

func (n *node) GetKey() interface{} {
	return *(*interface{})(unsafe.Pointer(&eface{intType, unsafe.Pointer(&n.data)}))
}

func (n *node) GetData() (key interface{}) {
	*(*eface)(unsafe.Pointer(&key)) = eface{intType, unsafe.Pointer(&n.data)}
	return
}

func (n *node) CopyData(src Iterator) {
	n.data = src.GetKey().(int)
}

func (t *Tree) Most(ch int) Iterator {
	return t.pointer2iterator(t.most(ch))
}

func (t *Tree) Leftmost() Iterator {
	return t.pointer2iterator(t.leftmost())
}
func (t *Tree) leftmost() unsafe.Pointer {
	var root = t.root()
	if sameNode(root, t.end()) {
		return root
	}
	for !sameNode(t.getChild(root, 0), t.end()) {
		root = t.getChild(root, 0)
	}
	return root
}

func (t *Tree) Rightmost() Iterator {
	return t.pointer2iterator(t.rightmost())
}
func (t *Tree) rightmost() unsafe.Pointer {
	var root = t.root()
	if sameNode(root, t.end()) {
		return root
	}
	for !sameNode(t.getChild(root, 1), t.end()) {
		root = t.getChild(root, 1)
	}
	return root
}

//check wheather the tree satisfy the rule of red-black tree
func (t *Tree) Check() (int, int) {
	if !sameNode(t.root(), t.end()) && t.getColor(t.root()) != black {
		panic("not black root")
	}
	return t.check(t.root())
}
func (t *Tree) check(root unsafe.Pointer) (l int, size int) {
	if sameNode(root, t.end()) {
		return 0, 0
	}
	//fmt.Printf("l:%p r:%p p:%p s:%p c:%s v:%d\n", root.getChild(0), root.getChild(1), root.getParent(), root, root.getColor().String(), root.GetKey())
	for i := 0; i < 2; i++ {
		if !sameNode(t.getChild(root, i), t.end()) {
			if t.getColor(root) == red && t.getColor(t.getChild(root, i)) == red {
				panic("linked red node")
			} else if !sameNode(t.getParent(t.getChild(root, i)), root) {
				panic("tree error")
			} else if i == 0 && t.compare(t.getKeyPointer(t.getChild(root, i)), t.getKeyPointer(root)) > 0 {
				panic("order error")
			} else if i == 1 && t.compare(t.getKeyPointer(t.getChild(root, i)), t.getKeyPointer(root)) < 0 {
				panic("order error")
			} else if t.compare(t.getKeyPointer(t.getChild(root, i)), t.getKeyPointer(root)) == 0 && t.unique { //unique set can't equal
				panic("order equal error")
			}
		}
	}
	var a, s1 = t.check(t.getChild(root, 0))
	var b, s2 = t.check(t.getChild(root, 1))
	if a != b {
		panic("path length of black not equal")
	}
	if t.getColor(root) == black {
		return a + 1, s1 + s2 + 1
	}
	return a, s1 + s2 + 1
}

func testTree(t *testing.T, length int, unique bool) {
	// init
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1000}
	var max = rand.Int()%length + 1
	var intSlice1K = make([]int, length)
	for i := range intSlice1K {
		intSlice1K[i] = rand.Int() % max
	}
	//fmt.Println("offsetNode:", offsetNode)
	var tree = NewTree(unique)
	var count = make(map[int]int, len(intSlice1K))
	// test empty tree and empty tree Begin and End
	if !tree.Empty() {
		t.Fatal("empty")
	}
	//fmt.Println(tree.Begin(), tree.End())
	if tree.Begin() != tree.End() {
		t.Fatal("empty tree begin and end error")
	}

	//test Treer
	var _ Treer = tree
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
		if tree.Begin().GetKey() != minVal || tree.Most(0) != tree.Leftmost() {
			t.Fatal("leftmost error")
		}
		if tree.End().Last().GetKey() != maxVal || tree.Most(1) != tree.Rightmost() {
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
	if tree.Begin().GetKey() != minVal {
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
		if tree.Most(0) != tree.Leftmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.Most(0), tree.Leftmost(), tree.End(), tree.Size())
		}
		if tree.Most(1) != tree.Rightmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.Most(1), tree.Rightmost(), tree.End(), tree.Size())
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
		if tree.Most(0) != tree.Leftmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.Most(0), tree.Leftmost(), tree.End(), tree.Size())
		}
		if tree.Most(1) != tree.Rightmost() {
			t.Fatalf("tree most error,%p %p %p %d\n", tree.Most(1), tree.Rightmost(), tree.End(), tree.Size())
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
func TestTree(t *testing.T) {
	for i := 0; i < 500; i++ {
		testTree(t, i+1, i%2 == 0)
	}
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

// BenchmarkIntSetInsert-4   	 3000000	       564 ns/op	      64 B/op	       1 allocs/op
func BenchmarkIntSetInsert(t *testing.B) {
	var set = NewTree(true)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		_, _ = set.Insert(rand.Int())
	}
}

func BenchmarkIntSetInsertWithPool(t *testing.B) {
	var set = NewTreeWithPool(true)
	var rand = benchRand
	for i := 0; i < t.N; i++ {
		_, _ = set.Insert(rand.Int())
	}
}

// BenchmarkIntSetErase-4   	 5000000	       370 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIntSetErase(t *testing.B) {
	var set = NewTree(true)
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

func BenchmarkIntSetEraseWithPool(t *testing.B) {
	var set = NewTreeWithPool(true)
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

// BenchmarkIntSetFind-4   	 5000000	       289 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIntSetFind(b *testing.B) {
	var set = NewTree(true)
	b.StopTimer()
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Find(keys[i])
	}
}

func BenchmarkIntSetFindWithPool(b *testing.B) {
	var set = NewTreeWithPool(true)
	b.StopTimer()
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Find(keys[i])
	}
}

func BenchmarkIntSetGetKey(b *testing.B) {
	var a = &node{data: 1}
	for i := 0; i < b.N; i++ {
		_ = a.GetKey()
	}
}
