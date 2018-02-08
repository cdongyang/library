package map_test

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
)

func TestMapHeapSize1E6(t *testing.T) {
	testMapHeapSize(1e6)
}
func TestMapHeapSize1E5(t *testing.T) {
	testMapHeapSize(1e5)
}

func testMapHeapSize(n int) {
	debug.SetGCPercent(-1)
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
	debug.SetGCPercent(100)
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
