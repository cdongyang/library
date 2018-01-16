package typeAssert_test

import (
	"testing"
	"unsafe"
)

func BenchmarkUnsafePointer(b *testing.B) {
	var a interface{} = 1
	for i := 0; i < b.N; i++ {
		_ = unsafe.Pointer(&a)
	}
}

func BenchmarkUnsafePointerAssert(b *testing.B) {
	var a interface{} = 1
	for i := 0; i < b.N; i++ {
		_ = *(*[2]uintptr)(unsafe.Pointer(&a))
	}
}

func BenchmarkUnsafePointerAssertCompare(b *testing.B) {
	var a, c interface{} = 1, 2
	for i := 0; i < b.N; i++ {
		ap := *(*[2]uintptr)(unsafe.Pointer(&a))
		cp := *(*[2]uintptr)(unsafe.Pointer(&c))
		_ = (ap[0] == cp[0])
	}
}
