package rbtree

type Maper interface {
	Size() int
	Empty() bool
	Begin() MapIterator
	End() MapIterator
	EndNode() MapIterator
	Find(interface{}) MapIterator
	LowerBound(interface{}) MapIterator
	UpperBound(interface{}) MapIterator
	Insert(interface{}) bool
	Erase(interface{}) bool
	EraseIterator(Iterator) bool
	Clear()
}

type MapIterator interface {
	Iterator
	GetValue() Valuer
	SetValue(Valuer)
}

type MapNode struct {
	RBTreeNode
	Keyer
	Valuer
}

func (node *MapNode) GetKey() Keyer {
	return node.Keyer
}

func (node *MapNode) SetKey(key Keyer) {
	node.Keyer = key
}

func (node *MapNode) GetValue() Valuer {
	return node.Valuer
}

func (node *MapNode) SetValue(value Valuer) {
	node.Valuer = value
}

func (node *MapNode) Copy(iter Iterator) {
	node.Keyer = iter.(*MapNode).Keyer
	node.Valuer = iter.(*MapNode).Valuer
}

type Map struct {
	RBTree
}

func NewMap() *Map {
	return &Map{*NewCustomRBTree(
		true,
		(*MapNode)(nil),
		func(key Keyer) Iterator {
			return &MapNode{RBTreeNode{}, key.GetKey(), key.(KeyValuer).GetValue()}
		},
		func(Iterator) {
		},
		func(a Iterator, b Iterator) bool {
			return a.(*MapNode) == b.(*MapNode)
		},
	)}
}

func (m *Map) Begin() MapIterator {
	return m.RBTree.Begin().(*MapNode)
}

func (m *Map) End() MapIterator {
	return m.RBTree.End().(*MapNode)
}

func (m *Map) EndNode() MapIterator {
	return m.RBTree.End().(*MapNode)
}

func (m *Map) Find(elem KeyValuer) MapIterator {
	return m.RBTree.Find(elem).(*MapNode)
}

func (m *Map) LowerBound(elem KeyValuer) MapIterator {
	return m.RBTree.LowerBound(elem).(*MapNode)
}

func (m *Map) UpperBound(elem KeyValuer) MapIterator {
	return m.RBTree.UpperBound(elem).(*MapNode)
}

type MultiMap struct {
	Map
}

func NewMultiMap(
	newElem func(Keyer) Iterator,
	deleteElem func(Iterator)) *MultiMap {
	return &MultiMap{Map{*NewCustomRBTree(
		false,
		(*MapNode)(nil),
		newElem,
		deleteElem,
		func(a Iterator, b Iterator) bool {
			return a.(*MapNode) == b.(*MapNode)
		},
	)}}
}

func (ms *MultiMap) EqualRange(elem KeyValuer) (MapIterator, MapIterator) {
	return ms.LowerBound(elem), ms.UpperBound(elem)
}

func (ms *MultiMap) Count(elem KeyValuer) (c int) {
	var beg, end Iterator = ms.LowerBound(elem), ms.UpperBound(elem)
	for ; beg != end; beg = beg.Next(beg) {
		c++
	}
	return c
}
