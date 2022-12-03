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

	groups := map[int]map[int]string{}
	i := 0
	j := 0
	for scanner.Scan() {
		_, prs := groups[i]
		if prs == false {
			groups[i] = map[int]string{}
		}
		groups[i][j] = scanner.Text()
		if j == 2 {
			i++
			j = 0
		} else {
			j++
		}
	}

	for _, bags := range groups {
		var duplicate string
		var priority int
		checked := map[string]bool{}

		for _, c := range bags[0] {
			if checked[string(c)] {
				continue
			}

			checked[string(c)] = true
			if s.Contains(bags[1], string(c)) && s.Contains(bags[2], string(c)) {
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

		fmt.Println(bags[0], bags[1], bags[2])
		fmt.Println(duplicate, priority, total)
	}
	fmt.Println(total)

	check(scanner.Err())
}
