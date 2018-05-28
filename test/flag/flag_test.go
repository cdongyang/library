package flag_test

import (
	"flag"
	"os"
	"testing"
)

func TestFlag(t *testing.T) {
	t.Log(os.Args)
	t.Log("-----------------------------")
	flag.Parse()
	t.Log(os.Args)
	t.Log("-----------------------------")
	flag.Visit(func(f *flag.Flag) {
		t.Log(f.Name, f.Value, f.DefValue, f.Usage)
	})
	t.Log("-----------------------------")
	flag.VisitAll(func(f *flag.Flag) {
		t.Log(f.Name, f.Value, f.DefValue, f.Usage)
	})
}
