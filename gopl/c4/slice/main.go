package main

import "fmt"

func main() {
	// slices()
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// reserve(a[:])
	// fmt.Println(a)

	// makeSlices()

	// tryAppend()
	// tryAppendInt()

	// tryNonempty()

	// tryRemove()

	// tryTrialReverse()
	// tryTrialRotate()
	tryRemoveSameStr()
}

func slices() {
	months := [...]string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	Q2 := months[4:7]     // len = 3,cap = 9
	summer := months[6:9] // len = 3,cap = 7
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// 切片操作超出cap(s)的上限将导致一个panic异常，
	// 但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大：
	// fmt.Println(summer[:20]) // panic: runtime error: slice bounds out of range [:20] with capacity 7
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

	// 除了和nil相等比较外，一个nil值的slice的行为和其它任意0长度的slice一样

	// 通常我们并不知道append调用是否导致了内存的重新分配
	// 也不能确认新的slice和原始的slice是否引用的是相同的底层数组空间
	// 同样，我们不能确认在原先的slice上的操作是否会影响到新的slice

	// slice并不是一个纯粹的引用类型，它实际上是一个类似下面结构体的聚合类型：
	type IntSlice struct {
		ptr      *int
		len, cap int
	}
}

func reserve(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func makeSlices() {
	a := make([]int, 10) // 整个数组的view
	fmt.Println(a)
	b := make([]int, 10, 20) // 数组的前len个元素，但是容量将包含整个数组
	fmt.Println(b)
}

func tryAppend() {
	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)
	if zLen <= cap(x) {
		// There is room to grow.Extend the slice.
		z = x[:zLen]
	} else {
		// There is insufficient space.Allocate a new array.
		// Grow by doubling,for amortized linear complexity.
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z[len(x):], y) // a build-in function
	}

	return z
}

func tryAppendInt() {
	var x, y []int
	// for i := 0; i < 10; i++ {
	y = appendInt(x, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("cap=%d\t%v\n", cap(y), y)
	// x = y
	// }
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func tryNonempty() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func stackOperations() {
	stack := []string{}
	stack = append(stack, "a")   // push
	top := stack[len(stack)-1]   // top of stack
	stack = stack[:len(stack)-1] // pop

	fmt.Println(top)
}

func removeSequence(slice []string, i int) []string {
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])

	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeNonSequence(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func tryRemove() {
	s := []int{5, 6, 7, 8, 9}
	// fmt.Println(removeSequence(strconv.I s, 2))
	fmt.Println(removeNonSequence(s, 2))
}

// trial  -------------------------------------------------------------------
// *：对指针操作，取得值
// &：对值操作，取得指针
func trialReverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func tryTrialReverse() {
	s := []int{5, 6, 7, 8, 9}
	trialReverse(&s)
	fmt.Println(s)
}

func trialRotate(r []int) {
	s := r[:]
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func tryTrialRotate() {
	s := []int{5, 6, 7, 8, 9}
	trialRotate(s)
	fmt.Println(s)
}

func removeSameStr(s []string) []string {
	r := s[:]
	for i, j := 0, 1; j <= len(r); {
		if r[i] == r[j] {
			r = removeSequence(r, j)
			fmt.Println(r)
		} else {
			j++
		}
		i++
	}

	return r
}

func tryRemoveSameStr() {
	s := [...]string{"a", "b", "b", "c", "c", "d"}
	fmt.Println(removeSameStr(s[:]))
}
