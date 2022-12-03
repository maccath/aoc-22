package util

import (
	"bufio"
	"os"
	"path/filepath"
)

func FileScanner(fp string) *bufio.Scanner {
	pwd, _ := os.Getwd()
	f, err := os.Open(filepath.Join(pwd, fp))
	Check(err)

	return bufio.NewScanner(f)
}
