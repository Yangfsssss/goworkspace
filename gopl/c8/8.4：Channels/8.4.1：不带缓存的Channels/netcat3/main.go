package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// 8.4.1. 不带缓存的Channels
// 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，
// 直到另一个goroutine在相同的Channels上执行接收操作，
// 当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
// 反之，如果接收操作先发生，那么接收者goroutine也将阻塞，
// 直到有另一个goroutine在相同的Channels上执行发送操作。
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		// io.Copy copies from src to dst until either EOF is reached on src or an error occurs.
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	// 关闭读和写方向的网络连接。
	// 关闭网络连接中的写方向的连接将导致server程序收到一个文件（end-of-file）结束的信号。
	// 关闭网络连接中读方向的连接将导致后台goroutine的io.Copy函数调用返回一个“read from closed connection”（“从关闭的连接读”）类似的错误
	// conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	fmt.Println("mustCopy done")
}
