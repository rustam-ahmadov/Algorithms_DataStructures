package main

import (
	"fmt"
	"time"
)

// 0 <= n <= 30

var memo []int = make([]int, 30)

func fib(n int) int {
	if n <= 1 {
		return n
	}
    if memo[n]!=0 {
        return memo[n]
    }
	memo[n] = fib(n-1) + fib(n-2)
    return memo[n]
}

func main(){
    start:= time.Now()
    fmt.Println(fib(10))
    end:= time.Now()
    fmt.Printf("time %s\n", end.Sub(start))
}
