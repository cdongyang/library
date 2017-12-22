package rbtree

type IntSetNode struct {
	RBTreeNode
	key int
}

func (s *IntSetNode) GetKey() interface{} {
	return s.key
}
func (s *IntSetNode) SetKey(key interface{}) {
	s.key = key.(int)
}

func (s *IntSetNode) Copy(des, src Iterator) {
	des.(*IntSetNode).SetKey(src.(*IntSetNode).GetKey())
}

type IntSet struct {
	RBTree
}

func NewIntSetNode(elem int) *IntSetNode {
	return &IntSetNode{key: elem}
}

func NewIntSet() *IntSet {
	return &IntSet{*NewCustomRBTree(
		true,
		(*IntSetNode)(nil),
		func(a, b Iterator) int {
			return a.(*IntSetNode).key - b.(*IntSetNode).key
		},
		func(elem interface{}) Iterator {
			return &IntSetNode{key: elem.(int)}
		},
		func(Iterator) {
		},
	)}
}

func NewCustomIntSet(
	newElem func(interface{}) Iterator,
	deleteElem func(Iterator)) *IntSet {
	return &IntSet{*NewCustomRBTree(
		true,
		(*IntSetNode)(nil),
		func(a, b Iterator) int {
			return a.(*IntSetNode).key - b.(*IntSetNode).key
		},
		newElem,
		deleteElem,
	)}
}

func (s *IntSet) Begin() SetIterator {
	return s.RBTree.Begin().(*IntSetNode)
}

func (s *IntSet) End() SetIterator {
	return s.RBTree.End().(*IntSetNode)
}

func (s *IntSet) EndNode() SetIterator {
	return s.RBTree.EndNode().(*IntSetNode)
}

func (s *IntSet) Find(elem interface{}) SetIterator {
	return s.RBTree.Find(elem).(*IntSetNode)
}

func (s *IntSet) LowerBound(elem interface{}) SetIterator {
	return s.RBTree.LowerBound(elem).(*IntSetNode)
}

func (s *IntSet) UpperBound(elem interface{}) SetIterator {
	return s.RBTree.UpperBound(elem).(*IntSetNode)
}
