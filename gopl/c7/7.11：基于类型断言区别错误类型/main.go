package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// func IsExist(err error) bool

// func IsNotExist(err error) bool
// func IsPermission(err error) bool

// 处理I/O错误的逻辑可能一个和另一个平台非常的不同
// func IsNotExist(err error) bool {
// 	return strings.Contains(err.Error(), "no such file or directory")
// }

// PathError records an error and a operation and file path that caused it
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

// var ErrNotExist = errors.New("file doesn't exist")
var ErrNotExist = errors.New("no such file or directory")

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		fmt.Println("ok: ", ok)
		err = pe.Err
	}

	fmt.Println(err)
	return err == syscall.ENOENT || err == ErrNotExist
}

func main() {
	_, err := os.Open("/no/such/file")
	// fmt.Println(os.IsNotExist(err))
	err = &PathError{"open", "/no/such/file", err}
	fmt.Println(IsNotExist(err))
}
