package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Enter a float number: ")
	fmt.Scan(&input)
	float, err := strconv.ParseFloat(input, 64)
	if err == nil {
		fmt.Printf("%.0f", math.Trunc(float))
	} else {
		fmt.Println(err)
	}
}
