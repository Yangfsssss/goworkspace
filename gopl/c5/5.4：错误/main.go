package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// 将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息。
// 如果导致失败的原因只有一个，额外的返回值可以是一个布尔值

// 有少部分函数在发生错误时，仍然会返回一些有用的返回值

// 使用该函数添加额外的前缀上下文信息到原始错误信息。
// 当错误最终由main函数处理时，错误信息应提供清晰的从原因到后果的因果链

// 要注意错误信息表达的一致性，即相同的函数或同包内的同一组函数返回的错误在构成和处理方式上是相似的

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}

		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// 如果错误发生后，程序无法继续运行，我们就可以采用第三种策略：输出错误信息并结束程序。
// 需要注意的是，这种策略只应在main中执行。
// 对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。

// Go中大部分函数的代码结构几乎相同，首先是一系列的初始检查，防止错误发生，之后是函数的实际逻辑
func main() {
	if err := WaitForServer("http://bad.gopl.io"); err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		log.Fatalf("Site is down: %v\n", err)
		os.Exit(1)
	}

	// 输出错误信息，不中断
	log.Println("Site is up.")

	// 标准错误流输出错误信息
	// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)

	// 直接忽略错误
	// fmt.Errorf("Site is down: %v\n", err)

	// 任何由文件结束引起的读取失败都返回同一个错误——io.EOF
	in := bufio.NewReader(os.Stdin)
	for {
		_, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("read failed: %v\n", err)
		}
	}
}
