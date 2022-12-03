package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const rock string = "rock"
	const paper string = "paper"
	const scissors string = "scissors"

	const win string = "win"
	const lose string = "lose"
	const draw string = "draw"

	plays := make(map[string]string)
	plays["X"] = rock
	plays["A"] = rock
	plays["Y"] = paper
	plays["B"] = paper
	plays["Z"] = scissors
	plays["C"] = scissors

	scores := make(map[string]int)
	scores[rock] = 1
	scores[paper] = 2
	scores[scissors] = 3
	scores[win] = 6
	scores[lose] = 0
	scores[draw] = 3

	outcomes := map[string]map[string]string{
		rock:     {},
		paper:    {},
		scissors: {},
	}
	// Them - You
	outcomes[rock][rock] = draw
	outcomes[rock][paper] = win
	outcomes[rock][scissors] = lose
	outcomes[paper][rock] = lose
	outcomes[paper][paper] = draw
	outcomes[paper][scissors] = win
	outcomes[scissors][rock] = win
	outcomes[scissors][paper] = lose
	outcomes[scissors][scissors] = draw

	f, err := os.Open("/Users/katy/Developer/aoc-22/day2/input")
	check(err)

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		round := scanner.Text()

		items := s.Split(round, " ")

		opponent := plays[items[0]]
		you := plays[items[1]]
		outcome := outcomes[opponent][you]

		total = total + scores[you] + scores[outcome]

		fmt.Println(opponent, "against", you)
		fmt.Println("you", outcome)
		fmt.Println(scores[you], "+", scores[outcome])
		fmt.Println(total)
	}

	fmt.Println("Total score:", total)

	check(scanner.Err())
}
