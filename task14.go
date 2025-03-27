package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(ch chan int) {
	for item := range ch {
		fmt.Println("Consumed:", item)
		time.Sleep(time.Millisecond * 700)
	}
}

func main() {
	ch := make(chan int, 2)
	
	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second * 5)
}
