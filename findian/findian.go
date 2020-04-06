package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	lower := strings.ToLower(input)
	if strings.HasPrefix(lower, "i") &&
		strings.HasSuffix(lower, "n") &&
		strings.Contains(lower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
