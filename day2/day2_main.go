package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func main() {
	input := ReadInput()
	safeCount := 0
	for _, value := range input {
		split := strings.Fields(value)
		if len(split) == 0 {
			continue
		}
		safeCount += IsSafe(split)
	}
	fmt.Println("The number of safe reports is: ", safeCount)
}

func IsSafe(report []string) int {
	firstValue, _ := strconv.Atoi(report[0])
	lastValue, _ := strconv.Atoi(report[1])
	isIncreasing := false

	if firstValue < lastValue {
		isIncreasing = true
	}

	if CheckSafeDifference(firstValue, lastValue) == false {
		return 0
	}

	for _, value := range report[2:] {
		currValue, _ := strconv.Atoi(value)
		isSafeDifference := CheckSafeDifference(lastValue, currValue)
		if isSafeDifference == false {
			return 0
		}

		if lastValue < currValue {
			if !isIncreasing {
				return 0
			}
		} else {
			if isIncreasing {
				return 0
			}
		}

		lastValue = currValue
	}
	fmt.Println(report, "Is Safe")
	return 1
}

func CheckSafeDifference(a int, b int) bool {
	difference := a - b
	if difference < 0 {
		difference *= -1
	}
	if difference > 3 || difference < 1 {
		return false
	}

	return true
}
