// package main

// import "fmt"

// type Student struct {
// 	rollno int
// 	name   string
// 	marks  int
// }

//	func main() {
//		var s1 Student = Student{201, "Vignesh", 98}
//		fmt.Println(s1)
//		fmt.Println(s1.name)
//		var s2 Student=Student{rollno:202,marks:97}
//		fmt.Println(s2)
//		fmt.Println(s2.name)
//	}
// package main

// import "fmt"
// type User struct{
// 	Name string
// 	DOB uint64
// }
// func main(){
// 	var u1 =User{"Vignesh",07102001}
// 	fmt.Println(u1)

// }
package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Vignesh",
			last:  "G K",
		},
		ltk: true,
	}
	p2 := person{
		first: "Vasudevan",
		last:  "M",
	}
	fmt.Println(sa1)
	fmt.Println(p2)

}
