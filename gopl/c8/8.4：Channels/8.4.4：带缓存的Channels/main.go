package main

import (
	"c1fetch"
	"fmt"
)

func main() {
	// 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，
	// 接收操作则是从队列的头部删除元素。
	// 如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。
	// 相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

	// 无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存channel，这些操作是解耦的
	ch := make(chan string, 3)

	// 无阻塞发送
	ch <- "a"
	ch <- "b"
	ch <- "c"
	// 阻塞
	ch <- "d"

	// 获取容量
	fmt.Println(cap(ch))
	// 获取元素个数
	fmt.Println(len(ch))
}

func mirroredQueue() string {
	response := make(chan string, 3)

	go func() {
		value, err := c1fetch.SingleFetch("asia.gopl.io")
		if err != nil {
			response <- err.Error()
		}
		response <- string(value)
	}()

	go func() {
		value, err := c1fetch.SingleFetch("europe.gopl.io")
		if err != nil {
			response <- err.Error()
		}
		response <- string(value)
	}()

	go func() {
		value, err := c1fetch.SingleFetch("america.gopl.io")
		if err != nil {
			response <- err.Error()
		}
		response <- string(value)
	}()

	return <-response
}
