package rbtree_test

import (
	"sort"
	"testing"

	"github.com/cdongyang/library/algorithm"
	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

func testMap(t *testing.T, length int, unique bool) {
	// init
	var rand = randint.Rand{First: 23456, Add: 12345, Mod: 1000}
	var max = rand.Int()%length + 1
	var intSlice1K = make([]int, length)
	for i := range intSlice1K {
		intSlice1K[i] = rand.Int() % max
	}
	//fmt.Println("offsetNode:", offsetNode)
	var tree = rbtree.NewMap(rbtree.CompareInt)
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
	var _ rbtree.Treer = tree
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
		_, ok := tree.Insert(rbtree.Pair{Key: val, Value: true})
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
		_, ok := tree.Insert(rbtree.Pair{Key: val, Value: true})
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
		_, _ = tree.Insert(rbtree.Pair{Key: val, Value: true})
	}
	tree.Clear()
	if tree.Size() != 0 || !tree.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}
func TestMap(t *testing.T) {
	for i := 0; i < 500; i++ {
		testMap(t, i+1, i%2 == 0)
	}
}
