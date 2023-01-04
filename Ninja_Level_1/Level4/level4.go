package main

import "fmt"

type vignesh int

var v vignesh

func main() {
	var v = 45
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	v = 42
	fmt.Println(v)
}
