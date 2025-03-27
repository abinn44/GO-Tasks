package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) PrintDetails() {
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "John Doe", Age: 30}
    person.PrintDetails()
}
