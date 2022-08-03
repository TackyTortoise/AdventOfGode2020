package main

import (
	helpers "AoG/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	ins   string
	value int
}

func (i instruction) String() string {
	return fmt.Sprintf("Ins: %s, V: %d", i.ins, i.value)
}

func handleInstructions(instructions []instruction) (int, bool) {
	acc := 0
	visited := []int{}
	for i := 0; ; {
		// Instructions successfully completed
		if i >= len(instructions) {
			return acc, true
		}

		ins := instructions[i]
		// If we already encountered instruction, we are in a loop
		for _, v := range visited {
			if v == i {
				return acc, false
			}
		}

		// Handle instruction
		visited = append(visited, i)
		switch ins.ins {
		case "acc":
			acc += ins.value
			i++
		case "jmp":
			i += ins.value
		case "nop":
			i++
		}
	}
}

func findFix(instructions []instruction) int {
	for i, ins := range instructions {
		// Ignore instructions that we won't change
		if ins.ins == "acc" {
			continue
		}

		// Copy instruction
		copied := make([]instruction, len(instructions))
		copy(copied, instructions)
		// Swap current instruction
		if ins.ins == "jmp" {
			copied[i].ins = "nop"
		} else if ins.ins == "nop" {
			// Jump 0 would not do anything, ignore
			if ins.value == 0 {
				continue
			}
			copied[i].ins = "jmp"
		} else {
			log.Fatal("Something went wrong.")
		}

		// Run instructions
		r, s := handleInstructions(copied)
		// If instruction finished, success!
		if s {
			return r
		}
	}
	log.Fatal("Infinite loop")
	return -1
}

func parseInstructions(lines []string) []instruction {
	instructions := []instruction{}
	for _, l := range lines {
		split := strings.Split(l, " ")
		instr := split[0]
		value, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		new := instruction{instr, value}
		instructions = append(instructions, new)
	}
	return instructions
}

func Day8Part1() {
	lines := helpers.ReadFileLines("./input/08.txt")
	instructions := parseInstructions(lines)
	result, _ := handleInstructions(instructions)
	fmt.Printf("Result: %d", result)
}

func Day8Part2() {
	lines := helpers.ReadFileLines("./input/08.txt")
	instructions := parseInstructions(lines)
	result := findFix(instructions)
	fmt.Printf("Result: %d", result)
}
