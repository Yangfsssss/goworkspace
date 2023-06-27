package main

import (
	"bufio"
	"fmt"
	"os"
)

func echo1 ()string{
	var s,sep string
	for i:=0;i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = "  "
	}

	return s
}

func echo2 (){
	// var s,sep string
	for i,v:=range os.Args {
		fmt.Println(i,v)
		fmt.Println()
	}
}

func dup1(){
	counts:=make(map[string]int)
	input:=bufio.NewScanner(os.Stdin)
}

func main(){
	// 1.2
	// s:= echo1()
	// fmt.Println(s)
	// echo2()


}