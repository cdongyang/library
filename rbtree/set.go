package rbtree

import (
	"unsafe"
)

// SetNode is the iterator of set, but it's not thread safe,
// if the node was erase from set, calling it's method may panic
type SetNode struct {
	n _node
}

// GetData get the data of SetNode, but you should not hold the return interface{},
// instead, you should do type assert immediately when after call this method
func (n SetNode) GetData() interface{} {
	return n.n.tree.getKey(n.n.node)
}

// Next return the next node of current node.
// it will panic if current node equal to set.End().
func (n SetNode) Next() SetNode {
	return SetNode{n.n.Next()}
}

// Last return the last node of current node.
// it will panic if current node equal to set.Begin().
func (n SetNode) Last() SetNode {
	return SetNode{n.n.Last()}
}

// GetSet return the Set that current node belong to.
func (n SetNode) GetSet() *Set {
	return (*Set)(unsafe.Pointer(n.n.tree))
}

type Set struct {
	tree
}

// NewSet return a unique set with data type and compare func,
// the return set has been executed init func.
func NewSet(data interface{}, compare func(a, b interface{}) int) *Set {
	var s = &Set{}
	s.Init(true, data, compare)
	return s
}

// NewSet return a not unique set with data type and compare func,
// the return set has been executed init func.
func NewMultiSet(data interface{}, compare func(a, b interface{}) int) *Set {
	var s = &Set{}
	s.Init(false, data, compare)
	return s
}

func (s *Set) pack(n _node) SetNode {
	return SetNode{n: n}
}

// Init init the set, function NewSet and NewMultiSet will calle it,
// only the first call of this function will have an affect on set
func (s *Set) Init(unique bool, data interface{}, compare func(a, b interface{}) int) {
	s.tree.Init(unique, data, nil, compare)
}

// Begin return the first SetNode of set,
// if set is empty, it return set.End()
func (s *Set) Begin() SetNode {
	return s.pack(s.tree.Begin())
}

// End represent the end of set,but it isn't a real node
func (s *Set) End() SetNode {
	return s.pack(s.tree.End())
}

func (s *Set) EqualRange(data interface{}) (beg, end SetNode) {
	a, b := s.tree.EqualRange(data)
	return s.pack(a), s.pack(b)
}

// EraseNode erase a SetNode from tree,
// if SetNode has been erased, calling will panic
func (s *Set) EraseNode(n SetNode) {
	s.tree.EraseNode(n.n)
}

func (s *Set) EraseNodeRange(beg, end SetNode) (count int) {
	return s.tree.EraseNodeRange(beg.n, end.n)
}

func (s *Set) Find(data interface{}) SetNode {
	return s.pack(s.tree.Find(data))
}

func (s *Set) Insert(data interface{}) (SetNode, bool) {
	n, ok := s.tree.Insert(data, nil)
	return s.pack(n), ok
}

func (s *Set) LowerBound(data interface{}) SetNode {
	return s.pack(s.tree.LowerBound(data))
}

func (s *Set) UpperBound(data interface{}) SetNode {
	return s.pack(s.tree.UpperBound(data))
}

func (s *Set) Count(data interface{}) (count int) {
	return s.tree.Count(data)
}

func (s *Set) Erase(data interface{}) (count int) {
	return s.tree.Erase(data)
}
