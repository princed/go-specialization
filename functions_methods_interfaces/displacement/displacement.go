package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return a*math.Pow(t, 2)/2 + v0*t + s0
	}
}

func main() {
	var (
		a, v0, s0, t float64
	)
	fmt.Print("Enter acceleration, initial velocity, and initial displacement separated by spaces: ")
	fmt.Scan(&a, &v0, &s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	fmt.Println(fn(t))
}
