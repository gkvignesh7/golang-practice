// package main

// import "fmt"

// func main() {
// 	v := map[string]int{
// 		"Vignesh": 07,
// 		"Vasu":    29,
// 		"Saran":   05,
// 		"Roshan":  11,
// 	}
// 	fmt.Println(v)
// 	fmt.Println(v["Saran"])
// 	x, ok := v["Sakthi_Sumanth"]
// 	fmt.Println(x)
// 	fmt.Println(ok)
// 	if b, ok := v["Vignesh"]; ok {
// 		fmt.Println("This is Okie ", b)

// 	}
// 	v["Sakthi"] = 9
// 	fmt.Println(v)
// 	delete(v, "Vignesh")
// 	delete(v, "Vasu")
// 	fmt.Println(v)

// }
// package main

// import "fmt"

// func main() {
// 	v := map[string]int{
// 		"Vivo":    200,
// 		"Samsung": 500,
// 		"Redmi ":  400,
// 		"OnePlus": 100,
// 		"Oppo":    300,
// 	}
// 	fmt.Println(v)
// 	fmt.Println(v["Vivo"])
// 	v["RealMe"] = 600
// 	fmt.Println(v)
// 	v["Apple"] = 1000
// 	fmt.Println(v)
// 	delete(v, "RealMe")
// 	fmt.Println(v)

// }
package main

import "fmt"

func main() {
	v := map[int]string{
		1: "Apple",
		2: "Redmi",
		3: "Samsung",
		4: "Oneplus Nord CE",
		5: "Vivo",
	}
	fmt.Println(v)
	fmt.Println(v[6])
	v[6] = "Oppo"
	fmt.Println(v)
}
