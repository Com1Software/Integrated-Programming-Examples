package main

import (
	"C"
	"fmt"
)

//export PrintInt
func PrintInt(x int) {
	fmt.Println("Hello world")
	fmt.Println(x)
}

func main() {}
