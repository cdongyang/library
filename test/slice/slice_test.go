package slice_test

import "testing"

// copy from package runtime
var One = []int64{1}

func TestAppendSliceGrowth(t *testing.T) {
	var x []int64
	check := func(want int) {
		if cap(x) != want {
			t.Errorf("len=%d, cap=%d, want cap=%d", len(x), cap(x), want)
		}
	}

	check(0)
	want := 1
	for i := 1; i <= 100; i++ {
		tmp := append(x, One...)
		t.Errorf("x: len=%d, cap=%d, want cap=%d", len(x), cap(x), want)
		t.Errorf("tmp: len=%d, cap=%d, want cap=%d", len(tmp), cap(tmp), want)
		x = tmp
		check(want)
		if i&(i-1) == 0 {
			want = 2 * i
		}
	}
}
