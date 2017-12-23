package rbtree

type Seter interface {
	Size() int
	Empty() bool
	Begin() SetIterator
	End() SetIterator
	EndNode() SetIterator
	Find(Keyer) SetIterator
	LowerBound(Keyer) SetIterator
	UpperBound(Keyer) SetIterator
	Insert(Keyer) bool
	Erase(Keyer) bool
	EraseIterator(Iterator) bool
	Clear()
}

type SetIterator interface {
	Iterator
}

type SetNode struct {
	RBTreeNode
	Keyer
}

func (node *SetNode) GetKey() Keyer {
	return node.Keyer
}

func (node *SetNode) SetKey(key Keyer) {
	node.Keyer = key
}

func (node *SetNode) Copy(src Iterator) {
	node.Keyer = src.(*SetNode).Keyer
}

type Set struct {
	RBTree
}

func NewSetNode(elem Keyer) *SetNode {
	return &SetNode{Keyer: elem}
}

func NewSet() *Set {
	return &Set{*NewCustomRBTree(
		true,
		(*SetNode)(nil),
		func(elem Keyer) Iterator {
			return &SetNode{Keyer: elem}
		},
		func(Iterator) {
		},
		func(a Iterator, b Iterator) bool {
			return a.(*SetNode) == b.(*SetNode)
		},
	)}
}

func NewCustomSet(
	newElem func(Keyer) Iterator,
	deleteElem func(Iterator)) *Set {
	return &Set{*NewCustomRBTree(
		true,
		(*SetNode)(nil),
		newElem,
		deleteElem,
		func(a Iterator, b Iterator) bool {
			return a.(*SetNode) == b.(*SetNode)
		},
	)}
}

func (s *Set) SameIter(a Iterator, b Iterator) bool {
	return a.(*SetNode) == b.(*SetNode)
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

func (s *Set) Find(elem Keyer) SetIterator {
	return s.RBTree.Find(elem).(*SetNode)
}

func (s *Set) LowerBound(elem Keyer) SetIterator {
	return s.RBTree.LowerBound(elem).(*SetNode)
}

func (s *Set) UpperBound(elem Keyer) SetIterator {
	return s.RBTree.UpperBound(elem).(*SetNode)
}

type MultiSet struct {
	Set
}

func NewMultiSet() *MultiSet {
	return &MultiSet{Set{*NewCustomRBTree(
		false,
		(*SetNode)(nil),
		func(elem Keyer) Iterator {
			return &SetNode{Keyer: elem}
		},
		func(Iterator) {
		},
		func(a Iterator, b Iterator) bool {
			return a.(*SetNode) == b.(*SetNode)
		},
	)}}
}

//need to rewirte
func (ms *MultiSet) Erase(elem Keyer) bool {
	return ms.RBTree.Erase(elem)
}

func (ms *MultiSet) EqualRange(elem Keyer) (SetIterator, SetIterator) {
	return ms.LowerBound(elem), ms.UpperBound(elem)
}

func (ms *MultiSet) Count(elem Keyer) (c int) {
	var beg, end Iterator = ms.LowerBound(elem), ms.UpperBound(elem)
	for ; beg != end; beg = beg.Next(beg) {
		c++
	}
	return c
}
