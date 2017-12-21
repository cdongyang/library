package rbtree

type Interface interface {
	Size() int
	Empty() bool
	Begin() Iterator
	End() Iterator
	EndNode() Iterator
	Find(interface{}) Iterator
	LowerBound(interface{}) Iterator
	UpperBound(interface{}) Iterator
	Insert(interface{}) bool
	Erase(interface{}) bool
	EraseIterator(Iterator) bool
	Clear()
}

type SetIterator interface {
	Iterator
	GetKey() interface{}
	SetKey(interface{})
}

type SetNode struct {
	RBTreeNode
	key interface{}
}

func (s *SetNode) GetKey() interface{} {
	return s.key
}
func (s *SetNode) SetKey(key interface{}) {
	s.key = key
}

func (s *SetNode) Copy(des, src Iterator) {
	des.(*SetNode).key = src.(*SetNode).key
}

type Set struct {
	RBTree
}

func NewSetNode(elem interface{}) *SetNode {
	return &SetNode{key: elem}
}

func NewSet(
	compare func(SetIterator, SetIterator) int,
	newElem func(interface{}) SetIterator,
	deleteElem func(SetIterator)) *Set {
	return &Set{*NewRBTree((*SetNode)(nil),
		func(a Iterator, b Iterator) int {
			return compare(a.(*SetNode), b.(*SetNode))
		},
		func(elem interface{}) Iterator {
			return newElem(elem)
		},
		func(iter Iterator) {
			deleteElem(iter.(*SetNode))
		},
	)}
}

func (s *Set) Begin() SetIterator {
	return s.RBTree.Begin().(*SetNode)
}

func (s *Set) End() SetIterator {
	return s.RBTree.End().(*SetNode)
}

func (s *Set) EndNode() SetIterator {
	return s.RBTree.EndNode().(*SetNode)
}

func (s *Set) Find(elem interface{}) SetIterator {
	return s.RBTree.Find(elem).(*SetNode)
}

func (s *Set) LowerBound(elem interface{}) SetIterator {
	return s.RBTree.LowerBound(elem).(*SetNode)
}

func (s *Set) UpperBound(elem interface{}) SetIterator {
	return s.RBTree.UpperBound(elem).(*SetNode)
}

func (s *Set) EraseIterator(iter SetIterator) bool {
	return s.RBTree.EraseIterator(iter)
}
