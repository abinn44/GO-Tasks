package main

import "fmt"

func modifyNumber(num *int) {
    *num = *num * 2
}

func main() {
    value := 10
    fmt.Println("Before:", value)
    
    modifyNumber(&value)
    
    fmt.Println("After:", value)
}
