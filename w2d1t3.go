package main

import "fmt"

func modify(num *int) {
    *num = *num * 2
}

func main() {
    value := 10
    fmt.Println("Before modification:", value)
    modify(&value)
    fmt.Println("After modification:", value)
}
