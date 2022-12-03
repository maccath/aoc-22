package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pwd, _ := os.Getwd()
	f, err := os.Open(filepath.Join(pwd, "day3", "input"))
	check(err)

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		contents := scanner.Text()

		part1 := contents[:len(contents)/2]
		part2 := contents[len(contents)/2:]

		var duplicate string
		var priority int
		checked := map[string]bool{}

		for _, c := range part1 {
			if checked[string(c)] {
				continue
			}

			checked[string(c)] = true
			if s.Contains(part2, string(c)) {
				duplicate = string(c)
				if int(c) > 96 { // lowercase
					priority = int(c) - 96
				} else { // uppercase
					priority = int(c) - 38
				}
				break
			}
		}

		total = total + priority

		fmt.Println(part1, part2, duplicate, priority, total)
	}
	fmt.Println(total)

	check(scanner.Err())
}
