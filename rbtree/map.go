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
	iter, ok := m.RBTree.insert(data, func(key unsafe.Pointer) int {
		return m.compare((*eface)((*eface)(unsafe.Pointer(&data)).pointer).pointer, key) //interface2pointer取Pair的第一个interface的pointer
	})
	return eface2iterator(iter), ok
}

func getMapNodeKeyPointer(p unsafe.Pointer) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&(*MapNode)(p).Key)).pointer
}

func NewMap(compare func(a, b unsafe.Pointer) int) *Map {
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
		getMapNodeKeyPointer,
		true,
	).(*Map)
}

func NewCustomMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewRBTreer(mp, header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		//SameMapNode,
		getMapNodeKeyPointer,
		true).(*Map)
}

func NewMultiMap(compare func(a, b unsafe.Pointer) int) *Map {
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
		getMapNodeKeyPointer,
		false,
	).(*Map)
}

func NewCustomMultiMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewRBTreer(mp, header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		//SameMapNode,
		getMapNodeKeyPointer,
		false).(*Map)
}
