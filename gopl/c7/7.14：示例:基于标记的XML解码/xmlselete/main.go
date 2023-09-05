package main

import (
	"bytes"
	"c1fetch"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("./xml.html")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
	// 	os.Exit(1)
	// }
	// dec := xml.NewDecoder(file)

	output := c1fetch.Fetch()
	r := bytes.NewReader(output[os.Args[1]])
	dec := xml.NewDecoder(r)

	selector := strings.Join(os.Args[2:], " ")
	fmt.Println("selector", selector)

	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.(xml.StartElement).Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			// if containsAll(stack, os.Args[1:]) {
			if containsAll(stack, strings.Split(selector, " ")) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y,in order
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}

		x = x[1:]
	}

	return false
}
