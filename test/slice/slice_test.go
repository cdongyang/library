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

func do(qs []int) {
	qs = append(qs, 7)
}

func equalSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAppendSlice(t *testing.T) {
	var limit = 5
	var qs = []int{1, 2, 3, 4, 5, 6}
	if len(qs) != 6 || cap(qs) != 6 {
		t.Fatal("not expected len or cap")
	}
	var batch = make([]int, limit)
	for i := 0; i < len(qs); i += limit {
		if i+limit < len(qs) {
			copy(batch, qs[i:i+limit])
		} else {
			batch = batch[:len(qs)-i]
			if cap(batch) != limit || len(batch) != len(qs)-i {
				t.Fatal("not expected len or cap", cap(batch), len(batch))
			}
			copy(batch, qs[i:])
		}
		do(batch)
		if i+limit < len(qs) {
			if !equalSlice(batch, qs[i:i+limit]) {
				t.Fatal("not equal slice")
			}
		} else {
			if !equalSlice(batch, qs[i:]) {
				t.Fatal("not equal slice")
			}
		}
	}
}

func TestSliceLen(t *testing.T) {
	a := make([]int, 10)
	t.Log(len(a), cap(a))
}
