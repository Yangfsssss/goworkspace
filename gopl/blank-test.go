package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type AreaCalculator interface {
	Area() float64
}

type Square float64

func (s Square) Area() float64 {
	return float64(s * s)
}

func main() {
	var shape AreaCalculator
	shape = Square(5)
	fmt.Println(shape.Area()) // 输出: 25

	var z io.Writer
	z.Write([]byte("hello"))
	fmt.Printf("%T\n", z)
	z = os.Stdout
	fmt.Printf("%T\n", z)
	z = new(bytes.Buffer)
	fmt.Printf("%T\n", z)
}
