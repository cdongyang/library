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

func getChild(node Iterator, ch int, offset uintptr) Iterator {
	return *(*Iterator)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&node)))[1] + offset + uintptr(ch*16)))
}
func ExampleGetChild() {
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	var offset = uintptr(unsafe.Pointer(node.(*SetNode))) - uintptr(unsafe.Pointer(&node.(*SetNode).child[0]))
	node.setChild(0, leftChild)
	var tmp = getChild(node, 0, offset)
	fmt.Println(SameSetNode(tmp, leftChild))
	// Output:
	//true
}

func BenchmarkUnsafeGetChild(b *testing.B) {
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	var offset = uintptr(unsafe.Pointer(node.(*SetNode))) - uintptr(unsafe.Pointer(&node.(*SetNode).child[0]))
	node.setChild(0, leftChild)
	for i := 0; i < b.N; i++ {
		var tmp = getChild(node, 0, offset)
		if !SameSetNode(tmp, leftChild) {
			b.Fatal("not same set node")
		}
	}
}

func BenchmarkGetChild(b *testing.B) {
	var node Iterator = &SetNode{}
	var leftChild Iterator = &SetNode{}
	node.setChild(0, leftChild)
	for i := 0; i < b.N; i++ {
		var tmp = node.getChild(0)
		if !SameSetNode(tmp, leftChild) {
			b.Fatal("not same set node")
		}
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
