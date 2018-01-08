package rbtree

type SetNode struct {
	RBTreeNode
	data interface{}
}

func (node *SetNode) Next() Iterator {
	return node.tree.Next(node)
}

func (node *SetNode) Last() Iterator {
	return node.tree.Last(node)
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

func sameSetNode(a Iterator, b Iterator) bool {
	aa, aok := a.(*SetNode)
	bb, bok := a.(*SetNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

func NewSet(compare func(interface{}, interface{}) int) *RBTree {
	return NewRBTree(
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		sameSetNode,
		true,
	)
}

func NewCustomSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *RBTree {
	return NewRBTree(newNode, deleteNode, compare, sameSetNode, true)
}

func NewMultiSet(compare func(interface{}, interface{}) int) *RBTree {
	return NewRBTree(
		func(data interface{}) Iterator {
			return &SetNode{data: data}
		},
		func(Iterator) {
		},
		compare,
		sameSetNode,
		false,
	)
}

func NewCustomMultiSet(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *RBTree {
	return NewRBTree(newNode, deleteNode, compare, sameSetNode, false)
}
