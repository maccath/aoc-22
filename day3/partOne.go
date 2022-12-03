package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	s "strings"
)

func findDuplicate(part1 string, part2 string) rune {
	var duplicate rune
	checked := map[string]bool{}

	for _, c := range part1 {
		if checked[string(c)] {
			continue
		}

		checked[string(c)] = true
		if s.Contains(part2, string(c)) {
			duplicate = c
			break
		}
	}

	return duplicate
}

func main() {
	scanner := util.FileScanner(filepath.Join("day3", "input"))

	total := 0

	for scanner.Scan() {
		contents := scanner.Text()

		part1 := contents[:len(contents)/2]
		part2 := contents[len(contents)/2:]

		duplicate := findDuplicate(part1, part2)
		priority := util.GetPriority(duplicate)

		total += priority

		fmt.Println("Item of type", string(duplicate), "is incorrectly arranged, with priority", priority)
	}

	fmt.Println("Total priority of incorrectly arranged items is", total)
}
