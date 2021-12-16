package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	cnt, err := conn.Write([]byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("发送数据成功，长度：", cnt)
	buffer := make([]byte, 1000)
	cnt, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("收到服务器数据：", string(buffer), "长度：", cnt)
	conn.Close()
}
