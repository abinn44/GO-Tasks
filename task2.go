package main

import (
"errors"
"fmt"
)

func calculate(num1, num2 float64, operator string)(float64, error){
	switch operator{
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0{
			return 0, errors.New("Division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("Invalid operator")
	}
}

func main(){
var num1, num2 float64
var operator string

fmt.Println("Enter the first number:")
fmt.Scan(&num1)
fmt.Println("Enter the operator ( + , - , * , / )")
fmt.Scan(&operator)
fmt.Println("Enter the second number")
fmt.Scan(&num2)

result, err := calculate(num1,num2, operator)
if err != nil{
	fmt.Println("Error:", err)
	return
}
fmt.Printf("%.2f\n",result)
}
