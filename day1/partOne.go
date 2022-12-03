package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	"strconv"
)

func main() {
	var max int
	var total int

	scanner := util.FileScanner(filepath.Join("day1", "input"))

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		emptyLine := num == 0

		if !emptyLine {
			total += num
			continue
		}

		if total > max {
			max = total
		}
		total = 0
	}

	fmt.Println("The top elf is carrying", max, "calories")
}
