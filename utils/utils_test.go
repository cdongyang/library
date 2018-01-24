package utils_test

import (
	"testing"

	"github.com/cdongyang/library/utils"
)

func getLowbit(x int) int {
	var res = 1
	for x != 0 {
		if (x & 1) == 1 {
			return res
		}
		x >>= 1
		res <<= 1
	}
	return res
}

func TestLowbit(t *testing.T) {
	for i := 1; i < (1 << 8); i++ {
		if utils.Lowbit(i) != getLowbit(i) {
			t.Fatalf("%d: want:%d real: %d", i, utils.Lowbit(i), getLowbit(i))
		}
	}
}
