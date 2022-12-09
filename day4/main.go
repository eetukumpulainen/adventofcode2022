package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	fullyOverlapping := 0
	overlapping := 0
	for scanner.Scan() {

		// Inspect assignment pairs
		assignments := strings.Split(scanner.Text(), ",")

		ass1 := strings.Split(assignments[0], "-")
		ass2 := strings.Split(assignments[1], "-")

		// In day1 we used fmt.Sscan to convert
		// Here we convert using strconv.Atoi
		// Would be good to know if this could be simplified or even comletely avoided
		x1, err1 := strconv.Atoi(ass1[0])
		x2, err2 := strconv.Atoi(ass1[1])
		y1, err3 := strconv.Atoi(ass2[0])
		y2, err4 := strconv.Atoi(ass2[1])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Fatal()
		}

		// Part one
		// Fully overlapping
		if (x1 >= y1 && x2 <= y2) || (y1 >= x1 && y2 <= x2) {
			fullyOverlapping++
		}

		// Part two
		// Overlapping (fully or partially)
		if (x1 <= y1 && y1 <= x2) || (x1 <= y2 && y2 <= x2) || (y1 <= x1 && x1 <= y2) || (y1 <= x2 && x2 <= y2) {
			overlapping++
		}
	}

	fmt.Println("Fully overlapping pairs:", fullyOverlapping)
	fmt.Println("Overlapping pairs (fully or partially):", overlapping)
	// Part1 correct answer 651
	// Part2 correct answer 956
}
