package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构定义方法
// 值传递
func (node Node) Print() {
	fmt.Println(node.Value)
}

// 指针传递
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored.")
		return
	}
	node.Value = value
}

// 值接收者 vs 指针接收者
//：要改变内容必须使用指针接收者
//：结构过大也考虑使用指针接收者
//：考虑一致性：如有指针接收者，最好都是指针接收者
//：值接收者是go语言特有（js是指针）
//：值/指针接收者均可接收值/指针

// 工厂函数
func CreateNode(value int) *Node {
	// 局部变量
	return &Node{Value: value}
}
