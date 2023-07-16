package main

import (
	"fmt"
	"time"
)

// 协程 Coroutine
//：轻量级“线程”
//：“非抢占式”多任务处理，由携程主动交出控制权
//：编译器/解释器/虚拟机层面的多任务，非操作系统层面
//：多个协程可能在一个或多个线程上运行
// top
// go -race

// goroutine的定义
//：任何函数只需加上go就能送给调度器运行
//：不需要在定义时区分是否是异步函数
//：调度器在合适的点进行切换
//：使用-race来检测数据访问冲突

// goroutine可能的切换点
//：I/O，select
//：channel
//：等待锁
//：函数调用（有时）
//：runtime.Gosched()
// 只是参考，不能保证切换，不能保证在其他地方不切换

func main() {
	// var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", i)
				// a[i]++
				// 很少用到
				// runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Minute)
	// fmt.Println(a)
}
