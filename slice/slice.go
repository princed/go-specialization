package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)
	for {
		fmt.Print("Enter an integer: ")
		var input string
		fmt.Scan(&input)
		if strings.ToLower(input) == "x" {
			break
		}
		integer, err := strconv.Atoi(input)
		if err == nil {
			slice = append(slice, integer)
			sort.Slice(slice, func(i, j int) bool {
				return slice[i] < slice[j]
			})
		}
		fmt.Println(slice)
	}
}
