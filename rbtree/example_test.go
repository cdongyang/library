package rbtree_test

import (
	"sort"
	"testing"
	"unsafe"

	"github.com/cdongyang/library/rbtree"
)

type IntSetNode struct {
	n rbtree.SetNode
}

func (n IntSetNode) GetData() int {
	return n.n.GetData().(int)
}

func (n IntSetNode) Next() IntSetNode {
	return IntSetNode{n: n.n.Next()}
}

func (n IntSetNode) Last() IntSetNode {
	return IntSetNode{n: n.n.Next()}
}

func (n IntSetNode) GetSet() *IntSet {
	return (*IntSet)(unsafe.Pointer(n.GetSet()))
}

type IntSet struct {
	set rbtree.Set
}

func NewIntSet(compare func(a, b int) int) *IntSet {
	var s = &IntSet{}
	s.Init(true, compare)
	return s
}

func NewMultiIntSet(compare func(a, b int) int) *IntSet {
	var s = &IntSet{}
	s.Init(false, compare)
	return s
}

func (s *IntSet) pack(n rbtree.SetNode) IntSetNode {
	return IntSetNode{n: n}
}

func (s *IntSet) Init(unique bool, compare func(a, b int) int) {
	s.set.Init(unique, int(0), func(a, b interface{}) int {
		return compare(a.(int), b.(int))
	})
}

func (s *IntSet) Begin() IntSetNode {
	return s.pack(s.set.Begin())
}

func (s *IntSet) End() IntSetNode {
	return s.pack(s.set.End())
}

func (s *IntSet) EqualRange(data int) (beg, end IntSetNode) {
	a, b := s.set.EqualRange(data)
	return s.pack(a), s.pack(b)
}

func (s *IntSet) EraseNode(n IntSetNode) {
	s.set.EraseNode(n.n)
}

func (s *IntSet) EraseNodeRange(beg, end IntSetNode) (count int) {
	return s.set.EraseNodeRange(beg.n, end.n)
}

func (s *IntSet) Find(data int) IntSetNode {
	return s.pack(s.set.Find(data))
}

func (s *IntSet) Insert(data int) (IntSetNode, bool) {
	// if data type is not direct interface, use NoescapeInterface to avoid data escape to heap,
	// thus reduce heap objects.
	// but if you don't know wheather which type is direct interface, don't use NoescapeInterface !!!
	n, ok := s.set.Insert(rbtree.NoescapeInterface(data))
	return s.pack(n), ok
}

func (s *IntSet) LowerBound(data int) IntSetNode {
	return s.pack(s.set.LowerBound(data))
}

func (s *IntSet) UpperBound(data int) IntSetNode {
	return s.pack(s.set.UpperBound(data))
}

func (s *IntSet) Count(data int) (count int) {
	return s.set.Count(data)
}

func (s *IntSet) Erase(data int) (count int) {
	return s.set.Erase(data)
}

func (s *IntSet) Size() int {
	return s.set.Size()
}

func TestIntSet(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
		var cpSlice = make([]int, len(slice))
		copy(cpSlice, slice)
		sort.IntSlice(cpSlice).Sort()
		s := NewIntSet(func(a, b int) int {
			return a - b
		})
		for _, val := range slice {
			s.Insert(val)
		}
		for it, i := s.Begin(), 0; it != s.End(); it = it.Next() {
			if it.GetData() != cpSlice[i] {
				t.Fatal(it.GetData(), cpSlice[i])
			}
			i++
		}
		for it := s.Begin(); it != s.End(); {
			tmp := it.GetData()
			it = it.Next()
			s.Erase(tmp)
		}
		if s.Size() != 0 {
			t.Fatal(s.Size())
		}
	})
	t.Run("escape", func(t *testing.T) {
		var x = 1
		s := NewIntSet(func(a, b int) int { return a - b })
		n := testing.AllocsPerRun(1000, func() {
			var tmp = x
			s.Insert(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("insert escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Find(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("find escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Erase(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("erase escape", n)
		}
	})
}

type intSet struct { // hide node
	set rbtree.Set
}

func NewintSet(compare func(a, b int) int) *intSet {
	var s = &intSet{}
	s.set.Init(true, int(0), func(a, b interface{}) int {
		return compare(a.(int), b.(int))
	})
	return s
}

func (s *intSet) Insert(data int) (ok bool) {
	_, ok = s.set.Insert(rbtree.NoescapeInterface(data))
	return ok
}

func (s *intSet) Find(data int) (ok bool) {
	return s.set.Find(data) == s.set.End()
}

func (s *intSet) Erase(data int) (ok bool) {
	return s.set.Erase(data) == 1
}

func (s *intSet) Range(f func(data int) bool) {
	for it := s.set.Begin(); it != s.set.End(); it = it.Next() {
		ok := f(it.GetData().(int))
		if !ok {
			return
		}
	}
}

func (s *intSet) Size() int {
	return s.set.Size()
}

func TestIntSetNoIterator(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
		var cpSlice = make([]int, len(slice))
		copy(cpSlice, slice)
		sort.IntSlice(cpSlice).Sort()
		s := NewintSet(func(a, b int) int {
			return a - b
		})
		for _, val := range slice {
			s.Insert(val)
		}
		i := 0
		s.Range(func(data int) bool {
			if data != cpSlice[i] {
				t.Fatal(data, cpSlice[i])
				return false
			}
			i++
			return true
		})
		s.Range(func(data int) bool {
			if !s.Erase(data) {
				t.Fatal(data)
				return false
			}
			return true
		})
		if s.Size() != 0 {
			t.Fatal(s.Size())
		}
	})
	t.Run("escape", func(t *testing.T) {
		var x = 1
		s := NewintSet(func(a, b int) int { return a - b })
		n := testing.AllocsPerRun(1000, func() {
			var tmp = x
			s.Insert(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("insert escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Find(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("find escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Erase(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("erase escape", n)
		}
	})
}
