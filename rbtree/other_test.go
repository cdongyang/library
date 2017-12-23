package rbtree_test

import (
	"fmt"
	"reflect"

	"github.com/cdongyang/library/rbtree"
)

func ExampleInterface() {
	type node = rbtree.SetNode
	type Iterator = rbtree.Iterator
	var null Iterator
	var root *node
	var iter Iterator = root
	var NULL = root
	fmt.Println(iter, root)
	fmt.Println(iter == nil, root == nil, iter == root)
	fmt.Println(reflect.DeepEqual(iter, root), reflect.DeepEqual(iter, nil))
	fmt.Println(iter == null, root == null, null == nil)
	fmt.Println(iter == NULL, root == NULL, NULL == nil)
	fmt.Println(iter == (*node)(nil))
	switch iter.(type) {
	case *node:
		fmt.Println("*node")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("other")
	}
	//== judge type first,then judge value
	// Output:
	//<nil> <nil>
	//false true true
	//true false
	//false false true
	//true true true
	//true
	//*node
}

func ExampleNullFunc() {
	defer func() {
		var err = recover()
		fmt.Println(err)
	}()
	type a struct {
		out func(int) int
	}
	var b = &a{out: func(i int) int {
		fmt.Println("b:", i)
		return i
	}}
	fmt.Println("b.out:", b.out(10))
	var c a
	fmt.Println(c.out(10))
	// Output:
	//b: 10
	//b.out: 10
	//runtime error: invalid memory address or nil pointer dereference
}

func testFunc(a int) int {
	return a
}

type testStruct struct {
	a  int
	a0 int64
	a1 int64
	a2 int64
}

func (t *testStruct) testFunc(a int) int {
	return a
}

func (t testStruct) testFunc1(a int) int {
	return a
}

type testInterface interface {
	testFunc(int) int
}

func ExampleInterfaceInherit() {
	type testInterface1 interface {
		testInterface
		testFunc1(int) int
	}
	var b testInterface = &testStruct{}
	if b.(testInterface1).testFunc1(100) == 100 {
		fmt.Println("type assert success!")
	}
	// Output: type assert success!
}
