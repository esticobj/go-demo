package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Cal struct {
}

type CalParams struct {
	Num1, Num2 float64
}

func (p *Cal) Multiply(params CalParams, res *float64) error {
	*res = params.Num1 * params.Num2
	return nil
}

func main() {
	cal := new(Cal)
	err := rpc.Register(cal)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
