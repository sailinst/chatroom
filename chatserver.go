package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	//"strings"
)

const (
	ip   = "127.0.0.1" // IP地址
	port = 7777        // 进程端口号
)

var ConnMap map[string]*net.TCPConn //字典：存取客户端发来的聊天信息

func main() { // chat -host=chat-server-ip -myname=zhang3
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPConn)
	tcpAddr = &net.TCPAddr{net.ParseIP(ip), port, ""}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("端口监听失败：", err.Error())
	}
	defer listener.Close()
	Server(listener)
}

func Server(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("客户端连接异常：" + err.Error()) // 打印错误信息
			continue
		}
		fmt.Println("客户端连接来自：" + conn.RemoteAddr().String()) // 打印连接信息
		fmt.Println("当前时间：", time.Now().Format("2006-01-02 15:04:05 PM"))
		ConnMap[conn.RemoteAddr().String()] = conn
		go tcpPipe(conn)
	}
}
func tcpPipe(conn *net.TCPConn) {
	defer func() {
		fmt.Println("用户" + conn.RemoteAddr().String() + "断开连接...")
		fmt.Println("当前时间：", time.Now().Format("2006-01-02 15:04:05 PM"))
		conn.Close() // 客户端退出时，关闭conn连接 ，即结束一个goroutine
	}()
	reader := bufio.NewReader(conn)
	for {
		//  读取数据，以换行符结束。
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取客户端数据失败：", err.Error())
			return
		}
		// 输出客户端消息
		fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
		msg := conn.RemoteAddr().String() + "说:" + string(message)
		b := []byte(msg)
		// range 方法遍历所有客户端，并发送消息
		for _, conn := range ConnMap {
			conn.Write(b)
		}

	}
}
