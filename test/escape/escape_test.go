package escape_test

import "testing"

type cmp interface {
	compareMethod(a, b interface{}) int
}

type Comp struct {
	Compare func(a, b interface{}) int
}

func (c *Comp) compareMethod(a, b interface{}) int {
	var aa, bb int
	switch t := a.(type) {
	case int:
		aa = t
	case *int:
		aa = *t
	default:
		panic("not int or *int")
	}
	switch t := b.(type) {
	case int:
		bb = t
	case *int:
		bb = *t
	default:
		panic("not int or *int")
	}
	return aa - bb
}

func compareFunc(a, b interface{}) int {
	var aa, bb int
	switch t := a.(type) {
	case int:
		aa = t
	case *int:
		aa = *t
	default:
		panic("not int or *int")
	}
	switch t := b.(type) {
	case int:
		bb = t
	case *int:
		bb = *t
	default:
		panic("not int or *int")
	}
	return aa - bb
}

func TestMethodEscape(t *testing.T) {
	var n float64
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		for i := 0; i < 100; i++ {
			comp.compareMethod(2, 3)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		for i := 0; i < 100; i++ {
			comp.compareMethod(1, 2)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		var a = 1
		for i := 0; i < 100; i++ {
			comp.compareMethod(a, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		var a = 5
		for i := 0; i < 100; i++ {
			comp.compareMethod(a, a)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		for i := 0; i < 100; i++ {
			var t = i
			comp.compareMethod(t, t)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		for i := 0; i < 100; i++ {
			comp.compareMethod(i, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		var a = [2]int{1, 2}
		for i := 0; i < 100; i++ {
			comp.compareMethod(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{}
		var a = []int{1, 2}
		for i := 0; i < 100; i++ {
			comp.compareMethod(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
}

func TestFuncEscape(t *testing.T) {
	var n float64
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 100; i++ {
			compareFunc(2, 3)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 100; i++ {
			compareFunc(1, 2)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var a = 1
		for i := 0; i < 100; i++ {
			compareFunc(a, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var a = 5
		for i := 0; i < 100; i++ {
			compareFunc(a, a)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 100; i++ {
			var t = i
			compareFunc(t, t)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		for i := 0; i < 100; i++ {
			compareFunc(i, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var a = [2]int{1, 2}
		for i := 0; i < 100; i++ {
			compareFunc(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var a = []int{1, 2}
		for i := 0; i < 100; i++ {
			compareFunc(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
}

func TestStructFuncEscape(t *testing.T) {
	var n float64
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			comp.Compare(2, 3)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			comp.Compare(1, 2)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = 1
		for i := 0; i < 100; i++ {
			comp.Compare(a, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = 5
		for i := 0; i < 100; i++ {
			comp.Compare(a, a)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			var t = i
			comp.Compare(t, t)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			comp.Compare(i, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = [2]int{1, 2}
		for i := 0; i < 100; i++ {
			comp.Compare(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = []int{1, 2}
		for i := 0; i < 100; i++ {
			comp.Compare(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
}

func TestInterfaceMethodEscape(t *testing.T) {
	var n float64
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		for i := 0; i < 100; i++ {
			comp.compareMethod(2, 3)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		for i := 0; i < 100; i++ {
			comp.compareMethod(1, 2)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		var a = 1
		for i := 0; i < 100; i++ {
			comp.compareMethod(a, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		var a = 5
		for i := 0; i < 100; i++ {
			comp.compareMethod(a, a)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		for i := 0; i < 100; i++ {
			var t = i
			comp.compareMethod(t, t)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		for i := 0; i < 100; i++ {
			comp.compareMethod(i, i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		var a = [2]int{1, 2}
		for i := 0; i < 100; i++ {
			comp.compareMethod(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp cmp = &Comp{}
		var a = []int{1, 2}
		for i := 0; i < 100; i++ {
			comp.compareMethod(a[0], a[1])
		}
	})
	t.Log(n, "allocs")
}

func TestStructFuncPointerEscape(t *testing.T) {
	var n float64
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			comp.Compare(2, 3)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			comp.Compare(1, 2)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = 1
		for i := 0; i < 100; i++ {
			comp.Compare(&a, &i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = 5
		for i := 0; i < 100; i++ {
			comp.Compare(&a, &a)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			var t = i
			comp.Compare(&t, &t)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		for i := 0; i < 100; i++ {
			comp.Compare(&i, &i)
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = [2]int{1, 2}
		for i := 0; i < 100; i++ {
			comp.Compare(&a[0], &a[1])
		}
	})
	t.Log(n, "allocs")
	n = testing.AllocsPerRun(1000, func() {
		var comp = &Comp{compareFunc}
		var a = []int{1, 2}
		for i := 0; i < 100; i++ {
			comp.Compare(&a[0], &a[1])
		}
	})
	t.Log(n, "allocs")
}

/*
----go1.9.2
=== RUN   TestMethodEscape
--- PASS: TestMethodEscape (0.04s)
        escape_test.go:63: 0 allocs
        escape_test.go:70: 0 allocs
        escape_test.go:78: 0 allocs
        escape_test.go:86: 0 allocs
        escape_test.go:94: 0 allocs
        escape_test.go:101: 0 allocs
        escape_test.go:109: 0 allocs
        escape_test.go:117: 0 allocs
=== RUN   TestFuncEscape
--- PASS: TestFuncEscape (0.03s)
        escape_test.go:127: 0 allocs
        escape_test.go:133: 0 allocs
        escape_test.go:140: 0 allocs
        escape_test.go:147: 0 allocs
        escape_test.go:154: 0 allocs
        escape_test.go:160: 0 allocs
        escape_test.go:167: 0 allocs
        escape_test.go:174: 0 allocs
=== RUN   TestStructFuncEscape
--- PASS: TestStructFuncEscape (0.33s)
        escape_test.go:185: 0 allocs
        escape_test.go:192: 0 allocs
        escape_test.go:200: 199 allocs
        escape_test.go:208: 200 allocs
        escape_test.go:216: 198 allocs
        escape_test.go:223: 198 allocs
        escape_test.go:231: 200 allocs
        escape_test.go:239: 200 allocs
=== RUN   TestInterfaceMethodEscape
--- PASS: TestInterfaceMethodEscape (0.61s)
        escape_test.go:250: 1 allocs
        escape_test.go:257: 1 allocs
        escape_test.go:265: 200 allocs
        escape_test.go:273: 201 allocs
        escape_test.go:281: 199 allocs
        escape_test.go:288: 199 allocs
        escape_test.go:296: 201 allocs
        escape_test.go:304: 201 allocs
=== RUN   TestStructFuncPointerEscape
--- PASS: TestStructFuncPointerEscape (0.23s)
        escape_test.go:315: 0 allocs
        escape_test.go:322: 0 allocs
        escape_test.go:330: 2 allocs
        escape_test.go:338: 1 allocs
        escape_test.go:346: 100 allocs
        escape_test.go:353: 1 allocs
        escape_test.go:361: 1 allocs
				escape_test.go:369: 1 allocs

----go1.7
=== RUN   TestMethodEscape
--- PASS: TestMethodEscape (0.04s)
        escape_test.go:63: 0 allocs
        escape_test.go:70: 0 allocs
        escape_test.go:78: 0 allocs
        escape_test.go:86: 0 allocs
        escape_test.go:94: 0 allocs
        escape_test.go:101: 0 allocs
        escape_test.go:109: 0 allocs
        escape_test.go:117: 0 allocs
=== RUN   TestFuncEscape
--- PASS: TestFuncEscape (0.04s)
        escape_test.go:127: 0 allocs
        escape_test.go:133: 0 allocs
        escape_test.go:140: 0 allocs
        escape_test.go:147: 0 allocs
        escape_test.go:154: 0 allocs
        escape_test.go:160: 0 allocs
        escape_test.go:167: 0 allocs
        escape_test.go:174: 0 allocs
=== RUN   TestStructFuncEscape
--- PASS: TestStructFuncEscape (0.08s)
        escape_test.go:185: 200 allocs
        escape_test.go:192: 200 allocs
        escape_test.go:200: 200 allocs
        escape_test.go:208: 200 allocs
        escape_test.go:216: 200 allocs
        escape_test.go:223: 200 allocs
        escape_test.go:231: 200 allocs
        escape_test.go:239: 200 allocs
=== RUN   TestInterfaceMethodEscape
--- PASS: TestInterfaceMethodEscape (0.08s)
        escape_test.go:250: 201 allocs
        escape_test.go:257: 201 allocs
        escape_test.go:265: 201 allocs
        escape_test.go:273: 201 allocs
        escape_test.go:281: 201 allocs
        escape_test.go:288: 201 allocs
        escape_test.go:296: 201 allocs
        escape_test.go:304: 201 allocs
=== RUN   TestStructFuncPointerEscape
--- PASS: TestStructFuncPointerEscape (0.04s)
				escape_test.go:315: 200 allocs
        escape_test.go:322: 200 allocs
        escape_test.go:330: 2 allocs
        escape_test.go:338: 1 allocs
        escape_test.go:346: 100 allocs
        escape_test.go:353: 1 allocs
        escape_test.go:361: 1 allocs
        escape_test.go:369: 1 allocs
		interface{} 会持有一个unsafe.Pointer,所以当值赋值给interface{}时有可能引起变量从栈内存逃逸到堆内存,
		引起内存分配,降低运行速度
		interface{} 函数调用传参 栈变量逃逸到堆 测试:
		推论:
			1.struct调用Method或者直接调用函数不会引起栈变量逃逸
			2.struct func element调用或者 interface method调用必会引起栈变量逃逸
			3.go1.9.2相对go1.7优化了常量的内存逃逸和0值变量的内存逃逸
			4.第2点条件下栈变量指针传值时也必定引起变量逃逸
*/
