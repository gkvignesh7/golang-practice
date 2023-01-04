package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divides by 4 is ,the remainder is %v", i, i%4)
	}
}
