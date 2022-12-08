package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*

Outcomes
A/X (rock) > C/Z (scissors)
C/Z (scissors) > B/Y (paper)
B/Y (paper) > A/X (rock)

Point system

You choose:
X -> 1pt
Y -> 2pts
Z -> 3pts

Outcome:
Win -> 6pts
Draw -> 3pts
Lose -> 0pts

*/

func main() {
	args := os.Args

	filename := args[1]

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	matching := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	outcomes := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	totalScore := 0

	for scanner.Scan() {
		score := 0
		game := strings.Split(scanner.Text(), " ")

		score += outcomes[game[1]]

		// Draw or win
		if game[0] == matching[game[1]] {
			score += 3
		} else if (game[0] == "A" && game[1] == "Y") || (game[0] == "B" && game[1] == "Z") || (game[0] == "C" && game[1] == "X") {
			score += 6
		}
		//fmt.Println(game, score)
		totalScore += score
	}

	fmt.Println("Part 1:", totalScore)
	part2(filename)
}

/*

Part 2 point system

Outcomes:
X -> Lose
Y -> Draw
Z -> Win

*/
func part2(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	outcomes := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	scores := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	totalScore := 0

	for scanner.Scan() {
		score := 0
		game := strings.Split(scanner.Text(), " ")

		score += scores[game[1]]

		// draw, win or lose
		if game[1] == "Y" {
			score += outcomes[game[0]]
		} else if game[1] == "Z" {
			switch first := game[0]; first {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		} else {
			switch first := game[0]; first {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		}

		totalScore += score
	}

	fmt.Println("Part 2:", totalScore)
}
