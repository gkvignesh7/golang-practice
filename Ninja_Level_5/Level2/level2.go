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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)
	fmt.Println(p1)
	fmt.Println(p2)

}
