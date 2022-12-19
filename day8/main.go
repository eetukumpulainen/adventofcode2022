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
	fmt.Println("Day 8")
	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	grid := [][]int{}

	for scanner.Scan() {
		stringrow := strings.Split(scanner.Text(), "")
		introw := []int{}
		for _, v := range stringrow {
			vint, err := strconv.Atoi(v)

			if err != nil {
				log.Fatal(err)
			}

			introw = append(introw, vint)
		}
		grid = append(grid, introw)
	}

	// Part1
	visibles, scenicScore := calculateVisibles(grid)
	fmt.Println("VISIBLES:", visibles)

	// Part 2
	// 1699908 Too high
	fmt.Println("TOP SCENIC SCORE:", scenicScore)
}

func calculateVisibles(grid [][]int) (int, int) {
	visibles := 0
	highestScenicScore := 0

	for i, v := range grid {

		// compare row item to all in the same row or column
		for j, vv := range v {
			height := vv
			visibleDirections := 4
			leftScore := j
			rightScore := len(v) - 1 - j
			aboveScore := i
			belowScore := len(grid) - 1 - i

			// compare with the ones left of it, starting from the closest
			for left := j - 1; left >= 0; left-- {
				if v[left] >= height {
					visibleDirections--
					leftScore = j - left
					break
				}
			}

			// compare with the ones right to it
			for right := j + 1; right < len(v); right++ {
				if v[right] >= height {
					visibleDirections--
					rightScore = right - j
					break
				}
			}

			// compare with the ones above it, starting from the closest
			for above := i - 1; above >= 0; above-- {
				if grid[above][j] >= height {
					visibleDirections--
					aboveScore = i - above
					break
				}
			}

			// compare with the ones below it
			for below := i + 1; below < len(grid); below++ {
				if grid[below][j] >= height {
					visibleDirections--
					belowScore = below - i
					break
				}
			}

			if visibleDirections > 0 {
				visibles++
			}

			scenicScore := leftScore * rightScore * aboveScore * belowScore
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return visibles, highestScenicScore
}
