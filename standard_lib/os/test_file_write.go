package main

import (
	"fmt"
	"os"
)

func write() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	f.Write([]byte("Hello golang"))
	f.Close()
}

func writeString() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	f.WriteString(" Hello java")
	f.Close()
}

func writeAt() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	f.WriteAt([]byte("test"), 3)
	f.Close()
}

func main() {
	// write()
	// writeString()
	writeAt()
}
