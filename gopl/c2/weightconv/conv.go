package weightconv

func GToO(g Gram) Ounce {
	return Ounce(g * 0.035274)
}

func OToG(o Ounce) Gram {
	return Gram(o * 28.349523125)
}

func GToP(g Gram) Pound {
	return Pound(g * 0.00220462)
}

func PToG(p Pound) Gram {
	return Gram(p * 453.59237)
}

func OToP(o Ounce) Pound {
	return Pound(o * 0.0625)
}

func PToO(p Pound) Ounce {
	return Ounce(p * 16)
}
