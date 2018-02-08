package runtimecp

import (
	"unsafe"
)

func UnsafeAlloc(size uintptr) unsafe.Pointer {
	bs := make([]byte, size)
	return (*Slice)(unsafe.Pointer(&bs)).Array
}

func EfaceOf(eface interface{}) Eface {
	return *(*Eface)(unsafe.Pointer(&eface))
}

func Eface2Interface(eface Eface) interface{} {
	return *(*interface{})(unsafe.Pointer(&eface))
}

func TypeOf(eface interface{}) *Type {
	return (*Eface)(Noescape(unsafe.Pointer(&eface))).Type
}

// IsDirectIface reports whether t is stored directly in an interface value.
func IsDirectIface(t *Type) bool {
	return t.Kind&KindDirectIface != 0
}

// copy from package runtime
// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input.  noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//go:nosplit
func Noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func NoescapeInterface(x interface{}) interface{} {
	return *(*interface{})(Noescape(unsafe.Pointer(&x)))
}

func add(x unsafe.Pointer, n uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(x) + n)
}

func GetCompare(a, b unsafe.Pointer, size uintptr) func(a, b unsafe.Pointer) int {
	lowbit := size & (-size)
	switch lowbit {
	case 1:
		return Compare8
	case 2:
		return Compare16
	case 4:
		return Compare32
	case 8:
		return Compare64
	case 16:
		return Compare128
	default:
		if lowbit > 16 {
			return Compare128
		}
	}
	return nil
}

func Compare8(a, b unsafe.Pointer) int {
	aa, bb := *(*int8)(a), *(*int8)(b)
	if aa > bb {
		return 1
	} else if aa < bb {
		return -1
	}
	return 0
}

func Compare16(a, b unsafe.Pointer) int {
	aa, bb := *(*int16)(a), *(*int16)(b)
	if aa > bb {
		return 1
	} else if aa < bb {
		return -1
	}
	return 0
}

func Compare32(a, b unsafe.Pointer) int {
	aa, bb := *(*int32)(a), *(*int32)(b)
	if aa > bb {
		return 1
	} else if aa < bb {
		return -1
	}
	return 0
}

func Compare64(a, b unsafe.Pointer) int {
	aa, bb := *(*int64)(a), *(*int64)(b)
	if aa > bb {
		return 1
	} else if aa < bb {
		return -1
	}
	return 0
}

func Compare128(a, b unsafe.Pointer) int {
	aa, bb := *(*[2]int64)(a), *(*[2]int64)(b)
	if aa[0] > bb[0] || aa[0] == bb[0] && aa[1] > bb[1] {
		return 1
	} else if aa[0] < bb[0] || aa[0] == bb[0] && aa[1] < bb[1] {
		return -1
	}
	return 0
}
