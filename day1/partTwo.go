package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	topNum := 3
	elves := make([]int, topNum)
	var carrying int

	scanner := util.FileScanner(filepath.Join("day1", "input"))

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		emptyLine := num == 0

		if !emptyLine {
			carrying += num
			continue
		}

		lastItem := len(elves) - 1
		if carrying > elves[lastItem] {
			elves = append(elves[:lastItem], carrying)
			sort.Sort(sort.Reverse(sort.IntSlice(elves)))
		}

		carrying = 0
	}

	var total int
	for _, val := range elves {
		total += val
	}

	fmt.Println("The top", topNum, "elves are carrying", elves, "calories")
	fmt.Println("The total calories carried by the elves is", total)
}
