package rbtree

import (
	"log"
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

var intType = interface2type(1)

func NewTree(unique bool) *tree {
	var tree = &tree{}
	tree.init(unique, intType, nil, CompareInt)
	return tree
	//fmt.Printf("header: %p header: %+v\n", header, header)
}

func (t *tree) Most(ch int) *node {
	return t.most(ch)
}

func (t *tree) Leftmost() *node {
	return t.getmost(0)
}
func (t *tree) Rightmost() *node {
	return t.getmost(1)
}
func (t *tree) getmost(ch int) *node {
	var root = t.root()
	if root == t.end() {
		return root
	}
	if root == nil {
		log.Fatal(t)
	}
	for root.child[ch] != t.end() {
		if root.child[ch] == nil {
			log.Fatal(root)
		}
		root = root.child[ch]
	}
	return root
}

//check wheather the tree satisfy the rule of red-black tree
func (t *tree) Check() (int, int) {
	if t.root() != t.end() && t.root().color != black {
		panic("not black root")
	}
	return t.check(t.root())
}

func (t *tree) check(root *node) (l int, size int) {
	if root == t.end() {
		return 0, 0
	}
	//fmt.Printf("l:%p r:%p p:%p s:%p c:%s v:%d\n", root.getChild(0), root.getChild(1), root.getParent(), root, root.getColor().String(), root.GetKey())
	for i := 0; i < 2; i++ {
		if root.child[i] != t.end() {
			cmp := t.compare(root.child[i].GetKey(), root.GetKey())
			if root.color == red && root.child[i].color == red {
				panic("linked red node")
			} else if root.child[i].parent != root {
				panic("tree error")
			} else if i == 0 && cmp > 0 {
				panic("order error")
			} else if i == 1 && cmp < 0 {
				panic("order error")
			} else if cmp == 0 && t.unique {
				panic("order equal error")
			}
		}
	}
	var a, s1 = t.check(root.child[0])
	var b, s2 = t.check(root.child[1])
	if a != b {
		panic("path length of black not equal")
	}
	if root.color == black {
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

	var iter = tree.End()
	var index int

	//test empty tree LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter != tree.End() {
		t.Fatal("empty tree LowerBound error")
	}
	//test empty tree UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter != tree.End() {
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
			t.Fatal("rightmost error", tree.End().Last().GetKey(), maxVal)
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
	if iter == tree.End() && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
		//ok
	} else {
		t.Fatal("LowerBound error")
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	index = algorithm.UpperBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if iter == tree.End() && index == len(sortSlice) {
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
	for it := tree.Begin(); it != tree.End(); it = it.Next() {
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
		if tree.Begin() == it {
			break
		}
	}
	//test Find method
	iter = tree.Find(intSlice1K[0])
	index = sort.SearchInts(sortSlice, intSlice1K[0])
	if iter == tree.End() && index == len(sortSlice) {
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
	//test EqualRange and EraseNodeRange
	for _, val := range intSlice1K {
		//fmt.Println("erase", val)
		beg, end := tree.EqualRange(val)
		num := tree.EraseNodeRange(beg, end)
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
	var testN = 500
	t.Run("unique", func(t *testing.T) {
		for i := 0; i < testN; i++ {
			testTree(t, i+1, true)
		}
	})
	t.Run("not unique", func(t *testing.T) {
		for i := 0; i < testN; i++ {
			testTree(t, i+1, false)
		}
	})
	t.Run("unique 1e4", func(t *testing.T) {
		testTree(t, 1e4, true)
	})
	t.Run("not unique 1e4", func(t *testing.T) {
		testTree(t, 1e4, false)
	})
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

// BenchmarkIntTreeInsert-4   	 3000000	       564 ns/op	      64 B/op	       1 allocs/op
func BenchmarkIntTreeInsert(b *testing.B) {
	var set = NewTree(true)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		_, _ = set.Insert(rand.Int())
	}
}

// BenchmarkIntTreeErase-4   	 5000000	       370 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIntTreeErase(b *testing.B) {
	var set = NewTree(true)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Erase(keys[i])
	}
}

// BenchmarkIntTreeFind-4   	 5000000	       289 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIntTreeFind(b *testing.B) {
	var set = NewTree(true)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Find(keys[i])
	}
}

func TestNewNode(t *testing.T) {
	var tree = NewTree(true)
	var size = 8
	for i := 0; i < 128; i++ {
		node := tree.newNode(i)
		if node.GetKey() != i {
			t.Fatal(node.GetKey(), i)
		}
		if node.child[0] != tree.end() || node.child[1] != tree.end() || node.parent != tree.end() ||
			node.parent != tree.end() || node.color != red {
			t.Fatal(node)
		}
		size--
		if tree.qsize != size {
			t.Fatal(i, tree.qsize, size)
		}
		if (i+1) >= 8 && (i&(i+1)) == 0 {
			size = tree.size
		}
	}
}

func CompareInt(a, b interface{}) int {
	return a.(int) - b.(int)
}

func TestNode(t *testing.T) {
	t.Run("int set", func(t *testing.T) {
		var tree = &tree{}
		tree.init(true, intType, nil, CompareInt)
		node := tree.newNode(10)
		t.Logf("%p %p %p %p %p %p %p %v\n", &node.child[0], &node.child[1], &node.parent, &node.tree,
			&node.color, interface2pointer(node.GetKey()), tree.dataSize)
	})
}
