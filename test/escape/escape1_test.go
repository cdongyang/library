package escape_test

import (
	"testing"
	"unsafe"
)

// copy from /runtime/stubs.go
// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input.  noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func noescapeInterface(iface interface{}) interface{} {
	return *(*interface{})(noescape(unsafe.Pointer(&iface)))
}

func typeAssert(iface interface{}) {
	switch v := iface.(type) {
	case int:
		_ = v
	}
}

func typeAssertNotInline(iface interface{}) {
	switch v := iface.(type) {
	case int:
		_ = v
	default:
		escape(iface)
	}
}

func escape(_ ...interface{}) {
}

func TestEscape(t *testing.T) {
	n := testing.AllocsPerRun(1000, func() {
		var iface interface{}
		for i := 0; i < 10; i++ {
			iface = i // i 的作用域是括号内, 当循环结束时iface必须依然持有i的值, 所以iface = i会escape
		}
		_ = iface.(int)
	})
	if n == 0 {
		t.Fatal(n)
	}
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 10; i++ {
			var iface interface{} = i
			_ = iface.(int)
		}
	})
	if n != 0 {
		t.Fatal(n)
	}
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 10; i++ {
			var iface interface{} = i
			switch v := iface.(type) {
			case int:
				_ = v
			}
		}
	})
	if n != 0 {
		t.Fatal(n)
	}
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 10; i++ {
			func(iface interface{}) {
				switch v := iface.(type) {
				case int:
					_ = v
				}
			}(i)
		}
	})
	if n == 0 {
		t.Fatal(n)
	}
	// 与上面一样的函数, 通过直接定义调用iface会escape
	// 通过函数调用编译器inline之后不会escape
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 10; i++ {
			typeAssert(i)
		}
	})
	if n != 0 {
		t.Fatal(n)
	}
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 10; i++ {
			go typeAssert(i)
		}
	})
	if n == 0 {
		t.Fatal(n)
	}
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 10; i++ {
			typeAssert(i)
		}
	})
	if n != 0 {
		t.Fatal(n)
	}
}
