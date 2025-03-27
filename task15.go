package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		mutex.Lock()
		counter++
		fmt.Println("Consumed:", item, "Counter:", counter)
		mutex.Unlock()
		time.Sleep(time.Millisecond * 700)
	}
}

func main() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(ch)
	go consumer(ch, &wg)

	wg.Wait()
}
