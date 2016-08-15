1. 工程名： chatroom

2. 主要功能：实现多人在线聊天的功能。

3. 包括两个文件：chatserver.go和chatclient.go

4. chatserver.go
   导入net包，用tcp协议监听localhost下7777端口，当有用户连接时，打印用户信息和接入时间，并开始一个客户端发送消息的goroutine线程，接收客户端传进来的消息，打印该消息并将该消息发送给所有客户端。直到读取客户端返回为字段为"exit"||"退出"时，打印当前断开客户端的相关信息，并结束该goroutine。
5. chatclient.go
   用tcp协议的方式连通chatserver,若连接成功，返回“连接成功”。
   提示：输入昵称name（条件:name不为空），若name输入无误，打印登录的name和登录系统的时间time，开始群聊。
   开始一个goroutine(读取用户的聊天信息)：onMessageRecived
   
   a. 当用户输入信息为空时，打印提示消息:输入异常：发送内容不能为空，请重新输入:
   
   b. 当用户输入信息是"exit" 和 "退出"时，打印提示消息：“退出成功，欢迎下次使用！”。
   
   c.否则,输出信息:“[用户名]说：+消息”。并将消息传给服务端。
