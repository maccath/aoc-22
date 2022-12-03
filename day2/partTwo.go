package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/janken"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	s "strings"
)

func main() {
	plays := make(map[string]string)
	plays["A"] = janken.Rock
	plays["B"] = janken.Paper
	plays["C"] = janken.Scissors
	plays["X"] = janken.Lose
	plays["Y"] = janken.Draw
	plays["Z"] = janken.Win

	scanner := util.FileScanner(filepath.Join("day2", "input"))

	var total int
	for scanner.Scan() {
		round := scanner.Text()

		items := s.Split(round, " ")

		opponent := plays[items[0]]
		outcome := plays[items[1]]
		play := janken.Fix(opponent, outcome)
		score := janken.Score(play, outcome)

		total += score

		fmt.Println("You must", outcome, "against", opponent)
		fmt.Println("You play", play, "and score", score)
	}

	fmt.Println(total)
}
