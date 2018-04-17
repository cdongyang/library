package rbtree

import (
	"unsafe"
)

// MapNode is the iterator of Map, but it's not thread safe,
// if the node was erase from map, calling it's method may panic
type MapNode struct {
	n _node
}

// GetKey get the key of MapNode, but you should not hold the return interface{},
// instead, you should do type assert immediately when after call this method
func (n MapNode) GetKey() interface{} {
	return n.n.tree.getKey(n.n.node)
}

// GetVal get the value of MapNode, but you should not hold the return interface{},
// instead, you should do type assert immediately when after call this method
func (n MapNode) GetVal() interface{} {
	return n.n.tree.getVal(n.n.node)
}

func (n MapNode) GetData() (key, val interface{}) {
	return n.GetKey(), n.GetVal()
}

func (n MapNode) SetVal(val interface{}) {
	n.n.tree.setVal(n.n.node, val)
}

func (n MapNode) Next() MapNode {
	return MapNode{n.n.Next()}
}

func (n MapNode) Last() MapNode {
	return MapNode{n.n.Last()}
}

func (n MapNode) GetMap() *Map {
	return (*Map)(unsafe.Pointer(n.n.tree))
}

type Map struct {
	tree
}

func NewMap(key, val interface{}, compare func(a, b interface{}) int) *Map {
	var m = &Map{}
	m.Init(true, key, val, compare)
	return m
}

func NewMultiMap(key, val interface{}, compare func(a, b interface{}) int) *Map {
	var s = &Map{}
	s.Init(false, key, val, compare)
	return s
}

func (s *Map) pack(n _node) MapNode {
	return MapNode{n: n}
}

func (s *Map) Init(unique bool, key, val interface{}, compare func(a, b interface{}) int) {
	s.tree.Init(unique, key, val, compare)
}

func (s *Map) Begin() MapNode {
	return s.pack(s.tree.Begin())
}

func (s *Map) End() MapNode {
	return s.pack(s.tree.End())
}

func (s *Map) EqualRange(key interface{}) (beg, end MapNode) {
	a, b := s.tree.EqualRange(key)
	return s.pack(a), s.pack(b)
}

func (s *Map) EraseNode(n MapNode) {
	s.tree.EraseNode(n.n)
}

func (s *Map) EraseNodeRange(beg, end MapNode) (count int) {
	return s.tree.EraseNodeRange(beg.n, end.n)
}

func (s *Map) Find(key interface{}) MapNode {
	return s.pack(s.tree.Find(key))
}

func (s *Map) Insert(key interface{}, val interface{}) (MapNode, bool) {
	n, ok := s.tree.Insert(key, val)
	return s.pack(n), ok
}

func (s *Map) LowerBound(key interface{}) MapNode {
	return s.pack(s.tree.LowerBound(key))
}

func (s *Map) UpperBound(key interface{}) MapNode {
	return s.pack(s.tree.UpperBound(key))
}

func (s *Map) Count(key interface{}) (count int) {
	return s.tree.Count(key)
}

func (s *Map) Erase(key interface{}) (count int) {
	return s.tree.Erase(key)
}
