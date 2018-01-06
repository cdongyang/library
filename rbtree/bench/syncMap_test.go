package bench_test

import (
	"sync"
	"testing"
)

func BenchmarkSyncMapInsert1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var mp = sync.Map{}
	for i := 0; i < b.N; i++ {
		mp.Store(IntKey(rand.Int()), true)
	}
	memStats()
}

func BenchmarkSyncMapErase1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = sync.Map{}
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp.Store(IntKey(keys[i]), true)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mp.Delete(IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSyncMapFind1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = sync.Map{}
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp.Store(IntKey(keys[i]), true)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = mp.Load(IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSyncMapInsertAndErase1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N/10)
	var insertn, erasen int
	var mp = sync.Map{}
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			mp.Store(IntKey(keys[insertn]), true)
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			mp.Delete(IntKey(keys[erasen]))
			erasen++
			i++
		}
	}
	memStats()
}
