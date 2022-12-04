package main

import (
	"fmt"
	"github.com/maccath/aoc-22/internal/pkg/util"
	"path/filepath"
	"strconv"
	s "strings"
)

func main() {
	scanner := util.FileScanner(filepath.Join("day4", "input"))

	var count int
	for scanner.Scan() {
		contents := scanner.Text()
		splits := s.Split(s.Replace(contents, ",", "-", -1), "-")

		a, _ := strconv.Atoi(splits[0])
		b, _ := strconv.Atoi(splits[1])
		x, _ := strconv.Atoi(splits[2])
		y, _ := strconv.Atoi(splits[3])

		if (x >= a && x <= b) || (y <= b && y >= a) || (a >= x && a <= y) || (b <= y && b >= x) {
			count++
			fmt.Println("The pair", contents, "overlap each other")
		}
	}

	fmt.Println("There are total", count, "pairs which overlap")
}
