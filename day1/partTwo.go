package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	"strconv"
)

func main() {
	scanner := util.FileScanner(filepath.Join("day1", "input"))

	var a [3]int

	tot := 0
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if int(num) != 0 {
			// Increment the total held by current elf
			tot = int(num) + tot
			continue
		}

		if tot > a[0] {
			a[2] = a[1]
			a[1] = a[0]
			a[0] = tot
		} else if tot > a[1] {
			a[2] = a[1]
			a[1] = tot
		} else if tot > a[2] {
			a[2] = tot
		}

		// Onto the next elf
		tot = 0
	}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[0] + a[1] + a[2])
}
