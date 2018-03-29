package tcp_test

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// ➜  src netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'
// FIN_WAIT2 10000	客户端关闭后服务端等待30s才Close, 回复FIN报文
// CLOSE_WAIT 10000 服务端收到FIN报文, 处于CLOSE_WAIT状态, 等待30s
// TIME_WAIT 1
// ESTABLISHED 5
// SYN_SENT 5
//
// after 30s
// ➜  src netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'
// TIME_WAIT 10000		客户端发出关闭ACK后进入TIME_WAIT, 如果服务端没收到ACK会重发FIN+ACK, 客户端仍能回应
// ESTABLISHED 8
// SYN_SENT 5
func TestTcp(t *testing.T) {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err.Error())
	}
	defer listener.Close()
	go func() { // 客户端建立连接后直接关闭
		for i := 0; i < 10000; i++ {
			conn, err := net.Dial("tcp", "127.0.0.1:8000")
			if err != nil {
				panic(err.Error())
			}
			conn.Close() // 发送客户端请求关闭FIN
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		go func(conn net.Conn) {
			var buf = make([]byte, 1000)
			defer conn.Close()        // 发送服务端关闭FIN报文
			for i := 0; i < 30; i++ { // 服务端等待30s后关闭
				time.Sleep(time.Second)
				n, err := conn.Read(buf)
				fmt.Printf("%v %v\n", n, err)
			}
			fmt.Println("Close")
		}(conn)
	}
}
