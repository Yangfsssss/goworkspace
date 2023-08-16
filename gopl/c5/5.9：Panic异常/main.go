package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// defer函数的调用在释放堆栈信息之前
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte

	// 但只有在panic时才能访问到f的堆栈
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

// defer stack
//：defer 3
//：defer 2
//：defer 1

// f(3)
// f(2)
// f(1)

// panic之后函数中断执行，之后的defer也不会入栈
func testDeferAfterPanic() {
	defer func() {
		fmt.Println("defer statement")
	}()

	fmt.Println("before panic")
	panic("something went wrong")

	defer func() {
		fmt.Println("after defer statement")
	}()
}
