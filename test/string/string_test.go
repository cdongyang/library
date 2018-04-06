package string_test

import (
	"fmt"
	"os"
	"testing"
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
