package rbtree

import (
	"unsafe"
)

func nodeP2setNodeP(n *node) *SetNode {
	return (*SetNode)(unsafe.Pointer(n))
}

func setNodeP2node(n *SetNode) *node {
	return (*node)(unsafe.Pointer(n))
}

type SetNode struct {
	node node
}

func (n *SetNode) GetData() interface{} {
	return n.node.GetData()
}

func (n *SetNode) Next() *SetNode {
	return nodeP2setNodeP(n.node.Next())
}

func (n *SetNode) Last() *SetNode {
	return nodeP2setNodeP(n.node.Last())
}

func NewSet(data interface{}, compare func(a, b interface{}) int) *Set {
	return newSet(true, data, compare)
}

func NewMultiSet(data interface{}, compare func(a, b interface{}) int) *Set {
	return newSet(false, data, compare)
}

func newSet(unique bool, data interface{}, compare func(a, b interface{}) int) *Set {
	var s = &Set{}
	s.init(unique, interface2type(data), nil, compare)
	return s
}

type Set struct {
	tree
}

func (s *Set) Begin() *SetNode {
	return nodeP2setNodeP(s.tree.Begin())
}

func (s *Set) End() *SetNode {
	return nodeP2setNodeP(s.tree.End())
}

func (s *Set) Find(data interface{}) *SetNode {
	return nodeP2setNodeP(s.tree.Find(data))
}

func (s *Set) Insert(data interface{}) (*SetNode, bool) {
	node, ok := s.tree.Insert(data)
	return nodeP2setNodeP(node), ok
}

func (s *Set) EraseNode(node *SetNode) {
	s.tree.Erase(setNodeP2node(node))
}

func (s *Set) LowerBound(data interface{}) *SetNode {
	return nodeP2setNodeP(s.tree.LowerBound(data))
}

func (s *Set) UpperBound(data interface{}) *SetNode {
	return nodeP2setNodeP(s.tree.UpperBound(data))
}

func (s *Set) EqualRange(data interface{}) (beg, end *SetNode) {
	p1, p2 := s.tree.EqualRange(data)
	return nodeP2setNodeP(p1), nodeP2setNodeP(p2)
}

func (s *Set) EraseNodeRange(beg, end *SetNode) (count int) {
	return s.tree.EraseNodeRange(setNodeP2node(beg), setNodeP2node(end))
}
