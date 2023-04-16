package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	go func() {
		n1 := 44
		fibN1 := fib(n1)
		fmt.Printf("\rFibonacci(%d) = %d\n", n1, fibN1)
	}()
	go func() {
		n2 := 45
		fibN2 := fib(n2)
		fmt.Printf("\rFibonacci(%d) = %d\n", n2, fibN2)
	}()
	time.Sleep(10 * time.Second)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
