package main

import (
	"fmt"
)
func duplicates(nums []int) []int {
	seen := make(map[int]bool)
	unique := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	return unique
}

func main() {
	numbers := []int{1, 8, 8, 3, 4, 4, 5, 3, 3, 7}
	fmt.Println("Original slice:", numbers)
	fmt.Println("Without duplicates:", duplicates(numbers))
}