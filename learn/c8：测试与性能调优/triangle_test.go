package main

import (
	"math"
	"testing"
)

// 传统测试
//：测试数据和测试逻辑混在一起
//：出错信息不明确
//：一旦一个数据出错测试全部结束

// 表格驱动测试
//：分离的测试数据和测试逻辑
//：明确的出错信息
//：可以部分失败
//：go语言的语法使得我们更易实践表格驱动测试

func calTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 0},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calTriangle(%d,%d); "+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
