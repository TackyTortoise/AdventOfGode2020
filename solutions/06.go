package main

import (
	helpers "AoG/helpers"
	"fmt"
)

func parseGroups(input []string) [][]string {
	groups := [][]string{}
	curGroup := []string{}
	for _, l := range input {
		if len(l) == 0 {
			groups = append(groups, curGroup)
			curGroup = []string{}
			continue
		}
		curGroup = append(curGroup, l)
	}
	groups = append(groups, curGroup)
	return groups
}

func countAnyoneAnswers(groups [][]string) int {
	result := 0
	for _, g := range groups {
		counts := map[rune]int{}
		for _, a := range g {
			for _, c := range a {
				counts[c]++
			}
		}
		result += len(counts)
	}
	return result
}

func countEveryoneAnsers(groups [][]string) int {
	result := 0
	for _, g := range groups {
		counts := map[rune]int{}
		for _, a := range g {
			for _, c := range a {
				counts[c]++
			}
		}
		for _, c := range counts {
			if c == len(g) {
				result++
			}
		}
	}
	return result
}

func Day6Part1() {
	lines := helpers.ReadFileLines("./input/06.txt")
	groups := parseGroups(lines)
	fmt.Println(countAnyoneAnswers(groups))
}

func Day6Part2() {
	lines := helpers.ReadFileLines("./input/06.txt")
	groups := parseGroups(lines)
	fmt.Println(countEveryoneAnsers(groups))
}
