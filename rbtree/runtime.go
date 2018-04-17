package rbtree

import (
	"unsafe"
)

// 将value赋值给interface时编译器取这个value的runtime type和在堆上分配的value对象的指针绑定成一个iface/eface struct
// 将pointer赋值给interface时编译器取这个pointer的指向的值的runtime type并和这个pointer绑定成一个iface/eface struct
type eface struct {
	_type *_type
	p     unsafe.Pointer
}

type flag uintptr

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      uint8
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *unsafe.Pointer
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       int32
	ptrToThis int32
}

func unpackIface(x interface{}) eface {
	return *(*eface)(unsafe.Pointer(&x))
}

func packIface(x eface) interface{} {
	return *(*interface{})(unsafe.Pointer(&x))
}

func pack2Iface(typ *_type, p unsafe.Pointer) interface{} {
	return *(*interface{})(unsafe.Pointer(&eface{_type: typ, p: p}))
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

const (
	kindBool = 1 + iota
	kindInt
	kindInt8
	kindInt16
	kindInt32
	kindInt64
	kindUint
	kindUint8
	kindUint16
	kindUint32
	kindUint64
	kindUintptr
	kindFloat32
	kindFloat64
	kindComplex64
	kindComplex128
	kindArray
	kindChan
	kindFunc
	kindInterface
	kindMap
	kindPtr
	kindSlice
	kindString
	kindStruct
	kindUnsafePointer

	kindDirectIface = 1 << 5
	kindGCProg      = 1 << 6
	kindNoPointers  = 1 << 7
	kindMask        = (1 << 5) - 1
)

// isDirectIface reports whether t is stored directly in an interface value.
func isDirectIface(t *_type) bool {
	return t.kind&kindDirectIface != 0
}

func isNoPtrIface(t *_type) bool {
	return t.kind&kindNoPointers != 0
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func newmem(size uintptr) unsafe.Pointer {
	bytes := make([]byte, size)
	return unsafe.Pointer(&bytes[0])
}

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

func noescapeInterface(x interface{}) interface{} {
	return *(*interface{})(noescape(unsafe.Pointer(&x)))
}

// you can use this func to make the arguments of insert method don't escape to heap
// if it's not direct interface{} (make sure this)!
// you can find a example in file example_test.go.
// USE CAREFULLY!
func NoescapeInterface(x interface{}) interface{} {
	return *(*interface{})(noescape(unsafe.Pointer(&x)))
}
