package main

import "fmt"

// &：取值的指针
// *：取指针的值
// *[type]：指针类型，即&[anyValue]

func main() {
	// 调用指针方法
	//：1
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println("*r", *r)
	//：2
	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println("*pptr", *pptr)
	//：3
	pp := Point{1, 2}
	(&pp).ScaleBy(2)
	fmt.Println("*pp", pp)
	//：special case
	// 对于变量p来说，p.ScaleBy(2)和(&p).ScaleBy(2)是等价的（隐式取指针）
	p.ScaleBy(2)
	// 等同于
	(&p).ScaleBy(2)

	// 对于指针来说，pptr.ScaleBy(2)和(*pptr).ScaleBy(2)是等价的（隐式取值）
	pptr.ScaleBy(2)
	// 等同于
	(*pptr).ScaleBy(2)

	operations()
}

// 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，
// 第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；
// 第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝。

type Point struct {
	X, Y float64
}

// 一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// An IntList is a linked list of integers
// A nil *IntList represents the empty list
type IntList struct {
	Value int
	Tail  *IntList
}

func (list IntList) sum() int {
	if list.Tail == nil {
		return 0
	}

	return list.Value + list.Tail.sum()
}

type Values map[string][]string

func (v Values) get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}

	return ""
}

func (v Values) add(key, value string) {
	if vs := v[key]; len(vs) == 0 {
		v[key] = []string{value}
	} else {
		v[key] = append(vs, value)
	}
}

func operations() {
	m := Values{"lang": {"en"}}
	m.add("item", "1")
	m.add("item", "2")

	fmt.Println(m.get("lang"))
	fmt.Println(m.get("q"))
	fmt.Println(m.get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.get("item"))
	m.add("item", "3")
}
