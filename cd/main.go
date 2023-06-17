package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	str := "中国"
	for _, k := range str {
		fmt.Printf("%c", k)
	}

	//var s string
	s := "中国"
	for _, item := range s {
		fmt.Printf("%c", item)
	}
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go process(conn)

	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}
