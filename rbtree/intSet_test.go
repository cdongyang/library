package rbtree_test

import "testing"
import "github.com/cdongyang/library/rbtree"

func BenchmarkIntSetInsert(b *testing.B) {
	var rand = benchRand
	var intSet = rbtree.NewIntSet()
	for i:=0;i<b.N;i++ {
		intSet.Insert(rand.Int())
	}
}
