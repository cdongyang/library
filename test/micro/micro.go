package main

import (
	"fmt"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:   "v",
				EnvVar: "VVV",
			},
			cli.StringSliceFlag{
				Name:   "slice",
				EnvVar: "SLICE",
				Value:  &cli.StringSlice{},
			},
		),
		micro.Action(func(ctx *cli.Context) {
			fmt.Println(ctx.Args())
			fmt.Println(ctx.FlagNames())
			fmt.Println(ctx.String("v"))
			slice := ctx.StringSlice("slice")
			fmt.Println(slice)
			for i := 0; i < len(slice); i++ {
				fmt.Println(slice[i])
			}
		}),
	)
	service.Init()
}

/*
优先解析命令行参数，如果命令行没传入则用环境变量
➜  micro git:(master) ✗ VVV=bbb go run micro.go -v aaa
[]
[]
aaa
➜  micro git:(master) ✗ VVV=bbb go run micro.go
[]
[]
bbb
➜  micro git:(master) ✗ VVV=bbb go run micro.go -v aaa
[]
[]
aaa
*/
