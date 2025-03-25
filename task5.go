package main

import "fmt"
import "errors"

func division(a,b float64)(float64, error){
	if b == 0{
		return 0, errors.New("Divison by zero is not allowed")
	} 
	return a/b , nil
}

	func main(){
		a, b := 10.00, 0.00
		result, err := division(a,b)
		if err != nil{
			fmt.Println("Error", err)
		} else {
			fmt.Printf("result: %.2f\n", result)
		}
	}
