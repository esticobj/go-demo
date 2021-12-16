package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func mkdir() {
	err := os.MkdirAll("a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func rmFile() {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func rmdir() {
	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func rename() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
	os.Chdir("D:/")
	dir, err = os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

func readFile() {
	b, err := os.ReadFile("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("string(b): %v\n", string(b))
	}
}

func writeFile() {
	err := os.WriteFile("test2.txt", []byte("hello 测试"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func main() {
	// createFile()
	// mkdir()
	// rmFile()
	// rmdir()
	// rename()
	// pwd()
	// readFile()
	writeFile()
}
