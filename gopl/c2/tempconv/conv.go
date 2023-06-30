package tempconv

func CToF(c Celsius) Fahrenheit {
	// 类型转换，而非函数调用
	// 对每一个类型T，都有对应的类型转换操作T(x)
	// 只有两个类型的底层基础类型相同时才能进行
	// 数值/字符串/特定类型的slice也可以进行转换
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}
