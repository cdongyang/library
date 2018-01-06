package rbtree

type RBTreer interface {
	Size() int
	Empty() bool
	Begin() Iterator
	End() Iterator
	EndNode() Iterator
	Find(Keyer) Iterator
	LowerBound(Keyer) Iterator
	UpperBound(Keyer) Iterator
	Insert(Keyer) bool
	Erase(Keyer) bool
	EraseIterator(Iterator) bool
	Clear()
}

type Iterator interface {
	Next(Iterator) Iterator
	Last(Iterator) Iterator
	Copy(src Iterator)
	SetKey(Keyer)
	Compare(Keyer) int
	GetKey() Keyer

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

type Keyer interface {
	Compare(Keyer) int
}

//set can be map
//type KeyValuer interface {
//	Keyer
//	GetValue() interface{}
//}

// RBTreeNode is the node of RBTree,you can inherit RBTreeNode to build you own tree node like
//	type mytreeNode struct {
//		RBTreeNode
//		key interface{}
//	}
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

func (node *RBTreeNode) GetKey() Keyer {
	return nil
}

func (node *RBTreeNode) SetKey(Keyer) {

}

// Copy copy the key and val from src to des Iterator.
// inherit RBTreeNode must rewrite this func,otherwise it will copy nothing
func (node *RBTreeNode) Copy(src Iterator) {
}

// Next find the next Iterator follow by root
func (node *RBTreeNode) Next(root Iterator) Iterator {
	var null = node.tree.null
	var same = node.tree.sameIter
	if !same(root.rightChild(), null) {
		root = root.rightChild()
		for !same(root.leftChild(), null) {
			root = root.leftChild()
		}
		return root
	}
	//var root Iterator = node
	//error: current Iterator may not *RBTreeNode,then code next 2 line will never be equal
	//debug for a long time to find this error
	for !same(root.getParent(), null) && same(root.getParent().rightChild(), root) {
		root = root.getParent()
	}
	return root.getParent()
}

// Last find the last Iterator follow by root
func (node *RBTreeNode) Last(root Iterator) Iterator {
	var null = node.tree.null
	var same = node.tree.sameIter
	if !same(root.leftChild(), null) {
		root = root.leftChild()
		for !same(root.rightChild(), null) {
			root = root.rightChild()
		}
		return root
	}
	for !same(root.getParent(), null) && same(root.getParent().leftChild(), root) {
		root = root.getParent()
	}
	return root.getParent()
}

// RBTree is a red-black tree
type RBTree struct {
	unique     bool // node is unique?
	size       int
	header     *RBTreeNode
	null, root Iterator
	newElem    func(Keyer) Iterator
	deleteElem func(Iterator)

	// sameIter judge wheater two Iterator is equal,but the judge between interface is so slow.
	// just rewrite this func like:
	// return a.(*RBTreeNode) == b.(*RBTreeNode)
	sameIter func(Iterator, Iterator) bool
}

// NewCustomRBTree return a poiter of RBTree with the nil value of elem type Iterator,compare func,newElem func,deleteElem func
func NewCustomRBTree(
	unique bool,
	root Iterator,
	newElem func(Keyer) Iterator,
	deleteElem func(Iterator),
	sameIter func(Iterator, Iterator) bool) *RBTree {
	var header = &RBTreeNode{}
	header.parent = root
	header.left = header
	header.right = header
	var tree = &RBTree{
		unique:     unique,
		size:       0,
		null:       root,
		root:       root,
		header:     header,
		newElem:    newElem,
		deleteElem: deleteElem,
		sameIter:   sameIter,
	}
	header.tree = tree
	tree.newElem = func(elem Keyer) Iterator {
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

//Begin return the first Iterator of RBTree
func (t *RBTree) Begin() Iterator {
	var root = t.root
	for !t.sameIter(root, t.null) && !t.sameIter(root.leftChild(), t.null) {
		root = root.leftChild()
	}
	return root
}

//End return the Iterator represent as the end of tree
func (t *RBTree) End() Iterator {
	return t.null
}

//EndNode return the last node Iterator of RBTree
func (t *RBTree) EndNode() Iterator {
	var root = t.root
	for !t.sameIter(root, t.null) && !t.sameIter(root.rightChild(), t.null) {
		root = root.rightChild()
	}
	return root
}

func (t *RBTree) leftRoate(node Iterator) {
	var (
		tmp     = node.leftChild() // tmp maybe nil
		parent  = node.getParent()
		grandpa = parent.getParent()
		same    = t.sameIter
	)
	node.setLeftChild(parent)
	parent.setRightChild(tmp)

	if !same(tmp, t.null) {
		tmp.setParent(parent)
	}
	parent.setParent(node)
	node.setParent(grandpa)
	if same(grandpa, t.null) {
		t.root = node
		return
	}
	if same(grandpa.leftChild(), parent) {
		grandpa.setLeftChild(node)
	} else {
		grandpa.setRightChild(node)
	}
}

func (t *RBTree) rightRoate(node Iterator) {
	var (
		tmp     = node.rightChild()
		parent  = node.getParent()
		grandpa = parent.getParent()
		same    = t.sameIter
	)
	node.setRightChild(parent)
	parent.setLeftChild(tmp)
	if !same(tmp, t.null) {
		tmp.setParent(parent)
	}
	parent.setParent(node)
	node.setParent(grandpa)
	if same(grandpa, t.null) {
		t.root = node
		return
	}
	if same(grandpa.leftChild(), parent) {
		grandpa.setLeftChild(node)
	} else {
		grandpa.setRightChild(node)
	}
}

func (t *RBTree) insertAdjust(node Iterator) {
	var same = t.sameIter
	if same(node.getParent(), t.null) { //if is root,convert to black
		node.setColor(black)
		return
	}
	var parent = node.getParent()
	if parent.getColor() == black {
		return
	}
	var grandpa = parent.getParent()
	if !same(grandpa.leftChild(), t.null) && !same(grandpa.rightChild(), t.null) && grandpa.leftChild().getColor() == red && grandpa.rightChild().getColor() == red {
		grandpa.setColor(red)
		grandpa.leftChild().setColor(black)
		grandpa.rightChild().setColor(black)
		t.insertAdjust(grandpa)
		return
	}
	if same(grandpa.rightChild(), parent) {
		if same(parent.leftChild(), node) {
			t.rightRoate(node)
		} else {
			node = node.getParent()
		}
		t.leftRoate(node)
		//swap node.color node.left.color
		var tmp = node.leftChild().getColor()
		node.leftChild().setColor(node.getColor())
		node.setColor(tmp)
		return
	}
	if same(grandpa.leftChild(), parent) {
		if same(parent.rightChild(), node) {
			t.leftRoate(node)
		} else {
			node = node.getParent()
		}
		t.rightRoate(node)
		//swap root.color root.right.color
		var tmp = node.rightChild().getColor()
		node.rightChild().setColor(node.getColor())
		node.setColor(tmp)
		return
	}
}

//Find find elem from RBTree,if not exist return nil (O(logN))
func (t *RBTree) Find(elem Keyer) Iterator {
	var same = t.sameIter
	var root = t.root
	for {
		if same(root, t.null) {
			return root
		}
		var cmp = root.GetKey().Compare(elem)
		if cmp == 0 {
			return root
		} else if cmp > 0 {
			root = root.leftChild()
		} else {
			root = root.rightChild()
		}
	}
}

//LowerBound return the first Iterator not less than elem (O(logN))
func (t *RBTree) LowerBound(elem Keyer) Iterator {
	var same = t.sameIter
	var root = t.root
	var parent = t.null
	for {
		if same(root, t.null) {
			if same(parent, t.null) {
				return parent
			} else if parent.GetKey().Compare(elem) >= 0 {
				return parent
			}
			return parent.Next(parent)
		}
		parent = root
		if root.GetKey().Compare(elem) < 0 {
			root = root.rightChild()
		} else {
			root = root.leftChild()
		}
	}
}

//UpperBound return the first Iterator greater than elem (O(logN))
func (t *RBTree) UpperBound(elem Keyer) Iterator {
	var same = t.sameIter
	var root = t.root
	var parent = t.null
	for {
		if same(root, t.null) {
			if same(parent, t.null) {
				return parent
			} else if parent.GetKey().Compare(elem) > 0 {
				return parent
			}
			return parent.Next(parent)
		}
		parent = root
		if root.GetKey().Compare(elem) <= 0 {
			root = root.rightChild()
		} else {
			root = root.leftChild()
		}
	}
}

//Insert insert a new elem into RBRree (O(logN)),if elem has been in RBTree,return false,else return true
func (t *RBTree) Insert(elem Keyer) bool {
	var ok = t.insert(elem)
	if ok {
		t.size++
	}
	return ok
}

func (t *RBTree) insert(elem Keyer) bool {
	var same = t.sameIter
	var node = t.root
	if same(node, t.null) {
		t.root = t.newElem(elem)
		t.insertAdjust(t.root)
		return true
	}
	var nodePoiter *Iterator
	var parent = node
	for {
		parent = node
		switch cmp := node.GetKey().Compare(elem); {
		case cmp == 0:
			if t.unique { // if node exists,check tree is unique?
				return false
			}
			fallthrough
		case cmp > 0:
			nodePoiter = node.leftChildPoiter()
			node = node.leftChild()
		case cmp < 0:
			nodePoiter = node.rightChildPoiter()
			node = node.rightChild()
		}
		if same(node, t.null) {
			*nodePoiter = t.newElem(elem)
			(*nodePoiter).setParent(parent)
			t.insertAdjust(*nodePoiter)
			return true
		}
	}
}

//in the path of parent to node,a black node is gone
func (t *RBTree) eraseAdjust(node, parent Iterator) {
	var same = t.sameIter
	if same(parent, t.null) {
		//if node is root,convert to black
		if !same(node, t.null) {
			node.setColor(black)
		}
		return
	}
	if !same(node, t.null) && node.getColor() == red {
		//if node is red,conver to black
		node.setColor(black)
		return
	}
	var brother = t.null
	if same(parent.leftChild(), node) {
		brother = parent.rightChild()
	} else {
		brother = parent.leftChild()
	}
	var brotherLeftChild = brother.leftChild()
	var brotherRightChild = brother.rightChild()
	//parent is red
	if parent.getColor() == red {
		//parent is red,then brother must be black
		if (same(brotherLeftChild, t.null) || brotherLeftChild.getColor() == black) &&
			(same(brotherRightChild, t.null) || brotherRightChild.getColor() == black) {
			//brother's children both are black
			//swap brother.color parent.color
			var tmp = brother.getColor()
			brother.setColor(parent.getColor())
			parent.setColor(tmp)
			return
		}
		if same(parent.leftChild(), node) {
			if !same(brother, t.null) && !same(brotherLeftChild, t.null) && brotherLeftChild.getColor() == red {
				//brother's children are red and ?
				parent.setColor(black)
				t.rightRoate(brotherLeftChild)
				t.leftRoate(parent.rightChild())
				return
			}
			//brother's children are black and red
			t.leftRoate(brother)
			return
		} else {
			if !same(brother, t.null) && !same(brotherRightChild, t.null) && brotherRightChild.getColor() == red {
				//brother's children are ? and red
				parent.setColor(black)
				t.leftRoate(brotherRightChild)
				t.rightRoate(parent.leftChild())
				return
			}
			//brother's children are red and black
			t.rightRoate(brother)
			return
		}
	}
	//parent is black
	if !same(brother, t.null) && brother.getColor() == red {
		//brother is red,brother's children must be black
		//swap parent.color brother.color
		var tmp = parent.getColor()
		parent.setColor(brother.getColor())
		brother.setColor(tmp)
		if same(parent.leftChild(), node) {
			t.leftRoate(brother)
		} else {
			t.rightRoate(brother)
		}
		//after deal in the path of node and parent is still need a black node,adjust again
		t.eraseAdjust(node, parent)
		return
	}
	//brother is black
	if (same(brotherLeftChild, t.null) || brotherLeftChild.getColor() == black) &&
		(same(brotherRightChild, t.null) || brotherRightChild.getColor() == black) {
		//brother's children both are black
		brother.setColor(red)
		t.eraseAdjust(parent, parent.getParent())
		return
	}
	if same(parent.leftChild(), node) {
		if !same(brotherLeftChild, t.null) && brotherLeftChild.getColor() == red {
			//brother's children are red and ?
			brotherLeftChild.setColor(black)
			t.rightRoate(brotherLeftChild)
			t.leftRoate(parent.rightChild())
			return
		}
		//brother's children and black and red
		brotherRightChild.setColor(black)
		t.leftRoate(brother)
		return
	} else {
		if !same(brotherRightChild, t.null) && brotherRightChild.getColor() == red {
			brotherRightChild.setColor(black)
			//brother's children are ? and red
			t.leftRoate(brotherRightChild)
			t.rightRoate(parent.leftChild())
			return
		}
		//brother's children are red and black
		brotherLeftChild.setColor(black)
		t.rightRoate(brother)
		return
	}
}

//Erase erase all value elem from RBTree (O(logN)),if RBTree has elem and erase
func (t *RBTree) Erase(elem Keyer) bool {
	var it = t.Find(elem)
	if t.sameIter(it, t.null) {
		return false
	}
	return t.EraseIterator(it)
}

//EraseIterator erase the Iterator elem from RBTree (O(1))
//if RBTree has elem and erase success,return true,else return false
func (t *RBTree) EraseIterator(node Iterator) bool {
	var ok = t.eraseIterator(node)
	if ok {
		t.size--
	}
	return ok
}

func (t *RBTree) eraseIterator(node Iterator) bool {
	var same = t.sameIter
	if same(node, t.null) {
		return false
	}
	if !same(node.leftChild(), t.null) && !same(node.rightChild(), t.null) {
		// if node has two son,copy the value of (the rightmost child of left child) to node
		// then erase the rightmost child of left child
		var tmp = node.leftChild()
		for !same(tmp.rightChild(), t.null) {
			tmp = tmp.rightChild()
		}
		node.Copy(tmp)
		node = tmp
	}
	if node.getColor() == red {
		//this node is red,so it would'n be root,just erase it
		//because the deal before,node has no right child
		var child = node.leftChild()
		var parent = node.getParent()
		if !same(child, t.null) {
			child.setParent(parent)
		}
		if same(parent.leftChild(), node) {
			parent.setLeftChild(child)
		} else {
			parent.setRightChild(child)
		}
		t.deleteElem(node)
		return true
	}
	var child = t.null
	if !same(node.leftChild(), t.null) {
		child = node.leftChild()
	} else if !same(node.rightChild(), t.null) {
		child = node.rightChild()
	}
	var parent = node.getParent()
	if !same(child, t.null) {
		child.setParent(parent)
	}
	if same(parent, t.null) {
		t.root = child
	} else if same(parent.leftChild(), node) {
		parent.setLeftChild(child)
	} else {
		parent.setRightChild(child)
	}
	t.eraseAdjust(child, parent)
	t.deleteElem(node)
	return true
}

//Clear clear all the elem from RBTree (O(N))
func (t *RBTree) Clear() {
	t.size = 0
	t.clear(&t.root)
}

func (t *RBTree) clear(root *Iterator) {
	if t.sameIter(*root, t.null) {
		return
	}
	t.clear((*root).leftChildPoiter())
	t.clear((*root).rightChildPoiter())
	t.deleteElem(*root)
	*root = t.null
}
