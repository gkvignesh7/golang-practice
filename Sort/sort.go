// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		x := []int{4, 7, 8, 9, 2, 3, 5, 1, 6}
//		y := []string{"Vignesh", "Sudarshan", "Vinoth", "Ajith"}
//		sort.Strings(y)
//		sort.Ints(x)
//		fmt.Println(x)
//		fmt.Println(y)
//	}
package main

import "fmt" 
type person struct{
	first string
	age int
}
func main(){
	p1:=person{"James",32}
	p2:=person{"Moneypenny",27}
	p3:=person{"Q",64}
	p4:=person{"M",56}

	people:=[]person{p1,p2,p3,p4}
	fmt.Println(people)

}