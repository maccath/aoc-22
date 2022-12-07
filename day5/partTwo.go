package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	"strconv"
	s "strings"
	"unicode"
)

func main() {
	scanner := util.FileScanner(filepath.Join("day5", "input"))

	// Get stacks
	stacks := map[int][]string{}
	for scanner.Scan() {
		contents := scanner.Text()

		if contents == "" {
			break
		}

		for i := 1; i <= 9; i++ {
			j := ((i - 1) * 4) + 1
			letter := []rune(contents[j : j+1])[0]

			if !unicode.IsLetter(letter) {
				continue
			}

			stacks[i] = append(stacks[i], string(letter))
		}
	}

	fmt.Println("The starting stacks are", stacks)

	for scanner.Scan() {
		contents := s.Split(scanner.Text(), " ")

		move, _ := strconv.Atoi(contents[1])
		from, _ := strconv.Atoi(contents[3])
		to, _ := strconv.Atoi(contents[5])

		fmt.Println("Moving", move, "items from stack", from, "to stack", to)

		fmt.Println("Stack", from, "has", stacks[from])
		fmt.Println("Moving", move, "from", from, "-", stacks[from][0:move])
		fmt.Println("Stack", to, "has", stacks[to])

		newslice := make([]string, move)
		oldslice := make([]string, len(stacks[to]))
		copy(newslice, stacks[from][0:move])
		copy(oldslice, stacks[to][0:])

		stacks[to] = append(newslice, oldslice...)
		stacks[from] = stacks[from][move:len(stacks[from])]

		fmt.Println("Stack", to, "now has", stacks[to])
		fmt.Println("Stack", from, "now has", stacks[from])
	}

	fmt.Println("The final stacks are", stacks)
}
