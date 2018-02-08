package struct_test

import (
	"reflect"
	"testing"
	"unsafe"
)

type s1 struct {
	a bool
	b bool
	c bool
}

type s2 struct {
	a int8
	b int8
	c int8
}

type s3 struct {
	a int16
	b int16
	c int16
}

type s4 struct {
	a int32
	b int32
	c int32
}

type s5 struct {
	a int64
	b int64
	c int64
}

type s6 struct {
	a complex128
	b complex128
	c complex128
}

type s7 struct {
	a complex128
	b bool
	c int64
}

type s8 struct {
	a complex128
	b bool
	c int32
}

type s9 struct {
	a bool
	b int32
	c bool
}

//1___4
//1________
//8
type s10 struct {
	a bool
	b int32
	c bool
	d int64
}

//11__
//4
type s11 struct {
	a bool
	b bool
	c int32
}

type s12 struct {
	a bool
	b bool
	c int16
	d int32
}

type s13 struct {
	a s2
	b int8
}

type s14 struct {
	a s3
	b int16
}

type s15 struct {
	a s4
	b int32 //4
}

type s16 struct {
	a struct {
		a int64 // 8
		b int64 // 8
		c int32 // 4_
	}
	b int32 // 4_
}

func checkSize(t *testing.T, name string, want, real, rf uintptr) {
	if want != real || real != rf {
		t.Errorf("%s: want size: %d, real size: %d,reflect size:%d\n", name, want, real, rf)
	}
}

func TestStructSize(t *testing.T) {
	checkSize(t, "s1", 3, unsafe.Sizeof(s1{}), reflect.TypeOf(s1{}).Size())
	checkSize(t, "s2", 3, unsafe.Sizeof(s2{}), reflect.TypeOf(s2{}).Size())
	checkSize(t, "s3", 6, unsafe.Sizeof(s3{}), reflect.TypeOf(s3{}).Size())
	checkSize(t, "s4", 12, unsafe.Sizeof(s4{}), reflect.TypeOf(s4{}).Size())
	checkSize(t, "s5", 24, unsafe.Sizeof(s5{}), reflect.TypeOf(s5{}).Size())
	checkSize(t, "s6", 48, unsafe.Sizeof(s6{}), reflect.TypeOf(s6{}).Size())
	checkSize(t, "s7", 32, unsafe.Sizeof(s7{}), reflect.TypeOf(s7{}).Size())
	checkSize(t, "s8", 24, unsafe.Sizeof(s8{}), reflect.TypeOf(s8{}).Size())
	checkSize(t, "s9", 12, unsafe.Sizeof(s9{}), reflect.TypeOf(s9{}).Size())
	checkSize(t, "s10", 24, unsafe.Sizeof(s10{}), reflect.TypeOf(s10{}).Size())
	checkSize(t, "s11", 8, unsafe.Sizeof(s11{}), reflect.TypeOf(s11{}).Size())
	checkSize(t, "s12", 8, unsafe.Sizeof(s12{}), reflect.TypeOf(s12{}).Size())
	checkSize(t, "s13", 4, unsafe.Sizeof(s13{}), reflect.TypeOf(s13{}).Size())
	checkSize(t, "s14", 8, unsafe.Sizeof(s14{}), reflect.TypeOf(s14{}).Size())
	checkSize(t, "s15", 16, unsafe.Sizeof(s15{}), reflect.TypeOf(s15{}).Size())
	checkSize(t, "s16", 32, unsafe.Sizeof(s16{}), reflect.TypeOf(s16{}).Size())
}

func checkOffset(t *testing.T, name string, want, real, rf uintptr) {
	if want != real || real != rf {
		t.Errorf("%s: want offset: %d, real offset: %d,reflect offset:%d\n", name, want, real, rf)
	}
}

func getReflectOffset(s interface{}, name string) uintptr {
	sf, ok := reflect.ValueOf(s).Type().FieldByName(name)
	if !ok {
		panic("not exist " + name)
	}
	return sf.Offset
}

func TestUnsafe(t *testing.T) {
	checkOffset(t, "s7.b", 16, unsafe.Offsetof(s7{}.b), getReflectOffset(s7{}, "b"))
	checkOffset(t, "s7.c", 24, unsafe.Offsetof(s7{}.c), getReflectOffset(s7{}, "c"))
}
