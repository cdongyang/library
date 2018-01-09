/*
learn about red-black tree
	https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91
*/
package rbtree

type Iterator interface {
	Next() Iterator
	Last() Iterator
	GetData() (data interface{})
	GetKey() (key interface{})
	GetValue() (value interface{})
	SetValue(value interface{})
	CopyData(src Iterator)
	GetTree() RBTreer

	getChild(int) Iterator
	setChild(int, Iterator)
	getChildPoiter(int) *Iterator
	getParent() Iterator
	setParent(Iterator)
	getParentPoiter() *Iterator
	getColor() colorType
	setColor(colorType)
	setTree(RBTreer)
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
	panic("use base Next()")
	//return node.GetTree().Next(node)
}

// Last return last Iterator of this
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) Last() Iterator {
	panic("use base Last()")
	//return node.GetTree().Last(node)
}

// GetData get the data of this
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) GetData() interface{} {
	return nil
}

func (node *RBTreeNode) GetKey() interface{} {
	return nil
}

func (node *RBTreeNode) GetValue() interface{} {
	return nil
}

func (node *RBTreeNode) SetValue(interface{}) {
}

//CopyData copy the node data to this from src
// inherit RBTreeNode must rewrite this method
func (node *RBTreeNode) CopyData(src Iterator) {
}

func (node *RBTreeNode) getChild(ch int) Iterator {
	return node.child[ch]
}

func (node *RBTreeNode) setChild(ch int, child Iterator) {
	node.child[ch] = child
}

func (node *RBTreeNode) getChildPoiter(ch int) *Iterator {
	return &node.child[ch]
}

func (node *RBTreeNode) getParent() Iterator {
	return node.parent
}

func (node *RBTreeNode) setParent(parent Iterator) {
	node.parent = parent
}

func (node *RBTreeNode) getParentPoiter() *Iterator {
	return &node.parent
}

func (node *RBTreeNode) getColor() colorType {
	return node.color
}

func (node *RBTreeNode) setColor(color colorType) {
	node.color = color
}

func (node *RBTreeNode) GetTree() RBTreer {
	return node.tree
}

func (node *RBTreeNode) setTree(tree RBTreer) {
	node.tree = tree
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
}

type RBTree struct {
	header       Iterator
	size         int
	newNode      func(interface{}) Iterator
	deleteNode   func(Iterator)
	compare      func(interface{}, interface{}) int
	sameIterator func(Iterator, Iterator) bool
	unique       bool
}

func NewRBTree(newNode func(interface{}) Iterator,
	deleteNode func(Iterator),
	compare func(interface{}, interface{}) int,
	sameIterator func(Iterator, Iterator) bool,
	unique bool) *RBTree {
	var t = new(RBTree)
	t.header = newNode(nil)
	t.header.setTree(t)
	*t.mostPoiter(0) = t.End()
	*t.mostPoiter(1) = t.End()
	*t.rootPoiter() = t.End()
	t.header.setColor(red)
	t.size = 0
	t.newNode = func(data interface{}) Iterator {
		var it = newNode(data)
		it.setChild(0, t.End())
		it.setChild(1, t.End())
		it.setParent(t.End())
		it.setTree(t)
		it.setColor(red)
		return it
	}
	t.deleteNode = func(node Iterator) {
		node.setChild(0, nil)
		node.setChild(1, nil)
		node.setParent(nil)
		node.setTree(nil)
		deleteNode(node)
	}
	t.compare = compare
	t.sameIterator = sameIterator
	t.unique = unique
	return t
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
	t.clear(root.getChild(0))
	t.clear(root.getChild(1))
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
	if !t.sameIterator(node.getChild(ch), t.End()) {
		node = node.getChild(ch)
		for !t.sameIterator(node.getChild(ch^1), t.End()) {
			node = node.getChild(ch ^ 1)
		}
		return node
	}
	for !t.sameIterator(node.getParent(), t.End()) && t.sameIterator(node.getParent().getChild(ch), node) {
		node = node.getParent()
	}
	return node.getParent()
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
			root = root.getChild(0)
		case cmp > 0:
			root = root.getChild(1)
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
	var parent = root.getParent()
	for !t.sameIterator(root, t.End()) {
		parent = root
		switch cmp := compare(root.GetKey()); {
		case cmp == 0:
			if t.unique {
				return t.End(), false
			}
			fallthrough
		case cmp < 0:
			rootPoiter = root.getChildPoiter(0)
			root = root.getChild(0)
		case cmp > 0:
			rootPoiter = root.getChildPoiter(1)
			root = root.getChild(1)
		}
	}
	t.size++
	*rootPoiter = t.newNode(data)
	(*rootPoiter).setParent(parent)
	for ch := 0; ch < 2; ch++ {
		if t.sameIterator(parent, t.most(ch)) && t.sameIterator(parent.getChild(ch), *rootPoiter) {
			*t.mostPoiter(ch) = *rootPoiter
		}
	}
	t.insertAdjust(*rootPoiter)
	return *rootPoiter, true
}

//insert node is default red
func (t *RBTree) insertAdjust(node Iterator) {
	var parent = node.getParent()
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
	var grandpa = parent.getParent()
	var parentCh = 0
	if t.sameIterator(grandpa.getChild(1), parent) {
		parentCh = 1
	}

	var uncle = grandpa.getChild(parentCh ^ 1)
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
	if t.sameIterator(parent.getChild(1), node) {
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
	if !t.sameIterator(node.getChild(0), t.End()) && !t.sameIterator(node.getChild(1), t.End()) {
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
	if !t.sameIterator(node.getChild(0), t.End()) {
		child = node.getChild(0)
	} else if !t.sameIterator(node.getChild(1), t.End()) {
		child = node.getChild(1)
	}
	var parent = node.getParent()
	if !t.sameIterator(child, t.End()) {
		child.setParent(parent)
	}
	if t.sameIterator(parent, t.End()) {
		*t.rootPoiter() = child
	} else if t.sameIterator(parent.getChild(0), node) {
		parent.setChild(0, child)
	} else {
		parent.setChild(1, child)
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
	if t.sameIterator(parent.getChild(1), node) {
		nodeCh = 1
	}
	var brother = parent.getChild(nodeCh ^ 1)
	//after case 1 parent must not be empty node and after case 2 node must be black
	if parent.getColor() == red {
		//parent is red,brother must be black but can't be empty node,because the path has a black node more
		if t.getColor(brother.getChild(0)) == black && t.getColor(brother.getChild(1)) == black {
			//fmt.Println("case 3: erase")
			brother.setColor(red)
			parent.setColor(black)
			return
		}
		if !t.sameIterator(brother, t.End()) && t.getColor(brother.getChild(nodeCh)) == red {
			//fmt.Println("case 4: erase", nodeCh)
			parent.setColor(black)
			t.rotate(nodeCh^1, brother.getChild(nodeCh))
			t.rotate(nodeCh, parent.getChild(nodeCh^1))
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
	if t.getColor(brother.getChild(0)) == black && t.getColor(brother.getChild(1)) == black {
		//fmt.Println("case 7: erase")
		brother.setColor(red)
		t.eraseAdjust(parent, parent.getParent())
		return
	}
	if t.getColor(brother.getChild(nodeCh)) == red {
		//fmt.Println("case 8: erase", nodeCh)
		brother.getChild(nodeCh).setColor(black)
		t.rotate(nodeCh^1, brother.getChild(nodeCh))
		t.rotate(nodeCh, parent.getChild(nodeCh^1))
		return
	}
	//fmt.Println("case 9: erase", nodeCh)
	brother.getChild(nodeCh ^ 1).setColor(black)
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
			root = root.getChild(1)
		} else {
			root = root.getChild(0)
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
			root = root.getChild(1)
		} else {
			root = root.getChild(0)
		}
	}
}

//ch = 0:take node for center,left rotate parent down,node is parent's right child
//ch = 1:take node for center,right rotate parent down,node is parent's left child
func (t *RBTree) rotate(ch int, node Iterator) {
	var (
		tmp     = node.getChild(ch)
		parent  = node.getParent()
		grandpa = parent.getParent()
	)
	node.setChild(ch, parent)
	parent.setChild(ch^1, tmp)

	if !t.sameIterator(tmp, t.End()) {
		tmp.setParent(parent)
	}
	parent.setParent(node)
	node.setParent(grandpa)
	if t.sameIterator(grandpa, t.End()) {
		*t.rootPoiter() = node
		return
	}
	if t.sameIterator(grandpa.getChild(0), parent) {
		grandpa.setChild(0, node)
	} else {
		grandpa.setChild(1, node)
	}
}

func (t *RBTree) root() Iterator {
	return t.header.getParent()
}

func (t *RBTree) rootPoiter() *Iterator {
	return t.header.getParentPoiter()
}

//ch = 0: leftmost; ch = 1: rightmost
func (t *RBTree) most(ch int) Iterator {
	return t.header.getChild(ch)
}

//ch = 0: leftmostPoiter; ch = 1: rightmostPoiter
func (t *RBTree) mostPoiter(ch int) *Iterator {
	return t.header.getChildPoiter(ch)
}

func (t *RBTree) getColor(node Iterator) colorType {
	if !t.sameIterator(node, t.End()) {
		return node.getColor()
	}
	return black
}
