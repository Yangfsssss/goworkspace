package weightconv

import "fmt"

type Gram float64
type Pound float64
type Ounce float64

func (g Gram) String() string {
	return fmt.Sprintf("%ggram", g)
}

func (p Pound) String() string {
	return fmt.Sprintf("%gpound", p)
}

func (o Ounce) String() string {
	return fmt.Sprintf("%gounce", o)
}
