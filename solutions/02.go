package main

import (
	helpers "AoG/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parseRequirement(requirement string) (int, int, rune) {
	splits := strings.Split(requirement, " ")
	minMax := strings.Split(splits[0], "-")
	min, minErr := strconv.Atoi(minMax[0])
	if minErr != nil {
		log.Fatal(minErr)
	}
	max, maxErr := strconv.Atoi(minMax[1])
	if maxErr != nil {
		log.Fatal(maxErr)
	}
	return min, max, rune(splits[1][0])
}

func isValidPassword(input string) bool {
	splits := strings.Split(input, ": ")
	min, max, char := parseRequirement(splits[0])
	passwd := splits[1]
	count := 0
	for _, c := range passwd {
		if c == char {
			count++
		}
	}
	return count >= min && count <= max
}

func isValidPassword2(input string) bool {
	splits := strings.Split(input, ": ")
	p1, p2, char := parseRequirement(splits[0])
	passwd := splits[1]
	hasPos1 := passwd[p1-1] == byte(char)
	hasPos2 := passwd[p2-1] == byte(char)
	return (hasPos1 || hasPos2) && !(hasPos1 && hasPos2)
}

func Day2Part1() {
	lines := helpers.ReadFileLines("./input/02.txt")
	validCount := 0
	for _, line := range lines {
		if isValidPassword(line) {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func Day2Part2() {
	lines := helpers.ReadFileLines("./input/02.txt")
	validCount := 0
	for _, line := range lines {
		if isValidPassword2(line) {
			validCount++
		}
	}
	fmt.Println(validCount)
}
