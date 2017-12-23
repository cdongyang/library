package randint_test

import (
	"fmt"

	"github.com/cdongyang/library/randint"
)

var base = randint.Rand{Add: 5, First: 4, Mod: 9}
var copyBase = base

func ExampleInt() {
	var rand1 = base
	for i := 0; i < 2*rand1.Mod; i++ {
		fmt.Printf("%d ", rand1.Int())
		if copyBase != base {
			panic("base changed")
		}
	}
	// Output: 0 5 1 6 2 7 3 8 4 0 5 1 6 2 7 3 8 4
}
