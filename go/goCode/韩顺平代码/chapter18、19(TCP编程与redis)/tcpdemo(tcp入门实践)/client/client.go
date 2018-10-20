package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8881")
	if err != nil {
		fmt.Println("client dial err=", err)
		return 
	}
	//功能一：客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]

	for {

		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		//如果用户输入的是 exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出..")
			break
		}

		//再将line 发送给 服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)	
		}
	}
	

}

/*
客户端功能:
	1. 编写一个客户端端程序，能链接到 服务器端的 8888 端口
	2. 客户端可以发送单行数据，然后就退出
	3. 能通过终端输入数据(输入一行发送一行), 并发送给服务器端 []
	4. 在终端输入 exit,表示退出程序.
*/