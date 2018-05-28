package os_args_test

import (
	"flag"
	"os"
	"testing"
)

func TestArgs(t *testing.T) {
	for _, val := range os.Args {
		t.Log(val)
	}
	flag.Parse()
	t.Log("after parse")
	for _, val := range os.Args {
		t.Log(val)
	}
}
