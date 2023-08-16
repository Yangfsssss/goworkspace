package main

import (
	"fmt"
	"funcValue"

	"golang.org/x/net/html"
)

func main() {
	returnANonZeroValueWithoutReturn(3)
	defer func() {
		for v := recover(); v != nil; {
			fmt.Println(v)
		}
	}()
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: // "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	funcValue.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

func returnANonZeroValueWithoutReturn(n int) (result int) {
	defer func() {
		if p := recover(); p != nil {
			panic(p)
		}
	}()

	panic(n)
}
