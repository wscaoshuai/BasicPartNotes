package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		fmt.Printf("服务器的等待客户端%s 发送消息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //等待客户端conn发送消息，如果客户端没有write，则阻塞在此处
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return //!
		}
		fmt.Print(string(buf[:n])) //细节：不用Println;要转成string;同时是[:n]
	}
}

func main() {
	fmt.Println("服务器开始监听端口")
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //监听本地端口0.0.0.0:8888
	if err != nil {
		fmt.Println("Listen err=", err)
		return
	}
	defer listen.Close() //及时关闭

	for {
		fmt.Printf("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err=", err)

		} else {
			fmt.Printf("Accept suc con%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//增加一个协程，为客户端服务
		go process(conn)
	}

}
