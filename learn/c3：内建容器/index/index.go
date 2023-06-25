package main

import (
	"fmt"
	"unicode/utf8"
)

// 1，数组
// 指针修改原数组
func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main1() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// range类似map，不改变原数组
	// [10]int和[20]int是不同类型
	// go里的数组是值类型
	// 在go语言中一般不直接使用数组
	printArray(&arr1)
	printArray(&arr3)
}

// 2，切片(slice)
// ：slice本身没有数据，是对底层array的一个view
func updateSlice(s []int) {
	s[0] = 100
}

// slice修改原数组
func printArraySlice(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func extendSlice() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]

	// fmt.Println(s1)
	// fmt.Println(s2)

	fmt.Println("arr = ", arr)
	fmt.Printf("s1=%v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
}

func main2() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6]) // 2-6
	fmt.Println("arr[:6] = ", arr[:6])   // 0-6
	s1 := arr[2:]
	fmt.Println("s1 = ", s1) // 2-end
	s2 := arr[:]
	fmt.Println("s2 = ", s2) // all

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("After extendSlice()")
	extendSlice()
}

// 3，切片的操作
// 向slice添加元素
// ：添加元素时如果超过cap，系统会重新分配更大的底层数组
// ：由于值传递的关系，必须接收append的返回值
// ：s = append(s,val)
func addSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // [2, 3, 4, 5]
	s2 := s1[3:5]  // [5,6]
	printSlice(s2)
	s3 := append(s2, 10) // [5,6,10] --> 没有7
	printSlice(s3)
	// s4 and s5 no longer view arr
	s4 := append(s3, 11) // [5,6,10,11] --> 没有7
	printSlice(s4)

	s5 := append(s4, 12) // [5,6,10,11,12] --> 没有7
	printSlice(s5)

	// fmt.Println(s1, s2, s3, s4, s5)

}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}

func testSlice() {
	s := make([]int, 4, 6)
	printSlice(s)

	s1 := append(s, 1)
	printSlice(s)
	printSlice(s1)
}

func main() {
	// testSlice()
	// fmt.Println("After addSlice()")
	addSlice()

	// var s []int // Zero value for slice is nil
	// for i := 0; i < 100; i++ {
	// 	printSlice(s)
	// 	s = append(s, 2*i+1)
	// }
	// fmt.Println(s)

	// s1 := []int{2, 4, 6, 8}
	// printSlice(s1)

	// // 不知道值，只知道类型和长度
	// s2 := make([]int, 16)
	// printSlice(s2)

	// // 不知道值，知道类型、长度和容量
	// s3 := make([]int, 10, 32)
	// printSlice(s3)

	// fmt.Println("Copying slice")
	// copy(s2, s1)
	// printSlice(s2)

	// fmt.Println("Deleting elements from slice")
	// s2 = append(s2[:3], s2[4:]...)
	// printSlice(s2)

	// fmt.Println("Popping from front")
	// front := s2[0]
	// s2 = s2[1:]
	// fmt.Println(front)
	// printSlice(s2)

	// fmt.Println("Popping from back")
	// tail := s2[len(s2)-1]
	// s2 = s2[:len(s2)-1]
	// fmt.Println(tail)
	// printSlice(s2)
}

// 4，Map
// Map的操作
// ：创建：make(map[string]int)
// ：获取元素：map[key]
// ：key不存在时，获得value类型的初始值
// ：用value,ok := map[key]来判断是否存在key

// 遍历Map
//：使用range
//：不保证遍历顺序，如需顺序，需手动对key排序
// ：使用len获得元素个数

// Map的key
// ：map使用哈希表，必须可以比较相等
// ：除了slice，map，function的内建类型都可以作为key
//：如果Struct类型不包含上述字段，也可以作为key

func main4() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key doesn't exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}

// 5，Map例题
// 寻找最长不含有重复字符的子串
// ：对于每一个字母x：
// lastOccurred[x]不存在，或者 < start，无需操作
// lastOccurred[x] >= start -> 更新 start
// 更新 lastOccurred[x]，更新 maxLength
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, c := range []rune(s) {
		if lastI, ok := lastOccurred[c]; ok && lastI >= start {
			start = lastOccurred[c] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[c] = i
	}

	return maxLength
}

func main5() {
	// func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发灰会花飞灰"))
}

// 6，字符和字符串处理
// rune相当于go的char
//：使用range遍历pos，rune对
//：使用utf8.RuneCountInString()获得字符数量
//：使用len获得字节长度
//：使用[]byte()获得字节

// 其他字符串操作
//：Fields，Split，Join
//：Contains，Index
//：ToUpper，ToLower
//：Trim，TrimLeft，TrimRight

// “汉” ---> unicode(6C49) ---> bytes(E6 B1 89) --> utf8.DecodeRune(bytes) == rune("汉")
func main6() {
	s := "Yes汉爱慕课网!" // UTF-8
	fmt.Println(len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, c := range s { // c is a rune
		fmt.Printf("(%d %X)", i, c)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		c, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", c)
	}
	fmt.Println()

	for i, c := range []rune(s) {
		fmt.Printf("(%d %c)", i, c)
	}
	fmt.Println()

}
