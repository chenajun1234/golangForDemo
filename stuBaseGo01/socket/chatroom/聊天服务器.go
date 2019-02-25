package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

/*
1. 主协程 监听 处理用户连接
2  协程 添加上线用户  广播上线提醒到用户列表
3. 协程 处理客户端请求模块
4  协程 真正发消息 功能
5   协程  转发大家的消息

*/
type user struct {
	Ip   string
	Name string
	ch   chan string
}

var onlineUser map[string]user

var message = make(chan string)

func main() {

	listener, e := net.Listen("tcp", ":8000")
	if e != nil {
		fmt.Println("Listen异常:", e)
	}
	fmt.Println("建立监听")
	defer listener.Close()
	//新开一个协程，转发消息，只要有消息来了，遍历map, 给map每个成员都发送此消息
	go broadCastMsg()
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("Listen异常:", e)
			continue
		}
		go handerConn(conn)

	}

}

//广播消息到用户列表
func broadCastMsg() {
	//给map分配空间
	onlineUser = make(map[string]user)
	for {
		msg := <-message
		for _, us := range onlineUser {
			fmt.Println("广播消息到用户列表:", msg)
			us.ch <- msg
		}
	}
}
func MakeMsg(u user, msg string) (buf string) {
	buf = "[" + u.Ip + "]" + u.Name + ": " + msg
	return buf
}
func sendMsgToClient(conn net.Conn, u user) {
	for c := range u.ch {
		conn.Write([]byte(c))
	}
}
func handerConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	u := user{addr, addr, make(chan string)}
	onlineUser[addr] = u
	//发送消息到客户端
	go sendMsgToClient(conn, u)
	//广播某个在线
	message <- MakeMsg(u, "上线了")

	isQuit := make(chan bool)  //退出
	hasData := make(chan bool) //还在发送数据
	go func() {
		var reader *bufio.Reader = nil
		for {
			reader = bufio.NewReader(conn)
			msg, e := reader.ReadString('\n')
			if e != nil {
				isQuit <- true
				fmt.Println("bufio.NewReader异常", e)
			}
			msg = strings.TrimSpace(msg)
			//rename:XXX
			if msg == "who" {
				//谁在线
				for _, temp := range onlineUser {
					message <- MakeMsg(temp, "我在线")
				}
			} else if len(msg) > 6 && msg[:6] == "rename" {

				u.Name = strings.Split(msg, ":")[1]
				onlineUser[addr] = u
				conn.Write([]byte("rename ok"))
			} else {
				message <- MakeMsg(u, msg)
			}
			hasData <- true
		}
	}()
	for {
		select {
		case <-isQuit:
			//退出
			delete(onlineUser, addr)
			message <- MakeMsg(u, "退出聊天室")
			return
		case <-hasData:
		case <-time.After(time.Second * 30):
			delete(onlineUser, addr)
			message <- MakeMsg(u, "超时")
			return

		}
	}
}
