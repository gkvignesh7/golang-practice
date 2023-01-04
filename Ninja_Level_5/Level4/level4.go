package main

import "fmt"

func main() {
	s:=struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first :"Vignesh",
		friends:map[string] int{
			"Vasu":555,
			"q":777,
			"m":888,
		},
		favDrinks: []string{
			"Water",
			"Coke",
		},
	}
	fmt.Println(s)
}