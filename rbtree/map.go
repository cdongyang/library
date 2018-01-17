package rbtree

import "unsafe"

type Pair struct {
	Key, Value interface{}
}

func NewPair(key, value interface{}) Pair {
	return Pair{Key: key, Value: value}
}

//type of MapNode.data must be *Pair
type MapNode struct {
	RBTreeNode
	Pair
}

func (node *MapNode) Next() Iterator {
	return node.GetTree().(*Map).Next(node)
}

func (node *MapNode) Last() Iterator {
	return node.GetTree().(*Map).Last(node)
}

func (node *MapNode) CopyData(src Iterator) {
	node.Pair = src.(*MapNode).Pair
}

func (node *MapNode) GetData() interface{} {
	return node.Pair
}

func (node *MapNode) GetKey() interface{} {
	return node.Pair.Key
}

func (node *MapNode) GetValue() interface{} {
	return node.Pair.Value
}

func (node *MapNode) SetValue(value interface{}) {
	node.Pair.Value = value
}

type Map struct {
	RBTree
}

func (m *Map) Insert(data interface{}) (Iterator, bool) {
	_ = data.(Pair)
	iter, ok := m.RBTree.insert(data, func(key interface{}) int {
		return m.compare(data.(Pair).Key, key)
	})
	return iface2iterator(iter), ok
}

func SameMapNode(a Iterator, b Iterator) bool {
	aa, aok := a.(*MapNode)
	bb, bok := b.(*MapNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

func NewMap(compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewRBTreer(
		mp,
		header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		func(data interface{}) Iterator {
			return &MapNode{Pair: data.(Pair)}
		},
		func(Iterator) {
		},
		compare,
		//SameMapNode,
		true,
	).(*Map)
}

func NewCustomMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewRBTreer(mp, header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		//SameMapNode,
		true).(*Map)
}

func NewMultiMap(compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewRBTreer(
		mp,
		header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		func(data interface{}) Iterator {
			return &MapNode{Pair: data.(Pair)}
		},
		func(Iterator) {
		},
		compare,
		//SameMapNode,
		false,
	).(*Map)
}

func NewCustomMultiMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewRBTreer(mp, header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		//SameMapNode,
		false).(*Map)
}
