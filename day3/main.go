package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var alphabetNumber = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
}

func main() {
	args := os.Args

	filename := args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	priority := 0

	for scanner.Scan() {
		bothRucksacks := strings.Split(scanner.Text(), "")

		// Divide string in two equally long pieces
		first := strings.Join(bothRucksacks[:len(bothRucksacks)/2], "")
		second := strings.Join(bothRucksacks[len(bothRucksacks)/2:], "")

		// Find duplicate in string
		duplicatePriority := 0
		for _, r := range first {
			if strings.ContainsRune(second, r) {
				if unicode.IsLower(r) {
					duplicatePriority = alphabetNumber[r]
				} else {
					duplicatePriority = alphabetNumber[unicode.ToLower(r)] + 26
				}
				priority += duplicatePriority
				break
			}
		}
	}

	fmt.Println(priority)
	part2(filename)
}

func part2(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	threeElves := make([]string, 3)
	priority := 0
	index := 0
	for scanner.Scan() {
		threeElves[index] = scanner.Text()

		// Search the three strings for a match
		if index == 2 {
			for _, s := range threeElves[0] {
				if strings.ContainsRune(threeElves[1], s) && strings.ContainsRune(threeElves[2], s) {
					if unicode.IsLower(s) {
						priority += alphabetNumber[s]
					} else {
						priority += alphabetNumber[unicode.ToLower(s)] + 26
					}
					index = 0
					break
				}
			}
		} else {
			index++
		}
	}
	fmt.Println(priority)
}
