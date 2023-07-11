package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func maps() {
	// map[K]V
	// K必须是支持==比较运算符的类型

	// 创建空的map
	emptyMap := map[string]int{}
	fmt.Println(emptyMap)

	// 查找失败返回zero
	fmt.Println(emptyMap["age"]) // 0

	// 禁止对map元素取址,原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效
	// a := &emptyMap["age"]

	// map类型的零值是nil，向nil值的map存入元素会导致panic异常
	fmt.Println(emptyMap == nil) // true
	emptyMap["a"] = 1            // panic: assignment to entry in nil map

	// 使用第二个返回值判断元素是否存在
	age, ok := emptyMap["age"]
	if ok {
		fmt.Println(age)
	}
}

func sortMap() {
	ages := map[string]int{
		"carol": 20,
		"dave":  21,
		"alice": 18,
		"bob":   19,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	fmt.Println("before sort", names)
	sort.Strings(names)
	fmt.Println("after sort", names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, vx := range x {
		if vy, ok := y[k]; !ok || vy != vx {
			return false
		}
	}

	return true
}

func dedup() {
	// 这种忽略value的map当作一个字符串集合
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func sliceKey() {
	var m = make(map[string]int)

	k := func(list []string) string {
		return fmt.Sprintf("%q", list)
	}

	Add := func(list []string) {
		m[k(list)]++
	}

	Count := func(list []string) int {
		return m[k(list)]
	}

	Add([]string{"a", "b", "c"})
	Add([]string{"a", "b", "c"})
	Add([]string{"a", "b", "c"})
	Add([]string{"d", "e", "f"})
	Add([]string{"d", "e", "f"})
	fmt.Println(Count([]string{"a", "b", "c"}))
	fmt.Println(Count([]string{"d", "e", "f"}))
}

func charCount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charCount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func graph() {
	var graph = make(map[string]map[string]bool)

	addEdge := func(from, to string) {
		edges := graph[from]
		if edges == nil {
			edges = make(map[string]bool)
			graph[from] = edges
		}

		edges[to] = true
	}

	hasEdge := func(from, to string) bool {
		return graph[from][to]
	}

	addEdge("a", "b")
	addEdge("a", "c")
	fmt.Println(graph)
	fmt.Println(hasEdge("a", "c"))
	fmt.Println(hasEdge("d", "c"))
	fmt.Println(hasEdge("d", "c"))
}

func newCharCount() {
	counts := map[string]int{
		"字母": 0,
		"数字": 0,
		"汉字": 0,
	}

	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "newCharCount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r) && !unicode.Is(unicode.Han, r):
			counts["字母"]++
		case unicode.Is(unicode.Han, r):
			counts["汉字"]++
		case unicode.IsDigit(r):
			counts["数字"]++
		}
	}

	fmt.Printf("type\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func wordFreq() {
	counts := make(map[string]int)
	// invalid := 0

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}

	fmt.Printf("word\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func main() {
	// sortMap()
	// dedup()
	// sliceKey()
	// charCount()
	// graph()

	// newCharCount()
	wordFreq()
}
