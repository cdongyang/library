package switch_test

import (
	"testing"

	"github.com/cdongyang/library/randint"
)

var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}

func BenchmarkSwitch(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		switch {
		case i == 0:
			a = 1
		case i > 0:
			a = 2
		case i < 0:
			a = 3
		}
	}
	_ = a
}

func BenchmarkIf(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		if i == 0 {
			a = 1
		} else if i > 0 {
			a = 2
		} else if i < 0 {
			a = 3
		}
	}
	_ = a
}
