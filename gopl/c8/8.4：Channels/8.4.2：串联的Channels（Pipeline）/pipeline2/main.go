package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}

		// 当一个channel被关闭后，再向该channel发送数据将导致panic异常。
		// 当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。

		// 只有当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。
		// 不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。
		// 试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常
		close(naturals)
	}()

	// Squares
	go func() {
		// 使用range依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环。
		for x := range naturals {
			squares <- x * x
		}

		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
