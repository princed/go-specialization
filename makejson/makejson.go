package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fields := []string{"name", "address"}
	m := make(map[string]string)
	var input string

	for _, field := range fields {
		fmt.Printf("Enter a %s: ", field)
		fmt.Scan(&input)
		m[field] = input
	}

	bytes, _ := json.Marshal(m)
	fmt.Println(string(bytes))
}
