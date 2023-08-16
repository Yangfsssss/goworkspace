package main

import (
	"fmt"
	"methods"
)

func main() {
	operations()
}

func operations() {
	p := methods.Point{X: 1, Y: 2}
	q := methods.Point{X: 4, Y: 6}

	// 在一个包的API需要一个函数值、且调用方希望操作的是某一个绑定了对象的方法的话，方法“值”会非常实用
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	var origin methods.Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)

	distance := methods.Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*methods.Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)

	// 方法表达式见6.1
}
