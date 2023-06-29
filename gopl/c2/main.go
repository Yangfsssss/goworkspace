package main

import "fmt"

func names() {
	// 关键字
	// break      default       func     interface   select
	// case       defer         go       map         struct
	// chan       else          goto     package     switch
	// const      fallthrough   if       range       type
	// continue   for           import   return      var

	// 常量/类型/函数
	//内建常量: true false iota nil

	//内建类型: int int8 int16 int32 int64
	//									uint uint8 uint16 uint32 uint64 uintptr
	//									float32 float64 complex128 complex64
	//									bool byte rune string error

	//内建函数: make len cap new append copy close delete
	//									complex real imag
	//									panic recover

	// 声明种类
	// var、const、type、func

	// new()
}

var globalA *int
var globalB *int

func main() {
	testPointer()
}

func testPointer() {
	x := 10

	// 基础值的指针化
	globalA = &x
	globalB = &x
}

// 计算最大公约数
func gcd(x, y int) int {
	for y != 0 {
		// 参数重新赋值
		x, y = y, x%y
	}

	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	// 类型转换，而非函数调用
	// 对每一个类型T，都有对应的类型转换操作T(x)
	// 只有两个类型的底层基础类型相同时才能进行
	// 数值/字符串/特定类型的slice也可以进行转换
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 方法定义
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
