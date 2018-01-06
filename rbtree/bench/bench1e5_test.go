package bench_test

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"testing"

	"github.com/cdongyang/library/randint"
	"github.com/cdongyang/library/rbtree"
)

var mem runtime.MemStats
var outmem = false

func memStats() {
	if !outmem {
		return
	}
	runtime.ReadMemStats(&mem)
	fmt.Println("HeapAlloc:", mem.HeapAlloc, "HeapInuse:", mem.HeapInuse, "HeapObjects:", mem.HeapObjects, "HeapIdle", mem.HeapIdle, "HeapReleased", mem.HeapReleased, "HeapSys", mem.HeapSys)
	runtime.GC()
}

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

type IntKey = rbtree.IntKey

func (s *struct32b) Compare(s1 rbtree.Keyer) int {
	return s.a - s1.(*struct32b).a
}

func (s *struct32b) GetKey() rbtree.Keyer {
	return s
}

func BenchmarkSort1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
	}
	sort.IntSlice(keys).Sort()
	memStats()
}

func BenchmarkSetInsert1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var set = rbtree.NewSet()
	for i := 0; i < b.N; i++ {
		var tmp = IntKey(rand.Int())
		set.Insert(tmp)
	}
	memStats()
}

func BenchmarkSetErase1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(IntKey(keys[i]))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Erase(IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSetNext1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(IntKey(keys[i]))
	}
	var iter = set.Begin()
	b.StartTimer()
	for ; iter != set.End(); iter = iter.Next(iter) {
	}
	memStats()
}

func BenchmarkSetLast1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(IntKey(keys[i]))
	}
	var iter = set.EndNode()
	b.StartTimer()
	for ; iter != set.End(); iter = iter.Last(iter) {
	}
	memStats()
}

func BenchmarkSetFind1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(IntKey(keys[i]))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Find(IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSetLowerBound1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(IntKey(keys[i]))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.LowerBound(IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSetUpperBound1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		set.Insert(IntKey(keys[i]))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.UpperBound(IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSetInsertAndErase1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	var insertn, erasen int
	var set = rbtree.NewSet()
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			set.Insert(IntKey(keys[insertn]))
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			set.Erase(IntKey(keys[erasen]))
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSetInsertAndEraseWithPool1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var nodePool = sync.Pool{New: func() interface{} {
		return &rbtree.SetNode{}
	}}
	var keys = make([]int, b.N)
	var insertn, erasen int
	var set = rbtree.NewCustomSet(
		func(elem rbtree.Keyer) rbtree.Iterator {
			var iter = nodePool.Get().(*rbtree.SetNode)
			iter.SetKey(elem)
			return iter
		},
		func(iter rbtree.Iterator) {
			nodePool.Put(iter)
		},
	)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			set.Insert(IntKey(keys[insertn]))
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			set.Erase(IntKey(keys[erasen]))
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSysHashMapInsert1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var mp = make(map[IntKey]bool)
	for i := 0; i < b.N; i++ {
		mp[IntKey(rand.Int())] = true
	}
	memStats()
}

func BenchmarkSysHashMapErase1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = make(map[IntKey]bool)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[IntKey(keys[i])] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		delete(mp, IntKey(keys[i]))
	}
	memStats()
}

func BenchmarkSysHashMapFind1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = make(map[IntKey]bool)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[IntKey(keys[i])] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = mp[IntKey(keys[i])]
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndErase1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N/10)
	var insertn, erasen int
	var mp = make(map[IntKey]bool)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			mp[IntKey(keys[insertn])] = true
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			delete(mp, IntKey(keys[erasen]))
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndEraseWithBuf1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]int, b.N/10)
	var insertn, erasen int
	var mp = make(map[IntKey]bool, b.N)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn] = rand.Int()
			mp[IntKey(keys[insertn])] = true
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			delete(mp, IntKey(keys[erasen]))
			erasen++
			i++
		}
	}
	memStats()
}
