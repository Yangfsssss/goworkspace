package main

import "testing"

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
