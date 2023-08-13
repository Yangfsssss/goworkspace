package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type NodeType int32

type Attribute struct {
	Key, Value string
}

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

func visit(links []string, n *html.Node) []string {
	fmt.Println("n = ", n.Data)
	fmt.Println("links = ", links)

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// c := n.FirstChild
	// if c != nil {
	// 	links = visit(links, c)
	// }

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func findLinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLinks: %v\n", err)
		os.Exit(1)
	}

	links := visit(nil, doc)
	// links := recursiveVisit([]string{}, doc.FirstChild)
	// links := recordSameElementInHtml(doc)
	fmt.Println(links)
	// outline(links, doc)
}

// func outline(stack []string, n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		stack = append(stack, n.Data)
// 		fmt.Println(stack)
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		outline(stack, c)
// 	}
// }

func recursiveVisit(stack []string, n *html.Node) []string {
	if n == nil {
		return stack
	}

	fmt.Println("n = ", n.Data)
	fmt.Println("stack = ", stack)

	fmt.Println(n.Type == html.ElementNode)

	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
	}

	if n.FirstChild == nil {
		return recursiveVisit(stack, n.NextSibling)
	}

	return recursiveVisit(stack, n.FirstChild)
}

func recordSameElementInHtml(n *html.Node) map[string]int {
	counts := make(map[string]int)

	for _, e := range n.Data {
		if counts[string(e)] == 0 {
			counts[string(e)] = 1
		} else {
			counts[string(e)]++
		}
	}

	fmt.Printf("%v\n", counts)

	return counts
}

func main() {
	findLinks()
}
