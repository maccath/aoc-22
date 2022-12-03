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
	plays["X"] = janken.Rock
	plays["A"] = janken.Rock
	plays["Y"] = janken.Paper
	plays["B"] = janken.Paper
	plays["Z"] = janken.Scissors
	plays["C"] = janken.Scissors

	scanner := util.FileScanner(filepath.Join("day2", "input"))

	var total int
	for scanner.Scan() {
		round := scanner.Text()

		items := s.Split(round, " ")

		opponent := plays[items[0]]
		you := plays[items[1]]
		outcome := janken.Play(you, opponent)
		score := janken.Score(you, outcome)

		total += score

		fmt.Println("You play", you, "against", opponent)
		fmt.Println("You", outcome, "and score", score)
	}

	fmt.Println("Total score:", total)
}
