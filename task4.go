package main

import "fmt"

func calculate(a,b float64){
	fmt.Printf("sum: %.2f\n", a+b)
	fmt.Printf("Difference: %.2f\n", a-b)
	fmt.Printf("Product: %.2f\n", a*b)
	if b != 0{
		fmt.Printf("Quotient %.2f\n", a/b)
	} else {
		fmt.Printf("Division by zero is not allowed")
	}
	
}

func main(){
	calculate(10,0)
}