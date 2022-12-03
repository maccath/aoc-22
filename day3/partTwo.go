package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	s "strings"
)

func getDuplicate(bags [groupSize]string) rune {
	checked := map[string]bool{}
	var c rune
	var found bool

	for _, c := range bags[0] {
		if checked[string(c)] {
			continue
		}

		checked[string(c)] = true
		found = true
		for _, b := range bags[1:] {
			if !s.Contains(b, string(c)) {
				found = false
			}
		}

		if found {
			return c
		}
	}

	return c
}

const groupSize = 3

func main() {
	scanner := util.FileScanner(filepath.Join("day3", "input"))

	var total int
	var i int
	var bags [groupSize]string

	for scanner.Scan() {
		bags[i] = scanner.Text()

		if i < (groupSize - 1) {
			i++
			continue
		}

		duplicate := getDuplicate(bags)
		priority := util.GetPriority(duplicate)
		total += priority
		i = 0

		fmt.Println("The next group of", groupSize, "bags all contain a", string(duplicate), "item, with priority", priority)
	}

	fmt.Println("Total priority of all groups is", total)
}
