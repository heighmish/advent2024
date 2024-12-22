package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running Advent2024 Day 1")
	Part1()
	Part2()
}

func ReadInput() []string {
	file, err := os.Open("input.txt")
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

func ParseInputToListsOfInts(input []string) ([]int, []int) {
	var array1, array2 []int
	for _, value := range input {
		var split = strings.Fields(value)
		if len(split) == 0 {
			continue
		}
		array1 = append(array1, UnsafeCastToInt(split[0]))
		array2 = append(array2, UnsafeCastToInt(split[1]))
	}
	return array1, array2
}

func Part1() {
	var input = ReadInput()
	fmt.Println("Length of input read is: ", len(input))
	array1, array2 := ParseInputToListsOfInts(input)

	sort.Ints(array1)
	sort.Ints(array2)
	var distance = 0
	for i := range array1 {
		var dist = array1[i] - array2[i]
		if dist < 0 {
			dist *= -1
		}

		distance += dist
	}

	fmt.Println("Part 1: Answer is: ", distance)
}

func Part2() {
	var input = ReadInput()
	fmt.Println("Length of input read is: ", len(input))
	array1, array2 := ParseInputToListsOfInts(input)

	var counts = make(map[int]int)
	for _, v := range array2 {
		counts[v] += 1
	}

	var similarity = 0
	for _, v := range array1 {
		similarity += (v * counts[v])
	}

	fmt.Println("Part 2: Answer is: ", similarity)
}

func UnsafeCastToInt(value string) int {
	cast, err := strconv.Atoi(value)
	if err != nil {
		panic("Failed to convert string to int")
	}
	return cast
}
