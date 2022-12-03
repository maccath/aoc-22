package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	"strconv"
)

func main() {
	scanner := util.FileScanner(filepath.Join("day1", "input"))

	max := 0
	total := 0
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		// Onto the next elf
		if int(num) == 0 {
			if total > max {
				max = total
			}
			total = 0
			continue
		}

		// Increment the total held by current elf
		total = int(num) + total
	}

	fmt.Println(max)
}
