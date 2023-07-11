package main

import (
	"fmt"
	"time"
)

func structs() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee
	// 直接赋值
	dilbert.Salary -= 5000
	// 地址访问
	position := &dilbert.Position
	*position = "Senior" + *position

	// 点操作符结合指针
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proactive team player)"
	// 等于
	(*employeeOfTheMonth).Position += "(proactive team player)"

	var manager Employee
	manager.ID = 5
	manager.Position = "sh"
	dilbert.ManagerID = manager.ID

	var employees = [...]Employee{dilbert, manager}

	EmployeeById := func(id int) *Employee {
		for _, employee := range employees {
			if employee.ID == id {
				return &employee
			}
		}
		return nil
	}

	fmt.Println(EmployeeById(dilbert.ManagerID).Position)
	id := dilbert.ID
	EmployeeById(id).Salary = 0

	// 结构体字面值
	type Point struct {
		X, Y int
	}

	p1 := Point{1, 2} // 按顺序赋值
	p2 := Point{X: 1} // 按key赋值,没赋值到的key为zero

	Scale := func(p Point, factor int) Point {
		return Point{p.X * factor, p.Y * factor}
	}
	fmt.Println(Scale(p1, 2))
	fmt.Println(Scale(p2, 2))

	// 考虑效率的话，较大的结构体通常会用指针的方式传入和返回
	Bonus := func(e *Employee, percent int) int {
		return e.Salary * percent / 100
	}
	fmt.Println(Bonus(&dilbert, 10))

	// 如果要在函数内部修改结构体成员的话，用指针传入是必须的
	AwardAnnualRaise := func(e *Employee) {
		e.Salary = e.Salary * 105 / 100
	}
	AwardAnnualRaise(&dilbert)
	fmt.Println(dilbert)

	// 因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址：
	// pp := &Point{1, 2}
	// 等价于
	qq := new(Point)
	*qq = Point{1, 2}

	// 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的
}

func structs2() {
	// 结构体嵌入和匿名成员
	type Point struct {
		X, Y int
	}

	// 只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员
	// 匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针
	// 类似...和Object.assign()
	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	// w.Circle.Center.X = 8
	w.X = 8
	// w.Circle.Center.Y = 8
	w.Y = 8
	// w.Circle.Radius = 5
	w.Radius = 5
	w.Spokes = 20

	// 结构体字面值没有简短表示匿名成员的语法
	// w = Wheel{8, 8, 5, 20} // cannot use 8 (untyped int constant) as Circle value in struct literal
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // unknown field X in struct literal of type Wheel

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	// 或者
	w = Wheel{
		Circle: Circle{
			Point:  Point{8, 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	// %v参数包含的#副词，它表示用和Go语言类似的语法打印值。
	// 对于结构体类型来说，将包含每个成员的名字。
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)

	// 因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。
	// 同时，因为成员的名字是由其类型隐式地决定的，所以匿名成员也有可见性的规则约束。

	// 任何命名的类型都可以作为结构体的匿名成员
	// 简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法。
	// 实际上，外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型导出的全部的方法。
	// 这个机制可以用于将一些有简单行为的对象组合成有复杂行为的对象。组合是Go语言中面向对象编程的核心，我们将在6.3节中专门讨论。
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(appendValues(values[:0], root))
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// 新建一个树
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func main() {
	// structs()
	structs2()
	// Sort([]int{9, 3, 6, 1, 4})
}
