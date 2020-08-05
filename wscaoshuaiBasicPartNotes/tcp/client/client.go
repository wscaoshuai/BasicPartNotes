package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.199.142:8888")
	if err != nil {
		fmt.Println("client dial err= ", err)
		return
	}
	fmt.Println("conn scuess=", conn)

	reader := bufio.NewReader(os.Stdin) //客户端可以发送单行数据，然后退出(os.Stdin代表标准输入)

	for {
		line, err := reader.ReadString('\n') //从终端读取一行数据，并且发送给服务器
		if err != nil {
			fmt.Println("readerString err=", err)
		}

		line = strings.Trim(line, "\r\n") //如果用户输入exit就退出
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		n, err := conn.Write([]byte(line + "\n")) //再将line发送给服务器
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Printf("客户端发送了%d字节数据,并退出\n", n)
	}
}
