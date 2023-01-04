package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("No Printing")
	case true:
		fmt.Println("Printing")
	}
}