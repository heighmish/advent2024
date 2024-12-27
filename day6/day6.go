package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

type Guard struct {
	location  Point
	direction Direction
}

var (
	Up    = Direction{X: -1, Y: 0}
	Down  = Direction{X: 1, Y: 0}
	Left  = Direction{X: 0, Y: -1}
	Right = Direction{X: 0, Y: 1}
)

const (
	OBSTACLE    = 0
	EMPTY_SPACE = 1
	COMPLETE    = 2
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	seenSet := make(map[Point]bool)
	guard := InitStart(input)
	fmt.Printf("Starting at (%d,%d) Facing: %d\n", guard.location.X, guard.location.Y, guard.direction)
	for {
		// fmt.Printf("Guard at (%d,%d) Facing: %d\n", guard.location.X, guard.location.Y, guard.direction)
		seenSet[guard.location] = true
		nextMove := guard.DetermineNextMove()
		nextMoveType := IsValidSpace(input, nextMove)
		if nextMoveType == COMPLETE {
			break
		} else if nextMoveType == EMPTY_SPACE {
			guard.location = nextMove
		} else {
			guard.direction = guard.TurnRight()
		}
	}
	fmt.Printf("Part1 guard visited %d unique locations\n", len(seenSet))
}

func (g *Guard) TurnRight() Direction {
	if g.direction == Up {
		return Right
	} else if g.direction == Down {
		return Left
	} else if g.direction == Left {
		return Up
	} else {
		return Down
	}
}

func (g *Guard) DetermineNextMove() Point {
	return Point{
		X: g.location.X + g.direction.X,
		Y: g.location.Y + g.direction.Y,
	}

}

func IsValidSpace(input []string, point Point) int {
	if point.X < 0 || point.X > len(input)-1 || point.Y < 0 || point.Y > len(input[0])-1 {
		return COMPLETE
	}

	if input[point.X][point.Y] == '#' {
		return OBSTACLE
	}

	return EMPTY_SPACE
}

func InitStart(input []string) Guard {
	var guard Guard
	for iRow, row := range input {
		for iCol, char := range row {
			guard.location.X = iRow
			guard.location.Y = iCol
			if char == '^' {
				guard.direction = Up
				return guard
			} else if char == '<' {
				guard.direction = Left
				return guard
			} else if char == '>' {
				guard.direction = Right
				return guard
			} else if char == 'v' {
				guard.direction = Down
				return guard
			}
		}
	}

	panic("Unable to find starting position")
}
