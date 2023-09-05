package main

import (
	"bytes"
	"c1fetch"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 练习 7.17： 扩展xmlselect程序以便让元素不仅可以通过名称选择，也可以通过它们CSS风格的属性进行选择。

func main() {
	// launchCustomHtml()
	xmlSelect()
}

func xmlSelect() {
	if len(os.Args) < 2 {
		log.Fatal("Missing URL and selector arguments.")
	}

	url := os.Args[1]
	selector := strings.Join(os.Args[2:], " ")

	output, err := c1fetch.SingleFetch(url)
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(output)
	dec := xml.NewDecoder(r)

	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			// fmt.Println("match = ", matchesSelector(tok, selector))
			if matchesSelector(tok, selector) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, strings.Split(selector, " ")) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func matchesSelector(e xml.StartElement, selector string) bool {
	parts := strings.Split(selector, " ")
	// fmt.Println("e = ", e)
	// fmt.Println("parts = ", parts)

	for _, part := range parts {
		// select by id

		// fmt.Println("part = ", part)
		if strings.HasPrefix(part, "id=") {
			// fmt.Println("HasPrefix:", strings.HasPrefix(part, "id"))
			id := strings.TrimPrefix(part, "id=")
			// fmt.Println("id = ", strings.TrimPrefix(part, "id"))
			for _, attr := range e.Attr {
				if attr.Name.Local == "id" && attr.Value == id {
					return true
				}
			}
			// select by class
		} else if strings.HasPrefix(part, "class=") {
			class := strings.TrimPrefix(part, "class=")
			for _, attr := range e.Attr {
				if attr.Name.Local == "class" {
					classes := strings.Split(attr.Value, " ")
					for _, c := range classes {
						if c == class {
							return true
						}
					}
				}
			}
			// select by name
		} else if part == e.Name.Local {
			return true
		}
	}

	return false
}

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

// func launchCustomHtml() {
// 	htm, err := os.ReadFile("./7.17.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write(htm)
// 	}))
// }
