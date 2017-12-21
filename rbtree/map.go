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
	SetIterator
	GetValue() interface{}
	SetValue(interface{})
}

type MapNode struct {
	SetNode
	value interface{}
}

func (m *MapNode) GetValue() interface{} {
	return m.value
}

func (m *MapNode) SetValue(value interface{}) {
	m.value = value
}

func (m *MapNode) Copy(des, src Iterator) {
	des.(*MapNode).key = src.(*MapNode).key
	des.(*MapNode).value = src.(*MapNode).value
}

type Map struct {
	RBTree
}

func NewMap(
	compare func(MapIterator, MapIterator) int,
	newElem func(interface{}) Iterator,
	deleteElem func(Iterator)) *Map {
	return &Map{*NewCustomRBTree(
		true,
		(*MapNode)(nil),
		func(a, b Iterator) int {
			return compare(a.(*MapNode), b.(*MapNode))
		},
		newElem,
		deleteElem,
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

func (m *Map) Find(elem interface{}) MapIterator {
	return m.RBTree.Find(elem).(*MapNode)
}

func (m *Map) LowerBound(elem interface{}) MapIterator {
	return m.RBTree.LowerBound(elem).(*MapNode)
}

func (m *Map) UpperBound(elem interface{}) MapIterator {
	return m.RBTree.UpperBound(elem).(*MapNode)
}

type MultiMap struct {
	Map
}

func NewMultiMap(
	compare func(MapIterator, MapIterator) int,
	newElem func(interface{}) Iterator,
	deleteElem func(Iterator)) *MultiMap {
	return &MultiMap{Map{*NewCustomRBTree(
		false,
		(*MapNode)(nil),
		func(a, b Iterator) int {
			return compare(a.(*MapNode), b.(*MapNode))
		},
		newElem,
		deleteElem,
	)}}
}

func (ms *MultiMap) EqualRange(elem interface{}) (MapIterator, MapIterator) {
	return ms.LowerBound(elem), ms.UpperBound(elem)
}

func (ms *MultiMap) Count(elem interface{}) (c int) {
	var beg, end Iterator = ms.LowerBound(elem), ms.UpperBound(elem)
	for ; beg != end; beg = beg.Next(beg) {
		c++
	}
	return c
}
