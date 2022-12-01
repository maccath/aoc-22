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
