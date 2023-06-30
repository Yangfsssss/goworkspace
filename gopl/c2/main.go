package main

import (
	"bufio"
	"fmt"
	"os"
	"popcount"
	"strconv"
	"tempconv"
	"weightconv"
)

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

	// 作用域
	//：声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时属性
	//：一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时概念。
	//：句法块/（隐式）词法块/全局词法块
	// for/if/switch/...
	// 包级语法域：在包的源文件之间共享
}

var globalA *int
var globalB *int

func main() {
	// testPointer()

	// fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	// fmt.Println(tempconv.CToF(tempconv.BoilingC))

	// cf()
	// gpo()
	fmt.Println(popcount.PopCount(256))
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

func cf() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

func gpo() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, input.Text())
		}
	}

	for _, arg := range args {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		g := weightconv.Gram(w)
		p := weightconv.Pound(w)
		o := weightconv.Ounce(w)
		fmt.Printf("%s = %s, %s = %s,%s = %s\n", g, weightconv.GToP(g), p, weightconv.PToO(p), o, weightconv.OToP(o))
	}
}
