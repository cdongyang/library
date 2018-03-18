package escape_test

import (
	"reflect"
	"testing"
)

func TestReflectValueEscape(t *testing.T) {
	var n float64
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		_ = reflect.ValueOf(a)
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		_ = reflect.TypeOf(a)
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var f = reflect.ValueOf(compareFunc)
		var a = []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
		f.Call(a)
	})
	t.Log(n, "allocs")
}
