package rbtree_test

import (
	"sort"
	"strconv"
	"testing"

	"github.com/cdongyang/library/algorithm"
	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

func NewMap(unique bool) *rbtree.Map {
	var newFun = rbtree.NewMultiMap
	if unique {
		newFun = rbtree.NewMap
	}
	return newFun("", new(int), func(a, b interface{}) int {
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
}

func testMap(t *testing.T, length int, unique bool) {
	// init
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1000}
	var max = rand.Int()%length + 1
	var intSlice = make([]int, length)
	for i := range intSlice {
		intSlice[i] = (rand.Int() % max) + 1
	}
	max++
	var equal = func(n rbtree.MapNode, val int) bool {
		return n.GetKey() == strconv.Itoa(val) && *(n.GetVal().(*int)) == val
	}
	//fmt.Println("offsetNode:", offsetNode)
	var tree = NewMap(unique)
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

func TestMap(t *testing.T) {
	var testN = 500
	t.Run("unique", func(t *testing.T) {
		for i := 0; i < testN; i++ {
			testMap(t, i+1, true)
		}
	})
	t.Run("not unique", func(t *testing.T) {
		for i := 0; i < testN; i++ {
			testMap(t, i+1, false)
		}
	})
	t.Run("unique 1e4", func(t *testing.T) {
		testMap(t, 1e4, true)
	})
	t.Run("not unique 1e4", func(t *testing.T) {
		testMap(t, 1e4, false)
	})
}
