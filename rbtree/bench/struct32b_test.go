package bench_test

import (
	"sort"
	"sync"
	"testing"

	"github.com/cdongyang/library/rbtree"
)

type struct32b struct {
	a int
	b int
	c int
	d int
}

func BenchmarkSortStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].a < keys[j].a
	})
	memStats()
}

func BenchmarkSetInsertStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var set = rbtree.NewSet()
	for i := 0; i < b.N; i++ {
		set.Insert(&struct32b{a: rand.Int()})
	}
	memStats()
}

func BenchmarkSetEraseStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
		set.Insert(&keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Erase(&keys[i])
	}
	memStats()
}

func BenchmarkSetFindStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
		set.Insert(&keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Find(&keys[i])
	}
	memStats()
}

func BenchmarkSetLowerBoundStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
		set.Insert(&keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.LowerBound(&keys[i])
	}
	memStats()
}

func BenchmarkSetUpperBoundStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	var set = rbtree.NewSet()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
		set.Insert(&keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.UpperBound(&keys[i])
	}
	memStats()
}

func BenchmarkSetInsertAndEraseStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	var insertn, erasen int
	var set = rbtree.NewSet()
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn].a = rand.Int()
			set.Insert(&keys[insertn])
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			set.Erase(&keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSetInsertAndEraseWithPoolStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var nodePool = sync.Pool{New: func() interface{} {
		return &rbtree.SetNode{}
	}}
	var keys = make([]struct32b, b.N)
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
			keys[insertn].a = rand.Int()
			set.Insert(&keys[insertn])
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			set.Erase(&keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSysHashMapInsertStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var mp = make(map[*struct32b]bool)
	for i := 0; i < b.N; i++ {
		mp[&struct32b{a: rand.Int()}] = true
	}
	memStats()
}

func BenchmarkSysHashMapEraseStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	b.StopTimer()
	var mp = make(map[*struct32b]bool)
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
		mp[&keys[i]] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		delete(mp, &keys[i])
	}
	memStats()
}

func BenchmarkSysHashMapFindStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N)
	b.StopTimer()
	var mp = make(map[*struct32b]bool)
	for i := 0; i < b.N; i++ {
		keys[i].a = rand.Int()
		mp[&keys[i]] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = mp[&keys[i]]
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndEraseStruct1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N/10)
	var insertn, erasen int
	var mp = make(map[*struct32b]bool)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn].a = rand.Int()
			mp[&keys[insertn]] = true
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			delete(mp, &keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndEraseStructWithBuf1E5(b *testing.B) {
	b.N = 1e5
	var rand = benchRand
	var keys = make([]struct32b, b.N/10)
	var insertn, erasen int
	var mp = make(map[*struct32b]bool, b.N)
	for i := 0; i < b.N; i++ {
		insertn, erasen = 0, 0
		for j := 0; j < b.N/10; j++ {
			keys[insertn].a = rand.Int()
			mp[&keys[insertn]] = true
			insertn++
			i++
		}
		for j := 0; j < b.N/10; j++ {
			delete(mp, &keys[erasen])
			erasen++
			i++
		}
	}
	memStats()
}
