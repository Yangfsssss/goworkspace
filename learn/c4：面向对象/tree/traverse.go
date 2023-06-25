package tree

// 为结构定义的方法必须放在同一个包内
// 可以是不同文件

// 如何扩充系统类型或者别人的类型
//：定义别名
//：使用组合

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
