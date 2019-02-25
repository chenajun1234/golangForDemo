package main

import (
	"fmt"
	"net"
)

func handerConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println(addr)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read异常", err)
			return
		}
		s := string(buffer[:n])
		fmt.Printf("接收:%s,长度%d", s, len(s))

		conn.Write([]byte("接收到了"))
	}

}
func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8001")
	if e != nil {
		fmt.Println("listener异常", e)
		return
	}
	defer listener.Close()

	for { //等待多个链接
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("conn异常", e)
			return
		}
		defer conn.Close()
		go handerConn(conn)
	}
}
