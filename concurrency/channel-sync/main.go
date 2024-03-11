package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := make(chan int, 1)
	temp := 0
	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			temp += <-num
			temp += 1
		}
		num <- temp
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			num <- 1
		}
	}()
	wg.Wait()
	close(num) // Closing the channel as we're done with it
	finalValue := <-num
	fmt.Println(finalValue)
}
