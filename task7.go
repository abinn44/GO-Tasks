package main

import "fmt"

func minMax(arr []int) (int, int) {
    if len(arr) == 0 {
        fmt.Println("Array is empty")
        return 0, 0
    }
    
    max, min := arr[0], arr[0]
    
    for _, num := range arr {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    
    return max, min
}

func main() {
    arr := []int{10, 5, 2, 8, 15}
    max, min := minMax(arr)
    fmt.Println("Maximum:", max)
    fmt.Println("Minimum:", min)
}
