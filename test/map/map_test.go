package map_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"unsafe"

	"github.com/cdongyang/library/test"
)

func TestMapHeapSize1E6(t *testing.T) {
	testBufferMapHeapSize(1e6)
	testMapHeapSize(1e6)
}
func TestMapHeapSize1E5(t *testing.T) {
	testBufferMapHeapSize(1e5)
	testMapHeapSize(1e5)
}

func TestMapKeyValue(t *testing.T) {
	type struct32B struct {
		a [4]int
	}
	type struct40B struct {
		a [5]int
	}
	type struct64B struct {
		a [8]int
	}
	type struct128B struct {
		a [16]int
	}
	type struct256B struct {
		a [32]int
	}
	t.Run("key 32b", func(t *testing.T) {
		var mp = make(map[struct32B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct32B{[4]int{i}}] = i
		}
		test.MemStats()
	})
	t.Run("key 40b", func(t *testing.T) {
		var mp = make(map[struct40B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct40B{[5]int{i}}] = i
		}
		test.MemStats()
	})
	t.Run("key 64b", func(t *testing.T) {
		var mp = make(map[struct64B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct64B{[8]int{i}}] = i
		}
		test.MemStats()
	})
	t.Run("key 128b", func(t *testing.T) {
		var mp = make(map[struct128B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct128B{[16]int{i}}] = i
		}
		test.MemStats()
	})
	t.Run("key 256b", func(t *testing.T) {
		var mp = make(map[struct256B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct256B{[32]int{i}}] = i
		}
		test.MemStats()
	})
	t.Run("val 32b", func(t *testing.T) {
		var mp = make(map[int]struct32B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct32B{[4]int{i}}
		}
		test.MemStats()
	})
	t.Run("val 40b", func(t *testing.T) {
		var mp = make(map[int]struct40B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct40B{[5]int{i}}
		}
		test.MemStats()
	})
	t.Run("val 64b", func(t *testing.T) {
		var mp = make(map[int]struct64B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct64B{[8]int{i}}
		}
		test.MemStats()
	})
	t.Run("val 128b", func(t *testing.T) {
		var mp = make(map[int]struct128B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct128B{[16]int{i}}
		}
		test.MemStats()
	})
	t.Run("val 256b", func(t *testing.T) {
		var mp = make(map[int]struct256B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct256B{[32]int{i}}
		}
		test.MemStats()
	})
}

func testMapHeapSize(n int) {
	var mp = make(map[int]int)
	for i := 0; i < n; i++ {
		mp[i] = i
	}
	test.MemStats()
	for k := range mp {
		delete(mp, k)
	}
	test.MemStats()
	for i := 0; i < 2*n; i++ {
		mp[i] = i
	}
	test.MemStats()
	for k := range mp {
		delete(mp, k)
	}
	test.MemStats()
	for i := 0; i < 3*n; i++ {
		mp[i] = i
	}
	test.MemStats()
}

func testBufferMapHeapSize(n int) {
	var mp = make(map[int]int, n)
	for i := 0; i < n; i++ {
		mp[i] = i
	}
	test.MemStats()
	for k := range mp {
		delete(mp, k)
	}
	test.MemStats()
	for i := 0; i < 2*n; i++ {
		mp[i] = i
	}
	test.MemStats()
	for k := range mp {
		delete(mp, k)
	}
	test.MemStats()
	for i := 0; i < 3*n; i++ {
		mp[i] = i
	}
	test.MemStats()
}

func panicWarper(f func()) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	f()
}

func TestMapKeyType(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		panicWarper(func() {
			m := make(map[unsafe.Pointer]int)
			func() {
				a := 1
				m[unsafe.Pointer(&a)] = 1
			}()
			time.Sleep(time.Second)
			runtime.GC()
			for key, val := range m {
				fmt.Println(*(*int)(key), val)
			}
		})
	})
	type tstruct struct {
		a *int
	}
}

func BenchmarkMakeMap(b *testing.B) {
	test.MemStats()
	for i := 0; i < b.N; i++ {
		m := make(map[int]bool)
		m[1] = true
	}
	test.MemStats()
}

func TestMapSize(t *testing.T) {
	test.MemStats()
	intMap := make(map[int]int, 1e5)
	test.MemStats()
	intMap = nil
	time.Sleep(time.Second)
	test.MemStats()
	time.Sleep(time.Second)
	test.MemStats()
	_ = intMap
}

func TestMapDirectIface(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		const n = 10000000
		var keys = [n]byte{}
		for i := range keys {
			keys[i] = '1'
		}
		var mp = make(map[*byte]*byte, n)
		for i := range keys {
			mp[&keys[i]] = &keys[i]
		}
	})
	t.Run("var", func(t *testing.T) {
		const n = 10000000
		var mp = make(map[*byte]*byte, n)
		for i := 0; i < n; i++ {
			var t byte = '1'
			mp[&t] = &t
		}
	})
}
