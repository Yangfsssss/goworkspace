package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin  = 0
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 方法定义
func (c Celsius) String() string {
	fmt.Println("String")

	return fmt.Sprintf("%g°99", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
