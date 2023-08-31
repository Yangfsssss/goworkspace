package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

// 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：
// 主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
// 定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。
type TableSorter struct {
	Table     []Row
	Clicks    []int
	Ascending bool
}

type Row struct {
	name  string
	age   int
	grade float64
}

func (t TableSorter) Len() int {
	return len(t.Table)
}

func (t TableSorter) Less(i, j int) bool {
	for _, click := range t.Clicks {
		switch click {
		case 0:
			return t.Table[i].name < t.Table[j].name
		case 1:
			return t.Table[i].age < t.Table[j].age
		case 2:
			return t.Table[i].grade < t.Table[j].grade
		}
	}

	return false
}

func (t TableSorter) Swap(i, j int) {
	t.Table[i], t.Table[j] = t.Table[j], t.Table[i]
}

func printTable(table []Row) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "name", "age", "grade")
	fmt.Fprintf(tw, format, "-----", "-----", "-----")
	for _, t := range table {
		fmt.Fprintf(tw, format, t.name, t.age, t.grade)
	}

	tw.Flush() // calculate column widths and print table

}

func testTableSorter() {
	table := []Row{
		{name: "Alice", age: 30, grade: 3.14},
		{name: "Bob", age: 25, grade: 1.23},
		{name: "Charlie", age: 40, grade: 4.56},
	}

	clicks := []int{0}

	sorter := TableSorter{
		Table:  table,
		Clicks: clicks,
	}

	fmt.Println("Original table:")
	printTable(table)

	fmt.Println("Sorted table:")
	sort.Sort(sorter)
	printTable(table)
}

func main() {
	testTableSorter()
}
