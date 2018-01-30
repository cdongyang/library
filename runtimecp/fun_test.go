package runtimecp_test

import (
	"testing"
	"unsafe"

	. "github.com/cdongyang/library/runtimecp"
)

func read(x interface{}) {
	if *(*int)(EfaceOf(x).Data) != 1 {
		panic("not 1")
	}
	if x.(int) != 1 {
		panic("not 1")
	}
}

func readP(p unsafe.Pointer) {
	if *(*int)(p) != 1 {
		panic("not 1")
	}
}

func readE(e Eface) {
	if *(*int)(e.Data) != 1 {
		panic("not 1")
	}
}

func TestNoescape(t *testing.T) {
	var n float64
	var f func(interface{})
	var fp func(unsafe.Pointer)
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		read(a) // read inline, read only, no escape
	})
	if n != 0 {
		t.Fatal("should not escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		readP(unsafe.Pointer(&a)) // read inline, read only, no escape
	})
	if n != 0 {
		t.Fatal("should not escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		f = read
		f(a) // f not inline, a escape
	})
	if n != 1 {
		t.Fatal("should escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		fp = readP
		fp(unsafe.Pointer(&a)) // f not inline, a escape
	})
	if n != 1 {
		t.Fatal("should escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		f = read
		f(NoescapeInterface(a)) // use Noescape func to hide pointer from escape analysis
	})
	if n != 0 {
		t.Fatal("should not escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		fp = readP
		fp(Noescape(unsafe.Pointer(&a))) // use Noescape func to hide pointer from escape analysis
	})
	if n != 0 {
		t.Fatal("should not escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		var b interface{} = a
		_ = b
	})
	if n != 0 {
		t.Fatal("should not escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		func(p unsafe.Pointer) {
			func(b unsafe.Pointer) {
				for i := 0; i < 1000; i++ {
					if *(*int)(b) != 1 {
						t.Fatal("not 1")
					}
				}
			}(p)
		}(unsafe.Pointer(&a))
	})
	if n == 0 {
		t.Fatal("should escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		func(p unsafe.Pointer) {
			func(b unsafe.Pointer) {
				for i := 0; i < 1000; i++ {
					if *(*int)(b) != 1 {
						t.Fatal("not 1")
					}
				}
			}(p)
		}(Noescape(unsafe.Pointer(&a)))
	})
	if n != 0 {
		t.Fatal("should not escape", n)
	}
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		func(p unsafe.Pointer) {
			go func(b unsafe.Pointer) {
				for i := 0; i < 1000; i++ {
					if *(*int)(b) != 1 {
						t.Fatal("not 1")
					}
				}
			}(p)
		}(Noescape(unsafe.Pointer(&a)))
	})
	if n == 0 {
		t.Fatal("should escape", n)
	}
}

type struct32ber interface {
	fun1()
}

type struct32b struct {
	a, b, c, d int
}

func (s struct32b) fun1() {
}

func TestInterface(t *testing.T) {
	var a = 2
	var ap = &a
	eface := EfaceOf(a)
	pEface := EfaceOf(ap)
	ppEface := EfaceOf(&ap)
	if (*int)(eface.Data) == ap || *(*int)(eface.Data) != 2 || eface.Type.Kind != KindInt+KindNoPointers {
		t.Fatal(eface.Type)
	}
	if (*int)(pEface.Data) != ap || *(*int)(pEface.Data) != 2 || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if *(**int)(ppEface.Data) != ap || **(**int)(ppEface.Data) != 2 || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}
	var b = "b"
	var bp = &b
	eface = EfaceOf(b)
	pEface = EfaceOf(bp)
	ppEface = EfaceOf(&bp)
	if *(*string)(eface.Data) != "b" || eface.Type.Kind != KindString {
		t.Fatal(eface.Type)
	}
	if *(*string)(pEface.Data) != "b" || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**string)(ppEface.Data) != "b" || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var c interface{} = 2
	var cp = &c
	eface = EfaceOf(c)
	pEface = EfaceOf(cp)
	ppEface = EfaceOf(&cp)
	if *(*int)(eface.Data) != 2 || eface.Type.Kind != KindInt+KindNoPointers {
		t.Fatal(eface.Type)
	}
	if *(*interface{})(pEface.Data) != 2 || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**interface{})(ppEface.Data) != 2 || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var s32 = &struct32b{1, 2, 3, 4}
	var d struct32ber = s32
	var dp = &d
	eface = EfaceOf(d) // ConvI2E, eface.Data is s32
	pEface = EfaceOf(dp)
	ppEface = EfaceOf(&dp)
	if (*struct32b)(eface.Data) != s32 || eface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(eface.Type)
	}
	if *(*struct32ber)(pEface.Data) != d || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**struct32ber)(ppEface.Data) != d || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var e = *s32
	var ep = &e
	eface = EfaceOf(e)
	pEface = EfaceOf(ep)
	ppEface = EfaceOf(&ep)
	if *(*struct32b)(eface.Data) != e || eface.Type.Kind != KindStruct+KindNoPointers {
		t.Fatal(eface.Type)
	}
	if *(*struct32b)(pEface.Data) != e || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**struct32b)(ppEface.Data) != e || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var f = []int{1, 2, 3}
	var fp = &f
	eface = EfaceOf(f)
	pEface = EfaceOf(fp)
	ppEface = EfaceOf(&fp)
	if !equalIntSlice(*(*[]int)(eface.Data), f) || eface.Type.Kind != KindSlice {
		t.Fatal(eface.Type)
	}
	if !equalIntSlice(*(*[]int)(pEface.Data), f) || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if !equalIntSlice(**(**[]int)(ppEface.Data), f) || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var g = [3]int{1, 2, 3}
	var gp = &g
	eface = EfaceOf(g)
	pEface = EfaceOf(gp)
	ppEface = EfaceOf(&gp)
	if *(*[3]int)(eface.Data) != g || eface.Type.Kind != KindArray+KindNoPointers {
		t.Fatal(eface.Type)
	}
	if *(*[3]int)(pEface.Data) != g || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**[3]int)(ppEface.Data) != g || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var h = make(chan bool)
	var hp = &h
	eface = EfaceOf(h)
	pEface = EfaceOf(hp)
	ppEface = EfaceOf(&hp)
	if *(*chan bool)(unsafe.Pointer(&eface.Data)) != h || eface.Type.Kind != KindChan+KindDirectIface {
		t.Fatal(eface.Type)
	}
	if *(*chan bool)(pEface.Data) != h || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**chan bool)(ppEface.Data) != h || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var i = map[int]bool{1: true, 2: true, 4: true}
	var ip = &i
	eface = EfaceOf(i)
	pEface = EfaceOf(ip)
	ppEface = EfaceOf(&ip)
	if !equalMap(*(*map[int]bool)(unsafe.Pointer(&eface.Data)), i) || eface.Type.Kind != KindMap+KindDirectIface {
		t.Fatal(eface.Type)
	}
	if !equalMap(*(*map[int]bool)(pEface.Data), i) || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if !equalMap(**(**map[int]bool)(ppEface.Data), i) || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var j = func() {}
	var jp = &j
	eface = EfaceOf(j)
	pEface = EfaceOf(jp)
	ppEface = EfaceOf(&jp)
	if eface.Type.Kind != KindFunc+KindDirectIface {
		t.Fatal(eface.Type)
	}
	if pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var k = unsafe.Pointer(&j)
	var kp = &k
	eface = EfaceOf(k)
	pEface = EfaceOf(kp)
	ppEface = EfaceOf(&kp)
	if eface.Data != k || eface.Type.Kind != KindUnsafePointer+KindDirectIface {
		t.Fatal(eface.Type)
	}
	if *(*unsafe.Pointer)(pEface.Data) != k || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**unsafe.Pointer)(ppEface.Data) != k || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

	var l = complex(1.0, 2.0)
	var lp = &l
	eface = EfaceOf(l)
	pEface = EfaceOf(lp)
	ppEface = EfaceOf(&lp)
	if *(*complex128)(eface.Data) != l || eface.Type.Kind != KindComplex128+KindNoPointers {
		t.Fatal(eface.Type)
	}
	if *(*complex128)(pEface.Data) != l || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(pEface.Type)
	}
	if **(**complex128)(ppEface.Data) != l || pEface.Type.Kind != KindPtr+KindDirectIface {
		t.Fatal(ppEface.Type)
	}

}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalMap(a, b map[int]bool) bool {
	for k, v := range a {
		if val, ok := b[k]; !ok || val != v {
			return false
		}
	}
	return true
}
