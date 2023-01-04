// package main

// import "fmt"

// func main() {
// 	numb := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println(numb)
// 	fmt.Println(numb[1:4])
// 	fmt.Println(numb[:3])
// 	fmt.Println(numb[4:])
// 	var vignesh="teacoffee"
// 	fmt.Println(vignesh[:3])
// 	fmt.Println(vignesh[4:7])
// }

// package main

// import "fmt"

//	func main() {
//		gk := []string{"Apache", "NS", "Xpluse"}
//		fmt.Println(gk)
//		kv := []string{"TVS", "Bajaj", "Hero"}
//		fmt.Println(kv)
//		bk := [][]string{gk, kv}
//		fmt.Println(bk)
//	}
package main

import "fmt"

func main() {
	my := []string{"a", "b", "c", "d"}
	fmt.Println(my)
	my = append(my, "v")
	fmt.Println(my)
}
