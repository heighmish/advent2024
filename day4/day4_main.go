package main

import (
	"advent2024/utils"
	"fmt"
)

type Direction struct {
	X int
	Y int
}

var counter = 0
var counterPart2 = 0

func main() {
	input := utils.ReadLines("input.txt")

	Part1(input)
	Part2(input)

}

func Part2(input []string) {
	diagonals := []Direction{
		{X: -1, Y: 1},  // NE
		{X: -1, Y: -1}, // NW
		{X: 1, Y: 1},   // SE
		{X: 1, Y: -1},  // SW
	}

	for iRow, row := range input {
		for iCol, colValue := range row {
			if colValue == 'A' {
				mCount, sCount := 0, 0
				for _, dir := range diagonals {
					diagonalRune := SearchPart2(input, iRow+dir.X, iCol+dir.Y)
					if diagonalRune == 'M' {
						mCount++
					} else if diagonalRune == 'S' {
						sCount++
					}
				}
				if mCount == 2 && sCount == 2 {
					if SearchPart2(input, iRow+1, iCol+1) != SearchPart2(input, iRow-1, iCol-1) {
						counterPart2++
					}
				}

			}
		}
	}
	fmt.Println("Part 2: found matching searches: ", counterPart2)

}

func SearchPart2(input []string, rowIndex int, colIndex int) rune {

	if rowIndex < 0 || rowIndex > len(input)-1 || colIndex < 0 || colIndex > len(input[rowIndex])-1 {
		return 0
	}

	return rune(input[rowIndex][colIndex])
}

func Part1(input []string) {
	directions := []Direction{
		{X: -1, Y: -1}, // NW
		{X: -1, Y: 0},  // N
		{X: -1, Y: 1},  // NE
		{X: 0, Y: -1},  // W
		{X: 0, Y: 1},   // E
		{X: 1, Y: -1},  // SW
		{X: 1, Y: 0},   // S
		{X: 1, Y: 1},   // SE
	}

	for iRow, row := range input {
		for iCol, colValue := range row {
			if colValue == 'X' {
				// fmt.Println("Starting new search")
				for _, dir := range directions {
					Search(input, iRow+dir.X, iCol+dir.Y, 'M', dir)
				}
			}
		}
	}
	fmt.Println("Part 1: found matching searches: ", counter)
}

func Search(input []string, rowIndex int, colIndex int, searchChar rune, direction Direction) {
	// fmt.Printf("Searching rowIndex %d, colIndex %d, searchChar %c \n", rowIndex, colIndex, searchChar)

	if rowIndex < 0 || rowIndex > len(input)-1 || colIndex < 0 || colIndex > len(input[rowIndex])-1 {
		return
	}

	currCharacter := input[rowIndex][colIndex]

	if currCharacter != byte(searchChar) {
		return
	}

	if currCharacter == 'S' {
		counter++
		return
	}

	var nextSearch rune

	if searchChar == 'M' {
		nextSearch = 'A'
	} else if searchChar == 'A' {
		nextSearch = 'S'
	} else {
		panic("invalid searchChar received")
	}

	Search(input, rowIndex+direction.X, colIndex+direction.Y, nextSearch, direction)
}
