package rbtree_test

import (
	"sort"
	"testing"

	"github.com/cdongyang/library/algorithm"
	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

func testSet(t *testing.T, length int, unique bool) {
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
	)
	var set = rbtree.NewSet(compare)
	var count = make(map[int]int, len(intSlice1K))
	// test empty set and empty set Begin and End
	if !set.Empty() {
		t.Fatal("empty")
	}
	if set.Begin() != set.End() {
		t.Fatal("empty set begin and end error")
	}

	//test RBTreer
	var _ rbtree.RBTreer = set
	var iter = set.End()
	var index int

	//test empty set LowerBound method
	iter = set.LowerBound(intSlice1K[0])
	if !set.SameIterator(iter, set.End()) {
		t.Fatal("empty set LowerBound error")
	}
	//test empty set UpperBound method
	iter = set.UpperBound(intSlice1K[0])
	if !set.SameIterator(iter, set.End()) {
		t.Fatal("empty set UpperBound error")
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
		_, ok := set.Insert(val)
		if !set.Unique() && !ok || set.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if set.Begin().GetKey() != minVal {
			t.Fatal("leftmost error")
		}
		if set.End().Last().GetKey() != maxVal {
			t.Fatal("rightmost error")
		}
		if set.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if set.Count(val) != count[val] {
			t.Fatal("count error", set.Count(val), count[val])
		}
		_, size := set.Check()
		if size != set.Size() {
			t.Fatal("size error", size, set.Size())
		}
	}

	// test Compare
	if set.Compare(set.Begin().GetKey(), minVal) != 0 {
		t.Fatal("Compare error")
	}
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.IntSlice(sortSlice).Sort()

	//test LowerBound method
	iter = set.LowerBound(intSlice1K[0])
	index = algorithm.LowerBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if set.SameIterator(iter, set.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
		//ok
	} else {
		t.Fatal("LowerBound error")
	}
	//test UpperBound method
	iter = set.UpperBound(intSlice1K[0])
	index = algorithm.UpperBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if set.SameIterator(iter, set.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
	}
	//test Begin and EndNode method
	if set.Begin().GetKey() != sortSlice[0] {
		t.Fatal("begin error", set.Begin().GetKey(), sortSlice[0])
	}
	if set.End().Last().GetKey() != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", set.Begin().GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := set.Begin(); !set.SameIterator(it, set.End()); it = it.Next() {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go through error", it.GetKey(), sortSlice[i])
		}
		if set.Unique() {
			for i+1 < len(sortSlice) && sortSlice[i+1] == sortSlice[i] {
				i++
			}
		}
		i++
	}
	//test End and Last method
	i = len(sortSlice) - 1
	for it := set.End().Last(); ; it = it.Last() {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go back set error", it.GetKey(), sortSlice[i])
		}
		if set.Unique() {
			for i > 0 && sortSlice[i-1] == sortSlice[i] {
				i--
			}
		}
		i--
		if set.SameIterator(set.Begin(), it) {
			break
		}
	}
	//test Find method
	iter = set.Find(intSlice1K[0])
	index = sort.SearchInts(sortSlice, intSlice1K[0])
	if set.SameIterator(iter, set.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == sortSlice[index] {
		//ok
	} else {
		t.Fatal("Find error")
	}
	if set.Find(max) != set.End() {
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
		num := set.Erase(val)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if set.Count(val) != count[val] {
			t.Fatal("count error", set.Count(val), count[val])
		}
		_, size := set.Check()
		if size != set.Size() {
			t.Fatal("size error")
		}
	}
	//test Empty method
	if !set.Empty() {
		t.Fatal("empty error")
	}
	count = make(map[int]int)
	for _, val := range intSlice1K {
		_, ok := set.Insert(val)
		if !set.Unique() && !ok || set.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if set.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if set.Count(val) != count[val] {
			t.Fatal("count error", set.Count(val), count[val])
		}
		_, size := set.Check()
		if size != set.Size() {
			t.Fatal("size error", size, set.Size())
		}
	}
	//test EqualRange and EraseIteratorRange
	for _, val := range intSlice1K {
		//fmt.Println("erase", val)
		beg, end := set.EqualRange(val)
		num := set.EraseIteratorRange(beg, end)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if set.Count(val) != count[val] {
			t.Fatal("count error", set.Count(val), count[val])
		}
		_, size := set.Check()
		if size != set.Size() {
			t.Fatal("size error")
		}
	}
	if set.Size() != 0 || !set.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
	//test Clear method
	for _, val := range intSlice1K {
		_, _ = set.Insert(val)
	}
	set.Clear()
	if set.Size() != 0 || !set.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}
func TestSet(t *testing.T) {
	for i := 0; i < 500; i++ {
		testSet(t, i+1, i%2 == 0)
	}
}
