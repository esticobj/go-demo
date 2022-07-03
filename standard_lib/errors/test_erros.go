package main

import (
	"errors"
	"fmt"
)

func validParam(str string) (string, error) {
	if str == "" {
		return "", errors.New("参数不能为空")
	}
	return str, nil
}

func main() {
	str := ""
	rs, err := validParam(str)
	if err == nil {
		fmt.Printf("rs: %v\n", rs)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}
