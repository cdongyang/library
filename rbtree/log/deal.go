package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filename = "go1.7bench.log"
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	out, err := os.Create(filename + ".md")
	if err != nil {
		panic(err.Error())
	}
	defer out.Close()
	buf := bufio.NewReader(file)
	var strArr = []string{}
	var idArr = []int{}
	for i := 0; ; i++ {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}
		var s = string(line)
		strArr = append(strArr, s)
		if strings.HasPrefix(s, "Benchmark") {
			idArr = append(idArr, i)
			fmt.Fprint(out, "#")
		}
		fmt.Fprintln(out, s)
	}
	/*for i, val := range idArr {
		if i != 0 {
			fmt.Fprintln(out, strArr[val-6])
			fmt.Fprintln(out, strArr[val-5])
			fmt.Fprintln(out, strArr[val-4])
			fmt.Fprintln(out, strArr[val-3])
			fmt.Fprintln(out, strArr[val-2])
			fmt.Fprintln(out, strArr[val-1])
		}
		fmt.Fprintln(out, "#"+strArr[val])
	}
	fmt.Fprintln(out, strArr[len(strArr)-7])
	fmt.Fprintln(out, strArr[len(strArr)-6])
	fmt.Fprintln(out, strArr[len(strArr)-5])
	fmt.Fprintln(out, strArr[len(strArr)-4])
	fmt.Fprintln(out, strArr[len(strArr)-3])
	fmt.Fprintln(out, strArr[len(strArr)-2])
	fmt.Fprintln(out, strArr[len(strArr)-1])*/
}
