package main

import (
	helpers "AoG/helpers"
	"fmt"
	"log"
	"sort"
)

func splitRange(min, max int, upper bool) (int, int) {
	spread := max - min + 1
	if upper {
		min += spread / 2
	} else {
		max -= spread / 2
	}
	return min, max
}

func calcRow(input string) int {
	min := 0
	max := 127
	//fmt.Printf("Row min %d, max %d\n", min, max)
	for _, i := range input {
		var upper bool
		if i == 'F' {
			upper = false
		} else if i == 'B' {
			upper = true
		} else {
			log.Fatal("Invalid row input")
		}
		min, max = splitRange(min, max, upper)
		//fmt.Printf("Row min %d, max %d\n", min, max)
	}
	if min != max {
		log.Fatal("Did not reach sincle row in calculation")
	}
	//fmt.Printf("Row result: %d\n", min)
	return min
}

func calcColumn(input string) int {
	min := 0
	max := 7
	//fmt.Printf("Column min %d, max %d\n", min, max)
	for _, i := range input {
		var upper bool
		if i == 'L' {
			upper = false
		} else if i == 'R' {
			upper = true
		} else {
			log.Fatal("Invalid column input")
		}
		min, max = splitRange(min, max, upper)
		//fmt.Printf("Column min %d, max %d\n", min, max)
	}
	if min != max {
		log.Fatal("Did not reach sincle column in calculation")
	}
	//fmt.Printf("Column result: %d\n", min)
	return min
}

func calcSeatId(input string) int {
	rowPart := input[:7]
	columnPart := input[7:]
	row := calcRow(rowPart)
	column := calcColumn(columnPart)
	seatId := row*8 + column
	//fmt.Printf("Seat id: %d\n", seatId)
	return seatId
}

func Day5Part1() {
	lines := helpers.ReadFileLines("./input/05.txt")
	max := 0
	for _, l := range lines {
		id := calcSeatId(l)
		if id > max {
			max = id
		}
	}
	fmt.Printf("Max id: %d\n", max)
}

func Day5Part2() {
	lines := helpers.ReadFileLines("./input/05.txt")
	var ids []int
	for _, l := range lines {
		id := calcSeatId(l)
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for i, id := range ids {
		if i >= len(ids) {
			log.Fatal("Did not find result.")
		}
		if id == ids[i+1]-2 {
			fmt.Printf("Your id: %d", id+1)
			return
		}
	}
}
