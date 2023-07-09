package main

import (
	"fmt"
)

// panic
//：停止当前函数执行
//：一直向上返回，执行每一层的defer
//：如果没有遇见recover，程序退出

// recover
//：仅在defer调用中使用
//：获取panic的值
//：如果无法处理，可重新panic

// error vs panic
//：尽量用error
//：意料之中的：使用error。如：文件打不开
//：意料之外的：使用panic。如：数组越界

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(fmt.Sprintf("I don't known what to do:%v", r))
		}
	}()

	// b := 0
	// a := 5 / b
	// fmt.Println(a)
	// panic(errors.New("this is an error"))
	panic(123)

}

func main() {
	tryRecover()
}
