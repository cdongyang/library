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
	var (
		compare = func(a, b interface{}) int {
			return a.(int) - b.(int)
		}
	)
	var mp = rbtree.NewMap(compare)
	var count = make(map[int]int, len(intSlice1K))
	// test empty mp and empty mp Begin and End
	if !mp.Empty() {
		t.Fatal("empty")
	}
	if mp.Begin() != mp.End() {
		t.Fatal("empty mp begin and end error")
	}

	//test RBTreer
	var _ rbtree.RBTreer = mp
	var iter = mp.End()
	var index int

	//test empty mp LowerBound method
	iter = mp.LowerBound(intSlice1K[0])
	if !mp.SameIterator(iter, mp.End()) {
		t.Fatal("empty mp LowerBound error")
	}
	//test empty mp UpperBound method
	iter = mp.UpperBound(intSlice1K[0])
	if !mp.SameIterator(iter, mp.End()) {
		t.Fatal("empty mp UpperBound error")
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
		_, ok := mp.Insert(rbtree.Pair{Key: val, Value: true})
		if !mp.Unique() && !ok || mp.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if mp.Begin().GetKey() != minVal {
			t.Fatal("leftmost error")
		}
		if mp.End().Last().GetKey() != maxVal {
			t.Fatal("rightmost error")
		}
		if mp.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if mp.Count(val) != count[val] {
			t.Fatal("count error", mp.Count(val), count[val])
		}
		_, size := mp.Check()
		if size != mp.Size() {
			t.Fatal("size error", size, mp.Size())
		}
	}

	// test Compare
	if mp.Compare(mp.Begin().GetKey(), minVal) != 0 {
		t.Fatal("Compare error")
	}
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	sort.IntSlice(sortSlice).Sort()

	//test LowerBound method
	iter = mp.LowerBound(intSlice1K[0])
	index = algorithm.LowerBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if mp.SameIterator(iter, mp.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
		//ok
	} else {
		t.Fatal("LowerBound error")
	}
	//test UpperBound method
	iter = mp.UpperBound(intSlice1K[0])
	index = algorithm.UpperBound(algorithm.SearchIntSlice{sortSlice, intSlice1K[0]})
	if mp.SameIterator(iter, mp.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == intSlice1K[0] {
	}
	//test Begin and EndNode method
	if mp.Begin().GetKey() != sortSlice[0] {
		t.Fatal("begin error", mp.Begin().GetKey(), sortSlice[0])
	}
	if mp.End().Last().GetKey() != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", mp.Begin().GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := mp.Begin(); !mp.SameIterator(it, mp.End()); it = it.Next() {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go through error", it.GetKey(), sortSlice[i])
		}
		if mp.Unique() {
			for i+1 < len(sortSlice) && sortSlice[i+1] == sortSlice[i] {
				i++
			}
		}
		i++
	}
	//test End and Last method
	i = len(sortSlice) - 1
	for it := mp.End().Last(); ; it = it.Last() {
		if it.GetKey() != sortSlice[i] {
			t.Fatal("go back mp error", it.GetKey(), sortSlice[i])
		}
		if mp.Unique() {
			for i > 0 && sortSlice[i-1] == sortSlice[i] {
				i--
			}
		}
		i--
		if mp.SameIterator(mp.Begin(), it) {
			break
		}
	}
	//test Find method
	iter = mp.Find(intSlice1K[0])
	index = sort.SearchInts(sortSlice, intSlice1K[0])
	if mp.SameIterator(iter, mp.End()) && index == len(sortSlice) {
		//ok
	} else if iter.GetKey() == sortSlice[index] {
		//ok
	} else {
		t.Fatal("Find error")
	}
	if mp.Find(max) != mp.End() {
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
		num := mp.Erase(val)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if mp.Count(val) != count[val] {
			t.Fatal("count error", mp.Count(val), count[val])
		}
		_, size := mp.Check()
		if size != mp.Size() {
			t.Fatal("size error")
		}
	}
	//test Empty method
	if !mp.Empty() {
		t.Fatal("empty error")
	}
	count = make(map[int]int)
	for _, val := range intSlice1K {
		_, ok := mp.Insert(rbtree.Pair{Key: val, Value: true})
		if !mp.Unique() && !ok || mp.Unique() && ok == (count[val] != 0) {
			t.Fatal("insert error", ok, count[val], val)
		}
		if mp.Unique() {
			count[val] = 1
		} else {
			count[val]++
		}
		if mp.Count(val) != count[val] {
			t.Fatal("count error", mp.Count(val), count[val])
		}
		_, size := mp.Check()
		if size != mp.Size() {
			t.Fatal("size error", size, mp.Size())
		}
	}
	//test EqualRange and EraseIteratorRange
	for _, val := range intSlice1K {
		//fmt.Println("erase", val)
		beg, end := mp.EqualRange(val)
		num := mp.EraseIteratorRange(beg, end)
		if num != count[val] {
			t.Fatal("erase error", num, count[val], val)
		}
		delete(count, val)
		if mp.Count(val) != count[val] {
			t.Fatal("count error", mp.Count(val), count[val])
		}
		_, size := mp.Check()
		if size != mp.Size() {
			t.Fatal("size error")
		}
	}
	if mp.Size() != 0 || !mp.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
	//test Clear method
	for _, val := range intSlice1K {
		_, _ = mp.Insert(rbtree.Pair{Key: val, Value: true})
	}
	mp.Clear()
	if mp.Size() != 0 || !mp.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}
func TestMap(t *testing.T) {
	for i := 0; i < 500; i++ {
		testMap(t, i+1, i%2 == 0)
	}
}
