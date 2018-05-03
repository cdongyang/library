package rbtree

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"unsafe"

	"github.com/cdongyang/library/algorithm"
	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/test"
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
	var intSlice = make([]int, length)
	for i := range intSlice {
		intSlice[i] = (rand.Int() % max) + 1
	}
	max++
	//fmt.Println("offsetNode:", offsetNode)
	var tree = NewTree(unique)
	var count = make(map[int]int, len(intSlice))
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
	iter = tree.LowerBound(intSlice[0])
	if iter != tree.End() {
		t.Fatal("empty tree LowerBound error")
	}
	//test empty tree UpperBound method
	iter = tree.UpperBound(intSlice[0])
	if iter != tree.End() {
		t.Fatal("empty tree UpperBound error")
	}

	//test Insert,Begin,End method
	var maxVal, minVal = intSlice[0], intSlice[0]
	for _, val := range intSlice {
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
	var sortSlice = make([]int, len(intSlice))
	copy(sortSlice, intSlice)
	sort.IntSlice(sortSlice).Sort()

	//test LowerBound method
	iter = tree.LowerBound(intSlice[0])
	index = algorithm.LowerBound(algorithm.SearchIntSlice{sortSlice, intSlice[0]})
	if iter == tree.End() && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice[0] {
		//ok
	} else {
		t.Fatal("LowerBound error")
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice[0])
	index = algorithm.UpperBound(algorithm.SearchIntSlice{sortSlice, intSlice[0]})
	if iter == tree.End() && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice[0] {
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
	iter = tree.Find(intSlice[0])
	index = sort.SearchInts(sortSlice, intSlice[0])
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
	for _, val := range intSlice {
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
	for _, val := range intSlice {
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
	for _, val := range intSlice {
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

func NewTree2(unique bool) *tree {
	var tree = &tree{}
	tree.Init(unique, "", new(int), func(a, b interface{}) int {
		aa, err := strconv.Atoi(a.(string))
		if err != nil {
			panic(err.Error())
		}
		bb, err := strconv.Atoi(b.(string))
		if err != nil {
			panic(err.Error())
		}
		return aa - bb
	})
	return tree
}

func testTree2(t *testing.T, length int, unique bool) {
	// init
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1000}
	var max = rand.Int()%length + 1
	var intSlice = make([]int, length)
	for i := range intSlice {
		intSlice[i] = (rand.Int() % max) + 1
	}
	max++
	var equal = func(n _node, val int) bool {
		return n.GetKey() == strconv.Itoa(val) && *(n.GetVal().(*int)) == val
	}
	//fmt.Println("offsetNode:", offsetNode)
	var tree = NewTree2(unique)
	var count = make(map[int]int, len(intSlice))
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
	iter = tree.LowerBound(strconv.Itoa(intSlice[0]))
	if iter != tree.End() {
		t.Fatal("empty tree LowerBound error")
	}
	//test empty tree UpperBound method
	iter = tree.UpperBound(strconv.Itoa(intSlice[0]))
	if iter != tree.End() {
		t.Fatal("empty tree UpperBound error")
	}

	//test Insert,Begin,End method
	var maxVal, minVal = intSlice[0], intSlice[0]
	for _, val := range intSlice {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
		//fmt.Println("insert", val)
		var tmp = val
		_, ok := tree.Insert(strconv.Itoa(val), &tmp)
		if !tree.Unique() && !ok || tree.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if !equal(tree.Begin(), minVal) || tree.Most(0) != tree.Leftmost() {
			t.Fatalf("leftmost error %v %v %v %v\n", tree.Begin().GetKey(), tree.Begin().GetVal(), *(tree.Begin().GetVal().(*int)), minVal)
		}
		if !equal(tree.End().Last(), maxVal) || tree.Most(1) != tree.Rightmost() {
			t.Fatal("rightmost error", tree.End().Last().GetKey(), maxVal)
		}
		if tree.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if tree.Count(strconv.Itoa(val)) != count[val] {
			t.Fatal("count error", tree.Count(strconv.Itoa(val)), count[val])
		}
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}

	// test Compare
	if !equal(tree.Begin(), minVal) {
		t.Fatal("Compare error")
	}
	var sortSlice = make([]int, len(intSlice))
	copy(sortSlice, intSlice)
	sort.IntSlice(sortSlice).Sort()

	//test LowerBound method
	iter = tree.LowerBound(strconv.Itoa(intSlice[0]))
	index = algorithm.LowerBound(algorithm.SearchIntSlice{sortSlice, intSlice[0]})
	if iter == tree.End() && index == len(sortSlice) {
		//ok
	} else if equal(iter, intSlice[0]) {
		//ok
	} else {
		t.Fatal("LowerBound error")
	}
	//test UpperBound method
	iter = tree.UpperBound(strconv.Itoa(intSlice[0]))
	index = algorithm.UpperBound(algorithm.SearchIntSlice{sortSlice, intSlice[0]})
	if iter == tree.End() && index == len(sortSlice) {
		//ok
	} else if equal(iter, intSlice[0]) {
	}
	//test Begin and EndNode method
	if !equal(tree.Begin(), sortSlice[0]) {
		t.Fatal("begin error", tree.Begin().GetKey(), sortSlice[0])
	}
	if !equal(tree.End().Last(), sortSlice[len(sortSlice)-1]) {
		t.Fatal("endNode error", tree.Begin().GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); it != tree.End(); it = it.Next() {
		if !equal(it, sortSlice[i]) {
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
		if !equal(it, sortSlice[i]) {
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
	iter = tree.Find(strconv.Itoa(intSlice[0]))
	index = sort.SearchInts(sortSlice, intSlice[0])
	if iter == tree.End() && index == len(sortSlice) {
		//ok
	} else if equal(iter, sortSlice[index]) {
		//ok
	} else {
		t.Fatal("Find error")
	}
	if tree.Find(strconv.Itoa(max)) != tree.End() {
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
	for _, val := range intSlice {
		if tree.Most(0) != tree.Leftmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(0), tree.Leftmost(), tree.End(), tree.Size())
		}
		if tree.Most(1) != tree.Rightmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(1), tree.Rightmost(), tree.End(), tree.Size())
		}
		//fmt.Println("erase", val)
		num := tree.Erase(strconv.Itoa(val))
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if tree.Count(strconv.Itoa(val)) != count[val] {
			t.Fatal("count error", tree.Count(strconv.Itoa(val)), count[val])
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
	for _, val := range intSlice {
		var tmp = val
		_, ok := tree.Insert(strconv.Itoa(val), &tmp)
		if !tree.Unique() && !ok || tree.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if tree.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if tree.Count(strconv.Itoa(val)) != count[val] {
			t.Fatal("count error", tree.Count(strconv.Itoa(val)), count[val])
		}
		_, size := tree.Check()
		if size != tree.Size() {
			t.Fatal("size error", size, tree.Size())
		}
	}
	//test EqualRange and EraseNodeRange
	for _, val := range intSlice {
		if tree.Most(0) != tree.Leftmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(0), tree.Leftmost(), tree.End(), tree.Size())
		}
		if tree.Most(1) != tree.Rightmost() {
			t.Fatalf("tree most error,%v %v %v %d\n", tree.Most(1), tree.Rightmost(), tree.End(), tree.Size())
		}
		//fmt.Println("erase", val)
		beg, end := tree.EqualRange(strconv.Itoa(val))
		num := tree.EraseNodeRange(beg, end)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if tree.Count(strconv.Itoa(val)) != count[val] {
			t.Fatal("count error", tree.Count(strconv.Itoa(val)), count[val])
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

func TestTree2(t *testing.T) {
	var testN = 500
	t.Run("unique", func(t *testing.T) {
		for i := 0; i < testN; i++ {
			testTree2(t, i+1, true)
		}
	})
	t.Run("not unique", func(t *testing.T) {
		for i := 0; i < testN; i++ {
			testTree2(t, i+1, false)
		}
	})
	t.Run("unique 1e4", func(t *testing.T) {
		testTree2(t, 1e4, true)
	})
	t.Run("not unique 1e4", func(t *testing.T) {
		testTree2(t, 1e4, false)
	})
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

func CompareInt(a, b interface{}) int {
	return a.(int) - b.(int)
}

func ExampleMaxSpan() {
	var tree = NewTree(true)
	fmt.Println(tree.GetMaxSpan())
	tree.SetMaxSpan(0)
	fmt.Println(tree.GetMaxSpan())
	tree.SetMaxSpan(20)
	fmt.Println(tree.GetMaxSpan())
	tree.SetMaxSpan(35)
	fmt.Println(tree.GetMaxSpan())
	tree.SetMaxSpan(512)
	fmt.Println(tree.GetMaxSpan())
	// Output:
	//1024
	//8
	//16
	//32
	//512
}

func ExampleNoescapeInsert() {
	var tree = NewTree(false)
	var rand = benchRand
	n := testing.AllocsPerRun(1000, func() {
		_, ok := tree.Insert(noescapeInterface(rand.Int()), nil)
		if !ok {
			panic("insert error")
		}
	})
	fmt.Println(n)
	// Output:
	//0
}

func TestGC(t *testing.T) {
	test.MemStats("begin")
	t.Run("GC tree", func(t *testing.T) {
		s := NewSet(int(0), CompareInt)
		for i := 0; i < 1e5; i++ {
			s.Insert(i)
		}
		test.MemStats("1e5 node")
		s = nil
		test.MemStats("free")
	})
	t.Run("hode interface", func(t *testing.T) {
		s := NewSet(int(0), CompareInt)
		for i := 0; i < 1e5; i++ {
			s.Insert(i)
		}
		s.SetMaxSpan(1 << 20)
		test.MemStats("1e5 node")
		node := s.Find(int(1e5 - 1)).GetData()
		s = nil
		test.MemStats("hold node, free tree")
		node = nil
		_ = node
		test.MemStats("free node")
	})
	t.Run("hode interface map", func(t *testing.T) {
		m := NewMap(int(0), false, CompareInt)
		for i := 0; i < 1e5; i++ {
			m.Insert(i, true)
		}
		m.SetMaxSpan(1 << 20)
		test.MemStats("1e5 node")
		k, v := m.Find(1).GetData()
		m = nil
		test.MemStats("hold node, free tree")
		k, v = nil, nil
		_, _ = k, v
		test.MemStats("free node")
	})
}

func TestGetArrayPtrOfSliceValue(t *testing.T) {
	a := 10
	atype := reflect.TypeOf(a)
	avalue := reflect.ValueOf(a)
	as := reflect.MakeSlice(reflect.SliceOf(atype), 4, 4)
	asp := getArrayPtrOfSliceValue(as)
	as.Index(0).Set(avalue)
	as.Index(2).Set(avalue)
	t.Log(*(*[4]int)(asp))
	t.Log(pack2Iface(unpackIface(a)._type, arrayAt(asp, 2, 8)))
	if *(*int)(asp) != a {
		t.Logf("%+v\n%+v\n", as, *(*slice)((*eface)(unsafe.Pointer(&as)).p))
		t.Fatal(*(*int)(asp))
	}
	b := &a
	btype := reflect.TypeOf(b)
	bvalue := reflect.ValueOf(b)
	bs := reflect.MakeSlice(reflect.SliceOf(btype), 4, 4)
	bsp := getArrayPtrOfSliceValue(bs)
	bs.Index(0).Set(bvalue)
	bs.Index(2).Set(bvalue)
	t.Log(*(*[4]unsafe.Pointer)(bsp))
	t.Log(pack2Iface(unpackIface(b)._type, *(*unsafe.Pointer)(arrayAt(bsp, 2, 8))), b)
	if **(**int)(bsp) != a {
		t.Fatal(**(**int)(bsp))
	}
	tree := NewTree2(false)
	beg, ok := tree.Insert("10", b)
	if !ok {
		t.Fatal()
	}
	if beg.GetKey() != "10" || beg.GetVal() != b {
		t.Fatal(beg.GetKey(), beg.GetVal())
	}
}
