package structMethod_test

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
)

type struct32b struct {
	a int
	b int
	c int
	d int
}

func (s *struct32b) PoiterFun(int) {

}

func (s struct32b) ValueFun(int) {

}

type struct40b struct {
	a int
	b int
	c int
	d int
	e int
}

func (s *struct40b) PoiterFun(int) {

}

func (s struct40b) ValueFun(int) {

}

type structer interface {
	PoiterFun(int)
	ValueFun(int)
}

type valueStructer interface {
	ValueFun(int)
}

type poiterStructer interface {
	PoiterFun(int)
}

var mem runtime.MemStats
var outmem = true

func memStats() {
	if !outmem {
		return
	}
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

//HeapAlloc: 71688 HeapInuse: 344064 HeapObjects: 288 HeapIdle 376832 HeapReleased 0 HeapSys 720896
//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDefineStruct32b(b *testing.B) {
	debug.SetGCPercent(-1)
	for i := 0; i < b.N; i++ {
		var a struct32b
		_ = a
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 71624 HeapInuse: 344064 HeapObjects: 287 HeapIdle 376832 HeapReleased 0 HeapSys 720896
//1000000000	         1.90 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDefineStruct40b(b *testing.B) {
	debug.SetGCPercent(-1)
	for i := 0; i < b.N; i++ {
		var a struct40b
		_ = a
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 78088 HeapInuse: 360448 HeapObjects: 295 HeapIdle 294912 HeapReleased 0 HeapSys 655360
//2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bValueUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s struct32b
	for i := 0; i < b.N; i++ {
		s.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 73160 HeapInuse: 417792 HeapObjects: 286 HeapIdle 335872 HeapReleased 0 HeapSys 753664
//500000000	         3.03 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bValueUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s struct40b
	for i := 0; i < b.N; i++ {
		s.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 73048 HeapInuse: 360448 HeapObjects: 284 HeapIdle 393216 HeapReleased 0 HeapSys 753664
//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bPoiterUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct32b{}
	for i := 0; i < b.N; i++ {
		s.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 74808 HeapInuse: 360448 HeapObjects: 288 HeapIdle 360448 HeapReleased 0 HeapSys 720896
//500000000	         3.41 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bPoiterUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct40b{}
	for i := 0; i < b.N; i++ {
		s.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 73320 HeapInuse: 483328 HeapObjects: 288 HeapIdle 237568 HeapReleased 0 HeapSys 720896
//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bValueUsePoiterFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s struct32b
	for i := 0; i < b.N; i++ {
		s.PoiterFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 74712 HeapInuse: 360448 HeapObjects: 287 HeapIdle 360448 HeapReleased 0 HeapSys 720896
//2000000000	         0.38 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bValueUsePoiterFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s struct40b
	for i := 0; i < b.N; i++ {
		s.PoiterFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 72840 HeapInuse: 360448 HeapObjects: 283 HeapIdle 393216 HeapReleased 0 HeapSys 753664
//2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bPoiterUsePoiterFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct32b{}
	for i := 0; i < b.N; i++ {
		s.PoiterFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 76264 HeapInuse: 360448 HeapObjects: 290 HeapIdle 327680 HeapReleased 0 HeapSys 688128
//2000000000	         0.44 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bPoiterUsePoiterFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct40b{}
	for i := 0; i < b.N; i++ {
		s.PoiterFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 73656 HeapInuse: 442368 HeapObjects: 285 HeapIdle 311296 HeapReleased 0 HeapSys 753664
//2000000000	         0.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bPoiterInterfaceUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct32b{}
	var iface structer = s
	for i := 0; i < b.N; i++ {
		iface.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 76968 HeapInuse: 401408 HeapObjects: 290 HeapIdle 286720 HeapReleased 0 HeapSys 688128
//500000000	         3.04 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bPoiterInterfaceUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct40b{}
	var iface structer = s
	for i := 0; i < b.N; i++ {
		iface.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 74168 HeapInuse: 409600 HeapObjects: 287 HeapIdle 344064 HeapReleased 0 HeapSys 753664
//500000000	         3.12 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bPoiterInterfaceUsePoiterFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct32b{}
	var iface structer = s
	for i := 0; i < b.N; i++ {
		iface.PoiterFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 76536 HeapInuse: 368640 HeapObjects: 294 HeapIdle 352256 HeapReleased 0 HeapSys 720896
//BenchmarkStruct32bPoiterInterfaceUsePoiterCacheFun-4   	200000000	         5.83 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bPoiterInterfaceUsePoiterCacheFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct32b{}
	var iface structer = s
	var fun = iface.PoiterFun
	for i := 0; i < b.N; i++ {
		fun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 76248 HeapInuse: 352256 HeapObjects: 290 HeapIdle 368640 HeapReleased 0 HeapSys 720896
//500000000	         3.13 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bPoiterInterfaceUsePoiterFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s = &struct40b{}
	var iface structer = s
	for i := 0; i < b.N; i++ {
		iface.PoiterFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 78664 HeapInuse: 393216 HeapObjects: 294 HeapIdle 229376 HeapReleased 0 HeapSys 622592
//300000000	         3.93 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct32bValueInterfaceUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s struct32b
	var iface valueStructer = s
	for i := 0; i < b.N; i++ {
		iface.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

//HeapAlloc: 73528 HeapInuse: 385024 HeapObjects: 283 HeapIdle 368640 HeapReleased 0 HeapSys 753664
//300000000	         5.84 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct40bValueInterfaceUseValueFun(b *testing.B) {
	debug.SetGCPercent(-1)
	var s struct40b
	var iface valueStructer = s
	for i := 0; i < b.N; i++ {
		iface.ValueFun(0)
	}
	memStats()
	debug.SetGCPercent(100)
}

/*
	32byte作为small struct 和 large struct 的分界,不超过32byte的结构体会在cache缓存,参考自
	https://github.com/qyuhen/book/blob/master/gopher2015/qyuhen.pdf	第10页
	/usr/local/go/src/runtime/malloc.go runtime.mallocgc源码:
	// Allocate an object of size bytes.
	// Small objects are allocated from the per-P cache's free lists.
	// Large objects (> 32 kB) are allocated straight from the heap.
	func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer

	goos: linux
	goarch: amd64
	pkg: github.com/cdongyang/library/test/structMethod
	BenchmarkDefineStruct32b-4                              2000000000               0.38 ns/op            0 B/op          0 allocs/op
	BenchmarkDefineStruct40b-4                              1000000000               1.89 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bValueUseValueFun-4                    2000000000               0.38 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bValueUseValueFun-4                    500000000                3.12 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bPoiterUseValueFun-4                   2000000000               0.38 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bPoiterUseValueFun-4                   500000000                3.52 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bValueUsePoiterFun-4                   2000000000               0.39 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bValueUsePoiterFun-4                   2000000000               0.38 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bPoiterUsePoiterFun-4                  2000000000               0.38 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bPoiterUsePoiterFun-4                  2000000000               0.38 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bPoiterInterfaceUseValueFun-4          500000000                3.91 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bPoiterInterfaceUseValueFun-4          300000000                5.94 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bPoiterInterfaceUsePoiterFun-4         500000000                3.04 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bPoiterInterfaceUsePoiterFun-4         500000000                3.02 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct32bValueInterfaceUseValueFun-4           500000000                3.86 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct40bValueInterfaceUseValueFun-4           300000000                6.13 ns/op            0 B/op          0 allocs/op
	PASS

推论:
	整个过程没有对象的分配,即方法的调用不会创建新对象
	结构体调用指针方法或结构体指针调用方法都是很快的
	对值方法进行调用时会进行值复制,但large struct没有cache,值复制会稍微有点慢
	interface{}对方法的调用和large struct复制是差不多速度的
*/
