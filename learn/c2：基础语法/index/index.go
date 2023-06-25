package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

// 1，变量定义
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// variableZeroValue returns the zero value of an integer and an empty string.
//
// No parameters.
// No return values.
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// variableInitialValue prints out the initial values of declared variables.
//
// There are no parameters.
// Does not return anything.
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// variableTypeDeduction is a Go function that demonstrates the language's
// ability to automatically deduce the type of a variable based on the value
// assigned to it. It takes no parameters and has no return type since it
// simply prints the values of four variables to the console.
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, "def", true
	fmt.Println(a, b, c, s)
}

// variableShorter shortens the variable names and prints them.
//
// This function does not have any parameters.
// It does not return anything.
func variableShorter() {
	a, b, c, s := 3, 4, "def", true
	b = 5
	fmt.Println(a, b, c, s)
}

// main is the entry point of the program.
//
// It does not have any parameters or return values.
func main1() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	fmt.Println(aa, ss, bb)
}

// 2，内建变量类型
// bool,string
// (u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr,
// byte,rune
// float32,float64,complex64,complex128

func euler() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

func main2() {
	// euler()
	triangle()
}

// 3，常量和枚举
// const 数值可作为各种类型使用
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int

	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main3() {
	consts()
	enums()
}

// 变量定义要点回顾
//：变量类型写在变量名之后
//：编译器可推测变量类型
//：没有char，只有rune
//：原生支持复数类型

// 4，条件语句
// if
func contents() {
	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)

	// if
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// }

	// 类for(;;)形式
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Println(contents)
}

// switch
func grade(score int) string {
	g := ""
	// switch后可以没有表达式，类似switch(true)
	// 不需要break
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

// 定点数https://zhuanlan.zhihu.com/p/338588296
// 浮点数https://zhuanlan.zhihu.com/p/339949186

func main4() {
	contents()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(101), // panic
		grade(-3),  // panic
	)
}

// 5，循环
// for
// ：条件里不需要括号
// ：可以省略初始条件、递增条件、递增表达式
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	// 没有初始和递增条件，相当于while
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 便捷的死循环
	for {
		fmt.Println("forever")
	}
}

func main5() {
	// fmt.Println(
	// 	convertToBin(5),
	// 	convertToBin(13),
	// 	convertToBin(72387885),
	// 	convertToBin(0),
	// )

	printFile("abc.txt")
	forever()
}

// 6，函数
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		// 两个返回值的一个常用场景
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s  with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 整数运算
func div(a, b int) (q, r int) {
	q = a / b
	fmt.Println("q", q)
	// r = a % b
	return a / b, a % b
}

// 浮点数运算
func divFloat64(c, d float64) (m float64) {
	m = c / d
	fmt.Println("m", m)
	return m
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表，即剩余参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func main6() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// fmt.Println(eval(3, 4, "/"))
	// d, f := div(13, 3)
	// fmt.Println(d, f)
	// m := divFloat64(3, 4)
	// fmt.Println(m)

	// fmt.Println(apply(pow, 3, 4))

	// 匿名函数作为参数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}

// 7，指针
// ：指针不能运算
// ：只存在值传递
func pointer() {
	var a int = 2
	var pa *int = &a

	*pa = 3
	fmt.Println(a)
}

// 指针的示例：值交换
func swap(a, b *int) {
	*b, *a = *a, *b
}

// 正常版
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	// pointer()

	a, b := 3, 4
	// swap(&a, &b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
