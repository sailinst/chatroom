package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	ip   = "127.0.0.1"
	port = 7777
)

func onMessageRecived(conn *net.TCPConn) {
	read := bufio.NewReader(conn)
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			//quitSemaphore <- true
			break
		}
		fmt.Println(msg)
	}
}
func inputname() string {
	var name string
	fmt.Println("请输入您的用户名,回车加入群聊(如：zhang3)")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("输入有误：用户名不能为空，请重新输入：", err.Error())
		return name
	}
	//fmt.Println("注册完成！")
	return name

}

var name string = ""

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr = &net.TCPAddr{net.ParseIP(ip), port, ""}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("连接服务端失败：", err.Error())
	}
	defer conn.Close()
	fmt.Println("连接成功！")
	//data := make([]byte, 1024)
	//conn.Read(data)
	//fmt.Println(string(data))
	for {
		name = inputname()
		if len(name) != 0 {
			fmt.Println("您好："+name+"登录时间：", time.Now().Format("2006-01-02 15:04:05 PM"))
			break
		}

	}
	fmt.Println("开始群聊:-------------------------------------------------------------------->")
	go onMessageRecived(conn)
	// 聊天
	for {
		//var msg string   这种写法，当输入空格时，消息会断层。
		// 输入聊天内容，如果输入exit||退出,则退出聊天系统。
		var input *bufio.Reader
		input = bufio.NewReader(os.Stdin)
		msg, _ := input.ReadString('\n') // 输入一条消息， 以回车结束。
		//conn.Write([]byte(str0 + name))   // 在这里加上‘|’后 msg初始值不为空了
		//fmt.Scanln(&msg)
		if msg == "\n" {
			fmt.Println("输入异常：发送内容不能为空，请重新输入:")
			continue
		}
		if msg == "exit\n" || msg == "退出\n" {
			fmt.Println("退出成功，欢迎下次使用！")
			break
		}
		fmt.Println("字节数：", len(msg)-1) // -1 是因为读取msg时末尾加的'\n'符号占用一个字节
		fmt.Println(msg)
		str0 := "|"
		str1 := "|说："
		b := []byte([]byte(str0 + name + str1 + msg)) //消息头加上用户名
		//fmt.Println(b)
		conn.Write(b)
	}
}
