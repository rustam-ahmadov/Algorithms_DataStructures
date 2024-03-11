package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Program starts...")
	var wg sync.WaitGroup
	wg.Add(2)

	var mu1 sync.Mutex
	var mu2 sync.Mutex
	// Goroutine 1
	go func() {
		mu1.Lock()
		fmt.Println("mu1 locked in goroutine1")
		mu2.Lock()
		fmt.Println("mu2 locked in goroutine1")
		defer mu1.Unlock()
		defer mu2.Unlock()
	}()

	// Goroutine 2
	go func() {
		mu2.Lock()
		fmt.Println("mu2 locked in goroutine2")
		mu1.Lock()
		fmt.Println("mu1 locked in goroutine2")
		defer mu2.Unlock()
		defer mu1.Unlock()
	}()

	fmt.Println("Program waits...")
	wg.Wait()
}
