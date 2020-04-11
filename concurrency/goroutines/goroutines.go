package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const partsNum = 4

func main() {
	// Prompt user for an array of ints
	ints := getInput()
	// Split array into approximately equal slices
	parts := getParts(ints)

	// Sort each part a in goroutine
	wg := sync.WaitGroup{}
	wg.Add(partsNum)
	for _, p := range parts {
		go sortPart(p, &wg)
	}
	wg.Wait()

	// Merge sorted parts
	sortedInts := make([]int, len(ints))
	sortParts(parts, sortedInts)

	fmt.Println(sortedInts)
}

func getInput() (ints []int) {
	fmt.Print("Enter a sequence of integers delimited by a space: ")
	stdinScanner := bufio.NewScanner(os.Stdin)
	stdinScanner.Scan()
	wordScanner := bufio.NewScanner(strings.NewReader(stdinScanner.Text()))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		num, err := strconv.Atoi(wordScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}
	return
}

func getParts(ints []int) (parts [][]int) {
	intsLen := len(ints)
	partLen := intsLen / partsNum
	partRem := intsLen % partsNum
	parts = make([][]int, 4)

	for i := 0; i < partsNum; i++ {
		partStart := i * partLen
		partEnd := (i + 1) * partLen
		if partRem > 0 {
			partEnd++
			partRem--
		}
		parts[i] = ints[partStart:partEnd]

	}
	return
}

func sortPart(ints []int, wg *sync.WaitGroup) {
	fmt.Println("Sorting", ints)
	sort.Ints(ints)
	wg.Done()
}

func sortParts(parts [][]int, sortedInts []int) {
	nextIndexes := make([]int, partsNum)
	for i := 0; i < len(sortedInts); i++ {
		minV := math.MaxInt64
		minVPartIndex := 0
		for j, p := range parts {
			if nextIndexes[j] >= len(p) {
				continue
			}
			v := p[nextIndexes[j]]
			if v < minV {
				minV = v
				minVPartIndex = j
			}
		}
		nextIndexes[minVPartIndex]++
		sortedInts[i] = minV
	}
}
