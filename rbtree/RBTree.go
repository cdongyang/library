/*
learn about red-black tree
	https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91
*/
package rbtree

import (
	"unsafe"
)

type Iterator interface {
	Next() Iterator
	Last() Iterator
	GetData() (data interface{})
	GetKey() (key interface{})
	GetValue() (value interface{})
	SetValue(value interface{})
	CopyData(src Iterator)
	GetTree() RBTreer

	setTree(RBTreer)
	getColor() colorType
	setColor(colorType)
}

type colorType bool

const (
	red   = false
	black = true
)

type RBTreeNode struct {
	child  [2]Iterator
	parent Iterator
	tree   RBTreer
	color  colorType
}

// Next return next Iterator of this
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) Next() Iterator {
	return node.GetTree().Next(node)
}

// Last return last Iterator of this
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) Last() Iterator {
	return node.GetTree().Last(node)
}

// GetData get the data of this
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) GetData() interface{} {
	return nil
}

// GetKey get the compare key of this
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) GetKey() interface{} {
	return nil
}

// GetValue get the value of this
// map node should rewrite this method
func (node *RBTreeNode) GetValue() interface{} {
	return nil
}

// SetValue set the value of this
// map node should rewrite this method
func (node *RBTreeNode) SetValue(interface{}) {
}

//CopyData copy the node data to this from src
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) CopyData(src Iterator) {
}

func (node *RBTreeNode) getColor() colorType {
	return node.color
}

func (node *RBTreeNode) setColor(color colorType) {
	node.color = color
}

// GetTree get the RBTreer of this
// but you should set the tree when create new node
func (node *RBTreeNode) GetTree() RBTreer {
	return node.tree
}

func (node *RBTreeNode) setTree(tree RBTreer) {
	node.tree = tree
}

var (
	offsetChild [2]uintptr
	offsetParent,
	offsetTree,
	offsetColor uintptr
)

func init() {
	var node = &RBTreeNode{}
	offsetChild[0] = uintptr(unsafe.Pointer(&node.child[0])) - uintptr(unsafe.Pointer(node))
	offsetChild[1] = uintptr(unsafe.Pointer(&node.child[1])) - uintptr(unsafe.Pointer(node))
	offsetParent = uintptr(unsafe.Pointer(&node.parent)) - uintptr(unsafe.Pointer(node))
	offsetTree = uintptr(unsafe.Pointer(&node.tree)) - uintptr(unsafe.Pointer(node))
	offsetColor = uintptr(unsafe.Pointer(&node.color)) - uintptr(unsafe.Pointer(node))
	//fmt.Println(offsetChild[0], offsetChild[1], offsetParent, offsetTree, offsetColor)
}

func getIteratorPointer(node Iterator, offset uintptr) *Iterator {
	return (*Iterator)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&node)))[1] + offset))
}

type RBTreer interface {
	DeleteNode(node Iterator)
	SameIterator(a, b Iterator) bool
	Compare(key1, key2 interface{}) int
	Clear()
	Unique() bool
	Size() int
	Empty() bool
	Begin() Iterator
	End() Iterator
	Last(node Iterator) Iterator
	Next(node Iterator) Iterator
	Count(key interface{}) int
	EqualRange(key interface{}) (beg Iterator, end Iterator)
	Find(key interface{}) Iterator
	Insert(data interface{}) (Iterator, bool)
	Erase(key interface{}) int
	EraseIterator(node Iterator)
	EraseIteratorRange(beg Iterator, end Iterator) int
	LowerBound(key interface{}) Iterator
	UpperBound(key interface{}) Iterator

	init(
		tree RBTreer,
		header Iterator,
		nodeOffset uintptr,
		newNode func(interface{}) Iterator,
		deleteNode func(Iterator),
		compare func(interface{}, interface{}) int,
		sameIterator func(Iterator, Iterator) bool,
		unique bool,
	)
	getChild(Iterator, int) Iterator
	setChild(Iterator, int, Iterator)
	getParent(Iterator) Iterator
	setParent(Iterator, Iterator)
}

type RBTree struct {
	header       Iterator
	size         int
	newNode      func(interface{}) Iterator
	deleteNode   func(Iterator)
	compare      func(interface{}, interface{}) int
	sameIterator func(Iterator, Iterator) bool
	nodeOffset   uintptr
	unique       bool
}

func NewRBTreer(
	t RBTreer,
	header Iterator,
	nodeOffset uintptr,
	newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int,
	sameIterator func(Iterator, Iterator) bool,
	unique bool) RBTreer {
	t.init(t, header, nodeOffset, newNode, deleteNode, compare, sameIterator, unique)
	return t
}

func (t *RBTree) init(
	tree RBTreer,
	header Iterator,
	nodeOffset uintptr,
	newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int,
	sameIterator func(Iterator, Iterator) bool,
	unique bool) {
	t.header = header
	t.header.setTree(tree)
	*t.mostPoiter(0) = t.End()
	*t.mostPoiter(1) = t.End()
	*t.rootPoiter() = t.End()
	t.header.setColor(red)
	t.size = 0
	t.nodeOffset = nodeOffset
	t.newNode = func(data interface{}) Iterator {
		var node = newNode(data)
		t.setChild(node, 0, t.End())
		t.setChild(node, 1, t.End())
		t.setParent(node, t.End())
		node.setTree(tree)
		node.setColor(red)
		return node
	}
	t.deleteNode = func(node Iterator) {
		t.setChild(node, 0, nil)
		t.setChild(node, 1, nil)
		t.setParent(node, nil)
		node.setTree(nil)
		deleteNode(node)
	}
	t.compare = compare
	t.sameIterator = sameIterator
	if sameIterator == nil {
		t.sameIterator = unsafeSameIterator
	}
	//t.sameIterator = sameInterface
	t.unique = unique
}

func unsafeSameIterator(a, b Iterator) bool {
	ap := (*[2]uintptr)(unsafe.Pointer(&a))
	bp := (*[2]uintptr)(unsafe.Pointer(&b))
	return ap[1] == bp[1]
}

func sameInterface(a, b interface{}) bool {
	return a == b
}

func (t *RBTree) DeleteNode(node Iterator) {
	t.deleteNode(node)
}

func (t *RBTree) SameIterator(a, b Iterator) bool {
	return t.sameIterator(a, b)
}

func (t *RBTree) Compare(key1, key2 interface{}) int {
	return t.compare(key1, key2)
}

func (t *RBTree) Clear() {
	t.clear(t.root())
	*t.mostPoiter(0) = t.End()
	*t.mostPoiter(1) = t.End()
	*t.rootPoiter() = t.End()
	t.size = 0
}

func (t *RBTree) clear(root Iterator) {
	if t.sameIterator(root, t.End()) {
		return
	}
	t.clear(t.getChild(root, 0))
	t.clear(t.getChild(root, 1))
	t.DeleteNode(root)
}

func (t *RBTree) Unique() bool {
	return t.unique
}

func (t *RBTree) Size() int {
	return t.size
}

func (t *RBTree) Empty() bool {
	return t.size == 0
}

func (t *RBTree) Begin() Iterator {
	return t.most(0)
}

func (t *RBTree) End() Iterator {
	return t.header
}

func (t *RBTree) Next(node Iterator) Iterator {
	if t.sameIterator(node, t.End()) {
		panic("end of tree has no Next()")
	}
	if t.sameIterator(node, t.most(1)) {
		return t.End()
	}
	return t.next(1, node)
}

func (t *RBTree) Last(node Iterator) Iterator {
	if t.sameIterator(node, t.Begin()) {
		panic("begin of tree has no Last()")
	}
	if t.sameIterator(node, t.End()) {
		return t.most(1)
	}
	return t.next(0, node)
}

func (t *RBTree) next(ch int, node Iterator) Iterator {
	if !t.sameIterator(t.getChild(node, ch), t.End()) {
		node = t.getChild(node, ch)
		for !t.sameIterator(t.getChild(node, ch^1), t.End()) {
			node = t.getChild(node, ch^1)
		}
		return node
	}
	for !t.sameIterator(t.getParent(node), t.End()) && t.sameIterator(t.getChild(t.getParent(node), ch), node) {
		node = t.getParent(node)
	}
	return t.getParent(node)
}

func (t *RBTree) Count(key interface{}) (count int) {
	if t.unique {
		if t.sameIterator(t.Find(key), t.End()) {
			return 0
		}
		return 1
	}
	var beg = t.LowerBound(key)
	for !t.sameIterator(beg, t.End()) && t.compare(beg.GetKey(), key) == 0 {
		beg = beg.Next()
		count++
	}
	return count
}

func (t *RBTree) EqualRange(key interface{}) (Iterator, Iterator) {
	return t.LowerBound(key), t.UpperBound(key)
}

func (t *RBTree) Find(key interface{}) Iterator {
	var root = t.root()
	for {
		if t.sameIterator(root, t.End()) {
			return root
		}
		switch cmp := t.compare(key, root.GetKey()); {
		case cmp == 0:
			return root
		case cmp < 0:
			root = t.getChild(root, 0)
		case cmp > 0:
			root = t.getChild(root, 1)
		}
	}
}

func (t *RBTree) Insert(data interface{}) (Iterator, bool) {
	return t.insert(data, func(key interface{}) int {
		return t.compare(data, key)
	})
}
func (t *RBTree) insert(data interface{}, compare func(interface{}) int) (Iterator, bool) {
	var root = t.root()
	var rootPoiter = t.rootPoiter()
	if t.sameIterator(root, t.End()) {
		t.size++
		*rootPoiter = t.newNode(data)
		t.insertAdjust(*rootPoiter)
		*t.mostPoiter(0) = *rootPoiter
		*t.mostPoiter(1) = *rootPoiter
		return *rootPoiter, true
	}
	var parent = t.getParent(root)
	for !t.sameIterator(root, t.End()) {
		parent = root
		switch cmp := compare(root.GetKey()); {
		case cmp == 0:
			if t.unique {
				return t.End(), false
			}
			fallthrough
		case cmp < 0:
			rootPoiter = t.getChildPointer(root, 0)
			root = t.getChild(root, 0)
		case cmp > 0:
			rootPoiter = t.getChildPointer(root, 1)
			root = t.getChild(root, 1)
		}
	}
	t.size++
	*rootPoiter = t.newNode(data)
	t.setParent((*rootPoiter), parent)
	for ch := 0; ch < 2; ch++ {
		if t.sameIterator(parent, t.most(ch)) && t.sameIterator(t.getChild(parent, ch), *rootPoiter) {
			*t.mostPoiter(ch) = *rootPoiter
		}
	}
	t.insertAdjust(*rootPoiter)
	return *rootPoiter, true
}

//insert node is default red
func (t *RBTree) insertAdjust(node Iterator) {
	var parent = t.getParent(node)
	if t.sameIterator(parent, t.End()) {
		//fmt.Println("case 1: insert")
		//node is root,set black
		node.setColor(black)
		return
	}
	if parent.getColor() == black {
		//fmt.Println("case 2: insert")
		//if parent is black,do nothing
		return
	}

	//parent is red,grandpa can't be empty and color is black
	var grandpa = t.getParent(parent)
	var parentCh = 0
	if t.sameIterator(t.getChild(grandpa, 1), parent) {
		parentCh = 1
	}

	var uncle = t.getChild(grandpa, parentCh^1)
	if !t.sameIterator(uncle, t.End()) && uncle.getColor() == red {
		//fmt.Println("case 3: insert")
		//uncle is red
		parent.setColor(black)
		grandpa.setColor(red)
		uncle.setColor(black)
		t.insertAdjust(grandpa)
		return
	}

	var childCh = 0
	if t.sameIterator(t.getChild(parent, 1), node) {
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
	parent.setColor(black)
	grandpa.setColor(red)
	t.rotate(parentCh^1, parent)
}

func (t *RBTree) Erase(key interface{}) (count int) {
	if t.unique {
		var iter = t.Find(key)
		if t.sameIterator(iter, t.End()) {
			return 0
		}
		t.EraseIterator(iter)
		return 1
	}
	var beg = t.LowerBound(key)
	for !t.sameIterator(beg, t.End()) && t.compare(key, beg.GetKey()) == 0 {
		var tmp = beg.Next()
		t.EraseIterator(beg)
		beg = tmp
		count++
	}
	return count
}

func (t *RBTree) EraseIterator(node Iterator) {
	if t.sameIterator(node, t.End()) {
		panic("can't erase empty node")
	}
	t.size--
	if !t.sameIterator(t.getChild(node, 0), t.End()) && !t.sameIterator(t.getChild(node, 1), t.End()) {
		//if node has two child,it's last node must has no more than one child,copy to node and erase last node
		var tmp = node.Last()
		node.CopyData(tmp)
		node = tmp
	}
	//adjust leftmost and rightmost
	for ch := 0; ch < 2; ch++ {
		if t.sameIterator(t.most(ch), node) {
			if ch == 0 {
				*t.mostPoiter(ch) = node.Next()
			} else {
				*t.mostPoiter(ch) = node.Last()
			}
		}
	}
	var child = t.End()
	if !t.sameIterator(t.getChild(node, 0), t.End()) {
		child = t.getChild(node, 0)
	} else if !t.sameIterator(t.getChild(node, 1), t.End()) {
		child = t.getChild(node, 1)
	}
	var parent = t.getParent(node)
	if !t.sameIterator(child, t.End()) {
		t.setParent(child, parent)
	}
	if t.sameIterator(parent, t.End()) {
		*t.rootPoiter() = child
	} else if t.sameIterator(t.getChild(parent, 0), node) {
		t.setChild(parent, 0, child)
	} else {
		t.setChild(parent, 1, child)
	}
	if node.getColor() == black { //if node is red,just erase,otherwise adjust
		t.eraseAdjust(child, parent)
		//fmt.Println("eraseAdjust:")
	}
	t.deleteNode(node)
	return
}

func (t *RBTree) eraseAdjust(node, parent Iterator) {
	if t.sameIterator(parent, t.End()) {
		//node is root
		//fmt.Println("case 1: erase")
		if !t.sameIterator(node, t.End()) {
			node.setColor(black)
		}
		return
	}
	if t.getColor(node) == red {
		//node is red,just set black
		//fmt.Println("case 2: erase")
		node.setColor(black)
		return
	}
	var nodeCh = 0
	if t.sameIterator(t.getChild(parent, 1), node) {
		nodeCh = 1
	}
	var brother = t.getChild(parent, nodeCh^1)
	//after case 1 parent must not be empty node and after case 2 node must be black
	if parent.getColor() == red {
		//parent is red,brother must be black but can't be empty node,because the path has a black node more
		if t.getColor(t.getChild(brother, 0)) == black && t.getColor(t.getChild(brother, 1)) == black {
			//fmt.Println("case 3: erase")
			brother.setColor(red)
			parent.setColor(black)
			return
		}
		if !t.sameIterator(brother, t.End()) && t.getColor(t.getChild(brother, nodeCh)) == red {
			//fmt.Println("case 4: erase", nodeCh)
			parent.setColor(black)
			t.rotate(nodeCh^1, t.getChild(brother, nodeCh))
			t.rotate(nodeCh, t.getChild(parent, nodeCh^1))
			return
		}
		//fmt.Println("case 5: erase")
		t.rotate(nodeCh, brother)
		return
	}
	//parent is black
	if t.getColor(brother) == red {
		//brother is red, it's children must be black
		//fmt.Println("case 6: erase")
		brother.setColor(black)
		parent.setColor(red)
		t.rotate(nodeCh, brother)
		t.eraseAdjust(node, parent) //goto redParent then end
		return
	}
	//brother is black
	if t.getColor(t.getChild(brother, 0)) == black && t.getColor(t.getChild(brother, 1)) == black {
		//fmt.Println("case 7: erase")
		brother.setColor(red)
		t.eraseAdjust(parent, t.getParent(parent))
		return
	}
	if t.getColor(t.getChild(brother, nodeCh)) == red {
		//fmt.Println("case 8: erase", nodeCh)
		t.getChild(brother, nodeCh).setColor(black)
		t.rotate(nodeCh^1, t.getChild(brother, nodeCh))
		t.rotate(nodeCh, t.getChild(parent, nodeCh^1))
		return
	}
	//fmt.Println("case 9: erase", nodeCh)
	t.getChild(brother, nodeCh^1).setColor(black)
	t.rotate(nodeCh, brother)
}

func (t *RBTree) EraseIteratorRange(beg, end Iterator) (count int) {
	for !t.sameIterator(beg, end) {
		var tmp = beg.Next()
		t.EraseIterator(beg)
		beg = tmp
		count++
	}
	return count
}

func (t *RBTree) LowerBound(key interface{}) Iterator {
	var root = t.root()
	var parent = t.End()
	for {
		if root == t.End() {
			if t.sameIterator(parent, t.End()) {
				return parent
			} else if t.compare(key, parent.GetKey()) <= 0 {
				return parent
			}
			return parent.Next()
		}
		parent = root
		if t.compare(key, root.GetKey()) > 0 {
			root = t.getChild(root, 1)
		} else {
			root = t.getChild(root, 0)
		}
	}
}

func (t *RBTree) UpperBound(key interface{}) Iterator {
	var root = t.root()
	var parent = t.End()
	for {
		if root == t.End() {
			if t.sameIterator(parent, t.End()) {
				return parent
			} else if t.compare(key, parent.GetKey()) < 0 {
				return parent
			}
			return parent.Next()
		}
		parent = root
		if t.compare(key, root.GetKey()) >= 0 {
			root = t.getChild(root, 1)
		} else {
			root = t.getChild(root, 0)
		}
	}
}

//ch = 0:take node for center,left rotate parent down,node is parent's right child
//ch = 1:take node for center,right rotate parent down,node is parent's left child
func (t *RBTree) rotate(ch int, node Iterator) {
	var (
		tmp     = t.getChild(node, ch)
		parent  = t.getParent(node)
		grandpa = t.getParent(parent)
	)
	t.setChild(node, ch, parent)
	t.setChild(parent, ch^1, tmp)

	if !t.sameIterator(tmp, t.End()) {
		t.setParent(tmp, parent)
	}
	t.setParent(parent, node)
	t.setParent(node, grandpa)
	if t.sameIterator(grandpa, t.End()) {
		*t.rootPoiter() = node
		return
	}
	if t.sameIterator(t.getChild(grandpa, 0), parent) {
		t.setChild(grandpa, 0, node)
	} else {
		t.setChild(grandpa, 1, node)
	}
}

func (t *RBTree) root() Iterator {
	return t.getParent(t.header)
}

func (t *RBTree) rootPoiter() *Iterator {
	return t.getParentPointer(t.header)
}

//ch = 0: leftmost; ch = 1: rightmost
func (t *RBTree) most(ch int) Iterator {
	return t.getChild(t.header, ch)
}

//ch = 0: leftmostPoiter; ch = 1: rightmostPoiter
func (t *RBTree) mostPoiter(ch int) *Iterator {
	return t.getChildPointer(t.header, ch)
}

func (t *RBTree) getColor(node Iterator) colorType {
	if !t.sameIterator(node, t.End()) {
		return node.getColor()
	}
	return black
}

func (t *RBTree) getChild(node Iterator, ch int) Iterator {
	return *getIteratorPointer(node, t.nodeOffset+offsetChild[ch])
}

func (t *RBTree) getChildPointer(node Iterator, ch int) *Iterator {
	return getIteratorPointer(node, t.nodeOffset+offsetChild[ch])
}

func (t *RBTree) setChild(node Iterator, ch int, child Iterator) {
	*getIteratorPointer(node, t.nodeOffset+offsetChild[ch]) = child
}

func (t *RBTree) getParent(node Iterator) Iterator {
	return *getIteratorPointer(node, t.nodeOffset+offsetParent)
}

func (t *RBTree) getParentPointer(node Iterator) *Iterator {
	return getIteratorPointer(node, t.nodeOffset+offsetParent)
}

func (t *RBTree) setParent(node Iterator, parent Iterator) {
	*getIteratorPointer(node, t.nodeOffset+offsetParent) = parent
}
