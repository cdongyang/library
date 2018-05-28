package pointer_test

import "testing"

func TestPointerEqual(t *testing.T) {
	var sp1 *string
	var sp2 = (*string)(nil)
	t.Log(sp1 == sp2, sp1 == nil, sp2 == nil)
	type node struct {
		a int
	}
	var np1 *node
	var np2 = (*node)(nil)
	t.Log(np1 == np2, np1 == nil, np2 == nil)
}
