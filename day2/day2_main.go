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

func Part1(input []string) {
	safeCount := 0
	for _, value := range input {
		split := strings.Fields(value)
		if len(split) == 0 {
			continue
		}
		safeCount += IsSafe(split)
	}
	fmt.Println("Part1: the number of safe reports is: ", safeCount)

}

func Part2(input []string) {
	safeCount := 0
	for _, value := range input {
		split := strings.Fields(value)
		if len(split) == 0 {
			continue
		}
		safeCount += Part2IsSafe(split)
	}
	fmt.Println("Part2: the number of safe reports is: ", safeCount)
}

func Part2IsSafe(report []string) int {
	// [0, 1, 2, 3, 4]
	// i = 0, 0:0, 1:5 [1, 2, 3, 4]
	// i = 1, 0:1, 2:5 [0, 2, 3, 4]
	// i = 2, 0:2, 3:5 [0, 1, 3, 4]
	// i = 3, 0:3, 4:5 [0, 1, 2, 4]
	// i = 4, 0:4, 5:5 [0, 1, 2, 3]
	for i := range report {
		list := append(report[:i:i], report[i+1:]...)
		fmt.Println("Checking if list is safe: ", list)
		if IsSafe(list) == 1 {
			return 1
		}
	}
	return 0
}

func main() {
	input := ReadInput()
	Part1(input)
	Part2(input)
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
	// fmt.Println(report, "Is Safe")
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
