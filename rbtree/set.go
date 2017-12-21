package rbtree

type Seter interface {
	Size() int
	Empty() bool
	Begin() SetIterator
	End() SetIterator
	EndNode() SetIterator
	Find(interface{}) SetIterator
	LowerBound(interface{}) SetIterator
	UpperBound(interface{}) SetIterator
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
	des.(*SetNode).SetKey(src.(*SetNode).GetKey())
}

type Set struct {
	RBTree
}

func NewSetNode(elem interface{}) *SetNode {
	return &SetNode{key: elem}
}

func NewSet(compare func(SetIterator, SetIterator) int) *Set {
	return &Set{*NewCustomRBTree(
		true,
		(*SetNode)(nil),
		func(a, b Iterator) int {
			return compare(a.(*SetNode), b.(*SetNode))
		},
		func(elem interface{}) Iterator {
			return &SetNode{key: elem}
		},
		func(Iterator) {
		},
	)}
}

func NewCustomSet(
	compare func(SetIterator, SetIterator) int,
	newElem func(interface{}) Iterator,
	deleteElem func(Iterator)) *Set {
	return &Set{*NewCustomRBTree(
		true,
		(*SetNode)(nil),
		func(a, b Iterator) int {
			return compare(a.(*SetNode), b.(*SetNode))
		},
		newElem,
		deleteElem,
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

type MultiSet struct {
	Set
}

func NewMultiSet(compare func(SetIterator, SetIterator) int) *MultiSet {
	return &MultiSet{Set{*NewCustomRBTree(
		false,
		(*SetNode)(nil),
		func(a, b Iterator) int {
			return compare(a.(*SetNode), b.(*SetNode))
		},
		func(elem interface{}) Iterator {
			return &SetNode{key: elem}
		},
		func(Iterator) {

		},
	)}}
}

func (ms *MultiSet) EqualRange(elem interface{}) (SetIterator, SetIterator) {
	return ms.LowerBound(elem), ms.UpperBound(elem)
}

func (ms *MultiSet) Count(elem interface{}) (c int) {
	var beg, end Iterator = ms.LowerBound(elem), ms.UpperBound(elem)
	for ; beg != end; beg = beg.Next(beg) {
		c++
	}
	return c
}
