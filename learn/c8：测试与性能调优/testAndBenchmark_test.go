package testAndBenchmark

import (
	"math"
	"testing"
)

// go test -coverprofile=c.out
// less c.out
// go tool cover
// go tool cover -html=c.out
// go test -bench .
// go test -bench . -cpuprofile=cpu.out
// go tool pprof cpu.out
// web
// quit

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

// var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	// stores last occurred pos +1
	// 0 means not seen
	// for i := range lastOccurred {
	// 	lastOccurred[i] = 0
	// }
	start := 0
	maxLength := 0
	for i, c := range []rune(s) {
		if lastI, ok := lastOccurred[c]; ok && lastI >= start {
			// if lastI := lastOccurred[c]; lastI > start {
			start = lastOccurred[c] + 1
			// start = lastI
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		// lastOccurred[c] = i + 1
		lastOccurred[c] = i
	}

	return maxLength
}

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese cases
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发灰会花飞灰", 0},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)

		if actual != tt.ans {
			t.Errorf("Got %d for input %s; "+" expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s; "+" expected %d", actual, s, ans)
		}
	}
}
