package rbtree

import (
	"sync"
)

type SafeSetIterator interface {
	Iterator
}

type SafeSetNode struct {
	RBTreeNode
	Keyer
}

func (node *SafeSetNode) GetKey() Keyer {
	return node.Keyer
}

func (node *SafeSetNode) SetKey(key Keyer) {
	node.Keyer = key
}

func (node *SafeSetNode) Copy(src Iterator) {
	node.Keyer = src.(*SafeSetNode).Keyer
}

type SafeSet struct {
	RBTree
	rw sync.RWMutex
}

func NewSafeSetNode(elem Keyer) *SafeSetNode {
	return &SafeSetNode{Keyer: elem}
}

func NewSafeSet() *SafeSet {
	return &SafeSet{*NewCustomRBTree(
		true,
		(*SafeSetNode)(nil),
		func(elem Keyer) Iterator {
			return &SafeSetNode{Keyer: elem}
		},
		func(Iterator) {
		},
		func(a Iterator, b Iterator) bool {
			return a.(*SafeSetNode) == b.(*SafeSetNode)
		},
	), sync.RWMutex{}}
}

func NewCustomSafeSet(
	newElem func(Keyer) Iterator,
	deleteElem func(Iterator)) *SafeSet {
	return &SafeSet{*NewCustomRBTree(
		true,
		(*SafeSetNode)(nil),
		newElem,
		deleteElem,
		func(a Iterator, b Iterator) bool {
			return a.(*SafeSetNode) == b.(*SafeSetNode)
		},
	), sync.RWMutex{}}
}

func (s *SafeSet) SameIter(a Iterator, b Iterator) bool {
	return a.(*SafeSetNode) == b.(*SafeSetNode)
}

func (s *SafeSet) Begin() SafeSetIterator {
	return s.RBTree.Begin().(*SafeSetNode)
}

func (s *SafeSet) End() SafeSetIterator {
	return s.RBTree.End().(*SafeSetNode)
}

func (s *SafeSet) EndNode() SafeSetIterator {
	return s.RBTree.EndNode().(*SafeSetNode)
}

func (s *SafeSet) Find(elem Keyer) SafeSetIterator {
	return s.RBTree.Find(elem).(*SafeSetNode)
}

func (s *SafeSet) LowerBound(elem Keyer) SafeSetIterator {
	return s.RBTree.LowerBound(elem).(*SafeSetNode)
}

func (s *SafeSet) UpperBound(elem Keyer) SafeSetIterator {
	return s.RBTree.UpperBound(elem).(*SafeSetNode)
}

type SafeMultiSet struct {
	SafeSet
}

func NewSafeMultiSet() *SafeMultiSet {
	return &SafeMultiSet{SafeSet{*NewCustomRBTree(
		false,
		(*SafeSetNode)(nil),
		func(elem Keyer) Iterator {
			return &SafeSetNode{Keyer: elem}
		},
		func(Iterator) {
		},
		func(a Iterator, b Iterator) bool {
			return a.(*SafeSetNode) == b.(*SafeSetNode)
		},
	), sync.RWMutex{}}}
}

//need to rewirte
func (ms *SafeMultiSet) Erase(elem Keyer) bool {
	return ms.RBTree.Erase(elem)
}

func (ms *SafeMultiSet) EqualRange(elem Keyer) (SafeSetIterator, SafeSetIterator) {
	return ms.LowerBound(elem), ms.UpperBound(elem)
}

func (ms *SafeMultiSet) Count(elem Keyer) (c int) {
	var beg, end Iterator = ms.LowerBound(elem), ms.UpperBound(elem)
	for ; beg != end; beg = beg.Next(beg) {
		c++
	}
	return c
}
