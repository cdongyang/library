package rbtree

import (
	"unsafe"
)

type Pair struct {
	Key, Value interface{}
}

func NewPair(key, value interface{}) Pair {
	return Pair{Key: key, Value: value}
}

//type of MapNode.data must be *Pair
type MapNode struct {
	Node
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
	Tree
}

func (m *Map) Insert(data interface{}) (Iterator, bool) {
	var pair = *(*Pair)(interface2pointer(data))
	iter, ok := m.Tree.insert(data, interface2pointer(pair.Key))
	return m.pointer2iterator(iter), ok
}

func getMapNodeKeyPointer(p unsafe.Pointer) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&(*MapNode)(p).Key)).pointer
}

func NewMap(compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewTreer(
		mp,
		header,
		uintptr(unsafe.Pointer(&header.Node))-uintptr(unsafe.Pointer(header)),
		func(data interface{}) Iterator {
			return &MapNode{Pair: data.(Pair)}
		},
		func(Iterator) {
		},
		compare,
		getMapNodeKeyPointer,
		true,
	).(*Map)
}

func NewCustomMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewTreer(mp, header,
		uintptr(unsafe.Pointer(&header.Node))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		getMapNodeKeyPointer,
		true).(*Map)
}

func NewMultiMap(compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewTreer(
		mp,
		header,
		uintptr(unsafe.Pointer(&header.Node))-uintptr(unsafe.Pointer(header)),
		func(data interface{}) Iterator {
			return &MapNode{Pair: data.(Pair)}
		},
		func(Iterator) {
		},
		compare,
		getMapNodeKeyPointer,
		false,
	).(*Map)
}

func NewCustomMultiMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	var header = &MapNode{}
	return NewTreer(mp, header,
		uintptr(unsafe.Pointer(&header.Node))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		getMapNodeKeyPointer,
		false).(*Map)
}
