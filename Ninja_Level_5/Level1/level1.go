package main

import "fmt"

type person struct {
	first   string
	last    string
	favFlav []string
}

func main() {
	p1 := person{
		first: "Vignesh",
		last:  "G K ",
		favFlav: []string{
			"coke", "Ice Cream",
		},
	}
	p2 := person{
		first: "Vasudevan",
		last:  "M",
		favFlav: []string{
			"Soda", "Juice",
		},
	}
	fmt.Println(p1)
	fmt.Println(p2)

}
