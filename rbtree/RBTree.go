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
}

type colorType bool

const (
	red   = false
	black = true
)

type rtype struct {
	size       uintptr
	ptrdata    uintptr        // number of bytes in the type that can contain pointers
	hash       uint32         // hash of type; avoids computation in hash tables
	tflag      uint8          // extra type information flags
	align      uint8          // alignment of variable with this type
	fieldAlign uint8          // alignment of struct field with this type
	kind       uint8          // enumeration for C
	alg        unsafe.Pointer // algorithm table
	gcdata     *byte          // garbage collection data
	str        uint32         // string form
	ptrToThis  uint32         // type for pointer to this type, may be zero
}

type eface struct {
	itype   *rtype
	pointer unsafe.Pointer
}

// embed RBTreeNode should not use pointer, because the operation of RBTree treate is as value
type RBTreeNode struct {
	child  [2]eface
	parent eface
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

type RBTreer interface {
	DeleteNode(node Iterator)
	SameIterator(a, b Iterator) bool
	Compare(key1, key2 unsafe.Pointer) int
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
		compare func(unsafe.Pointer, unsafe.Pointer) int,
		getKeyPointer func(unsafe.Pointer) unsafe.Pointer,
		unique bool,
	)
}

type RBTree struct {
	header     eface
	size       int
	newNode    func(interface{}) Iterator
	deleteNode func(Iterator)
	compare    func(unsafe.Pointer, unsafe.Pointer) int
	//compareFun unsafe.Pointer
	getKeyPointer func(unsafe.Pointer) unsafe.Pointer
	nodeOffset    uintptr
	unique        bool
}

func NewRBTreer(
	t RBTreer,
	header Iterator,
	nodeOffset uintptr,
	newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(unsafe.Pointer, unsafe.Pointer) int,
	getKeyPointer func(unsafe.Pointer) unsafe.Pointer,
	unique bool) RBTreer {
	t.init(t, header, nodeOffset, newNode, deleteNode, compare,
		getKeyPointer,
		unique)
	return t
}

var nilIface = eface{nil, nil}

func (t *RBTree) init(
	tree RBTreer,
	header Iterator,
	nodeOffset uintptr,
	newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(unsafe.Pointer, unsafe.Pointer) int,
	getKeyPointer func(unsafe.Pointer) unsafe.Pointer,
	unique bool) {
	t.nodeOffset = nodeOffset
	header.setTree(tree)
	t.header = iterator2eface(header)
	t.setColor(t.header, red)
	*t.mostPoiter(0) = t.end()
	*t.mostPoiter(1) = t.end()
	*t.rootPoiter() = t.end()
	t.size = 0
	t.newNode = func(data interface{}) Iterator {
		var node = newNode(data)
		var nodeIface = iterator2eface(node)
		t.setChild(nodeIface, 0, t.end())
		t.setChild(nodeIface, 1, t.end())
		t.setParent(nodeIface, t.end())
		node.setTree(tree)
		t.setColor(nodeIface, red)
		return node
	}
	t.deleteNode = func(node Iterator) {
		var nodeIface = iterator2eface(node)
		t.setChild(nodeIface, 0, nilIface)
		t.setChild(nodeIface, 1, nilIface)
		t.setParent(nodeIface, nilIface)
		node.setTree(nil)
		deleteNode(node)
	}
	t.compare = compare
	//t.compareFun = unsafe.Pointer(&compare)
	t.getKeyPointer = getKeyPointer
	t.unique = unique
}

func unsafeSameIterator(a, b Iterator) bool {
	return (*eface)(unsafe.Pointer(&a)).pointer == (*eface)(unsafe.Pointer(&b)).pointer
}

func sameIface(a, b eface) bool {
	return a.pointer == b.pointer
}

func (t *RBTree) DeleteNode(node Iterator) {
	t.deleteNode(node)
}

func (t *RBTree) SameIterator(a, b Iterator) bool {
	return unsafeSameIterator(a, b)
	//return sameIface(a, b)
}

func (t *RBTree) Compare(key1, key2 unsafe.Pointer) int {
	return t.compare(key1, key2)
}

func (t *RBTree) Clear() {
	t.clear(t.root())
	*t.mostPoiter(0) = t.end()
	*t.mostPoiter(1) = t.end()
	*t.rootPoiter() = t.end()
	t.size = 0
}

func (t *RBTree) clear(root eface) {
	if sameIface(root, t.end()) {
		return
	}
	t.clear(t.getChild(root, 0))
	t.clear(t.getChild(root, 1))
	t.DeleteNode(eface2iterator(root))
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
	return eface2iterator(t.begin())
}

func (t *RBTree) begin() eface {
	return t.most(0)
}

func (t *RBTree) End() Iterator {
	return eface2iterator(t.end())
}

func (t *RBTree) end() eface {
	return t.header
}

func (t *RBTree) Next(node Iterator) Iterator {
	return eface2iterator(t.next(iterator2eface(node)))
}
func (t *RBTree) next(node eface) eface {
	if sameIface(node, t.end()) {
		panic("end of tree has no Next()")
	}
	if sameIface(node, t.most(1)) {
		return t.end()
	}
	return t.gothrough(1, node)
}

func (t *RBTree) Last(node Iterator) Iterator {
	return eface2iterator(t.last(iterator2eface(node)))
}
func (t *RBTree) last(node eface) eface {
	if sameIface(node, t.begin()) {
		panic("begin of tree has no Last()")
	}
	if sameIface(node, t.end()) {
		return t.most(1)
	}
	return t.gothrough(0, node)
}

func (t *RBTree) gothrough(ch int, node eface) eface {
	if !sameIface(t.getChild(node, ch), t.end()) {
		node = t.getChild(node, ch)
		for !sameIface(t.getChild(node, ch^1), t.end()) {
			node = t.getChild(node, ch^1)
		}
		return node
	}
	for !sameIface(t.getParent(node), t.end()) && sameIface(t.getChild(t.getParent(node), ch), node) {
		node = t.getParent(node)
	}
	return t.getParent(node)
}

func (t *RBTree) Count(key interface{}) (count int) {
	var keyPointer = interface2pointer(key)
	if t.unique {
		if sameIface(t.find(keyPointer), t.end()) {
			return 0
		}
		return 1
	}
	var beg = t.lowerBound(keyPointer)
	for !sameIface(beg, t.end()) && t.compare(t.getKeyPointer(beg.pointer), keyPointer) == 0 {
		beg = t.next(beg)
		count++
	}
	return count
}

func (t *RBTree) EqualRange(key interface{}) (Iterator, Iterator) {
	var keyPointer = interface2pointer(key)
	return t.LowerBound(keyPointer), t.UpperBound(keyPointer)
}

func (t *RBTree) Find(key interface{}) Iterator {
	return eface2iterator(t.find(noescape(interface2pointer(key))))
}
func (t *RBTree) find(keyPointer unsafe.Pointer) eface {
	var root = t.root()
	for {
		if sameIface(root, t.end()) {
			return root
		}
		switch cmp := t.compare(keyPointer, t.getKeyPointer(root.pointer)); {
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
	iter, ok := t.insert(data, interface2pointer(data))
	return eface2iterator(iter), ok
}
func (t *RBTree) insert(data interface{}, key unsafe.Pointer) (eface, bool) {
	var root = t.root()
	var rootPoiter = t.rootPoiter()
	if sameIface(root, t.end()) {
		t.size++
		*rootPoiter = iterator2eface(t.newNode(data))
		t.insertAdjust(*rootPoiter)
		*t.mostPoiter(0) = *rootPoiter
		*t.mostPoiter(1) = *rootPoiter
		return *rootPoiter, true
	}
	var parent = t.getParent(root)
	for !sameIface(root, t.end()) {
		parent = root
		switch cmp := t.compare(key, t.getKeyPointer(root.pointer)); {
		case cmp == 0:
			if t.unique {
				return t.end(), false
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
	*rootPoiter = iterator2eface(t.newNode(data))
	t.setParent((*rootPoiter), parent)
	for ch := 0; ch < 2; ch++ {
		if sameIface(parent, t.most(ch)) && sameIface(t.getChild(parent, ch), *rootPoiter) {
			*t.mostPoiter(ch) = *rootPoiter
		}
	}
	t.insertAdjust(*rootPoiter)
	return *rootPoiter, true
}

//insert node is default red
func (t *RBTree) insertAdjust(node eface) {
	var parent = t.getParent(node)
	if sameIface(parent, t.end()) {
		//fmt.Println("case 1: insert")
		//node is root,set black
		t.setColor(node, black)
		return
	}
	if t.getColor(parent) == black {
		//fmt.Println("case 2: insert")
		//if parent is black,do nothing
		return
	}

	//parent is red,grandpa can't be empty and color is black
	var grandpa = t.getParent(parent)
	var parentCh = 0
	if sameIface(t.getChild(grandpa, 1), parent) {
		parentCh = 1
	}

	var uncle = t.getChild(grandpa, parentCh^1)
	if !sameIface(uncle, t.end()) && t.getColor(uncle) == red {
		//fmt.Println("case 3: insert")
		//uncle is red
		t.setColor(parent, black)
		t.setColor(grandpa, red)
		t.setColor(uncle, black)
		t.insertAdjust(grandpa)
		return
	}

	var childCh = 0
	if sameIface(t.getChild(parent, 1), node) {
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
	t.setColor(parent, black)
	t.setColor(grandpa, red)
	t.rotate(parentCh^1, parent)
}

func (t *RBTree) Erase(key interface{}) (count int) {
	var keyPointer = noescape(interface2pointer(key))
	if t.unique {
		var iter = t.find(keyPointer)
		if sameIface(iter, t.end()) {
			return 0
		}
		t.eraseIterator(iter)
		return 1
	}
	var beg = t.lowerBound(keyPointer)
	for !sameIface(beg, t.end()) && t.compare(keyPointer, t.getKeyPointer(beg.pointer)) == 0 {
		var tmp = t.next(beg)
		t.eraseIterator(beg)
		beg = tmp
		count++
	}
	return count
}

func (t *RBTree) EraseIterator(node Iterator) {
	t.eraseIterator(iterator2eface(node))
}
func (t *RBTree) eraseIterator(node eface) {
	if sameIface(node, t.end()) {
		panic("can't erase empty node")
	}
	t.size--
	if !sameIface(t.getChild(node, 0), t.end()) && !sameIface(t.getChild(node, 1), t.end()) {
		//if node has two child,it's last node must has no more than one child,copy to node and erase last node
		var tmp = t.last(node)
		eface2iterator(node).CopyData(eface2iterator(tmp))
		node = tmp
	}
	//adjust leftmost and rightmost
	for ch := 0; ch < 2; ch++ {
		if sameIface(t.most(ch), node) {
			if ch == 0 {
				*t.mostPoiter(ch) = t.next(node)
			} else {
				*t.mostPoiter(ch) = t.last(node)
			}
		}
	}
	var child = t.end()
	if !sameIface(t.getChild(node, 0), t.end()) {
		child = t.getChild(node, 0)
	} else if !sameIface(t.getChild(node, 1), t.end()) {
		child = t.getChild(node, 1)
	}
	var parent = t.getParent(node)
	if !sameIface(child, t.end()) {
		t.setParent(child, parent)
	}
	if sameIface(parent, t.end()) {
		*t.rootPoiter() = child
	} else if sameIface(t.getChild(parent, 0), node) {
		t.setChild(parent, 0, child)
	} else {
		t.setChild(parent, 1, child)
	}
	if t.getColor(node) == black { //if node is red,just erase,otherwise adjust
		t.eraseAdjust(child, parent)
		//fmt.Println("eraseAdjust:")
	}
	t.deleteNode(eface2iterator(node))
	return
}

func (t *RBTree) eraseAdjust(node, parent eface) {
	if sameIface(parent, t.end()) {
		//node is root
		//fmt.Println("case 1: erase")
		if !sameIface(node, t.end()) {
			t.setColor(node, black)
		}
		return
	}
	if t.mustGetColor(node) == red {
		//node is red,just set black
		//fmt.Println("case 2: erase")
		t.setColor(node, black)
		return
	}
	var nodeCh = 0
	if sameIface(t.getChild(parent, 1), node) {
		nodeCh = 1
	}
	var brother = t.getChild(parent, nodeCh^1)
	//after case 1 parent must not be empty node and after case 2 node must be black
	if t.getColor(parent) == red {
		//parent is red,brother must be black but can't be empty node,because the path has a black node more
		if t.mustGetColor(t.getChild(brother, 0)) == black && t.mustGetColor(t.getChild(brother, 1)) == black {
			//fmt.Println("case 3: erase")
			t.setColor(brother, red)
			t.setColor(parent, black)
			return
		}
		if !sameIface(brother, t.end()) && t.mustGetColor(t.getChild(brother, nodeCh)) == red {
			//fmt.Println("case 4: erase", nodeCh)
			t.setColor(parent, black)
			t.rotate(nodeCh^1, t.getChild(brother, nodeCh))
			t.rotate(nodeCh, t.getChild(parent, nodeCh^1))
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
		t.setColor(brother, black)
		t.setColor(parent, red)
		t.rotate(nodeCh, brother)
		t.eraseAdjust(node, parent) //goto redParent then end
		return
	}
	//brother is black
	if t.mustGetColor(t.getChild(brother, 0)) == black && t.mustGetColor(t.getChild(brother, 1)) == black {
		//fmt.Println("case 7: erase")
		t.setColor(brother, red)
		t.eraseAdjust(parent, t.getParent(parent))
		return
	}
	if t.mustGetColor(t.getChild(brother, nodeCh)) == red {
		//fmt.Println("case 8: erase", nodeCh)
		t.setColor(t.getChild(brother, nodeCh), black)
		t.rotate(nodeCh^1, t.getChild(brother, nodeCh))
		t.rotate(nodeCh, t.getChild(parent, nodeCh^1))
		return
	}
	//fmt.Println("case 9: erase", nodeCh)
	t.setColor(t.getChild(brother, nodeCh^1), black)
	t.rotate(nodeCh, brother)
}

func (t *RBTree) EraseIteratorRange(beg, end Iterator) (count int) {
	return t.eraseIteratorRange(iterator2eface(beg), iterator2eface(end))
}
func (t *RBTree) eraseIteratorRange(beg, end eface) (count int) {
	for !sameIface(beg, end) {
		var tmp = t.next(beg)
		t.eraseIterator(beg)
		beg = tmp
		count++
	}
	return count
}

func (t *RBTree) LowerBound(key interface{}) Iterator {
	return eface2iterator(t.lowerBound(noescape(interface2pointer(key))))
}
func (t *RBTree) lowerBound(keyPointer unsafe.Pointer) eface {
	var root = t.root()
	var parent = t.end()
	for {
		if root == t.end() {
			if sameIface(parent, t.end()) {
				return parent
			} else if t.compare(keyPointer, t.getKeyPointer(parent.pointer)) <= 0 {
				return parent
			}
			return t.next(parent)
		}
		parent = root
		if t.compare(keyPointer, t.getKeyPointer(root.pointer)) > 0 {
			root = t.getChild(root, 1)
		} else {
			root = t.getChild(root, 0)
		}
	}
}

func (t *RBTree) UpperBound(key interface{}) Iterator {
	return eface2iterator(t.upperBound(noescape(interface2pointer(key))))
}
func (t *RBTree) upperBound(keyPointer unsafe.Pointer) eface {
	var root = t.root()
	var parent = t.end()
	for {
		if root == t.end() {
			if sameIface(parent, t.end()) {
				return parent
			} else if t.compare(keyPointer, t.getKeyPointer(parent.pointer)) < 0 {
				return parent
			}
			return t.next(parent)
		}
		parent = root
		if t.compare(keyPointer, t.getKeyPointer(root.pointer)) >= 0 {
			root = t.getChild(root, 1)
		} else {
			root = t.getChild(root, 0)
		}
	}
}

//ch = 0:take node for center,left rotate parent down,node is parent's right child
//ch = 1:take node for center,right rotate parent down,node is parent's left child
func (t *RBTree) rotate(ch int, node eface) {
	var (
		tmp     = t.getChild(node, ch)
		parent  = t.getParent(node)
		grandpa = t.getParent(parent)
	)
	t.setChild(node, ch, parent)
	t.setChild(parent, ch^1, tmp)

	if !sameIface(tmp, t.end()) {
		t.setParent(tmp, parent)
	}
	t.setParent(parent, node)
	t.setParent(node, grandpa)
	if sameIface(grandpa, t.end()) {
		*t.rootPoiter() = node
		return
	}
	if sameIface(t.getChild(grandpa, 0), parent) {
		t.setChild(grandpa, 0, node)
	} else {
		t.setChild(grandpa, 1, node)
	}
}

func (t *RBTree) root() eface {
	return t.getParent(t.header)
}

func (t *RBTree) rootPoiter() *eface {
	return t.getParentPointer(t.header)
}

//ch = 0: leftmost; ch = 1: rightmost
func (t *RBTree) most(ch int) eface {
	return t.getChild(t.header, ch)
}

//ch = 0: leftmostPoiter; ch = 1: rightmostPoiter
func (t *RBTree) mostPoiter(ch int) *eface {
	return t.getChildPointer(t.header, ch)
}

func (t *RBTree) mustGetColor(node eface) colorType {
	if !sameIface(node, t.end()) {
		return t.getColor(node)
	}
	return black
}

func (t *RBTree) getChild(node eface, ch int) eface {
	return *getIteratorPointer(node, t.nodeOffset+offsetChild[ch])
}

func (t *RBTree) getChildPointer(node eface, ch int) *eface {
	return getIteratorPointer(node, t.nodeOffset+offsetChild[ch])
}

func (t *RBTree) setChild(node eface, ch int, child eface) {
	*getIteratorPointer(node, t.nodeOffset+offsetChild[ch]) = child
}

func (t *RBTree) getParent(node eface) eface {
	return *getIteratorPointer(node, t.nodeOffset+offsetParent)
}

func (t *RBTree) getParentPointer(node eface) *eface {
	return getIteratorPointer(node, t.nodeOffset+offsetParent)
}

func (t *RBTree) setParent(node eface, parent eface) {
	*getIteratorPointer(node, t.nodeOffset+offsetParent) = parent
}

func (t *RBTree) getColor(node eface) colorType {
	return *getColorPointer(node, t.nodeOffset+offsetColor)
}

func (t *RBTree) setColor(node eface, color colorType) {
	*getColorPointer(node, t.nodeOffset+offsetColor) = color
}

/*func (t *RBTree) compare(a, b unsafe.Pointer) int {
	fun := (*func(a, b unsafe.Pointer) int)(t.compareFun)
	return (*fun)(a, b)
}*/

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

func getIteratorPointer(node eface, offset uintptr) *eface {
	return (*eface)(unsafe.Pointer(uintptr(node.pointer) + offset))
}

func getColorPointer(node eface, offset uintptr) *colorType {
	return (*colorType)(unsafe.Pointer(uintptr(node.pointer) + offset))
}

func iterator2eface(node Iterator) eface {
	return *(*eface)(unsafe.Pointer(&node))
}

func eface2iterator(node eface) Iterator {
	return *(*Iterator)(unsafe.Pointer(&node))
}

func interface2eface(node interface{}) eface {
	return *(*eface)(unsafe.Pointer(&node))
}

func eface2interface(node eface) interface{} {
	return *(*interface{})(unsafe.Pointer(&node))
}

func interface2pointer(a interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&a)).pointer
}

func CompareInt(a, b unsafe.Pointer) int {
	return *(*int)(a) - *(*int)(b)
}

// copy from package runtime
// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input.  noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
