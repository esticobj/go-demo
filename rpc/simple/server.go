package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type RectCal struct {
}

type RectParams struct {
	Weight, Height int
}

func (p *RectCal) Area(params RectParams, res *int) error {
	*res = params.Weight * params.Height
	return nil
}

func (p *RectCal) Circle(params RectParams, res *int) error {
	*res = (params.Weight + params.Height) * 2
	return nil
}
func main() {
	rect := new(RectCal)
	regErr := rpc.Register(rect)
	if regErr != nil {
		log.Fatal("注册服务异常", regErr)
	}
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("监听服务异常", err)
	}

}
