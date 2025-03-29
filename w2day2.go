package main

import "fmt"

func main(){

n := new([]int)
fmt.Println("n = ", *n)

*n = make([]int, 5)
(*n)[0] = 100
(*n)[1] = 200
fmt.Println("n = ", *n)

m := make([]int, 4)
fmt.Println("m = ",m)
m[0] = 300
m[1] = 400
fmt.Println("m = ", m)

}