package main

import (
	"AoG/helpers"
	"fmt"
	"log"
	"math"
	"strconv"
)

func getDirection(dir byte) (int, int) {
	switch dir {
	case 'N':
		return 0, -1
	case 'E':
		return 1, 0
	case 'S':
		return 0, 1
	case 'W':
		return -1, 0
	default:
		panic("invalid direction")
	}
}

type ship struct {
	dir int
	pos intPair
}

func (s *ship) move(ins byte, num int) {
	directions := []byte{'N', 'E', 'S', 'W'}

	if ins == 'R' {
		amount := num / 90
		s.dir = (s.dir + amount) % len(directions)
		return
	} else if ins == 'L' {
		amount := num / 90
		s.dir -= amount
		for s.dir < 0 {
			s.dir += len(directions)
		}
		return
	}

	currDir := ins
	if ins == 'F' {
		currDir = directions[s.dir]
	}
	x, y := getDirection(currDir)
	s.pos.first += x * num
	s.pos.second += y * num
}

type wayPoint intPair

// Moves relative position to ship
func (w *wayPoint) move(dir byte, num int) {
	x, y := getDirection(dir)
	w.first += x * num
	w.second += y * num
}

// Rotate waypoint, always assume origin
func (w *wayPoint) rotate(dir byte, degrees int) {
	amount := degrees / 90
	localX, localY := w.first, w.second

	// Rotate
	for i := 0; i < amount; i++ {
		if dir == 'R' {
			oX := localX
			localX = -localY
			localY = oX
		} else if dir == 'L' {
			oX := localX
			localX = localY
			localY = -oX
		} else {
			panic("Invalid rotation")
		}
	}
	w.first = localX
	w.second = localY
}

func Day12Part1() {
	lines := helpers.ReadFileLines("./input/12.txt")
	s := ship{dir: 1}
	for _, l := range lines {
		ins := l[0]
		num, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatal(err)
		}
		s.move(ins, num)
	}
	result := math.Abs(float64(s.pos.first)) + math.Abs(float64(s.pos.second))
	fmt.Println("Result: ", result)
}

func Day12Part2() {
	lines := helpers.ReadFileLines("./input/12.txt")
	shipPos := intPair{}
	wp := wayPoint{first: 10, second: -1}
	for _, l := range lines {
		ins := l[0]
		num, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatal(err)
		}
		if ins == 'F' {
			shipPos.first += num * wp.first
			shipPos.second += num * wp.second
		} else if ins == 'R' || ins == 'L' {
			wp.rotate(ins, num)
		} else {
			wp.move(ins, num)
		}
	}
	result := math.Abs(float64(shipPos.first)) + math.Abs(float64(shipPos.second))
	fmt.Println("Result: ", result)
}
