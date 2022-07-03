package main

import (
	"fmt"
	"io"
	"os"
)

func read() {
	f, err := os.Open("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	buf := make([]byte, 10)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("buf: %v\n", string(buf))
	}
	f.Close()
}

func readAt() {
	f, err := os.Open("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	buf := make([]byte, 10)
	n, _ := f.ReadAt(buf, 3)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("buf: %v\n", string(buf))
	f.Close()
}

func readDir() {
	de, err := os.ReadDir("standard_lib")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	for _, item := range de {
		fmt.Printf("item.IsDir(): %v\n", item.IsDir())
		fmt.Printf("item.Name(): %v\n", item.Name())
	}
}

func readSeek() {
	f, err := os.Open("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	f.Seek(-4, 2)
	buf := make([]byte, 10)
	n, err := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("buf: %v\n", string(buf))
}
func main() {
	// read()
	// readAt()
	// readDir()
	readSeek()
}
