package rbtree

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestRBTreeNode(t *testing.T) {
	var setNode Iterator = &SetNode{}
	fmt.Printf("setNode: %p, RBTreeNode: %p, data: %p\n", setNode.(*SetNode), &setNode.(*SetNode).RBTreeNode, &setNode.(*SetNode).data)
	fmt.Printf("RBTreeNode.child: %p,RBTreeNode.parent: %p\n", &setNode.(*SetNode).RBTreeNode.child, &setNode.(*SetNode).RBTreeNode.parent)
	fmt.Printf("RBTreeNode.tree: %p,RBTreeNode.color: %p\n", &setNode.(*SetNode).RBTreeNode.tree, &setNode.(*SetNode).RBTreeNode.color)
	elem := reflect.TypeOf(setNode).Elem()
	for i := 0; i < elem.NumField(); i++ {
		secElem := elem.Field(i)
		fmt.Println(secElem.Name, secElem.Offset, secElem.Index, secElem.Anonymous, secElem.PkgPath, secElem.Type.String())
	}
	// Output:
	//
}

func ExampleNode() {
	var setNode Iterator = &SetNode{}
	var ptr = unsafe.Pointer(setNode.(*SetNode))
	var copyNode = *((*Iterator)(ptr))
	fmt.Println(SameSetNode(copyNode, setNode))
	fmt.Printf("setNode: %p,Iterator Pointer: %x\n", setNode.(*SetNode), (*(*[2]uintptr)(unsafe.Pointer(&ptr)))[1])
	fmt.Println(uintptr(unsafe.Pointer(setNode.(*SetNode))) == (*(*[2]uintptr)(unsafe.Pointer(&ptr)))[1])
	// Output:
	//false
}

/*
func ExampleGetChild() {
	var set = NewSet(nil)
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	*getIteratorPointer(node, offsetChild[0]) = leftChild
	var tmp = getChild(&set.RBTree, node, 0)
	fmt.Println(SameSetNode(tmp, leftChild))
	// Output:
	//true
}

func BenchmarkUnsafeGetChild(b *testing.B) {
	var set = NewSet(nil)
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	*getIteratorPointer(node, offsetChild[0]) = leftChild
	for i := 0; i < b.N; i++ {
		_ = getChild(&set.RBTree, node, 0)
	}
}
*/
func getIfaceUnsafePointer(iface interface{}) unsafe.Pointer {
	return ((*[2]unsafe.Pointer)(unsafe.Pointer(&iface)))[1]
}

func getIfaceUintPtr(iface interface{}) uintptr {
	return ((*[2]uintptr)(unsafe.Pointer(&iface)))[1]
}

func BenchmarkGetUnsafePointer(b *testing.B) {
	var iface interface{} = 1
	for i := 0; i < b.N; i++ {
		_ = getIfaceUnsafePointer(iface)
	}
}

func BenchmarkGetUintPtr(b *testing.B) {
	var iface interface{} = 1
	for i := 0; i < b.N; i++ {
		_ = getIfaceUintPtr(iface)
	}
}

func getChild(t *RBTree, node eface, ch int) eface {
	return *getIteratorPointer(node, t.nodeOffset+offsetChild[ch])
}

/*
func BenchmarkSetGetChild(b *testing.B) {
	var set = NewSet(nil)
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	*getIteratorPointer(node, offsetChild[0]) = leftChild
	for i := 0; i < b.N; i++ {
		_ = set.getChild(node, 0)
	}
}

func BenchmarkGetIteratorPointer(b *testing.B) {
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	*getIteratorPointer(node, offsetChild[0]) = leftChild
	for i := 0; i < b.N; i++ {
		_ = *getIteratorPointer(node, offsetChild[0])
	}
}
*/
// the size to store interface{} is 16B
// type eface struct {
// 	_type *_type
// 	data  unsafe.Pointer
// }
//
// type iface struct {
// 	tab  *itab
// 	data unsafe.Pointer
// }

func ExampleNodeSize() {
	fmt.Println(reflect.TypeOf(RBTreeNode{}).Size())
	fmt.Println(reflect.TypeOf(SetNode{}).Size())
	fmt.Println(reflect.TypeOf(MapNode{}).Size())
	//Output:
	//72
	//88
	//104
}

func ExampleSameSetNode() {
	var a Iterator = &SetNode{data: 1}
	var b Iterator = &SetNode{data: 2}
	var c Iterator = &SetNode{data: 1}
	fmt.Println(SameSetNode(a, b))
	fmt.Println(SameSetNode(a, a))
	fmt.Println(SameSetNode(a, c))
	// Output:
	//false
	//true
	//false
}

//BenchmarkNewSetNode-4   	300000000	         4.90 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNewSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &SetNode{data: 1}
	}
}

//BenchmarkNewHeapSetNode-4   	20000000	        85.1 ns/op	      96 B/op	       1 allocs/op
func BenchmarkNewHeapSetNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var node *SetNode
		func() {
			node = new(SetNode)
		}()
		node.data = 1
	}
}

//BenchmarkNewNode-4   	20000000	       125 ns/op	      96 B/op	       1 allocs/op
func BenchmarkSetNewNode(b *testing.B) {
	var set = NewSet(CompareInt)
	for i := 0; i < b.N; i++ {
		_ = set.newNode(0)
	}
}

//BenchmarkNewPoiterNode-4   	20000000	       123 ns/op	      96 B/op	       1 allocs/op
func BenchmarkSetNewPoiterNode(b *testing.B) {
	var set = NewSet(CompareInt)
	var a = 0
	var c = &a
	for i := 0; i < b.N; i++ {
		_ = set.newNode(c)
	}
}

// BenchmarkRBTreeCompare-4                200000000                8.67 ns/op            0 B/op          0 allocs/op
func BenchmarkRBTreeCompare(b *testing.B) {
	var set = NewSet(CompareInt)
	var a, c = set.newNode(0).(*SetNode), set.newNode(1).(*SetNode)
	for i := 0; i < b.N; i++ {
		_ = set.RBTree.compare(set.getKeyPointer(unsafe.Pointer(a)), set.getKeyPointer(unsafe.Pointer(c)))
	}
}

// BenchmarkGetKey-4   	2000000000	         0.46 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetKey(b *testing.B) {
	var set = NewSet(CompareInt)
	var a = set.newNode(1).(*SetNode)
	for i := 0; i < b.N; i++ {
		_ = set.getKeyPointer(unsafe.Pointer(a))
	}
}

// BenchmarkIteratorGetKey-4   	500000000	         3.49 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIteratorGetKey(b *testing.B) {
	var set = NewSet(CompareInt)
	var a = set.newNode(1)
	for i := 0; i < b.N; i++ {
		_ = set.getKeyPointer(interface2pointer(a))
	}
}

// BenchmarkEfaceTransform-4   	500000000	         3.07 ns/op	       0 B/op	       0 allocs/op
func BenchmarkEfaceTransform(b *testing.B) {
	var compare = func(a, b interface{}) {
	}
	var a, c interface{} = 1, 2
	for i := 0; i < b.N; i++ {
		compare(a, c)
	}
}

func BenchmarkIfaceTransform(b *testing.B) {
	var compare = func(a, b Iterator) {
	}
	var a, c Iterator = &SetNode{data: 1}, &SetNode{data: 2}
	for i := 0; i < b.N; i++ {
		compare(a, c)
	}
}

// BenchmarkCompare-4   	200000000	         6.75 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompare(b *testing.B) {
	var set = NewSet(CompareInt)
	var a, c = set.newNode(0).(*SetNode), set.newNode(1).(*SetNode)
	for i := 0; i < b.N; i++ {
		_ = set.compare(interface2pointer(a.data), interface2pointer(c.data))
	}
}

// BenchmarkUnsafeCompareInt-4   	500000000	         3.82 ns/op	       0 B/op	       0 allocs/op
func BenchmarkUnsafeCompareInt(b *testing.B) {
	var set = NewSet(CompareInt)
	var a, c interface{} = 1, 2
	for i := 0; i < b.N; i++ {
		_ = set.compare(interface2pointer(a), interface2pointer(c))
	}
}

// BenchmarkCompareInt-4   	300000000	         3.69 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompareInt(b *testing.B) {
	var set = NewSet(CompareInt)
	var a, c = 1, 2
	for i := 0; i < b.N; i++ {
		_ = set.compare(unsafe.Pointer(&a), unsafe.Pointer(&c))
	}
}

// BenchmarkCompareIteratorGetKey-4   	100000000	        12.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompareIteratorGetKey(b *testing.B) {
	var set = NewSet(CompareInt)
	var a, c = set.newNode(0), set.newNode(1)
	for i := 0; i < b.N; i++ {
		_ = set.compare(interface2pointer(a.GetKey()), interface2pointer(c.GetKey()))
	}
}

// BenchmarkSetCompareInt-4   	500000000	         3.61 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSetCompareInt(b *testing.B) {
	var set = NewSet(CompareInt)
	var a, c = 1, 2
	for i := 0; i < b.N; i++ {
		_ = set.compare(unsafe.Pointer(&a), unsafe.Pointer(&c))
	}
}

// BenchmarkRBTreeCompareInt-4   	500000000	         3.89 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRBTreeCompareInt(b *testing.B) {
	var set = NewSet(CompareInt)
	var tree = &set.RBTree
	var a, c = 1, 2
	for i := 0; i < b.N; i++ {
		_ = tree.compare(unsafe.Pointer(&a), unsafe.Pointer(&c))
	}
}

// BenchmarkCompareInt-4   	2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
// go test -run=^$ -gcflags '-l' -benchmem -v -bench=^BenchmarkCompareInt$
// BenchmarkCompareInt-4           500000000                3.03 ns/op            0 B/op          0 allocs/op
func BenchmarkRawCompareInt(b *testing.B) { //不内联时速度和interface method/struct func element调用差不多
	var a, c = 1, 2
	for i := 0; i < b.N; i++ {
		_ = CompareInt(unsafe.Pointer(&a), unsafe.Pointer(&c))
	}
}

func compareInt(a, b interface{}) int {
	return a.(int) - b.(int)
}

// BenchmarkAssertCompareInt-4   	2000000000	         0.40 ns/op	       0 B/op	       0 allocs/op
// go test -run=^$ -gcflags '-l' -benchmem -v -bench=^BenchmarkAssertCompareInt$
// BenchmarkAssertCompareInt-4     50000000                26.3 ns/op             0 B/op          0 allocs/op
func BenchmarkAssertCompareInt(b *testing.B) { //不内联时类型断言代价明显增高
	var a, c interface{} = 1, 2
	for i := 0; i < b.N; i++ {
		_ = compareInt(a, c)
	}
}

// BenchmarkGetKeyPointer-4   	500000000	         3.34 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetKeyPointer(b *testing.B) {
	var set = NewSet(CompareInt)
	var setNode = set.newNode(0)
	var pointer = iterator2eface(setNode).pointer
	for i := 0; i < b.N; i++ {
		_ = set.getKeyPointer(pointer)
	}
}

// SameSetNode judge wheather two Iterator is the same node, but not equal node
// type assert to *SetNode to speed up
func SameSetNode(a, b Iterator) bool {
	aa, aok := a.(*SetNode)
	bb, bok := b.(*SetNode)
	if aok && bok {
		return aa == bb
	}
	return a == b
}

// BenchmarkSameSetNode-4   	2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSameSetNode(b *testing.B) {
	var a, c Iterator = &SetNode{data: 1}, &SetNode{data: 2}
	for i := 0; i < b.N; i++ {
		_ = SameSetNode(a, c)
	}
}

// BenchmarkSameInterface-4   	200000000	         7.32 ns/op	       0 B/op	       0 allocs/op
func sameInterface(a, b interface{}) bool {
	return a == b
}
func BenchmarkSameInterface(b *testing.B) {
	var a, c Iterator = &SetNode{data: 1}, &SetNode{data: 2}
	for i := 0; i < b.N; i++ {
		_ = sameInterface(a, c)
	}
}

// BenchmarkUnsafeSameNode-4   	2000000000	         1.57 ns/op	       0 B/op	       0 allocs/op
func BenchmarkUnsafeSameNode(b *testing.B) {
	var a, c Iterator = &SetNode{data: 1}, &SetNode{data: 2}
	for i := 0; i < b.N; i++ {
		_ = unsafeSameIterator(a, c)
	}
}

// BenchmarkSetSameNode-4   	300000000	         5.44 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSetSameNode(b *testing.B) {
	var a, c = iterator2eface(&SetNode{data: 1}), iterator2eface(&SetNode{data: 2})
	for i := 0; i < b.N; i++ {
		_ = sameIface(a, c)
	}
}

// BenchmarkIface2Iterator-4   	2000000000	         0.82 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIface2Iterator(b *testing.B) {
	var a Iterator = &SetNode{}
	for i := 0; i < b.N; i++ {
		_ = iterator2eface(a)
	}
}

// BenchmarkIterator2Iface-4   	2000000000	         0.93 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIterator2Iface(b *testing.B) {
	var a = iterator2eface(Iterator(&SetNode{}))
	for i := 0; i < b.N; i++ {
		_ = eface2iterator(a)
	}
}

func ExampleGetKey() {
	var a, b Iterator = &SetNode{data: 1}, &SetNode{data: 2}
	_, _ = a, b
	method, ok := reflect.TypeOf(a).MethodByName("GetKey")
	if !ok {
		panic("no method GetKey")
	}
	value := method.Func.Call([]reflect.Value{reflect.ValueOf(a)})
	fmt.Println(value[0].Interface())
	// Output:
	// 1
}

// BenchmarkTypeAssert-4   	2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
// go test -run=^$ -gcflags '-l' -benchmem -v -bench=^BenchmarkTypeAssert$
// BenchmarkTypeAssert-4           100000000               15.0 ns/op             0 B/op          0 allocs/op
func AssertInt(a interface{}) int {
	return a.(int)
}
func BenchmarkTypeAssert(b *testing.B) {
	var a interface{} = 1
	for i := 0; i < b.N; i++ {
		_ = AssertInt(a)
	}
}

// 函数不被内联时interface{}断言很慢,用unsafe.Pointer强制转换优化
// 结构体调用method很快,但调用"函数成员"或者interface调用method很慢
// 小函数容易在编译是被内联,因而并发测试直接调用时会特别快,通过结构体"函数成员"调用会很慢,测试时可以用 -gcflags '-l'禁止内联

func (t *RBTree) escape1(a interface{}) {
	_ = a
}
func (t *RBTree) escape2(a interface{}) {
	_ = interface2pointer(a)
}
func (t *RBTree) escape3(a interface{}) {
	p := interface2pointer(a)
	t.compare(p, p)
}
func (t *RBTree) escape4(a interface{}) {
	p := interface2pointer(a)
	CompareInt(p, p)
}
func (t *RBTree) escape5(a interface{}) {
	p := noescape(interface2pointer(a))
	CompareInt(p, p)
}
func ExampleEscape() {
	var n float64
	var t = NewMap(CompareInt)
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		t.escape1(a)
	})
	fmt.Println(n, "allocs/op")
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		t.escape2(a)
	})
	fmt.Println(n, "allocs/op")
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		t.escape3(a)
	})
	fmt.Println(n, "allocs/op")
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		t.escape4(a)
	})
	fmt.Println(n, "allocs/op")
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		t.escape5(a)
	})
	fmt.Println(n, "allocs/op")
	// Output:
	//0 allocs/op
	//0 allocs/op
	//1 allocs/op
	//0 allocs/op
	//0 allocs/op
}

func ExampleMapEscape() {
	var pair = Pair{1, 2}
	var t = NewMap(CompareInt)
	var n float64
	t.Insert(pair)
	t.Insert(Pair{3, 4})
	n = testing.AllocsPerRun(10, func() {
		if _, ok := t.Insert(pair); ok {
			panic("repeat insert")
		}
	})
	fmt.Println(n, "allocs/op")
	n = testing.AllocsPerRun(10, func() {
		if t.Find(1) == t.End() {
			panic("not found")
		}
	})
	fmt.Println(n, "allocs/op")
	n = testing.AllocsPerRun(10, func() {
		if t.Erase(1) != 1 {
			panic("not erase")
		}
		t.Insert(pair)
	})
	fmt.Println(n, "allocs/op")
	// Output:
	//0 allocs/op
	//0 allocs/op
	//0 allocs/op
}

func getType(a interface{}) *rtype {
	return (*eface)(unsafe.Pointer(&a)).itype
}

func ExampleType() {
	var (
		i   int    = 1
		i32 int32  = 1
		i64 int64  = 1
		s   string = "1"
	)
	fmt.Println(getType(int(1)) == getType(i))
	fmt.Println(getType(int32(1)) == getType(i32))
	fmt.Println(getType(int64(1)) == getType(i64))
	fmt.Println(getType(string("1")) == getType(s))
	// Output:
	//true
	//true
	//true
}
