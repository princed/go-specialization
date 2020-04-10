package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(ints []int, i int) {
	ints[i], ints[i+1] = ints[i+1], ints[i]
}

func BubbleSort(ints []int) {
	for swapped := true; swapped == true; {
		swapped = false
		for i := 0; i < len(ints)-1; i++ {
			if ints[i+1] < ints[i] {
				Swap(ints, i)
				swapped = true
			}
		}
	}
}

func main() {
	fmt.Print("Enter a sequence of integers delimited by a space: ")
	stdinScanner := bufio.NewScanner(os.Stdin)
	stdinScanner.Scan()
	wordScanner := bufio.NewScanner(strings.NewReader(stdinScanner.Text()))
	wordScanner.Split(bufio.ScanWords)
	var ints []int
	for wordScanner.Scan() {
		num, err := strconv.Atoi(wordScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}
	BubbleSort(ints)
	fmt.Println(ints)
}
