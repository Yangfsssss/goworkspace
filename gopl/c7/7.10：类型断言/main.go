package main

import (
	"byteCounter"
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout

	// 类型断言：x.(T)
	// 第一种情况：如果T是具体类型，类型断言检查x的动态类型是否和T相等，
	// 如果是，类型断言的结果是x的动态值，类型是T
	f := w.(*os.File) // success: f == os.Stdout
	// 如果不是，抛出panic
	c := w.(*bytes.Buffer) // panic:interface hold *os.File,not *bytes.Buffer

	// 第二种情况：如果T是接口类型，类型断言检查x的动态类型是否满足T。
	// 如果是，则结果仍是具有相同动态类型和值部分的接口值，此时结果的类型转变为T
	// 换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），
	// 但是它保留了接口值内部的动态类型和值的部分。
	var y io.Writer
	y = os.Stdout
	ry := y.(io.ReadWriter) // success:*os.File has both Read and Write
	y = new(byteCounter.ByteCounter)
	ry = y.(io.ReadWriter) // panic:*byteCounter.ByteCounter has no Read method

	// 如果x的值是nil，无论T是什么类型，类型断言都会失败
	y = ry             // io.ReadWriter is assignable to io.Writer
	y = ry.(io.Writer) // fails only if ry == nil

	// 经常地，对一个接口值的动态类型我们是不确定的，并且我们更愿意去检验它是否是一些特定的类型。
	// 使用ok可以避免类型断言失败引发的panic
	// 如果失败，第一个结果等于被断言类型的零值
	var z io.Writer = os.Stdout
	f, ok := z.(*os.File)      // success:ok,f == os.Stdout
	b, ok := z.(*bytes.Buffer) // failure:!ok z != *bytes.Buffer,b == nil

	// 声明了一个同名的新的本地变量，外层原来的w不会被改变
	if w, ok := w.(*os.File); ok {
		// ...
	}

	// 补充chatGPT的解释：
	// 当我们使用一个接口类型的变量时，变量的值会被分为两个部分：动态类型和动态值。
	// 动态类型指的是变量实际所持有的值的类型，而动态值指的是变量实际所持有的值本身。
	// 在类型断言中，我们需要检查一个变量的动态类型是否符合某个特定的类型，
	// 如果符合，我们就可以使用类型断言后的变量来访问该类型的方法或属性。
	//如果我们使用类型断言将一个接口类型转换为一个具体类型，那么类型断言会从变量的动态值中提取具体的值，
	// 并将其转换为所需的类型。这意味着我们可以直接使用类型断言后的变量来访问该类型的方法或属性，因为变量的动态值已经是具体类型了。
	// 如果我们使用类型断言将一个接口类型转换为另一个接口类型，那么类型断言并不会从变量的动态值中提取具体的值，
	// 而是保留了原始变量的动态类型和动态值。这意味着我们可以使用类型断言后的变量来访问新的接口类型的方法或属性，但是变量的动态值仍然是原始类型的值。
}
