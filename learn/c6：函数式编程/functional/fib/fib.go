package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"tree"
)

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type FuncTreeNode struct {
	node *tree.Node
}

func (myNode *FuncTreeNode) Traverse() {
	node := FuncTreeNode{myNode.node}

	node.TraverseFunc(func(n *FuncTreeNode) {
		n.node.Print()
	})

	fmt.Println()
}

func (myNode *FuncTreeNode) TraverseFunc(f func(*FuncTreeNode)) {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := FuncTreeNode{myNode.node.Left}
	right := FuncTreeNode{myNode.node.Right}

	left.TraverseFunc(f)
	f(myNode)
	right.TraverseFunc(f)
}

func tryNode() {
	var root tree.Node
	fmt.Println(root)

	// 不论地址还是结构本身，一律使用.来访问成员
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	myRoot := FuncTreeNode{&root}
	myRoot.Traverse()
}

func main() {
	// f := fibonacci()

	// printFileContents(f)
	tryNode()
}
