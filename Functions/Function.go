// package main

// import "fmt"

// func main() {
// 	var a int = 100
// 	var b int = 200

//		var ret int
//		ret = max(a, b)
//		fmt.Println(ret)
//	}
//
//	func max(num1,num2 int) int{
//		var result int
//		if(num1>num2){
//			result=num1
//		}else{
//			result=num2
//		}
//		return result
//	}
// package main

// import "fmt"

//	func main(){
//		var b int =5
//		var c int =6
//		if(b>c){
//			fmt.Println("b is greater")
//		}else{
//			fmt.Println("c is greater")
//		}
//	}
// package main

// import "fmt"
// func main(){
// 	var a int =50
// 	var b int =123
// 	var c int
// 	c=sum(a,b)
// 	fmt.Println(c)
// }
//  func sum(a,b int) int{
// 	return a+b
//  }

// package main

// import "fmt"
//
//	func sum(a,b int) int{
//		return a+b
//	}
//
//	func main(){
//		var a int =153
//		var b int =127
//		var c int
//		c=sum(a,b)
//		fmt.Println(c)
//	}
// package main

// import "fmt"

// var y = 54
// var x = 32

//	func main() {
//		var c int =07
//		fmt.Println(c)
//		age()
//	}
//
//	func age() {
//		fmt.Println("My Age is 21!!!")
//		fmt.Println(x, y,)
//	}
package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func main() {
	var a int
	var b int
	var c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c = sum(a, b)
	fmt.Println("Sum of two number is:")
	fmt.Println(c)

}
