package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	args := os.Args

	filename := args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	calories := 0
	mostCalories := 0
	topThreeCalories := []int{0, 0, 0}
	for scanner.Scan() {
		cal := 0
		txt := scanner.Text()
		if txt != "" {
			fmt.Sscan(txt, &cal)
			calories = calories + cal
		} else {
			if calories > mostCalories {
				mostCalories = calories
			}

			// Part 2
			// Keep three biggest calories on sorted slice
			// Compare new calories to the min calories in slice and update if needed
			if calories > topThreeCalories[0] {
				topThreeCalories[0] = calories
				sort.Ints(topThreeCalories)
			}

			calories = 0
		}
	}

	// After scanner closes, the last item on the input is ignored, so it's handled here
	if calories > mostCalories {
		mostCalories = calories
	}

	// Part 2
	// Keep three biggest calories on sorted slice
	// Compare new calories to the min calories in slice and update if needed
	if calories > topThreeCalories[0] {
		topThreeCalories[0] = calories
		sort.Ints(topThreeCalories)
	}

	fmt.Println(mostCalories)
	fmt.Println(topThreeCalories)

	sum := 0
	for _, v := range topThreeCalories {
		sum = sum + v
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
