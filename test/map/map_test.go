package map_test

import (
	"fmt"
	"runtime"
	"testing"
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
		memStats()
	})
	t.Run("key 40b", func(t *testing.T) {
		var mp = make(map[struct40B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct40B{[5]int{i}}] = i
		}
		memStats()
	})
	t.Run("key 64b", func(t *testing.T) {
		var mp = make(map[struct64B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct64B{[8]int{i}}] = i
		}
		memStats()
	})
	t.Run("key 128b", func(t *testing.T) {
		var mp = make(map[struct128B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct128B{[16]int{i}}] = i
		}
		memStats()
	})
	t.Run("key 256b", func(t *testing.T) {
		var mp = make(map[struct256B]int)
		for i := 0; i < 1e5; i++ {
			mp[struct256B{[32]int{i}}] = i
		}
		memStats()
	})
	t.Run("val 32b", func(t *testing.T) {
		var mp = make(map[int]struct32B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct32B{[4]int{i}}
		}
		memStats()
	})
	t.Run("val 40b", func(t *testing.T) {
		var mp = make(map[int]struct40B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct40B{[5]int{i}}
		}
		memStats()
	})
	t.Run("val 64b", func(t *testing.T) {
		var mp = make(map[int]struct64B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct64B{[8]int{i}}
		}
		memStats()
	})
	t.Run("val 128b", func(t *testing.T) {
		var mp = make(map[int]struct128B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct128B{[16]int{i}}
		}
		memStats()
	})
	t.Run("val 256b", func(t *testing.T) {
		var mp = make(map[int]struct256B)
		for i := 0; i < 1e5; i++ {
			mp[i] = struct256B{[32]int{i}}
		}
		memStats()
	})
}

func testMapHeapSize(n int) {
	var mp = make(map[int]int)
	for i := 0; i < n; i++ {
		mp[i] = i
	}
	memStats()
	for k := range mp {
		delete(mp, k)
	}
	memStats()
	for i := 0; i < 2*n; i++ {
		mp[i] = i
	}
	memStats()
	for k := range mp {
		delete(mp, k)
	}
	memStats()
	for i := 0; i < 3*n; i++ {
		mp[i] = i
	}
	memStats()
}

func testBufferMapHeapSize(n int) {
	var mp = make(map[int]int, n)
	for i := 0; i < n; i++ {
		mp[i] = i
	}
	memStats()
	for k := range mp {
		delete(mp, k)
	}
	memStats()
	for i := 0; i < 2*n; i++ {
		mp[i] = i
	}
	memStats()
	for k := range mp {
		delete(mp, k)
	}
	memStats()
	for i := 0; i < 3*n; i++ {
		mp[i] = i
	}
	memStats()
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
