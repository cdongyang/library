package string_test

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

func constString() {
	fmt.Fprintln(os.Stdout, "long long long long long string")
}

func TestString(t *testing.T) {
	n := testing.AllocsPerRun(1000, constString)
	if n != 0 {
		t.Fatal(n)
	}
}

func BenchmarkConstString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		constString()
	}
}

func ExampleStr() {
	a := 2
	c := (*string)(unsafe.Pointer(&a))
	*c = "44" // a只有八字节,写入时string的后八字节len会越界,结果不可预测
	fmt.Println(*c)
	fmt.Println(a)
	fmt.Println((*struct {
		p   unsafe.Pointer
		len int
	})(unsafe.Pointer(&a)).len)
	fmt.Printf("&c:%p c:%p &a:%p len:%d\n", &c, c, &a, len(*c))
	//Output:
	//44
	//5458994
	//5458994
	//&c:0xc42000c058 c:0xc420016290 &a:0xc420016290 len:5458994
}
