package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic("Unable to open input.txt file")
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}
