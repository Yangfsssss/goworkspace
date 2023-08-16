package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// 对匿名函数采用defer机制，可以使其观察函数的返回值。
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

// 修改函数返回给调用者的返回值
// defer只能访问具名的返回值
// defer的执行时机在return之后，返回值之前
// 在Go的panic机制中，延迟函数的调用在释放堆栈信息之前
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	// 使用defer关闭
	defer func() {
		if closeErr := f.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	// if closeErr := f.Close(); closeErr == nil {
	// 	err = closeErr
	// }

	return local, n, err
}

func main() {
	// fmt.Println(double(3))
	fmt.Println(triple(3))
}
