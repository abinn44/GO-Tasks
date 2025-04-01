package main

import "fmt"

func FindMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	numbers := []int{32, 54, 70, 81, 30, 20}
	max := FindMax(numbers)
	fmt.Println("The maximum number in the array is:", max)
}
