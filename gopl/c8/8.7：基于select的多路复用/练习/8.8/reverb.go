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

		handleConn2(conn)
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

func handleConn2(c net.Conn) {
	// 设置超时时间为 10 秒
	timeout := 10 * time.Second

	// 创建一个计时器，用于超时检测
	timer := time.NewTimer(timeout)

	// 创建一个通道，用于接收客户端的输入
	input := make(chan string)

	// 启动一个 goroutine 读取客户端的输入并发送到 input 通道中
	go func() {
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			input <- scanner.Text()
		}
	}()

	for {
		select {
		case msg := <-input:
			// 如果收到客户端的输入，重置计时器
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(timeout)

			// 处理客户端的输入
			fmt.Println("Received:", msg)
			go echo(c, msg, 2*time.Second)

			// 回复客户端
			fmt.Fprintln(c, msg)

		case <-timer.C:
			// 超时，断开连接
			fmt.Println("Connection timed out")
			c.Close()
			return
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
