## 迭代器Iterator
```go
type Iterator interface {
	Next(Iterator) Iterator
	Last(Iterator) Iterator
	Copy(des Iterator, src Iterator)

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
```
Iterator是基础迭代器,其它的迭代器都是通过继承Iterator来实现红黑树的操作

```go
type RBTreeNode struct {
	left, right, parent Iterator
	color               bool
	tree                *RBTree
}
```
RBTreeNode是红黑树节点,是Iterator的实现,其它容器的节点里都含有RBTreeNode节点,从而隐藏上层节点实现红黑树操作

```go
type SetIterator interface {
	Iterator
	GetKey() interface{}
	SetKey(interface{})
}

type SetNode struct {
	RBTreeNode
	key interface{}
}
```
泛化Set的节点和迭代器

```go
type Seter interface {
	Size() int
	Empty() bool
	Begin() SetIterator
	End() SetIterator
	EndNode() SetIterator
	Find(interface{}) SetIterator
	LowerBound(interface{}) SetIterator
	UpperBound(interface{}) SetIterator
	Insert(interface{}) bool
	Erase(interface{}) bool
	EraseIterator(Iterator) bool
	Clear()
}
```