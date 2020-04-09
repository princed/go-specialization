package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter a name of a file containing series of names: ")
	var fileName string
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	names := make([]Name, 0, 5)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")
		names = append(names, Name{name[0], name[1]})
	}
	for _, name := range names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}
