package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Connecting countdown.")

	// 创建一个类型为 <-chan time.Time 的channel
	// 每隔1s发送一个当前的时间戳
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		// 每次执行阻塞1s
		<-tick
	}

	launch()
}

func launch() {
	fmt.Println("Launch!")
}
