package main

import (
	"fmt"
	"sync"
)

func printNumbers(start, step int, wg *sync.WaitGroup) {
	for i := start; i <= 10; i += step {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go printNumbers(1, 2, &wg) // Print odd numbers
	go printNumbers(2, 2, &wg) // Print even numbers
	
	wg.Wait()
}
