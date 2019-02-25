package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", ":8000")
	if e != nil {
		fmt.Println("conn异常:", e)
		return
	}
	defer conn.Close()
	go func() { //辅协程 轮询接收键盘输入字符
		for {
			reader := bufio.NewReader(os.Stdin)
			str, e := reader.ReadString('\n')
			if e != nil {
				fmt.Println("Stdin异常", e)
			}
			conn.Write([]byte(str))
		}
	}()

	//主协程 接收服务器发来的数据后 函数结束,辅协程也随着主函数结束而结束

	//接收服务器回复的数据
	//切片缓冲
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务器的请求
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印接收到的内容, 转换为字符串再打印
	}

}
