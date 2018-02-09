package rbtree

import (
	"errors"
	"unsafe"
)

var (
	ErrNotInTree  = errors.New("*node is not a node of this tree")
	ErrNoLast     = errors.New("begin of tree has no Last()")
	ErrNoNext     = errors.New("end of tree has no Next()")
	ErrEraseEmpty = errors.New("can't erase empty node")
)

//const nodeSize = unsafe.Offsetof(node{}.color) + unsafe.Sizeof(red)

const nodeSize = unsafe.Sizeof(node{})
const maxSpan = 512

type colorType bool

const (
	red   = false
	black = true
)

type node struct {
	child  [2]*node
	parent *node
	tree   *tree
	color  colorType
}

func (n *node) GetData() interface{} {
	return n.tree.getData(n)
}

func (n *node) GetKey() interface{} {
	return n.tree.getKey(n)
}

func (n *node) GetVal() interface{} {
	return n.tree.getVal(n)
}

func (n *node) SetVal(val interface{}) {
	n.tree.setVal(n, val)
}

func (n *node) Next() *node {
	return n.tree.next(n)
}

func (n *node) Last() *node {
	return n.tree.last(n)
}

type tree struct {
	header    node
	keyType   *_type
	valType   *_type
	size      int
	compare   func(a, b interface{}) int
	unique    bool
	directkey bool // true will store the key direct
	directval bool
	keySize   uintptr
	dataSize  uintptr
	nodeQ     unsafe.Pointer
	qsize     int
}

func (t *tree) init(unique bool, keyType, valType *_type, compare func(a, b interface{}) int) {
	t.header.child[0] = t.end()
	t.header.child[1] = t.end()
	t.header.parent = t.end()
	t.header.tree = t
	t.header.color = red
	t.unique = unique
	t.keyType = keyType
	t.valType = valType
	t.size = 0
	t.compare = compare
	t.nodeQ = nil
	t.qsize = 0

	if keyType == nil {
		panic("no key type")
	}
	t.directkey = !isDirectIface(keyType)
	if t.directkey {
		t.keySize = t.keyType.size
	} else {
		t.keySize = unsafe.Sizeof(unsafe.Pointer(&t))
	}
	t.dataSize = t.keySize
	if valType != nil {
		t.directval = !isDirectIface(valType)
		if t.directval {
			t.dataSize += t.valType.size
		} else {
			t.dataSize += unsafe.Sizeof(unsafe.Pointer(&t))
		}
	}
}

func (t *tree) getKeyPointer(node *node) unsafe.Pointer {
	return add(unsafe.Pointer(node), nodeSize)
}

func (t *tree) getValPointer(node *node) unsafe.Pointer {
	return add(unsafe.Pointer(node), nodeSize+t.keySize)
}

func (t *tree) getKey(node *node) interface{} {
	kp := t.getKeyPointer(node)
	if !t.directkey {
		kp = *(*unsafe.Pointer)(kp)
	}
	return eface2interface(eface{t.keyType, kp})
}

func (t *tree) getVal(node *node) interface{} {
	if t.valType == nil {
		panic("no value")
	}
	vp := t.getValPointer(node)
	if !t.directval {
		vp = *(*unsafe.Pointer)(vp)
	}
	return eface2interface(eface{t.valType, vp})
}

func (t *tree) setKey(node *node, key interface{}) {
	keyEface := interface2eface(key)
	if keyEface._type != t.keyType {
		panic("not same key type")
	}
	if t.directkey {
		memcopy(t.getKeyPointer(node), keyEface.p, t.keySize)
	} else {
		*(*unsafe.Pointer)(t.getKeyPointer(node)) = keyEface.p
	}
}

func (t *tree) setVal(node *node, val interface{}) {
	valEface := interface2eface(val)
	if valEface._type != t.valType {
		panic("not same value type")
	}
	if t.directval {
		memcopy(t.getValPointer(node), valEface.p, t.valType.size)
	} else {
		*(*unsafe.Pointer)(t.getValPointer(node)) = valEface.p
	}
}

func (t *tree) getData(node *node) interface{} {
	return t.getKey(node)
}

func (t *tree) copyNodeData(des, src *node) {
	memcopy(t.getKeyPointer(des), t.getKeyPointer(src), t.dataSize)
}

func (t *tree) newNode(data interface{}) *node {
	var fullNodeSize = t.dataSize + nodeSize
	if t.qsize == 0 {
		if t.size > maxSpan {
			t.qsize = maxSpan
		} else if t.size == 0 {
			t.qsize = 8 // begin at 8 node
		}
		t.nodeQ = newmem(uintptr(t.qsize) * fullNodeSize)
	}
	var res = (*node)(t.nodeQ)
	res.child[0] = t.end()
	res.child[1] = t.end()
	res.parent = t.end()
	res.tree = t
	res.color = red
	if t.valType != nil {
		t.setKey(res, data.(pair).key)
		t.setVal(res, data.(pair).val)
	} else {
		t.setKey(res, data)
	}
	t.nodeQ = getGCPointer(add(t.nodeQ, fullNodeSize), int(fullNodeSize))
	t.size++
	return res
}

func (t *tree) deleteNode(n *node) {
	n.child[0] = nil
	n.child[1] = nil
	n.parent = nil
	n.tree = nil
	t.size--
}

func (t *tree) Size() int {
	return t.size
}

func (t *tree) Unique() bool {
	return t.unique
}

func (t *tree) Empty() bool {
	return t.size == 0
}

func (t *tree) Begin() *node {
	return t.begin()
}

func (t *tree) begin() *node {
	return t.most(0)
}

func (t *tree) End() *node {
	return t.end()
}

func (t *tree) end() *node {
	return &t.header
}

func (t *tree) root() *node {
	return t.header.parent
}

func (t *tree) rootPointer() **node {
	return &t.header.parent
}

func (t *tree) most(ch int) *node {
	return t.header.child[ch]
}

func (t *tree) mostPointer(ch int) **node {
	return &t.header.child[ch]
}

func (t *tree) mustGetColor(node *node) colorType {
	if node == t.end() || node.color == black {
		return black
	}
	return red
}

func (t *tree) next(node *node) *node {
	if node == t.end() {
		panic(ErrNoNext)
	}
	if node == t.most(1) {
		return t.end()
	}
	return t.gothrough(1, node)
}

func (t *tree) last(node *node) *node {
	if node == t.begin() {
		panic(ErrNoLast)
	}
	if node == t.end() {
		return t.most(1)
	}
	return t.gothrough(0, node)
}

func (t *tree) gothrough(ch int, node *node) *node {
	if node.child[ch] != t.end() {
		node = node.child[ch]
		for node.child[ch^1] != t.end() {
			node = node.child[ch^1]
		}
		return node
	}
	for node.parent != t.end() && node.parent.child[ch] == node {
		node = node.parent
	}
	return node.parent
}

func (t *tree) Find(key interface{}) *node {
	return t.find(interface2noescape(key))
}

func (t *tree) find(key interface{}) *node {
	var root = t.root()
	for {
		if root == t.end() {
			return root
		}
		switch cmp := t.compare(key, root.GetKey()); {
		case cmp == 0:
			return root
		case cmp < 0:
			root = root.child[0]
		case cmp > 0:
			root = root.child[1]
		}
	}
}

func (t *tree) Insert(data interface{}) (*node, bool) {
	noescapedata := interface2noescape(data)
	noescapekey := interface2noescape(data)
	if t.valType != nil {
		noescapekey = interface2noescape(data.(pair).key)
	}
	return t.insert(noescapedata, noescapekey)
}

func (t *tree) insert(data interface{}, key interface{}) (*node, bool) {
	var root = t.root()
	var rootPointer = t.rootPointer()
	if root == t.end() {
		*rootPointer = t.newNode(data)
		t.insertAdjust(*rootPointer)
		*t.mostPointer(0) = *rootPointer
		*t.mostPointer(1) = *rootPointer
		return *rootPointer, true
	}
	var parent = root.parent
	for root != t.end() {
		parent = root
		switch cmp := t.compare(key, root.GetKey()); {
		case cmp == 0:
			if t.unique {
				return t.end(), false
			}
			fallthrough
		case cmp < 0:
			rootPointer = &root.child[0]
			root = *rootPointer
		case cmp > 0:
			rootPointer = &root.child[1]
			root = *rootPointer
		}
	}
	*rootPointer = t.newNode(data)
	(*rootPointer).parent = parent
	for ch := 0; ch < 2; ch++ {
		if parent == t.most(ch) && parent.child[ch] == *rootPointer {
			*t.mostPointer(ch) = *rootPointer
		}
	}
	t.insertAdjust(*rootPointer)
	return *rootPointer, true
}

//insert node is default red
func (t *tree) insertAdjust(node *node) {
	var parent = node.parent
	if parent == t.end() {
		//fmt.Println("case 1: insert")
		//node is root,set black
		node.color = black
		return
	}
	if parent.color == black {
		//fmt.Println("case 2: insert")
		//if parent is black,do nothing
		return
	}

	//parent is red,grandpa can't be empty and color is black
	var grandpa = parent.parent
	var parentCh = 0
	if grandpa.child[1] == parent {
		parentCh = 1
	}

	var uncle = grandpa.child[parentCh^1]
	if t.mustGetColor(uncle) == red {
		//fmt.Println("case 3: insert")
		//uncle is red
		parent.color = black
		grandpa.color = red
		uncle.color = black
		t.insertAdjust(grandpa)
		return
	}

	var childCh = 0
	if parent.child[1] == node {
		childCh = 1
	}
	if childCh != parentCh {
		//fmt.Println("case 4: insert")
		t.rotate(parentCh, node)
		var tmp = parent
		parent = node
		node = tmp
	}

	//fmt.Println("case 5: insert")
	parent.color = black
	grandpa.color = red
	t.rotate(parentCh^1, parent)
}

//ch = 0:take node for center,left rotate parent down,node is parent's right child
//ch = 1:take node for center,right rotate parent down,node is parent's left child
func (t *tree) rotate(ch int, node *node) {
	var (
		tmp     = node.child[ch]
		parent  = node.parent
		grandpa = parent.parent
	)
	node.child[ch] = parent
	parent.child[ch^1] = tmp

	if tmp != t.end() {
		tmp.parent = parent
	}

	parent.parent = node
	node.parent = grandpa

	if grandpa == t.end() {
		*t.rootPointer() = node
		return
	}
	if grandpa.child[0] == parent {
		grandpa.child[0] = node
	} else {
		grandpa.child[1] = node
	}
}

// Erase erase all the node keys equal to key in this tree and return the number of erase node
func (t *tree) Erase(key interface{}) (count int) {
	var noescapekey = interface2noescape(key)
	if t.unique {
		var iter = t.find(noescapekey)
		if iter == t.end() {
			return 0
		}
		t.eraseNode(iter)
		return 1
	}
	var beg = t.lowerBound(noescapekey)
	for beg != t.end() && t.compare(noescapekey, beg.GetKey()) == 0 {
		var tmp = t.next(beg)
		t.eraseNode(beg)
		beg = tmp
		count++
	}
	return count
}

func (t *tree) sameTree(node *node) bool {
	return t == node.tree
}

// EraseNode erase node from the tree
// if node is not in tree, it will panic
func (t *tree) EraseNode(node *node) {
	if !t.sameTree(node) {
		panic(ErrNotInTree)
	}
	t.eraseNode(node)
}
func (t *tree) eraseNode(node *node) {
	if node == t.end() {
		panic(ErrEraseEmpty)
	}
	if node.child[0] != t.end() && node.child[1] != t.end() {
		//if node has two child,it's last node must has no more than one child,copy to node and erase last node
		var tmp = t.last(node)
		t.copyNodeData(node, tmp)
		node = tmp
	}
	//adjust leftmost and rightmost
	for ch := 0; ch < 2; ch++ {
		if t.most(ch) == node {
			if ch == 0 {
				*t.mostPointer(ch) = t.next(node)
			} else {
				*t.mostPointer(ch) = t.last(node)
			}
		}
	}
	var child = t.end()
	if node.child[0] != t.end() {
		child = node.child[0]
	} else if node.child[1] != t.end() {
		child = node.child[1]
	}

	var parent = node.parent
	if child != t.end() {
		child.parent = parent
	}
	if parent == t.end() {
		*t.rootPointer() = child
	} else if parent.child[0] == node {
		parent.child[0] = child
	} else {
		parent.child[1] = child
	}
	if node.color == black { //if node is red,just erase,otherwise adjust
		t.eraseAdjust(child, parent)
		//fmt.Println("eraseAdjust:")
	}
	t.deleteNode(node)
	return
}

func (t *tree) eraseAdjust(node, parent *node) {
	if parent == t.end() {
		//node is root
		//fmt.Println("case 1: erase")
		if node != t.end() {
			node.color = black
		}
		return
	}
	if t.mustGetColor(node) == red {
		//node is red,just set black
		//fmt.Println("case 2: erase")
		node.color = black
		return
	}
	var nodeCh = 0
	if parent.child[1] == node {
		nodeCh = 1
	}
	var brother = parent.child[nodeCh^1]
	//after case 1 parent must not be empty node and after case 2 node must be black
	if parent.color == red {
		//parent is red,brother must be black but can't be empty node,because the path has a black node more
		if t.mustGetColor(brother.child[0]) == black && t.mustGetColor(brother.child[1]) == black {
			//fmt.Println("case 3: erase")
			brother.color = red
			parent.color = black
			return
		}
		if brother != t.end() && t.mustGetColor(brother.child[nodeCh]) == red {
			//fmt.Println("case 4: erase", nodeCh)
			parent.color = black
			t.rotate(nodeCh^1, brother.child[nodeCh])
			t.rotate(nodeCh, parent.child[nodeCh^1])
			return
		}
		//fmt.Println("case 5: erase")
		t.rotate(nodeCh, brother)
		return
	}
	//parent is black
	if t.mustGetColor(brother) == red {
		//brother is red, it's children must be black
		//fmt.Println("case 6: erase")
		brother.color = black
		parent.color = red
		t.rotate(nodeCh, brother)
		t.eraseAdjust(node, parent) //goto redParent then end
		return
	}
	//brother is black
	if t.mustGetColor(brother.child[0]) == black && t.mustGetColor(brother.child[1]) == black {
		//fmt.Println("case 7: erase")
		brother.color = red
		t.eraseAdjust(parent, parent.parent)
		return
	}
	if t.mustGetColor(brother.child[nodeCh]) == red {
		//fmt.Println("case 8: erase", nodeCh)
		brother.child[nodeCh].color = black
		t.rotate(nodeCh^1, brother.child[nodeCh])
		t.rotate(nodeCh, parent.child[nodeCh^1])
		return
	}
	//fmt.Println("case 9: erase", nodeCh)
	brother.child[nodeCh^1].color = black
	t.rotate(nodeCh, brother)
}

func (t *tree) LowerBound(key interface{}) *node {
	return t.lowerBound(key)
}

func (t *tree) lowerBound(key interface{}) *node {
	var root = t.root()
	var parent = t.end()
	for {
		if root == t.end() {
			if parent == t.end() {
				return parent
			} else if t.compare(key, parent.GetKey()) <= 0 {
				return parent
			}
			return t.next(parent)
		}
		parent = root
		if t.compare(key, root.GetKey()) > 0 {
			root = root.child[1]
		} else {
			root = root.child[0]
		}
	}
}

func (t *tree) UpperBound(key interface{}) *node {
	return t.upperBound(key)
}

func (t *tree) upperBound(key interface{}) *node {
	var root = t.root()
	var parent = t.end()
	for {
		if root == t.end() {
			if parent == t.end() {
				return parent
			} else if t.compare(key, parent.GetKey()) < 0 {
				return parent
			}
			return t.next(parent)
		}
		parent = root
		if t.compare(key, root.GetKey()) >= 0 {
			root = root.child[1]
		} else {
			root = root.child[0]
		}
	}
}

// Count return the num of node key equal to key in this tree
func (t *tree) Count(key interface{}) (count int) {
	if t.unique {
		if t.find(key) == t.end() {
			return 0
		}
		return 1
	}
	var beg = t.lowerBound(key)
	for beg != t.end() && t.compare(beg.GetKey(), key) == 0 {
		beg = t.next(beg)
		count++
	}
	return count
}

// EqualRange return the Node range of equal key node in this tree
func (t *tree) EqualRange(key interface{}) (beg, end *node) {
	return t.LowerBound(key), t.UpperBound(key)
}

// EraseNodeRange erase the given iterator range
// if the given range is not in this tree, it will panic with ErrNoInTree
// if end can get beg after multi Next method, it will panic with ErrNoLast
func (t *tree) EraseNodeRange(beg, end *node) (count int) {
	return t.eraseNodeRange(beg, end)
}
func (t *tree) eraseNodeRange(beg, end *node) (count int) {
	for beg != end {
		var tmp = t.next(beg)
		t.eraseNode(beg)
		beg = tmp
		count++
	}
	return count
}

func (t *tree) Clear() {
	t.clear(t.root())
	*t.mostPointer(0) = t.end()
	*t.mostPointer(1) = t.end()
	*t.rootPointer() = t.end()
	t.size = 0
}

func (t *tree) clear(root *node) {
	if root == t.end() {
		return
	}
	t.clear(root.child[0])
	t.clear(root.child[1])
	t.deleteNode(root)
}
