package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("/Users/katy/Developer/aoc-22/day1/input")
	check(err)

	scanner := bufio.NewScanner(f)

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

	check(scanner.Err())
}
