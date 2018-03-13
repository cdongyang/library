package rbtree

import (
	"fmt"
	"runtime"
	"sort"
	"testing"
	"time"

	"github.com/cdongyang/library/algorithm"
	"github.com/cdongyang/library/randint"
)

func (t colorType) String() string {
	if t == red {
		return "red"
	}
	return "black"
}

func (t *tree) Most(ch uintptr) _node {
	return t.pack(t.most(ch))
}

func (t *tree) Leftmost() _node {
	return t.pack(t.getmost(0))
}
func (t *tree) Rightmost() _node {
	return t.pack(t.getmost(1))
}
func (t *tree) getmost(ch uintptr) node {
	var root = t.root()
	if root == t.end() {
		return t.end()
	}
	for t.getChild(root, ch) != t.end() {
		root = t.getChild(root, ch)
	}
	return root
}

//check wheather the tree satisfy the rule of red-black tree
func (t *tree) Check() (int, int) {
	if t.mustGetColor(t.root()) != black {
		panic("not black root")
	}
	return t.check(t.root())
}

func (t *tree) check(root node) (l int, size int) {
	if root == t.end() {
		return 0, 0
	}
	for i := uintptr(0); i < 2; i++ {
		if t.getChild(root, i) != t.end() {
			cmp := t.compare(t.getKey(t.getChild(root, i)), t.getKey(root))
			if t.getColor(root) == red && t.getColor(t.getChild(root, i)) == red {
				panic("linked red node")
			} else if t.getParent(t.getChild(root, i)) != root {
				panic("tree error")
			} else if cmp == 0 && t.unique {
				panic("order equal error")
			} else if i == 0 && cmp > 0 {
				fmt.Println(i, t.getKey(t.getChild(root, i)), t.getKey(root))
				panic("order error")
			} else if i == 1 && cmp < 0 {
				fmt.Println(i, t.getKey(t.getChild(root, i)), t.getKey(root))
				panic("order error")
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

func NewTree(unique bool) *tree {
	var tree = &tree{}
	tree.Init(unique, int(0), nil, CompareInt)
	return tree
	//fmt.Printf("header: %p header: %+v\n", header, header)
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
		_, ok := tree.Insert(val, nil)
		if !tree.Unique() && !ok || tree.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if tree.Begin().GetKey() != minVal || tree.Most(0) != tree.Leftmost() {
			t.Fatal("leftmost error", tree.Begin().GetKey(), minVal)
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
		if tree.Most(0) != tree.Leftmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(0), tree.Leftmost(), tree.End(), tree.Size())
		}
		if tree.Most(1) != tree.Rightmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(1), tree.Rightmost(), tree.End(), tree.Size())
		}
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
	}
	//test Empty method
	if !tree.Empty() {
		t.Fatal("empty error")
	}
	count = make(map[int]int)
	for _, val := range intSlice1K {
		_, ok := tree.Insert(val, nil)
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
		if tree.Most(0) != tree.Leftmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(0), tree.Leftmost(), tree.End(), tree.Size())
		}
		if tree.Most(1) != tree.Rightmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(1), tree.Rightmost(), tree.End(), tree.Size())
		}
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
	}
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
		_, _ = set.Insert(rand.Int(), nil)
	}
}

// BenchmarkIntTreeErase-4   	 5000000	       370 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIntTreeErase(b *testing.B) {
	var set = NewTree(true)
	var keys = make([]int, b.N)
	var rand = benchRand
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		_, _ = set.Insert(keys[i], nil)
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
		_, _ = set.Insert(keys[i], nil)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Find(keys[i])
	}
}

func TestNewNode(t *testing.T) {
	var tree = NewTree(true)
	var size = 7
	for i := 1; i < 128; i++ {
		node := tree.pack(tree.newNode(i, nil))
		if node.GetKey() != i {
			t.Fatal(node.GetKey(), i)
		}
		func() {
			defer func() {
				err := recover()
				if err != ErrNoValue.Error() {
					t.Fatal(err)
				}
			}()
			node.GetVal()
		}()
		if tree.getChild(node.node, 0) != tree.end() || tree.getChild(node.node, 1) != tree.end() ||
			tree.getParent(node.node) != tree.end() || tree.getColor(node.node) != red {
			t.Fatal(node)
		}
		size--
	}
}

func CompareInt(a, b interface{}) int {
	return a.(int) - b.(int)
}

func ExampleSet() {
	s := NewSet(tmpInt, CompareInt)
	fmt.Println(s.ifacedata, s.directkey, s.directval)
	n, ok := s.Insert(1)
	if !ok {
		panic("insert error")
	}
	time.Sleep(time.Second)
	runtime.GC()
	fmt.Println(n.GetData())
	time.Sleep(time.Second)
	runtime.GC()
	fmt.Println(n.GetData())
	// Output:
	//true false false
	//1
	//1
}
