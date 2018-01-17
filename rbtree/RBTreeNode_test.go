package rbtree

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func ExampleRBTreeNode() {
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

func getIfaceUnsafePointer(iface interface{}) unsafe.Pointer {
	return unsafe.Pointer(((*[2]uintptr)(unsafe.Pointer(&iface)))[1])
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
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := 0; i < b.N; i++ {
		_ = set.newNode(0)
	}
}

//BenchmarkNewPoiterNode-4   	20000000	       123 ns/op	      96 B/op	       1 allocs/op
func BenchmarkSetNewPoiterNode(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a = 0
	var c = &a
	for i := 0; i < b.N; i++ {
		_ = set.newNode(c)
	}
}

// BenchmarkRBTreeCompare-4                200000000                8.67 ns/op            0 B/op          0 allocs/op
func BenchmarkRBTreeCompare(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a, c = set.newNode(0).(*SetNode), set.newNode(1).(*SetNode)
	for i := 0; i < b.N; i++ {
		_ = set.RBTree.compare(a.GetKey(), c.GetKey())
	}
}

// BenchmarkGetKey-4   	2000000000	         0.46 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGetKey(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a = set.newNode(1).(*SetNode)
	for i := 0; i < b.N; i++ {
		_ = a.GetKey()
	}
}

// BenchmarkIteratorGetKey-4   	500000000	         3.49 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIteratorGetKey(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a = set.newNode(1)
	for i := 0; i < b.N; i++ {
		_ = a.GetKey()
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
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a, c = set.newNode(0).(*SetNode), set.newNode(1).(*SetNode)
	for i := 0; i < b.N; i++ {
		_ = set.compare(a.GetKey(), c.GetKey())
	}
}

// BenchmarkCompareInt-4   	300000000	         4.60 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompareInt(b *testing.B) {
	var compare = func(a, b interface{}) int {
		return a.(int) - b.(int)
	}
	var a, c interface{} = 1, 2
	for i := 0; i < b.N; i++ {
		_ = compare(a, c)
	}
}

// BenchmarkCompareIteratorGetKey-4   	100000000	        12.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCompareIteratorGetKey(b *testing.B) {
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a, c = set.newNode(0), set.newNode(1)
	for i := 0; i < b.N; i++ {
		_ = set.compare(a.GetKey(), c.GetKey())
	}
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
	var set = NewSet(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	var a, c = set.newNode(0), set.newNode(1)
	var tree = &set.RBTree
	for i := 0; i < b.N; i++ {
		_ = tree.sameIterator(a, c)
	}
}
