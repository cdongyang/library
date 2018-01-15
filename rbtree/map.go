package rbtree

type Pair struct {
	key, value interface{}
}

func (p *Pair) GetKey() interface{} {
	return p.key
}

func (p *Pair) GetValue() interface{} {
	return p.value
}

func (p *Pair) SetValue(value interface{}) {
	p.value = value
}

func NewPair(key, value interface{}) *Pair {
	return &Pair{key: key, value: value}
}

//type of MapNode.data must be *Pair
type MapNode struct {
	SetNode
}

func (node *MapNode) Next() Iterator {
	return node.tree.Next(node)
}

func (node *MapNode) Last() Iterator {
	return node.tree.Next(node)
}

func (node *MapNode) CopyData(src Iterator) {
	node.data = src.(*MapNode).data
}

func (node *MapNode) GetKey() interface{} {
	return node.data.(*Pair).GetKey()
}

func (node *MapNode) GetValue() interface{} {
	return node.data.(*Pair).GetValue()
}

func (node *MapNode) SetValue(value interface{}) {
	node.data.(*Pair).SetValue(value)
}

type Map struct {
	RBTree
}

func (m *Map) Insert(data interface{}) (Iterator, bool) {
	return m.RBTree.insert(data, func(key interface{}) int {
		return m.compare(data.(*Pair).GetKey(), key)
	})
}

func SameMapNode(a Iterator, b Iterator) bool {
	aa, aok := a.(*MapNode)
	bb, bok := a.(*MapNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

func NewMap(compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	return NewRBTreer(
		mp,
		&MapNode{},
		func(data interface{}) Iterator {
			return &MapNode{SetNode{data: data}}
		},
		func(Iterator) {
		},
		compare,
		SameMapNode,
		true,
	).(*Map)
}

func NewCustomMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	return NewRBTreer(mp, &MapNode{}, newNode, deleteNode, compare, SameMapNode, true).(*Map)
}

func NewMultiMap(compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	return NewRBTreer(
		mp,
		&MapNode{},
		func(data interface{}) Iterator {
			return &MapNode{SetNode{data: data}}
		},
		func(Iterator) {
		},
		compare,
		SameMapNode,
		false,
	).(*Map)
}

func NewCustomMultiMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Map {
	var mp = &Map{}
	return NewRBTreer(mp, &MapNode{}, newNode, deleteNode, compare, SameMapNode, false).(*Map)
}
