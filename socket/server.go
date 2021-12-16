package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listner, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listner.Accept()
	if err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, 1000)
	cnt, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("读取：", string(buffer), "长度：", cnt)
	res := strings.ToUpper(string(buffer))
	cnt, err = conn.Write([]byte(res))
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
