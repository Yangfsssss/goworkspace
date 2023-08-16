package methods

import (
	"fmt"
	"math"
)

func main() {
	perim := Path{
		Point{1, 1},
		Point{5, 1},
		Point{5, 4},
		Point{1, 1},
	}
	// 省去包名和函数名
	fmt.Println(perim.Distance())
}

type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing,but as a method of the Point type
// p：方法的接收器（receiver），早期的面向对象语言留下的遗产将调用一个方法称为“向一个对象发送消息”。
// 可以使用其类型的第一个字母
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者interface
type Path []Point

// 对于一个给定的类型，其内部的方法都必须有唯一的方法名，但是不同的类型却可以有同样的方法名
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			// 编译器会根据方法的名字以及接收器来决定具体调用的是哪一个函数
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// 根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(q, p Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}
