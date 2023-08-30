package main

import "fmt"

type v interface {
	a() bool
	b() int
}

type myType struct{}

func (m *myType) a() bool {
	return true
}

func (m *myType) b() int {
	return 42
}

func main() {
	myVar := &myType{}
	fmt.Println(myVar.a())
	fmt.Println(myVar.b())
}
