package main

import (
	"fmt"
	"sort"
)

// 练习 7.10： sort.Interface类型也可以适用在其它地方。
// 编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。
// 假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。
func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}

func main() {
	// sort.IntSlice是一个整型切片类型的包装器，用于为整型切片提供排序功能
	fmt.Println(IsPalindrome(sort.IntSlice{1, 2, 3, 2, 1}))
}
