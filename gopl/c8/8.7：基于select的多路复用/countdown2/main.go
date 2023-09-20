package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	done := make(chan struct{})

	go func() {
		// 阻塞 goroutine，直到从标准输入读取到一个字节为止
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Connecting countdown. Press return to abort.")

	go func() {
		tick := time.Tick(1 * time.Second)
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
			// 每次执行阻塞1s
			<-tick
		}

		done <- struct{}{}
	}()

	// 阻塞select，直到某个分支被执行
	// 然后main结束
	select {
	case <-abort:
		fmt.Println("Launch aborted!")
	case <-done:
		launch()
	}
}

func launch() {
	fmt.Println("Launch!")
}
