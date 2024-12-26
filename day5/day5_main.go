package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	file, _ := os.Open(inputFile)

	reader := bufio.NewScanner(file)

	pageOrderingRules := false
	pageOrdering := make(map[string][]string)

	var reports [][]string

	for reader.Scan() {
		line := reader.Text()
		// fmt.Printf("Read line: %s\n", line)
		if line == "" {
			pageOrderingRules = true
			continue
		}

		if !pageOrderingRules {
			pair := strings.Split(line, "|")
			// fmt.Printf("Split line on '|' pair: %s\n", pair)
			pageOrdering[pair[0]] = append(pageOrdering[pair[0]], pair[1])
		} else {
			report := strings.Split(line, ",")
			// fmt.Println(report)
			if len(report) > 0 {
				reports = append(reports, report)
			}
		}
	}
	p1MiddleNumberSum := 0
	for _, report := range reports {
		if IsReportValid(report, pageOrdering) {
			if len(report)%2 == 0 {
				panic("EVEN REPORT LENGTH")
			} else {
				middleIndex := (len(report) / 2)
				conv, _ := strconv.Atoi(report[middleIndex])
				// fmt.Printf("Adding %d to answer \n", conv)
				p1MiddleNumberSum += conv
			}
		}
	}

	fmt.Println("Part 1: Solution", p1MiddleNumberSum)

	p2Sol := 0
	for _, report := range reports {
		if !IsReportValid(report, pageOrdering) {
			p2Sol += Part2(report, pageOrdering)
		}
	}

	fmt.Println("Part 2: Solution", p2Sol)
}

func IsReportValid(report []string, lookupTable map[string][]string) bool {
	// fmt.Println(report)
	for i, pageNumber := range report {
		lookup := lookupTable[pageNumber]
		// fmt.Println(lookup)
		for backwardsSearchPointer := i - 1; backwardsSearchPointer >= 0; backwardsSearchPointer-- {
			searchValue := report[backwardsSearchPointer]
			// fmt.Printf("Searching lookupTable %s for searchValue %s \n", lookup, searchValue)
			if slices.Index(lookup, searchValue) != -1 {
				return false
			}
		}
	}
	return true
}

func Part2(report []string, lookupTable map[string][]string) int {
	// fmt.Println("Starting new report...")
	// fmt.Println(report)
	for i := 0; i < len(report); i++ {
		lookup := lookupTable[report[i]]
		for backwardsSearchPointer := i - 1; backwardsSearchPointer >= 0; backwardsSearchPointer-- {
			searchValue := report[backwardsSearchPointer]
			searchIndex := slices.Index(lookup, searchValue)
			if searchIndex != -1 {
				// fmt.Printf("SearchValue: %s, i: %d, backwardsSearchPointer %d\n", searchValue, i, backwardsSearchPointer)
				for shift := i; shift > backwardsSearchPointer; shift-- {
					// fmt.Printf("Shifting %s with %s\n", report[shift], report[shift-1])
					tmp := report[shift]
					report[shift] = report[shift-1]
					report[shift-1] = tmp

					i = backwardsSearchPointer
					break
				}
			}
		}
	}
	// fmt.Println("Final report value...")
	// fmt.Println(report)

	middleIndex := (len(report) / 2)
	conv, _ := strconv.Atoi(report[middleIndex])
	return conv
}
