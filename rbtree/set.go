package rbtree

import "unsafe"

// SetNode is the node of Set,it implement Iterator
type SetNode struct {
	_node
	data interface{}
}

// Next return next Iterator of this
func (node *SetNode) Next() Iterator {
	return node.GetTree().(*Set).Next(node)
}

// Last return last Iterator of this
func (node *SetNode) Last() Iterator {
	return node.GetTree().(*Set).Last(node)
}

// GetData get the data of this
func (node *SetNode) GetData() interface{} {
	return node.data
}

// GetKey get the compare key of this
func (node *SetNode) GetKey() interface{} {
	return node.data
}

//CopyData copy the node data to this from src
func (node *SetNode) CopyData(src Iterator) {
	node.data = src.(*SetNode).data
}

var setNodeOffset uintptr

func init() {
	var header = &SetNode{}
	setNodeOffset = uintptr(unsafe.Pointer(&header._node)) - uintptr(unsafe.Pointer(header))
}

// Set is a set of data with red-black tree data struct, it implement Treer
// you can use the Unique method to find out wheather the Set key is unique
// you can use NewSet or NewCustomSet to create a unique Set
// you can use NewMultiSet or NewCustomSet to create a not unique Set
type Set struct {
	_tree
}

//var GetSetKeyPointerCount = 0

func getSetNodeKeyPointer(p unsafe.Pointer) unsafe.Pointer {
	//GetSetKeyPointerCount++
	return (*eface)(unsafe.Pointer(&(*SetNode)(p).data)).pointer
}

// NewSet create a new unique Set with compare func
// the compare func
//	return negative int when a < b
//	return 0 when a == b and
//	return positive int when a > b
func NewSet(compare func(a, b unsafe.Pointer) int) *Set {
	var set = &Set{}
	return NewTreer(
		set,
		&SetNode{},
		setNodeOffset,
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		getSetNodeKeyPointer,
		true,
	).(*Set)
}

// NewCustomSet create a new unique Set with newNode, deleteNode and compare func
// you can define you own func to create a new node and a new node
// such as with a sync.Pool to reduce the pressure of GC
func NewCustomSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Set {
	var set = &Set{}
	return NewTreer(set, &SetNode{}, setNodeOffset, newNode, deleteNode, compare, getSetNodeKeyPointer, true).(*Set)
}

// NewMultiSet create a new not unique Set with compare func
// the compare func
//	return negative int when a < b
//	return 0 when a == b and
//	return positive int when a > b
func NewMultiSet(compare func(a, b unsafe.Pointer) int) *Set {
	var set = &Set{}
	return NewTreer(
		set,
		&SetNode{},
		setNodeOffset,
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		getSetNodeKeyPointer,
		false,
	).(*Set)
}

// NewCustomMultiSet create a new not unique Set with newNode, deleteNode and compare func
// you can define you own func to create a new node and a new node
// such as with a sync.Pool to reduce the pressure of GC
func NewCustomMultiSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(a, b unsafe.Pointer) int) *Set {
	var set = &Set{}
	return NewTreer(set, &SetNode{}, setNodeOffset, newNode, deleteNode, compare, getSetNodeKeyPointer, false).(*Set)
}
