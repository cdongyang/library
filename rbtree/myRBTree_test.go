package rbtree_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/cdongyang/library/rbtree"
)

type RBTree = rbtree.RBTree
type RBTreeNode = rbtree.RBTreeNode
type Iterator = rbtree.Iterator

type node struct {
	RBTreeNode
	key int
}

func (n *node) GetKey() interface{} {
	return n.key
}
func (n *node) SetKey(k interface{}) {
	n.key = k.(int)
}

type mytree struct {
	RBTree
}

func myNewTree() {
}

func TestRBTree(t *testing.T) {
	testRBTree(t, 100)
}

func testRBTree(t *testing.T, length int) {
	var max = rand.Int()%length + 1
	var intSlice1K = make([]int, length)
	for i := range intSlice1K {
		intSlice1K[i] = rand.Int() % max
	}
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewRBTree((*node)(nil), compare, newElem, deleteElem)
	var exists = make(map[int]bool, len(intSlice1K))
	if !tree.Empty() {
		panic("empty")
	}
	var iter = tree.End()
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).GetKey() != intSlice1K[0] {
		if iter.(*node).GetKey().(int) < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.(*node).GetKey(), intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).GetKey().(int) < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.(*node).GetKey(), intSlice1K[0])
	}
	//test Insert method
	for _, val := range intSlice1K {
		ok := tree.Insert(val)
		if ok == exists[val] {
			panic("insert error")
		}
		exists[val] = true
	}
	var sortSlice = make([]int, len(intSlice1K))
	copy(sortSlice, intSlice1K)
	/*sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i] < sortSlice[j]
	})*/
	sort.IntSlice(sortSlice).Sort()
	var uniqueN = 1
	for i := range sortSlice {
		if i > 0 && sortSlice[i] != sortSlice[i-1] {
			sortSlice[uniqueN] = sortSlice[i]
			uniqueN++
		}
	}
	sortSlice = sortSlice[:uniqueN]
	//test LowerBound method
	iter = tree.LowerBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).GetKey() != intSlice1K[0] {
		if iter.(*node).GetKey().(int) < intSlice1K[0] {
			t.Fatal("LowerBound error", iter.(*node).GetKey(), intSlice1K[0])
		}
	}
	//test UpperBound method
	iter = tree.UpperBound(intSlice1K[0])
	if iter == tree.End() {
		//t.Fatal("UpperBound to End", intSlice1K[0], sortSlice)
	} else if iter.(*node).GetKey().(int) < intSlice1K[0] {
		t.Fatal("UpperBound error", iter.(*node).GetKey(), intSlice1K[0])
	}
	//test Begin and EndNode method
	if tree.Begin().(*node).GetKey() != sortSlice[0] {
		t.Fatal("begin error", tree.Begin().(*node).GetKey(), sortSlice[0])
	}
	if tree.EndNode().(*node).GetKey() != sortSlice[len(sortSlice)-1] {
		t.Fatal("endNode error", tree.Begin().(*node).GetKey(), sortSlice[len(sortSlice)-1])
	}
	//test Begin and End and Next method
	var i int
	for it := tree.Begin(); it != tree.End(); it = it.Next(it) {
		if it.(*node).GetKey() != sortSlice[i] {
			t.Fatal("go through error", it.(*node).GetKey(), sortSlice[i])
		}
		i++
	}
	//test EndNode and End and Last method
	i = len(sortSlice) - 1
	for it := tree.EndNode(); it != tree.End(); it = it.Last(it) {
		if it.(*node).GetKey() != sortSlice[i] {
			t.Fatal("go back tree error", it.(*node).GetKey(), sortSlice[i])
		}
		i--
	}
	//test Find method
	iter = tree.Find(intSlice1K[0])
	if iter == tree.End() {
		t.Fatal("find error", intSlice1K[0])
	} else if iter.(*node).GetKey() != intSlice1K[0] {
		t.Fatal("find not equal", intSlice1K[0])
	}
	if tree.Find(max) != tree.End() {
		t.Fatal("find max error", max)
	}
	//test Erase method
	for _, val := range intSlice1K {
		ok := tree.Erase(val)
		if ok != exists[val] {
			t.Fatal("erase error")
		}
		delete(exists, val)
	}
	//test Empty method
	if !tree.Empty() {
		t.Fatal("empty error")
	}
	exists = make(map[int]bool)
	for _, val := range intSlice1K {
		ok := tree.Insert(val)
		if ok == exists[val] {
			t.Fatal("insert error")
		}
		exists[val] = true
	}
	//test Clear method
	tree.Clear()
	if tree.Size() != 0 || !tree.Empty() {
		t.Fatal("clear error,size != 0 or not empty")
	}
}
