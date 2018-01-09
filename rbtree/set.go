package rbtree

type SetNode struct {
	RBTreeNode
	data interface{}
}

func (node *SetNode) Next() Iterator {
	return node.GetTree().Next(node)
}

func (node *SetNode) Last() Iterator {
	return node.GetTree().Last(node)
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

func sameSetNode(a, b Iterator) bool {
	aa, aok := a.(*SetNode)
	bb, bok := a.(*SetNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

type Set struct {
	RBTree
}

func NewSet(compare func(interface{}, interface{}) int) *Set {
	return &Set{*NewRBTree(
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		sameSetNode,
		true,
	)}
}

func NewCustomSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Set {
	return &Set{*NewRBTree(newNode, deleteNode, compare, sameSetNode, true)}
}

func NewMultiSet(compare func(interface{}, interface{}) int) *Set {
	return &Set{*NewRBTree(
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		sameSetNode,
		false,
	)}
}

func NewCustomMultiSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Set {
	return &Set{*NewRBTree(newNode, deleteNode, compare, sameSetNode, false)}
}
