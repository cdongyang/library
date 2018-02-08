package rbtree

import "unsafe"

type pair struct {
	key interface{}
	val interface{}
}

func nodeP2mapNodeP(n *node) *MapNode {
	return (*MapNode)(unsafe.Pointer(n))
}

func mapNodeP2node(n *MapNode) *node {
	return (*node)(unsafe.Pointer(n))
}

type MapNode struct {
	node
}

func (n *MapNode) GetData() (key interface{}, val interface{}) {
	p := n.node.GetData().(pair)
	return p.key, p.val
}

func (n *MapNode) Next() *MapNode {
	return nodeP2mapNodeP(n.node.Next())
}

func (n *MapNode) Last() *MapNode {
	return nodeP2mapNodeP(n.node.Last())
}

func NewMap(key, val interface{}, compare func(a, b interface{}) int) *Map {
	return newMap(true, key, val, compare)
}

func NewMultiMap(key, val interface{}, compare func(a, b interface{}) int) *Map {
	return newMap(false, key, val, compare)
}

func newMap(unique bool, key, val interface{}, compare func(a, b interface{}) int) *Map {
	var m = &Map{}
	m.init(unique, interface2type(key), interface2type(val), compare)
	return m
}

type Map struct {
	tree
}

func (m *Map) Begin() *MapNode {
	return nodeP2mapNodeP(m.tree.Begin())
}

func (m *Map) End() *MapNode {
	return nodeP2mapNodeP(m.tree.End())
}

func (m *Map) Find(key interface{}) *MapNode {
	return nodeP2mapNodeP(m.tree.Find(key))
}

func (m *Map) Insert(key interface{}, val interface{}) (*MapNode, bool) {
	node, ok := m.tree.Insert(pair{key, val})
	return nodeP2mapNodeP(node), ok
}

func (m *Map) EraseNode(node *MapNode) {
	m.tree.Erase(mapNodeP2node(node))
}

func (m *Map) LowerBound(key interface{}) *MapNode {
	return nodeP2mapNodeP(m.tree.LowerBound(key))
}

func (m *Map) UpperBound(key interface{}) *MapNode {
	return nodeP2mapNodeP(m.tree.UpperBound(key))
}

func (m *Map) EqualRange(key interface{}) (beg, end *MapNode) {
	p1, p2 := m.tree.EqualRange(key)
	return nodeP2mapNodeP(p1), nodeP2mapNodeP(p2)
}

func (m *Map) EraseNodeRange(beg, end *MapNode) (count int) {
	return m.tree.EraseNodeRange(mapNodeP2node(beg), mapNodeP2node(end))
}
