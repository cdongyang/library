package interfaceConvert

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"
)

type struct16b struct {
	a int64
	b int64
}

type struct24b struct {
	a int64
	b int64
	c int64
}

type struct32b struct {
	a int64
	b int64
	c int64
	d int64
}

func (s *struct32b) Fun() {

}

func (s *struct32b) Fun1() {

}

type struct32ber interface {
	Fun()
}

type struct32ber1 interface {
	struct32ber
	Fun1()
}

type struct40b struct {
	a int64
	b int64
	c int64
	d int64
	e int64
}

type struct48b struct {
	a int64
	b int64
	c int64
	d int64
	e int64
	f int64
}

func ExampleSize() {
	var s16 struct16b
	var s24 struct24b
	var s32 struct32b
	var s40 struct40b
	var s48 struct48b
	fmt.Println("sizeof struct16b", reflect.TypeOf(s16).Size())
	fmt.Println("sizeof struct24b", reflect.TypeOf(s24).Size())
	fmt.Println("sizeof struct32b", reflect.TypeOf(s32).Size())
	fmt.Println("sizeof struct40b", reflect.TypeOf(s40).Size())
	fmt.Println("sizeof struct48b", reflect.TypeOf(s48).Size())

	fmt.Println("sizeof int", reflect.TypeOf(int(0)).Size())
	fmt.Println("sizeof int32", reflect.TypeOf(int32(0)).Size())
	fmt.Println("sizeof int64", reflect.TypeOf(int64(0)).Size())
	// Output:
	//sizeof struct16b 16
	//sizeof struct24b 24
	//sizeof struct32b 32
	//sizeof struct40b 40
	//sizeof struct48b 48
	//sizeof int 8
	//sizeof int32 4
	//sizeof int64 8
}

var mem runtime.MemStats
var outmem = false

func memStats() {
	if !outmem {
		return
	}
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

//HeapAlloc: 73128 HeapInuse: 466944 HeapObjects: 287 HeapIdle 286720 HeapReleased 0 HeapSys 753664
//2000000000	         0.40 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertIntCompare(b *testing.B) {
	var it int
	for i := 0; i < b.N; i++ {
		it = i
		_ = it
	}
	memStats()
}

//HeapAlloc: 70008 HeapInuse: 344064 HeapObjects: 285 HeapIdle 409600 HeapReleased 0 HeapSys 753664
//2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDefineVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a = 10
		_ = a
	}
	memStats()
}

//HeapAlloc: 73928 HeapInuse: 352256 HeapObjects: 291 HeapIdle 335872 HeapReleased 0 HeapSys 688128
//2000000000	         0.40 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDefineHeapVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a = new(int)
		_ = a
	}
	memStats()
}

//HeapAlloc: 1213864 HeapInuse: 1490944 HeapObjects: 71433 HeapIdle 4374528 HeapReleased 0 HeapSys 5865472
//50000000	        32.5 ns/op	       8 B/op	       1 allocs/op
func BenchmarkInterfaceConvertInt(b *testing.B) { //1 allocs/op和主要耗时都在iface = i
	var iface interface{} = 1
	for i := 0; i < b.N; i++ {
		iface = i // i 的作用域是括号内, 当循环结束时iface必须依然持有i的值, 所以iface = i会escape
	}
	_ = iface.(int)
	memStats()
}

func BenchmarkInterfaceConvertIntNoescape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var iface interface{} = i
		_ = iface.(int)
	}
	memStats()
}

//HeapAlloc: 73992 HeapInuse: 368640 HeapObjects: 287 HeapIdle 352256 HeapReleased 0 HeapSys 720896
//2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertInt1(b *testing.B) {
	var t = 0
	var iface interface{} = t
	for i := 0; i < b.N; i++ {
		_ = iface.(int)
	}
	memStats()
}

//HeapAlloc: 1223768 HeapInuse: 1515520 HeapObjects: 71941 HeapIdle 3334144 HeapReleased 0 HeapSys 4849664
//50000000	        35.1 ns/op	       8 B/op	       1 allocs/op
func BenchmarkInterfaceConvertInt2(b *testing.B) {
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = i
	}
	_ = iface
	memStats()
}

//HeapAlloc: 73240 HeapInuse: 376832 HeapObjects: 287 HeapIdle 376832 HeapReleased 0 HeapSys 753664
//2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertIntPoiter(b *testing.B) {
	var iface interface{}
	var a = 0
	var intPoiter = &a
	for i := 0; i < b.N; i++ {
		iface = intPoiter
	}
	_ = iface
	memStats()
}

//HeapAlloc: 76936 HeapInuse: 507904 HeapObjects: 293 HeapIdle 147456 HeapReleased 0 HeapSys 655360
//2000000000	         0.40 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertStruct32bPoiter(b *testing.B) {
	var s = &struct32b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(*struct32b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 4080072 HeapInuse: 4366336 HeapObjects: 125055 HeapIdle 1269760 HeapReleased 0 HeapSys 5636096
//30000000	        53.5 ns/op	      32 B/op	       1 allocs/op
func BenchmarkInterfaceConvertStruct32b(b *testing.B) {
	var s = struct32b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(struct32b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 78744 HeapInuse: 393216 HeapObjects: 295 HeapIdle 262144 HeapReleased 0 HeapSys 655360
//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertStruct40bPoiter(b *testing.B) {
	var s = &struct40b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(*struct40b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 2498488 HeapInuse: 2785280 HeapObjects: 50715 HeapIdle 3112960 HeapReleased 0 HeapSys 5898240
//30000000	        64.8 ns/op	      48 B/op	       1 allocs/op
func BenchmarkInterfaceConvertStruct40b(b *testing.B) {
	var s = struct40b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(struct40b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 2553864 HeapInuse: 2834432 HeapObjects: 155016 HeapIdle 3063808 HeapReleased 0 HeapSys 5898240
//30000000	        52.8 ns/op	      16 B/op	       1 allocs/op
func BenchmarkInterfaceConvertStruct16b(b *testing.B) {
	var s = struct16b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(struct16b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 2169768 HeapInuse: 2449408 HeapObjects: 65647 HeapIdle 3448832 HeapReleased 0 HeapSys 5898240
//30000000	        49.5 ns/op	      32 B/op	       1 allocs/op
func BenchmarkInterfaceConvertStruct24b(b *testing.B) {
	var s = struct24b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(struct24b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 1576536 HeapInuse: 1867776 HeapObjects: 31505 HeapIdle 7176192 HeapReleased 0 HeapSys 9043968
//30000000	        60.0 ns/op	      48 B/op	       1 allocs/op
func BenchmarkInterfaceConvertStruct48b(b *testing.B) {
	var s = struct48b{}
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = s
		_ = iface.(struct48b)
	}
	_ = iface
	memStats()
}

//HeapAlloc: 78024 HeapInuse: 442368 HeapObjects: 295 HeapIdle 180224 HeapReleased 0 HeapSys 622592
//100000000	        13.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertToInterface(b *testing.B) {
	var s interface{} = &struct32b{}
	for i := 0; i < b.N; i++ {
		_ = s.(struct32ber)
	}
	memStats()
}

//100000000	        13.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertToInterface1(b *testing.B) {
	var s1 struct32ber1 = &struct32b{}
	var s struct32ber
	for i := 0; i < b.N; i++ {
		s = s1
	}
	_ = s
	memStats()
}

//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertToOriginConvertToInterface(b *testing.B) {
	var s1 struct32ber1 = &struct32b{}
	var s struct32ber
	for i := 0; i < b.N; i++ {
		s = s1.(*struct32b)
	}
	_ = s
	memStats()
}

//100000000	        13.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDifferenceInterfaceCompare(b *testing.B) {
	var s1 struct32ber1 = &struct32b{}
	var s struct32ber
	for i := 0; i < b.N; i++ {
		if s1 == s {
		}
	}
	_ = s
	memStats()
}

//BenchmarkSameInterfaceCompare-4   	2000000000	         0.76 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSameInterfaceDifferenceTypeCompare(b *testing.B) {
	var s1 struct32ber1 = &struct32b{}
	var s struct32ber1
	for i := 0; i < b.N; i++ {
		if s1 == s {
		}
	}
	_ = s
	memStats()
}

//BenchmarkSameInterfaceCompare1-4   	200000000	         6.80 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSameInterfaceNilTypeEqualCompare(b *testing.B) {
	var s1 struct32ber1 = (*struct32b)(nil)
	var s struct32ber1 = (*struct32b)(nil)
	for i := 0; i < b.N; i++ {
		if s1 == s {
		}
	}
	_ = s
	memStats()
}

//BenchmarkSameInterfaceCompare2-4   	2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSameInterfaceNilTypeCompare(b *testing.B) {
	var s struct32ber1
	for i := 0; i < b.N; i++ {
		if s == nil {
		}
	}
	_ = s
	memStats()
}

//BenchmarkSameInterfaceNotNilTypeEqualCompare-4   	200000000	         6.74 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSameInterfaceNotNilTypeEqualCompare(b *testing.B) {
	var s1 struct32ber1 = &struct32b{}
	var s struct32ber1 = &struct32b{}
	for i := 0; i < b.N; i++ {
		if s == s1 {
		}
	}
	_ = s
	memStats()
}

//BenchmarkSameInterfaceNotNilTypeEqualCompare1-4   	2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSameInterfaceNotNilTypeAssertEqualCompare1(b *testing.B) {
	var s1 struct32ber1 = &struct32b{}
	var s struct32ber1 = &struct32b{}
	for i := 0; i < b.N; i++ {
		if s.(*struct32b) == s1.(*struct32b) {
		}
	}
	_ = s
	memStats()
}

//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterfaceConvertToOriginEqual(b *testing.B) {
	var s1 struct32ber1 = (*struct32b)(nil)
	var s struct32ber = (*struct32b)(nil)
	for i := 0; i < b.N; i++ {
		if s1.(*struct32b) == s.(*struct32b) {

		}
	}
	_ = s
	memStats()
}

//为什么是一半heapObjects?
//HeapAlloc: 400074136 HeapInuse: 400441344 HeapObjects: 25000289 HeapIdle 835584 HeapReleased 0 HeapSys 401276928
//50000000	        26.8 ns/op	       8 B/op	       1 allocs/op
func BenchmarkAssignIntToInterfaceNOGC(b *testing.B) {
	debug.SetGCPercent(-1)
	var iface interface{}
	for i := 0; i < b.N; i++ {
		iface = i
	}
	_ = iface
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 73896 HeapInuse: 385024 HeapObjects: 286 HeapIdle 368640 HeapReleased 0 HeapSys 753664
//1000000000	         2.31 ns/op	       0 B/op	       0 allocs/op
func BenchmarkAssignIntPoiterToInterfaceNOGC(b *testing.B) {
	debug.SetGCPercent(-1)
	var iface interface{} = new(int)
	for i := 0; i < b.N; i++ {
		iface = &i
	}
	_ = iface
	memStats()
	debug.SetGCPercent(100)
}

// go test -benchmem -run=^$ -bench=AssertInt
// BenchmarkAssertInt-4   	2000000000	         0.43 ns/op	       0 B/op	       0 allocs/op
// go test -benchmem -gcflags '-l' -run=^$ -bench=AssertInt
// BenchmarkAssertInt-4    500000000                3.47 ns/op            0 B/op          0 allocs/op
func AssertInt(a interface{}) int {
	return a.(int)
}
func BenchmarkAssertInt(b *testing.B) {
	var a interface{} = 1
	for i := 0; i < b.N; i++ {
		_ = AssertInt(a)
	}
}

// go test -benchmem -run=^$ -bench=AssertIntPointer
// BenchmarkAssertIntPointer-4   	2000000000	         0.41 ns/op	       0 B/op	       0 allocs/op
// go test -benchmem -gcflags '-l' -run=^$ -bench=AssertIntPointer
// BenchmarkAssertIntPointer-4     500000000                3.85 ns/op            0 B/op          0 allocs/op
func AssertIntPointer(a interface{}) *int {
	return a.(*int)
}
func BenchmarkAssertIntPointer(b *testing.B) {
	var a interface{} = new(int)
	for i := 0; i < b.N; i++ {
		_ = AssertIntPointer(a)
	}
}

// go test -benchmem -run=^$ -bench=AssertStruct32
// BenchmarkAssertStruct32b-4      2000000000               0.38 ns/op            0 B/op          0 allocs/op
// go test -benchmem -gcflags '-l' -run=^$ -bench=AssertStruct32
// BenchmarkAssertStruct32b-4      300000000                4.32 ns/op            0 B/op          0 allocs/op
func AssertStruct32b(a interface{}) struct32b {
	return a.(struct32b)
}
func BenchmarkAssertStruct32b(b *testing.B) {
	var a interface{} = struct32b{}
	for i := 0; i < b.N; i++ {
		_ = AssertStruct32b(a)
	}
}

// go test -benchmem -run=^$ -bench=AssertStruct40
// BenchmarkAssertStruct40b-4      200000000                7.84 ns/op            0 B/op          0 allocs/op
// go test -benchmem -gcflags '-l' -run=^$ -bench=AssertStruct40
// BenchmarkAssertStruct40b-4      100000000               10.4 ns/op             0 B/op          0 allocs/op
func AssertStruct40b(a interface{}) struct40b {
	return a.(struct40b)
}
func BenchmarkAssertStruct40b(b *testing.B) {
	var a interface{} = struct40b{}
	for i := 0; i < b.N; i++ {
		_ = AssertStruct40b(a)
	}
}

// go test -benchmem -run=^$ -bench=AssertStruct40bPointer
// BenchmarkAssertStruct40bPointer-4       2000000000               0.38 ns/op            0 B/op          0 allocs/op
// go test -benchmem -gcflags '-l' -run=^$ -bench=AssertStruct40bPointer
// BenchmarkAssertStruct40bPointer-4       500000000                3.87 ns/op            0 B/op          0 allocs/op
func AssertStruct40bPointer(a interface{}) *struct40b {
	return a.(*struct40b)
}
func BenchmarkAssertStruct40bPointer(b *testing.B) {
	var a interface{} = &struct40b{}
	for i := 0; i < b.N; i++ {
		_ = AssertStruct40bPointer(a)
	}
}

// 编译器会优化变量在同一个函数的的断言,并且会内联小函数,但加上-gcflags '-l'禁止内联后会调用runtime.memmove复制值,速度变慢,如果结构体较大的话更慢

/*

http://legendtkl.com/2017/07/01/golang-interface-implement/
两个iface的比较会比较慢

goos: linux
goarch: amd64
pkg: github.com/cdongyang/library/test/interfaceConvert
BenchmarkInterfaceConvertIntCompare-4           2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkDefineVariable-4                       2000000000               0.39 ns/op            0 B/op          0 allocs/op
BenchmarkDefineHeapVariable-4                   2000000000               0.39 ns/op            0 B/op          0 allocs/op
BenchmarkInterfaceConvertInt-4                  50000000                40.9 ns/op             8 B/op          1 allocs/op
BenchmarkInterfaceConvertInt1-4                 2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkInterfaceConvertInt2-4                 50000000                40.6 ns/op             8 B/op          1 allocs/op
BenchmarkInterfaceConvertIntPoiter-4            2000000000               0.40 ns/op            0 B/op          0 allocs/op
BenchmarkInterfaceConvertStruct32bPoiter-4      2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkInterfaceConvertStruct32b-4            30000000                56.1 ns/op            32 B/op          1 allocs/op
BenchmarkInterfaceConvertStruct40bPoiter-4      2000000000               0.38 ns/op            0 B/op          0 allocs/op
BenchmarkInterfaceConvertStruct40b-4            30000000                67.4 ns/op            48 B/op          1 allocs/op
BenchmarkInterfaceConvertStruct16b-4            30000000                55.1 ns/op            16 B/op          1 allocs/op
BenchmarkInterfaceConvertStruct24b-4            30000000                56.2 ns/op            32 B/op          1 allocs/op
BenchmarkInterfaceConvertStruct48b-4            20000000                74.1 ns/op            48 B/op          1 allocs/op
BenchmarkInterfaceConvertToInterface-4          100000000               12.6 ns/op             0 B/op          0 allocs/op
BenchmarkAssignIntToInterfaceNOGC-4             50000000                24.7 ns/op             8 B/op          0 allocs/op
BenchmarkAssignIntPoiterToInterfaceNOGC-4       1000000000               2.29 ns/op            0 B/op          0 allocs/op
PASS

	interface{} 有一个element为unsafe.Pointer类型,持有对象的指针,值赋值给interface时易引起堆内存分配,指针赋值时也易引起栈变量逃逸到堆
	结构体源码:/usr/local/go/src/runtime/runtime2.go
	type iface struct {
		tab  *itab
		data unsafe.Pointer
	}

	type eface struct {
		_type *_type
		data  unsafe.Pointer
	}

推论:
	由值类型赋值给interface{}会导致 1 allocs/op,比直接指针赋值给interface{}慢很多,因为不需要重新在堆上分配空间
	由interface{}断言为原本的类型被内联并在变量定义的函数内跟普通int赋值差不多,但断言为其它interface{}较慢,不被内联时调用了runtime.memmove也较慢
	每个interface{}都会持有一个unsafe.Pointer,如果赋值给interface{}的是值而不是指针就会在heap新建一份值并将指针赋给interface{}
	堆内存动态分配内存速度比直接定义变量慢很多,并且会增大heapObjects数量,直接定义的变量在编译阶段就静态分配了内存到栈空间
	从一个自定义的interface到另一个自定义interface的转化较慢,无论是类型断言还是直接赋值
	比较底层且调用较频繁的函数或方法尽可能使用unsafe.Pointer代替interface{}

总结:
	interface{}虽然能实现泛化,但是以runtime开销作为代价的,在有些效率要求高并且调用频繁的情况下尽量使用具体类型
*/
