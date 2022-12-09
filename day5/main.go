package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Part one answer is TLNGFGMFN

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	matr := [][]string{}

	// Read stacks and build matrix
	for scanner.Scan() {

		// First we need to construct stacks
		ln := scanner.Text()

		ln = strings.ReplaceAll(ln, "[", " ")
		ln = strings.ReplaceAll(ln, "]", " ")
		arr := strings.Split(ln, "")

		if len(arr) == 0 {
			break
		}

		matr = append(matr, arr)

	}

	// Transpose matrix
	matr = transpose(matr)

	// Remove empty rows from matrix
	matr = remove_empty_rows(matr)

	// Copy matrix for part2
	matr2 := make([][]string, len(matr))
	for i := range matr {
		matr2[i] = make([]string, len(matr[i]))
		copy(matr2[i], matr[i])
	}

	// Read instructions and move elements accordingly
	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), " ")
		blocks, err := strconv.Atoi(ln[1])
		from := ln[3]
		to := ln[5]

		if err != nil {
			log.Fatal(err)
		}

		// Call move operation as many times as needed
		for i := 0; i < blocks; i++ {
			move(matr, from, to)
		}

		// Call move operation suitable for part2
		movePart2(matr2, blocks, from, to)
	}

	// Return the first elements of the matrix rows

	for _, v := range matr {
		fmt.Print(v[0])
	}
	fmt.Println()

	for _, v := range matr2 {
		fmt.Print(v[0])
	}
	fmt.Println()
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func remove_empty_rows(matr [][]string) [][]string {
	newmatr := [][]string{}
	for _, i := range matr {
		for _, j := range i {
			if j != "" && j != " " {
				newmatr = append(newmatr, i)
				break
			}
		}
	}
	return newmatr
}

func move(matr [][]string, from string, to string) {
	block := ""
	destinationArray := []string{}
	destinationIndex := 0
	for i, v := range matr {
		if v[len(v)-1] == from {
			// Memorize and remove first block
			for j, b := range v {
				if b != "" && b != " " {
					block = b
					matr[i] = append(matr[i][:j], matr[i][j+1:]...)
					break
				}
			}
		}
		if v[len(v)-1] == to {
			// Memorize right destination array
			destinationIndex = i
			destinationArray = v
		}
	}

	// Move memorized block into destination arrays first element and keep order
	matr[destinationIndex] = append([]string{block}, destinationArray...)

}

func movePart2(matr [][]string, q int, from string, to string) {
	blocks := []string{}
	destinationIndex := 0
	destinationArray := []string{}
	for i, v := range matr {
		if v[len(v)-1] == from {
			// Memorize and remove blocks
			// Get next block q times
			for count := 0; count < q; count++ {
				for j, b := range v {
					if b != "" && b != " " {
						blocks = append(blocks, b)
						matr[i] = append(matr[i][:j], matr[i][j+1:]...)
						break
					}
				}
			}
		}
		if v[len(v)-1] == to {
			// Memorize right destination array
			destinationIndex = i
			destinationArray = v
		}
	}

	// Move memorized blocks into destination array and keep order
	matr[destinationIndex] = append(blocks, destinationArray...)
}
