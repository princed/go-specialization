package main

import (
	"fmt"
	"time"
)

// Race conditions
//
// Race condition is a problem where the outcome of the program depends on the interleaving,
// i.e. order in which different concurrent threads (or goroutines or processes) are executed,
// which leads to non-deterministic program behaviour. Race conditions can only happen when
// different concurrent thread communicate with each other in any way,
// e.g. when they share the same variable.

func main() {
	for i := 0; i < 15; i++ {
		x := 0
		go setX(&x)
		go printX(&x)
		time.Sleep(time.Second)
	}
}

func printX(x *int) {
	fmt.Print(*x)
}

func setX(x *int) {
	*x++
}
