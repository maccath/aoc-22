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
	plays["A"] = rock
	plays["B"] = paper
	plays["C"] = scissors
	plays["X"] = lose
	plays["Y"] = draw
	plays["Z"] = win

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
	outcomes[rock][draw] = rock
	outcomes[rock][win] = paper
	outcomes[rock][lose] = scissors
	outcomes[paper][lose] = rock
	outcomes[paper][draw] = paper
	outcomes[paper][win] = scissors
	outcomes[scissors][win] = rock
	outcomes[scissors][lose] = paper
	outcomes[scissors][draw] = scissors

	f, err := os.Open("/Users/katy/Developer/aoc-22/day2/input")
	check(err)

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		round := scanner.Text()

		items := s.Split(round, " ")

		opponent := plays[items[0]]
		outcome := plays[items[1]]
		play := outcomes[opponent][outcome]

		total = total + scores[play] + scores[outcome]

		fmt.Println(outcome, "against", opponent)
		fmt.Println("play", play)
		fmt.Println(scores[play], "+", scores[outcome])
		fmt.Println(total)
	}

	fmt.Println(total)

	check(scanner.Err())
}
