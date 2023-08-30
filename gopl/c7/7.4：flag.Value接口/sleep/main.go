package main

import (
	"flag"
	"fmt"
	"time"
)

// flag.Duration可以将命令行参数中表示时间间隔的字符串值解析为 time.Duration 类型
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	// 当在程序中使用 flag 包定义了命令行参数后，需要调用 flag.Parse() 函数来解析命令行参数并将其赋值给对应的变量。
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println("Done.")
}
