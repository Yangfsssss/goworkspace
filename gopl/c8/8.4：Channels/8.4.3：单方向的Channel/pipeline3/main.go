package main

import "fmt"

// 当一个channel作为一个函数参数时，它一般总是被专门用于只发送或者只接收。
// 根据channel在函数/goroutine中的使用情况决定
// 类型chan<- int表示一个只发送int的channel，只能发送不能接收，称为只写通道。
// 相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送，称为只读通道。
// 只读通道和只写通道通常成对出现
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// naturals在counter中存储值，所以是只读通道
	go counter(naturals)
	// naturals在squares中释放值，所以是只写通道
	// squares在counter中储存值，所以是只读通道
	go squarer(squares, naturals)
	// squares在counter中释放值，所以是只写通道
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}

	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}

	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
