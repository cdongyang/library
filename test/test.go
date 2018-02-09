package test

import (
	"fmt"
	"runtime"
)

var mem runtime.MemStats
var OutPutmem = true

func MemStats(prefix ...string) {
	if !OutPutmem {
		return
	}
	runtime.ReadMemStats(&mem)
	fmt.Println(prefix, "HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}
