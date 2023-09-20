package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		// 阻塞 goroutine，直到从标准输入读取到一个字节为止
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Connecting countdown. Press return to abort.")

	tick := time.Tick(100 * time.Microsecond)
	for countdown := 1000; countdown > 0; countdown-- {
		fmt.Println(countdown)
		// 非阻塞select，有default分支，没有值就绪时不阻塞
		select {
		case <-tick:
			fmt.Println("tick calling")
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		default:
		}
	}

	launch()
}

func launch() {
	fmt.Println("Launch!")
}
