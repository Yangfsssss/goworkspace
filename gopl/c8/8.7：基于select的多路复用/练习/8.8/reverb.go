package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// 练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
func main() {
	listen()
}

func listen() {
	listen, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// abort := make(chan struct{})
	keep := make(chan struct{}, 1)

	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Println("Received:", input.Text())
		keep <- struct{}{}
		go echo(c, input.Text(), 2*time.Second)
	}

	ticker := time.NewTicker(1 * time.Second)

	select {
	case <-ticker.C:
		fmt.Println("Connection aborted!")
		c.Close()
		return
	case <-keep:
		fmt.Println("Connection kept!")
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
