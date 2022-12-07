package main

import (
	"bufio"
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
)

func main() {
	scanner := util.FileScanner(filepath.Join("day6", "input"))
	scanner.Split(bufio.ScanRunes)

	i := 1
	var queue []string
	var x string
	m := make(map[string]int)

	for scanner.Scan() {
		c := scanner.Text()

		if len(queue) == 4 {
			x, queue = queue[0], queue[1:]
			m[x] -= 1
		}
		queue = append(queue, c)
		m[c] += 1

		fmt.Println("Last four chars after", i, "are", queue)

		found := true
		for _, val := range queue {
			if m[val] > 1 {
				found = false
				break
			}
		}

		if len(queue) >= 4 && found {
			fmt.Println("Marker found after character", i)
			return
		}

		i++
	}
}
