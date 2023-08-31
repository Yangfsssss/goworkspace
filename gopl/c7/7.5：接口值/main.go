package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

// 接口值由动态类型和动态值组成
// 具体的类型和它的值被称为接口的动态类型和动态值
// 类型描述符即为接口的动态类型
func main() {
	// io.Writer是一个接口
	// 赋值给w时，w的初始类型和值都为nil
	// 可以通过w != nil来判断接口值是否为空
	var w io.Writer
	w.Write([]byte("hello")) // panic: runtime error: invalid memory address or nil pointer dereference
	// os.Stdout是一个值，类型是*os.File，它实现了io.Writer
	// 所以可以赋值给w
	// 此时w的类型是*os.File，值是os.Stdout
	w = os.Stdout

	// bytes.Buffer是一个结构体，类型是Buffer，也实现了io.Writer
	w = new(bytes.Buffer)

	// 重置了类型和值
	w = nil

	// 一个接口值可以持有任意大的动态值
	// 接口值可以比较。
	//两个接口值相等仅当它们都是nil值，或者它们的动态类型相同并且动态值也根据这个动态类型的==操作相等。
	// 因为接口值是可比较的，所以它们可以用在map的键或者作为switch语句的操作数。
	var x interface{} = time.Now()
	// 如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将它们进行比较就会失败并且panic:
	var y interface{} = []int{1, 2, 3}
	fmt.Println(y == y) // panic: comparing incomparable type []int
	// 在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的panic
	// 同样的风险也存在于使用接口作为map的键或者switch的操作数。只能比较你非常确定它们的动态值是可比较类型的接口值

	// 使用fmt包的%T获得接口值的动态类型
	// 在fmt包内部，使用反射来获取接口动态类型的名称
	var z io.Writer
	fmt.Printf("%T\n", z) // <nil>
	z = os.Stdout
	fmt.Printf("%T\n", z) // *os.File
	z = new(bytes.Buffer)
	fmt.Printf("%T\n", z) // *bytes.Buffer
}

const debug = true

func run() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)

	if debug {
		// ...
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("hello"))
	}
}
