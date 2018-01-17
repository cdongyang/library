package rbtree

import "unsafe"

type SetNode struct {
	RBTreeNode
	data interface{}
}

func (node *SetNode) Next() Iterator {
	return node.GetTree().(*Set).Next(node)
}

func (node *SetNode) Last() Iterator {
	return node.GetTree().(*Set).Last(node)
}

func (node *SetNode) GetData() interface{} {
	return node.data
}

func (node *SetNode) GetKey() interface{} {
	return node.data
}

func (node *SetNode) CopyData(src Iterator) {
	node.data = src.(*SetNode).data
}

// SameSetNode judge wheather two Iterator is the same node, but not equal node
// type assert to *SetNode to speed up
func SameSetNode(a, b Iterator) bool {
	aa, aok := a.(*SetNode)
	bb, bok := b.(*SetNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

// Set is a set of element
// you can use the Unique method to get wheather the Set element is unique
// you can use NewSet or NewCustomSet to create a unique Set
// you can use NewMultiSet or NewCustomSet to create a not unique Set
type Set struct {
	RBTree
}

// NewSet create a new unique Set with compare func
// the compare func return negative int when a < b, return 0 when a == b and return positive int when a > b
func NewSet(compare func(a, b interface{}) int) *Set {
	var set = &Set{}
	var header = &SetNode{}
	return NewRBTreer(
		set,
		header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		//SameSetNode,
		true,
	).(*Set)
}

// NewCustomSet create a new unique Set with newNode, deleteNode and compare func
// you can define you own func to create a new node and a new node
// such as with a sync.Pool to reduce the pressure of GC
func NewCustomSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Set {
	var set = &Set{}
	var header = &SetNode{}
	return NewRBTreer(set, header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		//SameSetNode,
		true).(*Set)
}

// NewMultiSet create a new not unique Set with compare func
// the compare func return negative int when a < b, return 0 when a == b and return positive int when a > b
func NewMultiSet(compare func(interface{}, interface{}) int) *Set {
	var set = &Set{}
	var header = &SetNode{}
	return NewRBTreer(
		set,
		header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		//SameSetNode,
		false,
	).(*Set)
}

// NewCustomMultiSet create a new not unique Set with newNode, deleteNode and compare func
// you can define you own func to create a new node and a new node
// such as with a sync.Pool to reduce the pressure of GC
func NewCustomMultiSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Set {
	var set = &Set{}
	var header = &SetNode{}
	return NewRBTreer(set, header,
		uintptr(unsafe.Pointer(&header.RBTreeNode))-uintptr(unsafe.Pointer(header)),
		newNode, deleteNode, compare,
		//SameSetNode,
		false).(*Set)
}
