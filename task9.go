package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float64)


	grades["Abin"] = 85.5
	grades["Ameen"] = 78.0
	grades["Shahid"] = 92.3

	grades["Ameen"] = 80.5

	fmt.Println("Student Grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %.2f\n", name, grade)
	}

	grades["Anwar"] = 88.7
	fmt.Println("\nAfter Adding Anwar:")
	for name, grade := range grades {
		fmt.Printf("%s: %.2f\n", name, grade)
	}


	delete(grades, "Abin")
	fmt.Println("\nAfter Removing Abin:")
	for name, grade := range grades {
		fmt.Printf("%s: %.2f\n", name, grade)
	}
}