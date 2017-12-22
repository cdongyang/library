package rbtree_test

import (
	"testing"

	"github.com/cdongyang/library/rbtree"
)

func BenchmarkRBTreeInsert1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	for i := 0; i < b.N; i++ {
		tree.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkRBTreeInsertWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeBuf = make([]node, b.N)
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= b.N {
					return &node{key: e}
				}
				nodeBuf[num].key = e
				num++
				return &nodeBuf[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	for i := 0; i < b.N; i++ {
		tree.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkSysHashMapInsert1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var mp = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		mp[rand.Int()] = true
	}
}

func BenchmarkSysHashMapInsertWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var mp = make(map[int]bool, b.N)
	for i := 0; i < b.N; i++ {
		mp[rand.Int()] = true
	}
}

func BenchmarkSysHashMapErase1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		delete(mp, keys[i])
	}
}

func BenchmarkSysHashMapEraseWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var keys = make([]int, b.N)
	b.StopTimer()
	var mp = make(map[int]bool, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		delete(mp, keys[i])
	}
}

func BenchmarkSetInsert1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var set = rbtree.NewSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		})
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkSetInsertWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var arr = make([]rbtree.SetNode, b.N)
	var num = 0
	var set = rbtree.NewCustomSet(
		func(a rbtree.SetIterator, b rbtree.SetIterator) int {
			return a.GetKey().(int) - b.GetKey().(int)
		},
		func(elem interface{}) rbtree.Iterator {
			if num >= b.N {
				return rbtree.NewSetNode(elem)
			}
			arr[num].SetKey(elem)
			num++
			return &arr[num-1]
		},
		func(iter rbtree.Iterator) {
		},
	)
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
	memStats()
}

func BenchmarkRBTreeErase1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	b.StopTimer()
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}

func BenchmarkRBTreeEraseWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	b.StopTimer()
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeBuf = make([]node, b.N)
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= b.N {
					return &node{key: e}
				}
				nodeBuf[num].key = e
				num++
				return &nodeBuf[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}

func BenchmarkRBTreeInsertAndErase1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				return &node{key: e}
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}

func BenchmarkRBTreeInsertAndEraseWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var (
		compare = func(a Iterator, b Iterator) int {
			return a.(*node).key - b.(*node).key
		}
		nodeBuf = make([]node, b.N)
		num     = 0
		newElem = func(elem interface{}) Iterator {
			switch e := elem.(type) {
			case node:
				return &e
			case *node:
				return e
			case int:
				if num >= b.N {
					return &node{key: e}
				}
				nodeBuf[num].key = e
				num++
				return &nodeBuf[num-1]
			default:
				panic("error new type")
			}
		}
		deleteElem = func(elem Iterator) {
		}
	)
	tree := rbtree.NewCustomRBTree(true, (*node)(nil), compare, newElem, deleteElem)
	var keys = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		tree.Insert(keys[i])
	}
	for i := 0; i < b.N; i++ {
		tree.Erase(keys[i])
	}
	memStats()
}

func BenchmarkSysHashMapInsertAndErase1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var keys = make([]int, b.N)
	var mp = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	for i := 0; i < b.N; i++ {
		delete(mp, keys[i])
	}
}

func BenchmarkSysHashMapInsertAndEraseWithBuf1E6(b *testing.B) {
	b.N = 1e6
	var rand = benchRand
	var keys = make([]int, b.N)
	var mp = make(map[int]bool, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Int()
		mp[keys[i]] = true
	}
	for i := 0; i < b.N; i++ {
		delete(mp, keys[i])
	}
}
