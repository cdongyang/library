## A red-black tree with an API similar to C++ STL's.
a high performance red-black tree with less heap objects.
## INSTALLTION
```
    go get github.com/cdongyang/library/rbtree
```

## EXAMPLE
more example please see example_test.go
```go
func ExampleMap() {
	var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
	// key type: int, value type: *int
	mp := rbtree.NewMap(int(0), new(int), func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := range slice {
		mp.Insert(slice[i], &slice[i])
	}
	var indexOf = func(p *int) int {
		for i := range slice {
			if &slice[i] == p {
				return i
			}
		}
		return -1
	}
	// iterator
	for it, i := mp.Begin(), 0; it != mp.End(); it = it.Next() {
		fmt.Println(it.GetKey(), indexOf(it.GetVal().(*int)))
		i++
	}
	//Output:
	//1 0
	//2 6
	//3 4
	//4 1
	//5 3
	//6 2
	//7 5
	//9 7
}
```

## TYPES AND FUNCTION
```go
func NoescapeInterface(x interface{}) interface{}
type Map
    func NewMap(key, val interface{}, compare func(a, b interface{}) int) *Map
    func NewMultiMap(key, val interface{}, compare func(a, b interface{}) int) *Map
    func (s *Map) Begin() MapNode
    func (s *Map) Count(key interface{}) (count int)
    func (t *Map) Empty() bool
    func (s *Map) End() MapNode
    func (s *Map) EqualRange(key interface{}) (beg, end MapNode)
    func (s *Map) Erase(key interface{}) (count int)
    func (s *Map) EraseNode(n MapNode)
    func (s *Map) EraseNodeRange(beg, end MapNode) (count int)
    func (s *Map) Find(key interface{}) MapNode
    func (t *Map) GetMaxSpan() uint32
    func (s *Map) Init(unique bool, key, val interface{}, compare func(a, b interface{}) int)
    func (s *Map) Insert(key interface{}, val interface{}) (MapNode, bool)
    func (s *Map) LowerBound(key interface{}) MapNode
    func (t *Map) SetMaxSpan(maxSpan uint32)
    func (t *Map) Size() int
    func (t *Map) Unique() bool
    func (s *Map) UpperBound(key interface{}) MapNode
type MapNode
    func (n MapNode) GetData() (key, val interface{})
    func (n MapNode) GetKey() interface{}
    func (n MapNode) GetMap() *Map
    func (n MapNode) GetVal() interface{}
    func (n MapNode) Last() MapNode
    func (n MapNode) Next() MapNode
    func (n MapNode) SetVal(val interface{})
type Set
    func NewMultiSet(data interface{}, compare func(a, b interface{}) int) *Set
    func NewSet(data interface{}, compare func(a, b interface{}) int) *Set
    func (s *Set) Begin() SetNode
    func (s *Set) Count(data interface{}) (count int)
    func (t *Set) Empty() bool
    func (s *Set) End() SetNode
    func (s *Set) EqualRange(data interface{}) (beg, end SetNode)
    func (s *Set) Erase(data interface{}) (count int)
    func (s *Set) EraseNode(n SetNode)
    func (s *Set) EraseNodeRange(beg, end SetNode) (count int)
    func (s *Set) Find(data interface{}) SetNode
    func (t *Set) GetMaxSpan() uint32
    func (s *Set) Init(unique bool, data interface{}, compare func(a, b interface{}) int)
    func (s *Set) Insert(data interface{}) (SetNode, bool)
    func (s *Set) LowerBound(data interface{}) SetNode
    func (t *Set) SetMaxSpan(maxSpan uint32)
    func (t *Set) Size() int
    func (t *Set) Unique() bool
    func (s *Set) UpperBound(data interface{}) SetNode
type SetNode
    func (n SetNode) GetData() interface{}
    func (n SetNode) GetSet() *Set
    func (n SetNode) Last() SetNode
    func (n SetNode) Next() SetNode
```

## MEMORY ALLOC
I use a slice of block memory to store node data. In addition, i store the unuse node in a two-dimension queue. when it needs a node, i pop from begin of queue, and push a node in queue when delete a node, so the node will reuse. And each block memory can store curSpan nodes, however, the curSpan is dynamic change following the tree size. If curSpan < maxSpan, curSpan = 1 << (high bit of tree size), if curSpan > maxSpan, curSpan = maxSpan, so the number of heap objects will be close to O(tree size / maxSpan) when tree size if so large.

## ATTENTION
Because of the strategy of memory alloc, the data of interface{} return by method GetKey(),GetVal() or GetData() will store in block memory, so we should do the type assert immediately when get this kind of interface{}. If not, don't hold it for a long time, otherwise the block memory will not collect by GC until you never hold the interface{}.What's more, you should only read the interface{} in compare function.

### 参考资料:
- [wiki红黑树讲解]( https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91)
- 侯捷 <<STL源码剖析>>
- 郝林 <<Go并发编程实战>>
- [qyuhen性能优化技巧](https://www.jianshu.com/u/44d32fdece77)
- [qyuhen 学习笔记](https://github.com/qyuhen/book)
- [深入理解 Go Interface](http://legendtkl.com/2017/06/12/understanding-golang-interface/)
- [Go Interface 源码剖析](http://legendtkl.com/2017/07/01/golang-interface-implement/)
- [Golang逃逸分析](https://gocn.io/article/355)
- sort,runtime,reflect包源码
- [golang 垃圾回收机制（算法）](https://lengzzz.com/note/gc-in-golang)
- [golang的垃圾回收（GC）机制]( http://blog.csdn.net/liangzhiyang/article/details/52670021)
