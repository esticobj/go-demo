package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type CalParams struct {
	Num1, Num2 float64
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	res := 0.0
	err = client.Call("Cal.Multiply", CalParams{10, 20}, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("res", res)
}
