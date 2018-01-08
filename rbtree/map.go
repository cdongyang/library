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

func (m *Map) NewNode(data interface{}) Iterator {
	_ = data.(*Pair)
	return m.RBTree.NewNode(data)
}

func (m *Map) Insert(data interface{}) (Iterator, bool) {
	return m.RBTree.insert(data, func(key interface{}) int {
		return m.compare(data.(*Pair).GetKey(), key)
	})
}

func sameMapNode(a Iterator, b Iterator) bool {
	aa, aok := a.(*MapNode)
	bb, bok := a.(*MapNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

func NewMap(compare func(interface{}, interface{}) int) *Map {
	return &Map{*NewRBTree(
		func(data interface{}) Iterator {
			return &MapNode{SetNode{data: data}}
		},
		func(Iterator) {
		},
		compare,
		sameMapNode,
		true,
	)}
}

func NewCustomMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Map {
	return &Map{*NewRBTree(newNode, deleteNode, compare, sameMapNode, true)}
}

func NewMultiMap(compare func(interface{}, interface{}) int) *Map {
	return &Map{*NewRBTree(
		func(data interface{}) Iterator {
			return &MapNode{SetNode{data: data}}
		},
		func(Iterator) {
		},
		compare,
		sameMapNode,
		false,
	)}
}

func NewCustomMultiMap(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int) *Map {
	return &Map{*NewRBTree(newNode, deleteNode, compare, sameMapNode, false)}
}
