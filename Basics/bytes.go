package main

import "fmt"

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b", kb, kb)
	fmt.Printf("%d\t\t%b", mb, mb)
	fmt.Printf("%d\t\t%b", gb, gb)
}
