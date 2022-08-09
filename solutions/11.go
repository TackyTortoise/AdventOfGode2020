package main

import (
	"AoG/helpers"
	"fmt"
	"strings"
)

type plane struct {
	width  int
	height int
	seats  string
}

func (p *plane) printSeats() {
	for i := 0; i < len(p.seats); i += p.width {
		fmt.Println(p.seats[i : i+p.width])
	}
}

func (p *plane) toIndex(x, y int) int {
	if x < 0 || y < 0 {
		return -1
	}
	return x + y*p.width
}

func (p *plane) toCoord(i int) (int, int) {
	return i % p.width, i / p.width
}

func (p *plane) getAtCoord(x, y int) byte {
	i := p.toIndex(x, y)
	return p.seats[i]
}

func (p *plane) getOccupiedNeigboorCount(i int) int {
	result := 0
	checkNeighbourResult := func(x, y int) {
		v := p.getAtCoord(x, y)
		if v == '#' {
			result++
		}
	}
	x, y := p.toCoord(i)
	// Left
	if x > 0 {
		checkNeighbourResult(x-1, y)
		// Left up
		if y > 0 {
			checkNeighbourResult(x-1, y-1)
		}
		// Left down
		if y < p.height-1 {
			checkNeighbourResult(x-1, y+1)
		}
	}
	// Up
	if y > 0 {
		checkNeighbourResult(x, y-1)
	}
	// Right
	if x < p.width-1 {
		checkNeighbourResult(x+1, y)
		// Right up
		if y > 0 {
			checkNeighbourResult(x+1, y-1)
		}
		// Right down
		if y < p.height-1 {
			checkNeighbourResult(x+1, y+1)
		}
	}
	// Down
	if y < p.height-1 {
		checkNeighbourResult(x, y+1)
	}
	return result
}

func (p *plane) isValidCoord(x, y int) bool {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return false
	}
	i := p.toIndex(x, y)
	return helpers.IsValidIndex(p.seats, i)
}

func (p *plane) getOccupiedNeigboorCountFar(i int) int {
	result := 0
	directions := []intPair{
		{-1, 0},  // L
		{1, 0},   // R
		{0, -1},  // U
		{0, 1},   // D
		{-1, -1}, // LU
		{-1, 1},  // LD
		{1, -1},  // RU
		{1, 1},   // RD
	}
	x, y := p.toCoord(i)
	for _, d := range directions {
		cX, cY := x+d.first, y+d.second
		for p.isValidCoord(cX, cY) {
			cI := p.toIndex(cX, cY)
			if p.seats[cI] == '#' {
				result++
				break
			} else if p.seats[cI] == 'L' {
				break
			}

			cX += d.first
			cY += d.second
		}
	}
	return result
}

func (p *plane) advanceStep(p2 bool) bool {
	newSeats := make([]byte, len(p.seats))
	anyChange := false

	emptyTh := 4
	if p2 {
		emptyTh = 5
	}

	for i, v := range p.seats {
		var nC int
		if !p2 {
			nC = p.getOccupiedNeigboorCount(i)
		} else {
			nC = p.getOccupiedNeigboorCountFar(i)
		}
		if v == 'L' && nC == 0 {
			newSeats[i] = '#'
			anyChange = true
		} else if v == '#' && nC >= emptyTh {
			newSeats[i] = 'L'
			anyChange = true
		} else {
			newSeats[i] = byte(v)
		}
	}
	p.seats = string(newSeats)
	return anyChange
}

func newPlane(input []string) *plane {
	return &plane{width: len(input[0]), height: len(input), seats: strings.Join(input, "")}
}

func Day11Part1() {
	lines := helpers.ReadFileLines("./input/11.txt")
	p := newPlane(lines)
	anyChange := true
	for anyChange {
		anyChange = p.advanceStep(false)
		p.printSeats()
	}
	result := strings.Count(p.seats, "#")
	fmt.Println("Result: ", result)
}

func Day11Part2() {
	lines := helpers.ReadFileLines("./input/11.txt")
	p := newPlane(lines)
	anyChange := true
	for anyChange {
		anyChange = p.advanceStep(true)
	}
	result := strings.Count(p.seats, "#")
	fmt.Println("Result: ", result)
}
