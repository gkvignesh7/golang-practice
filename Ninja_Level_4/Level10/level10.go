package main

import "fmt"

func main() {
	m := map[string][]string{
		"Vignesh": []string{"Cricket", "Hai"},
		"Vasu":    []string{"Volleyball", "Hello"},
		"Saran":   []string{"Football", "Bye"},
	}
	fmt.Println(m)
	m["Roshan"] = []string{"Bike", "Tata"}
	fmt.Println(m)
	delete(m, "Vignesh")
	fmt.Println(m)
}
