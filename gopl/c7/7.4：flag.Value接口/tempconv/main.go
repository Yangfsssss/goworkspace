package tempconv2

import (
	"flag"
	"fmt"
	"tempconv"
)

// type Celsius float64

type celsiusFlag struct {
	tempconv.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	// 用于从字符串中按照指定格式解析数据并将其赋值给变量
	// %s:字符串
	// %d:整数
	// %f:浮点数
	fmt.Sscanf(s, "%f%s", &value, &unit)

	fmt.Println("Set")

	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = tempconv.KToC(tempconv.Kelvin(value))
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	// 执行了tempconv.Celsius.String()
	f := celsiusFlag{value}
	// 想要解析命令行参数的值到一个自定义类型的变量时，可以使用 flag.CommandLine.Var 函数来实现
	// 执行了Set()
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
