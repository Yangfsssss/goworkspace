package main

import (
	"flag"
	"fmt"
	"tempconv2"
)

var temp = tempconv2.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
