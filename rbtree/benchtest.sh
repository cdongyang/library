rm cpu.out
go test -v -benchmem -bench="." -benchtime 10s -cpuprofile=cpu.out