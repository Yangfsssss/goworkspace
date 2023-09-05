package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// 练习 7.18： 使用基于标记的解码API，编写一个可以读取任意XML文档并构造这个文档所代表的通用节点树的程序。
// 节点有两种类型：CharData节点表示文本字符串，和 Element节点表示被命名的元素和它们的属性。每一个元素节点有一个子节点的切片。
func main() {
	getXmlNodeTree()
}

type Node interface{} // CharData or *Element

type Book struct {
	Title  string  `xml:"title"`
	Author string  `xml:"author"`
	Year   int     `xml:"year"`
	Price  float64 `xml:"price"`
}

type CharData string

type Element struct {
	Type     xml.Name `xml:"element"`
	Attr     []xml.Attr
	Children []*Node
}

func getXmlNodeTree() {
	xmlFile, err := os.Open("./7.18.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("xmlData = ", xmlData)

	// var root Node
	var root Book
	err = xml.Unmarshal(xmlData, &root)
	fmt.Println("root = ", root)
	if err != nil {
		fmt.Println("解码XML失败:", err)
		return
	}

	printNode(root, 0)
}

func printNode(node Node, indent int) {
	switch n := node.(type) {
	case CharData:
		fmt.Printf("%*sCharData: %s\n", indent*2, "", n)
	case Element:
		fmt.Printf("%*sElement: %s\n", indent*2, "", n.Type)

		for _, child := range n.Children {
			printNode(child, indent+1)
		}
	}
}
