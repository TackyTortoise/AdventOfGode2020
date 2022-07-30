package main

import (
	helpers "AoG/helpers"
	"fmt"
)

func plotCourse(field []string, h, v int) int {
	width := len(field[0])
	maxV := len(field)

	cX := h
	count := 0
	for y := v; y < maxV; y += v {
		currChar := field[y][cX%width]
		cX += h
		if currChar == '#' {
			count++
		}
	}
	return count
}

func Day3Part1() {
	lines := helpers.ReadFileLines("./input/03.txt")
	result := plotCourse(lines, 3, 1)
	fmt.Println(result)
}

func Day3Part2() {
	lines := helpers.ReadFileLines("./input/03.txt")
	a := plotCourse(lines, 1, 1)
	b := plotCourse(lines, 3, 1)
	c := plotCourse(lines, 5, 1)
	d := plotCourse(lines, 7, 1)
	e := plotCourse(lines, 1, 2)
	result := a * b * c * d * e
	fmt.Println(result)
}
