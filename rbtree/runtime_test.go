package rbtree

import (
	"testing"
	"time"
	"unsafe"

	"github.com/cdongyang/library/test"
)

func TestGetGCPointer(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		test.MemStats("init")
		var arr = make([]byte, 1e6)
		test.MemStats("make array")
		var p = *(*unsafe.Pointer)(unsafe.Pointer(&arr))
		arr = nil
		time.Sleep(time.Second)
		test.MemStats("free arr")
		var pp = getGCPointer(p, 1e5)
		p = nil
		time.Sleep(time.Second)
		test.MemStats("free p")
		pp = nil
		_ = pp
		time.Sleep(time.Second)
		test.MemStats("free pp")
		time.Sleep(time.Second)
		test.MemStats("Repeat free pp")
	})
	test.MemStats("end")
}
