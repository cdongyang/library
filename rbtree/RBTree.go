package rbtree

type Iterator interface {
	Next(Iterator) Iterator
	Last(Iterator) Iterator
	SetKey(interface{})
	GetKey() interface{}
	Copy(Iterator, Iterator)

	leftChild() Iterator
	leftChildPoiter() *Iterator
	rightChild() Iterator
	rightChildPoiter() *Iterator
	getParent() Iterator
	setLeftChild(Iterator)
	setRightChild(Iterator)
	setParent(Iterator)
	setColor(bool)
	getColor() bool
	init(*RBTree)
}

type RBTreeNode struct {
	left, right, parent Iterator
	color               bool
	tree                *RBTree
}

const black = true
const red = false

func (node *RBTreeNode) init(tree *RBTree) {
	node.tree = tree
	node.left = tree.null
	node.right = tree.null
	node.parent = tree.null
	node.color = red //default color of new node is red
}

func (node *RBTreeNode) leftChildPoiter() *Iterator {
	return &(node.left)
}

func (node *RBTreeNode) rightChildPoiter() *Iterator {
	return &(node.right)
}

func (node *RBTreeNode) leftChild() Iterator {
	return node.left
}

func (node *RBTreeNode) rightChild() Iterator {
	return node.right
}

func (node *RBTreeNode) getParent() Iterator {
	return node.parent
}

func (node *RBTreeNode) setLeftChild(left Iterator) {
	node.left = left
}

func (node *RBTreeNode) setRightChild(right Iterator) {
	node.right = right
}

func (node *RBTreeNode) setParent(parent Iterator) {
	node.parent = parent
}

func (node *RBTreeNode) setColor(color bool) {
	node.color = color
}

func (node *RBTreeNode) getColor() bool {
	return node.color
}

func (node *RBTreeNode) SetKey(interface{}) {
}

func (node *RBTreeNode) GetKey() interface{} {
	return nil
}

func (node *RBTreeNode) Copy(des, src Iterator) {
	des.SetKey(src.GetKey())
}

func (node *RBTreeNode) Next(root Iterator) Iterator {
	var null = node.tree.null
	if root.rightChild() != null {
		root = root.rightChild()
		for root.leftChild() != null {
			root = root.leftChild()
		}
		return root
	}
	//var root Iterator = node
	//error: current Iterator may not *RBTreeNode,then code next 2 line will never be equal
	//debug for a long time to find this error
	for root.getParent() != null && root.getParent().rightChild() == root {
		root = root.getParent()
	}
	return root.getParent()
}

func (node *RBTreeNode) Last(root Iterator) Iterator {
	var null = node.tree.null
	if root.leftChild() != null {
		root = root.leftChild()
		for root.rightChild() != null {
			root = root.rightChild()
		}
		return root
	}
	for root.getParent() != null && root.getParent().leftChild() == root {
		root = root.getParent()
	}
	return root.getParent()
}

type RBTree struct {
	size                   int
	null, root, begin, end Iterator
	compare                func(Iterator, Iterator) int
	newElem                func(interface{}) Iterator
	deleteElem             func(Iterator)
}

func NewRBTree(
	root Iterator,
	compare func(Iterator, Iterator) int,
	newElem func(interface{}) Iterator,
	deleteElem func(Iterator)) *RBTree {
	var tree = &RBTree{
		size:       0,
		null:       root,
		root:       root,
		begin:      root,
		end:        root,
		compare:    compare,
		newElem:    newElem,
		deleteElem: deleteElem,
	}
	tree.newElem = func(elem interface{}) Iterator {
		var iter = newElem(elem)
		iter.init(tree)
		return iter
	}
	return tree
}

//Size return the number of elem in RBTree
func (t *RBTree) Size() int {
	return t.size
}

//Empty return wheather the RBTree is empty
func (t *RBTree) Empty() bool {
	return t.size == 0
}

func (t *RBTree) Begin() Iterator {
	var root = t.root
	for root != t.null && root.leftChild() != t.null {
		root = root.leftChild()
	}
	return root
}

func (t *RBTree) End() Iterator {
	return t.null
}

func (t *RBTree) EndNode() Iterator {
	var root = t.root
	for root != t.null && root.rightChild() != t.null {
		root = root.rightChild()
	}
	return root
}

func (t *RBTree) LeftRoate(node Iterator) {
	var (
		tmp     = node.leftChild() // tmp maybe nil
		parent  = node.getParent()
		grandpa = parent.getParent()
	)
	node.setLeftChild(parent)
	parent.setRightChild(tmp)

	if tmp != t.null {
		tmp.setParent(parent)
	}
	parent.setParent(node)
	node.setParent(grandpa)
	if grandpa == t.null {
		t.root = node
		return
	}
	if grandpa.leftChild() == parent {
		grandpa.setLeftChild(node)
	} else {
		grandpa.setRightChild(node)
	}
}

func (t *RBTree) RightRoate(node Iterator) {
	var (
		tmp     = node.rightChild()
		parent  = node.getParent()
		grandpa = parent.getParent()
	)
	node.setRightChild(parent)
	parent.setLeftChild(tmp)
	if tmp != t.null {
		tmp.setParent(parent)
	}
	parent.setParent(node)
	node.setParent(grandpa)
	if grandpa == t.null {
		t.root = node
		return
	}
	if grandpa.leftChild() == parent {
		grandpa.setLeftChild(node)
	} else {
		grandpa.setRightChild(node)
	}
}

func (t *RBTree) insertAdjust(node Iterator) {
	if node.getParent() == t.null { //if is root,convert to black
		node.setColor(black)
		return
	}
	var parent = node.getParent()
	if parent.getColor() == black {
		return
	}
	var grandpa = parent.getParent()
	if grandpa.leftChild() != t.null && grandpa.rightChild() != t.null && grandpa.leftChild().getColor() == red && grandpa.rightChild().getColor() == red {
		grandpa.setColor(red)
		grandpa.leftChild().setColor(black)
		grandpa.rightChild().setColor(black)
		t.insertAdjust(grandpa)
		return
	}
	if grandpa.rightChild() == parent {
		if parent.leftChild() == node {
			t.RightRoate(node)
		} else {
			node = node.getParent()
		}
		t.LeftRoate(node)
		//swap node.color node.left.color
		var tmp = node.leftChild().getColor()
		node.leftChild().setColor(node.getColor())
		node.setColor(tmp)
		return
	}
	if grandpa.leftChild() == parent {
		if parent.rightChild() == node {
			t.LeftRoate(node)
		} else {
			node = node.getParent()
		}
		t.RightRoate(node)
		//swap root.color root.right.color
		var tmp = node.rightChild().getColor()
		node.rightChild().setColor(node.getColor())
		node.setColor(tmp)
		return
	}
}

//Find find elem from RBTree,if not exist return nil (O(logN))
func (t *RBTree) Find(elem interface{}) Iterator {
	var root = t.root
	var elemIter = t.newElem(elem)
	for {
		if root == t.null {
			return t.null
		}
		var cmp = t.compare(elemIter, root)
		if cmp == 0 {
			return root
		} else if cmp < 0 {
			root = root.leftChild()
		} else {
			root = root.rightChild()
		}
	}
}

//LowerBound return the first Iterator not less than elem (O(logN))
func (t *RBTree) LowerBound(elem interface{}) Iterator {
	var root = t.root
	var elemIter = t.newElem(elem)
	var parent = t.null
	for {
		if root == t.null {
			if parent == t.null {
				return parent
			} else if t.compare(elemIter, parent) <= 0 { //elem >= parent
				return parent
			}
			return parent.Next(parent)
		}
		parent = root
		if t.compare(elemIter, root) > 0 {
			root = root.rightChild()
		} else {
			root = root.leftChild()
		}
	}
}

//UpperBound return the first Iterator greater than elem (O(logN))
func (t *RBTree) UpperBound(elem interface{}) Iterator {
	var root = t.root
	var elemIter = t.newElem(elem)
	var parent = t.null
	for {
		if root == t.null {
			if parent == t.null {
				return parent
			} else if t.compare(elemIter, parent) < 0 {
				return parent
			}
			return parent.Next(parent)
		}
		parent = root
		if t.compare(elemIter, root) >= 0 {
			root = root.rightChild()
		} else {
			root = root.leftChild()
		}
	}
}

//Insert insert a new elem into RBRree (O(logN)),if elem has been in RBTree,return false,else return true
func (t *RBTree) Insert(elem interface{}) bool {
	var ok = t.insert(t.newElem(elem))
	if ok {
		t.size++
	}
	return ok
}

func (t *RBTree) insert(elem Iterator) bool {
	var node = t.root
	if node == t.null {
		t.root = elem
		t.insertAdjust(t.root)
		return true
	}
	var nodePoiter *Iterator
	var parent = node
	for {
		parent = node
		if t.compare(elem, node) == 0 {
			return false
		} else if t.compare(elem, node) < 0 {
			nodePoiter = node.leftChildPoiter()
			node = node.leftChild()
		} else {
			nodePoiter = node.rightChildPoiter()
			node = node.rightChild()
		}
		if node == t.null {
			*nodePoiter = elem
			(*nodePoiter).setParent(parent)
			t.insertAdjust(*nodePoiter)
			return true
		}
	}
}

//in the path of parent to node,a black node is gone
func (t *RBTree) eraseAdjust(node, parent Iterator) {
	if parent == t.null {
		//if node is root,convert to black
		if node != t.null {
			node.setColor(black)
		}
		return
	}
	if node != t.null && node.getColor() == red {
		//if node is red,conver to black
		node.setColor(black)
		return
	}
	var brother = t.null
	if parent.leftChild() == node {
		brother = parent.rightChild()
	} else {
		brother = parent.leftChild()
	}
	//parent is red
	if parent.getColor() == red {
		//parent is red,then brother must be black
		if (brother.leftChild() == t.null || brother.leftChild().getColor() == black) &&
			(brother.rightChild() == t.null || brother.rightChild().getColor() == black) {
			//brother's children both are black
			//swap brother.color parent.color
			var tmp = brother.getColor()
			brother.setColor(parent.getColor())
			parent.setColor(tmp)
			return
		}
		if parent.leftChild() == node {
			if brother != t.null && brother.leftChild() != t.null && brother.leftChild().getColor() == red {
				//brother's children are red and ?
				parent.setColor(black)
				t.RightRoate(brother.leftChild())
				t.LeftRoate(parent.rightChild())
				return
			}
			//brother's children are black and red
			t.LeftRoate(brother)
			return
		} else {
			if brother != t.null && brother.rightChild() != t.null && brother.rightChild().getColor() == red {
				//brother's children are ? and red
				parent.setColor(black)
				t.LeftRoate(brother.rightChild())
				t.RightRoate(parent.leftChild())
				return
			}
			//brother's children are red and black
			t.RightRoate(brother)
			return
		}
	}
	//parent is black
	if brother != t.null && brother.getColor() == red {
		//brother is red,brother's children must be black
		//swap parent.color brother.color
		var tmp = parent.getColor()
		parent.setColor(brother.getColor())
		brother.setColor(tmp)
		if parent.leftChild() == node {
			t.LeftRoate(brother)
		} else {
			t.RightRoate(brother)
		}
		//after deal in the path of node and parent is still need a black node,adjust again
		t.eraseAdjust(node, parent)
		return
	}
	//brother is black
	if (brother.leftChild() == t.null || brother.leftChild().getColor() == black) &&
		(brother.rightChild() == t.null || brother.rightChild().getColor() == black) {
		//brother's children both are black
		brother.setColor(red)
		t.eraseAdjust(parent, parent.getParent())
		return
	}
	if parent.leftChild() == node {
		if brother.leftChild() != t.null && brother.leftChild().getColor() == red {
			//brother's children are red and ?
			brother.leftChild().setColor(black)
			t.RightRoate(brother.leftChild())
			t.LeftRoate(parent.rightChild())
			return
		}
		//brother's children and black and red
		brother.rightChild().setColor(black)
		t.LeftRoate(brother)
		return
	} else {
		if brother.rightChild() != t.null && brother.rightChild().getColor() == red {
			brother.rightChild().setColor(black)
			//brother's children are ? and red
			t.LeftRoate(brother.rightChild())
			t.RightRoate(parent.leftChild())
			return
		}
		//brother's children are red and black
		brother.leftChild().setColor(black)
		t.RightRoate(brother)
		return
	}
}

//Erase erase all value elem from RBTree (O(logN)),if RBTree has elem and erase su
func (t *RBTree) Erase(elem interface{}) bool {
	var it = t.Find(elem)
	if it == t.null {
		return false
	}
	return t.EraseIterator(it)
}

//EraseIterator erase the Iterator elem from RBTree (O(logN))
//if RBTree has elem and erase success,return true,else return false
func (t *RBTree) EraseIterator(node Iterator) bool {
	var ok = t.eraseIterator(node)
	if ok {
		t.size--
	}
	return ok
}
func (t *RBTree) eraseIterator(node Iterator) bool {
	if node == t.null {
		return false
	}
	if node.leftChild() != t.null && node.rightChild() != t.null {
		// if node has two son,copy the value of (the rightmost child of left child) to node
		// then erase the rightmost child of left child
		var tmp = node.leftChild()
		for tmp.rightChild() != t.null {
			tmp = tmp.rightChild()
		}
		node.Copy(node, tmp)
		node = tmp
	}
	if node.getColor() == red {
		//this node is red,so it would'n be root,just erase it
		//because the deal before,node has no right child
		var child = node.leftChild()
		var parent = node.getParent()
		if child != t.null {
			child.setParent(parent)
		}
		if parent.leftChild() == node {
			parent.setLeftChild(child)
		} else {
			parent.setRightChild(child)
		}
		t.deleteElem(node)
		return true
	}
	var child = t.null
	if node.leftChild() != t.null {
		child = node.leftChild()
	} else if node.rightChild() != t.null {
		child = node.rightChild()
	}
	var parent = node.getParent()
	if child != t.null {
		child.setParent(parent)
	}
	if parent == t.null {
		t.root = child
	} else if parent.leftChild() == node {
		parent.setLeftChild(child)
	} else {
		parent.setRightChild(child)
	}
	t.eraseAdjust(child, parent)
	t.deleteElem(node)
	return true
}

//Clear erase all the elem from RBTree (O(N))
func (t *RBTree) Clear(root *Iterator) {
	t.size = 0
	t.clear(root)
}
func (t *RBTree) clear(root *Iterator) {
	if *root == t.null {
		return
	}
	t.clear((*root).leftChildPoiter())
	t.clear((*root).rightChildPoiter())
	t.deleteElem(*root)
	*root = t.null
}
