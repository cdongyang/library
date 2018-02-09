package gc_test

import (
	"bytes"
	"testing"
	"time"
	"unsafe"

	"github.com/cdongyang/library/test"
)

func TestSliceGC(t *testing.T) {
	var s = bytes.Repeat([]byte("1"), 1e6)
	test.MemStats()
	var b = s[1e5 : 2*1e5 : 2*1e5]
	test.MemStats()
	for _, val := range b {
		if val != '1' {
			t.Fatal()
		}
	}
	s = nil
	time.Sleep(time.Second)
	test.MemStats()
	for _, val := range b {
		if val != '1' {
			t.Fatal()
		}
	}
	b = nil
	time.Sleep(time.Second)
	test.MemStats()
	time.Sleep(time.Second)
	test.MemStats()
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func TestSliceGCPointer(t *testing.T) {
	test.MemStats("test 1")
	t.Run("1", func(t *testing.T) {
		var s = bytes.Repeat([]byte("1"), 1e6)
		test.MemStats()
		var b = *(*unsafe.Pointer)(unsafe.Pointer(&s))
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != '1' {
				t.Fatal()
			}
		}
		s = nil
		time.Sleep(time.Second)
		test.MemStats("free s")
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != '1' {
				t.Fatal()
			}
		}
		b = nil
		time.Sleep(time.Second)
		test.MemStats("free b")
		time.Sleep(time.Second)
		test.MemStats()
	})
	test.MemStats("test 2")
	/* b持有s的非直接指针, s被gc处理不受b影响, s被gc后b对内存的访问会core
	t.Run("2", func(t *testing.T) {
		var s = bytes.Repeat([]byte("1"), 1e6)
		test.MemStats()
		var b = *(*unsafe.Pointer)(add(unsafe.Pointer(&s), 1e5))
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != '1' {
				t.Fatal()
			}
		}
		s = nil
		time.Sleep(time.Second)
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != '1' {
				t.Fatal()
			}
		}
		b = nil
		time.Sleep(time.Second)
		test.MemStats()
		time.Sleep(time.Second)
		test.MemStats()
	})
	*/
	test.MemStats("test 4")
	t.Run("4", func(t *testing.T) {
		var s = bytes.Repeat([]byte("1"), 1e6)
		test.MemStats()
		var sp = *(*unsafe.Pointer)(unsafe.Pointer(&s))
		var b = getGCPointer(add(sp, 1e5), 1e5)
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != '1' {
				t.Fatal()
			}
		}
		s = nil
		time.Sleep(time.Second)
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != '1' {
				t.Fatal()
			}
		}
		time.Sleep(time.Second)
		test.MemStats()
		var tmp = *(*[]byte)(unsafe.Pointer(&struct {
			p   unsafe.Pointer
			len int
			cap int
		}{b, 1e5, 1e5}))
		b = nil
		tmp = nil
		_ = tmp
		time.Sleep(time.Second)
		test.MemStats()
		time.Sleep(time.Second)
		test.MemStats()
		s = bytes.Repeat([]byte("6"), 1e7)
		time.Sleep(time.Second)
		test.MemStats()
		s = nil
		time.Sleep(time.Second)
		test.MemStats()
	})
	test.MemStats("test 3")
	t.Run("3", func(t *testing.T) { // memory leak, s can't be free
		var cmpb byte = '3'
		var s = bytes.Repeat([]byte{cmpb}, 1e6)
		test.MemStats()
		var sp = *(*unsafe.Pointer)(unsafe.Pointer(&s))
		var b = getGCPointer(add(sp, 1e5), 1e5)
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != cmpb {
				t.Fatal()
			}
		}
		s = nil
		time.Sleep(time.Second)
		test.MemStats()
		for i := uintptr(0); i < 1e5; i++ {
			if *(*byte)(add(b, i)) != cmpb {
				t.Fatal()
			}
		}
		b = nil
		time.Sleep(time.Second)
		test.MemStats()
		time.Sleep(time.Second)
		test.MemStats()
		time.Sleep(time.Second)
		test.MemStats()
	})
	test.MemStats("end")
}

func getGCPointer(p unsafe.Pointer, size int) unsafe.Pointer {
	var tmp = *(*[]byte)(unsafe.Pointer(&struct {
		p   unsafe.Pointer
		len int
		cap int
	}{p, size, size}))
	return *(*unsafe.Pointer)(unsafe.Pointer(&tmp))
}
