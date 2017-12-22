package rbtree_test

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleInterface() {
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

//goos: linux
//goarch: amd64
//pkg: github.com/cdongyang/library/rbtree
//BenchmarkInterface-4                    20000000                83.9 ns/op             8 B/op          0 allocs/op
//BenchmarkInt-4                          2000000000               1.28 ns/op            0 B/op          0 allocs/op
//PASS

func BenchmarkInterface(b *testing.B) {
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = i
		_ = iface.(int)
	}
}

func BenchmarkInt(b *testing.B) {
	var it int
	for i := 0; i < b.N; i++ {
		it = i
		_ = it
	}
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

/*
type testStruct struct {
	a  int
	a0 int64
	a1 int64
	a2 int64
}
goos: linux
goarch: amd64
pkg: github.com/cdongyang/library/rbtree
BenchmarkFunc-4                 2000000000               1.27 ns/op            0 B/op          0 allocs/op
BenchmarkStructFunc-4           2000000000               1.29 ns/op            0 B/op          0 allocs/op
BenchmarkStructPoiterFunc-4     2000000000               1.29 ns/op            0 B/op          0 allocs/op
BenchmarkStructPoiterFunc1-4    2000000000               1.28 ns/op            0 B/op          0 allocs/op
BenchmarkInterfaceFunc-4        100000000               10.2 ns/op             0 B/op          0 allocs/op
PASS


type testStruct struct {
	a  int
	a0 int64
	a1 int64
	a2 int64
	a3 int64
}
goos: linux
goarch: amd64
pkg: github.com/cdongyang/library/rbtree
BenchmarkFunc-4                 2000000000               1.28 ns/op            0 B/op          0 allocs/op
BenchmarkStructFunc-4           100000000               11.8 ns/op             0 B/op          0 allocs/op
BenchmarkStructPoiterFunc-4     2000000000               1.27 ns/op            0 B/op          0 allocs/op
BenchmarkStructPoiterFunc1-4    100000000               10.7 ns/op             0 B/op          0 allocs/op
BenchmarkInterfaceFunc-4        100000000               10.3 ns/op             0 B/op          0 allocs/op
PASS
*/

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

//goos: linux
//goarch: amd64
//pkg: github.com/cdongyang/library/rbtree
//BenchmarkTypeAssertToOrigin-4           2000000000               1.29 ns/op            0 B/op          0 allocs/op
//BenchmarkTypeAssertToInterface-4        30000000                43.5 ns/op             0 B/op          0 allocs/op
//PASS

func BenchmarkTypeAssertToOrigin(b *testing.B) {
	var s interface{} = &testStruct{}
	for i := 0; i < b.N; i++ {
		_ = s.(*testStruct)
	}
}

func BenchmarkTypeAssertToInterface(b *testing.B) {
	var s interface{} = &testStruct{}
	for i := 0; i < b.N; i++ {
		_ = s.(testInterface)
	}
}

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFunc(0)
	}
}

func BenchmarkStructFunc(b *testing.B) {
	var t testStruct
	for i := 0; i < b.N; i++ {
		t.testFunc1(0)
	}
}

func BenchmarkStructPoiterFunc(b *testing.B) {
	var t = &testStruct{}
	for i := 0; i < b.N; i++ {
		t.testFunc(0)
	}
}

func BenchmarkStructPoiterFunc1(b *testing.B) {
	var t = &testStruct{}
	for i := 0; i < b.N; i++ {
		t.testFunc1(0)
	}
}

func BenchmarkInterfaceFunc(b *testing.B) {
	var iface testInterface = &testStruct{}
	for i := 0; i < b.N; i++ {
		iface.testFunc(0)
	}
}
