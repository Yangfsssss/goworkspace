package main

import "fmt"

// 接口被以两种不同的方式使用。
// 在第一个方式中，以io.Reader，io.Writer，fmt.Stringer，sort.Interface，http.Handler和error为典型，
// 一个接口的方法表达了实现这个接口的具体类型间的相似性，但是隐藏了代码的细节和这些具体类型本身的操作。
// 重点在于方法上，而不是具体的类型上。

// 第二个方式是利用一个接口值可以持有各种具体类型值的能力，将这个接口认为是这些类型的联合。
// 类型断言用来动态地区别这些类型，使得对每一种情况都不一样。
// 在这个方式中，重点在于具体的类型满足这个接口，而不在于接口的方法（如果它确实有一些的话），
// 并且没有任何的信息隐藏。我们将以这种方式使用的接口描述为discriminated unions（可辨识联合）。

func sqlQuote(x interface{}) string {
	// x.(type)是一种类型选择语法，用于判断一个接口类型变量的具体类型。它只能出现在switch语句中。
	// 虽然x的类型是interface{}，但是我们把它认为是一个int，uint，bool，string，和nil值的discriminated union（可识别联合）
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return fmt.Sprintf("'%s'", x)
	default:
		panic(fmt.Sprintf("unsupported type %T:%v", x, x))
	}
}

func main() {}
