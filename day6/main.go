package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	// Sequence start is indicated with 4 different characters
	// Report number of chars from the beginning of the stream to the end of the start sequence
	// Test result is 7, 5, 6, 10, 11

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		seq := scanner.Text()
		fmt.Println(seq)

		// Part one answer
		fmt.Println(findUniqueSequence(seq, 4))

		// Part two answer
		fmt.Println(findUniqueSequence(seq, 14))
	}
}

func findUniqueSequence(seq string, length int) int {
	index := 1
	prev := []rune{}
	for _, v := range seq {

		startSequence := true

		if len(prev) != length {
			startSequence = false
			prev = append(prev, v)
		} else {
			prev = append(prev[1:], v)

			for ind1, i := range prev {
				for ind2, j := range prev {
					if ind2 == ind1 {
						continue
					}
					if i == j {
						startSequence = false
						break
					}
				}
			}
		}

		if startSequence {
			return index
		}
		index++
	}
	return 0
}
