package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

// 数组是僵化的类型，因为数组的类型包含了僵化的长度信息
// 而且也没有任何添加或删除数组元素的方法
// 一般使用slice来替代数组
func main() {
	// array()
	// trySha256()
	// fmt.Println(differentCount(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
	makeHash()
}

func array() {
	// 指定索引和对应值列表的方式初始化
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "￡",
		RMB: "￥",
	}

	fmt.Println(symbol)
	fmt.Println(RMB, symbol[RMB])

	// 100个元素的数组r
	r := [...]int{99: -1}
	fmt.Println(r)

	// 如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	// d := [3]int{1, 2}
	// fmt.Println(a == d)
}

func trySha256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println([]byte("x"), []byte("X"))
	// %t副词参数是用于打印布尔型数据，
	// %T副词参数是用于显示一个值对应的数据类型。
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func orZero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func differentCount(c1, c2 [32]byte) (count int) {
	popCount := func(x, y byte) int {
		count := 0
		for i := 0; i < 8; i++ {
			xb := x % 2
			yb := y % 2
			if xb != yb {
				count++
			}

			x = x >> 1
			y = y >> 1
		}

		return count
	}

	for i := range c1 {
		count += popCount(c1[i], c2[i])
	}

	return
}

func makeHash() {
	// usage：
	// go run main.go --method=2
	//Enter a string: abcd
	//1165b3406ff0b52a3d24721f785462ca2276c9f454a116c2b2ba20171a7905ea5a026682eb659c4d5f115c363aa3c79b
	var method string

	flag.StringVar(&method, "method", "1", "编码方式（1 - SHA256，2 - SHA384，3 - SHA512）")
	flag.Parse()

	if method != "1" && method != "2" && method != "3" {
		log.Fatalf("method %s is not supported", method)
	}

	fmt.Println(method)

	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)
	var inputBytes = []byte(input)

	if method == "1" {
		fmt.Printf("%x\n", sha256.Sum256(inputBytes))
	} else if method == "2" {
		fmt.Printf("%x\n", sha512.Sum384(inputBytes))
	} else {
		fmt.Printf("%x\n", sha512.Sum512(inputBytes))
	}
}
