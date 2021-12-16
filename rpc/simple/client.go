package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type RectParams struct {
	Weight, Height int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("连接异常", err)
	}
	var res int
	callErr := client.Call("RectCal.Area", RectParams{10, 20}, &res)
	if callErr != nil {
		log.Println("调用异常", callErr)
	}
	fmt.Println("面积为", res)
	callErr = client.Call("RectCal.Circle", RectParams{10, 20}, &res)
	if callErr != nil {
		log.Println("调用异常", callErr)
	}
	fmt.Println("面积为", res)
	callErr = client.Call("RectCal.Other", RectParams{10, 20}, &res)
	if callErr != nil {
		log.Println("调用异常", callErr)
	}
	log.Println("End")
}
