package rbtree

import (
	"unsafe"
)

var pairType = interface2type(Pair{})

// Pair is the data of MapNode
type Pair struct {
	Key, Value interface{}
}

// NewPair create a new Pair with key and value
func NewPair(key, value interface{}) Pair {
	return Pair{Key: key, Value: value}
}

// MapNode is the node of Map,it implement Iterator
type MapNode struct {
	_node
	Pair
}

// Next return next Iterator of this
func (node *MapNode) Next() Iterator {
	return node.GetTree().(*Map).Next(node)
}

// Last return last Iterator of this
func (node *MapNode) Last() Iterator {
	return node.GetTree().(*Map).Last(node)
}

// GetData get the data of this
func (node *MapNode) GetData() interface{} {
	// do this to avoid heap alloc
	return eface2interface(eface{pairType, unsafe.Pointer(&node.Pair)})
}

// GetKey get the compare key of this
func (node *MapNode) GetKey() interface{} {
	return node.Pair.Key
}

// GetValue get the value of this
func (node *MapNode) GetValue() interface{} {
	return node.Pair.Value
}

// SetValue set the value of this
func (node *MapNode) SetValue(value interface{}) {
	node.Pair.Value = value
}

//CopyData copy the node data to this from src
func (node *MapNode) CopyData(src Iterator) {
	node.Pair = src.(*MapNode).Pair
}

var mapNodeOffset uintptr

func init() {
	var header = &MapNode{}
	mapNodeOffset = uintptr(unsafe.Pointer(&header._node)) - uintptr(unsafe.Pointer(header))
}

// Map is a set of key-value Pair with red-black tree data struct, it implement Treer
// you can use the Unique method to find out wheather the Map key is unique
// you can use NewMap or NewCustomMap to create a unique Map
// you can use NewMultiMap or NewCustomMap to create a not unique Map
type Map struct {
	_tree
}

// Insert is rewrite to get data key and type assert data to Pair
func (m *Map) Insert(data interface{}) (Iterator, bool) {
	iter, ok := m.Tree.insert(data, interface2pointer(data.(Pair).Key))
	return m.pointer2iterator(iter), ok
}

func getMapNodeKeyPointer(p unsafe.Pointer) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&(*MapNode)(p).Key)).pointer
}

// NewMap create a new unique Map with compare func
// the compare func
//	return negative int when a < b
//	return 0 when a == b and
//	return positive int when a > b
func NewMap(compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	return NewTreer(
		mp,
		&MapNode{},
		mapNodeOffset,
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

// NewCustomMap create a new unique Map with newNode, deleteNode and compare func
// you can define you own func to create a new node and a new node
// such as with a sync.Pool to reduce the pressure of GC
func NewCustomMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Map {
	var set = &Map{}
	return NewTreer(set, &MapNode{}, mapNodeOffset, newNode, deleteNode, compare, getMapNodeKeyPointer, true).(*Map)
}

// NewMultiMap create a new not unique Map with compare func
// the compare func
//	return negative int when a < b
//	return 0 when a == b and
//	return positive int when a > b
func NewMultiMap(compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	return NewTreer(
		mp,
		&MapNode{},
		mapNodeOffset,
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

// NewCustomMultiMap create a new not unique Map with newNode, deleteNode and compare func
// you can define you own func to create a new node and a new node
// such as with a sync.Pool to reduce the pressure of GC
func NewCustomMultiMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Map {
	var mp = &Map{}
	return NewTreer(mp, &MapNode{}, mapNodeOffset, newNode, deleteNode, compare, getMapNodeKeyPointer, false).(*Map)
}
