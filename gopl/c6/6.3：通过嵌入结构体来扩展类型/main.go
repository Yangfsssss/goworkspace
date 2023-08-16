package main

import (
	"fmt"
	"image/color"
	"methods"
	"sync"
)

func main() {

	operations()
}

type ColoredPoint struct {
	methods.Point // 内嵌结构体
	Color         color.RGBA
}

func operations() {
	var cp ColoredPoint
	cp.X = 1 // 简写形式
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{methods.Point{X: 1, Y: 1}, red}
	var q = ColoredPoint{methods.Point{X: 5, Y: 4}, blue}
	// 内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	// 等价实现
	// 	func (p ColoredPoint) Distance(q Point) float64 {
	//     return p.Point.Distance(q)
	// }

	//	func (p *ColoredPoint) ScaleBy(factor float64) {
	//	    p.Point.ScaleBy(factor)
	//	}
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

// 方法只能在命名类型（像Point）或者指向类型的指针上定义，
// 使用内嵌，可以给匿名struct类型来定义方法
func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()

	return v
}
