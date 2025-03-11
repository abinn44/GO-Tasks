// package main
// import "fmt"
// func main(){
// var n int
// fmt.Println("Enter a number")
// fmt.Scan(&n)
// if n % 2 == 0{
// 	fmt.Println("Number is even")
// } else {
// 	fmt.Println("Number is odd")
// }

// }


package main

import "fmt"

func main() {
	grade := "A"

	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good!")
		fallthrough // Executes next case even if it's not a match
	case "C":
		fmt.Println("Needs Improvement.")
	default:
		fmt.Println("Invalid grade.")
	}
}
