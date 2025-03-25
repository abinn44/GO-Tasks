package main

import (
	"fmt"
	"bridgeon/task/calculations"
)

func main() {
	a, b := 50.00, 5.00

	fmt.Printf("Sum: %.2f\n", calculations.Add(a, b))
	fmt.Printf("Difference: %.2f\n", calculations.Subtract(a, b))
	fmt.Printf("Product: %.2f\n", calculations.Multiply(a, b))

	quot, err := calculations.Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %.2f\n", quot)
	}
}
