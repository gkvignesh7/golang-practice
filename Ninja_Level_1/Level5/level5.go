package main

import "fmt"

type vignesh int

var x vignesh
var y int

func main() {
	var x = 45
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
