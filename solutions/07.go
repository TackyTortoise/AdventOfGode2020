package main

import (
	helpers "AoG/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Bag struct {
	color        string
	referencedIn []*Bag
	canContain   map[*Bag]int
}

func parseEmptyBags(in string) *Bag {
	s := strings.Split(in, " bags contain ")
	curCol := s[0]
	return &Bag{color: curCol, referencedIn: []*Bag{}, canContain: map[*Bag]int{}}
}

func linkBags(in string, bags map[string]*Bag) {
	if strings.Contains(in, "no other bags") {
		return
	}

	noDot := in[:len(in)-1]
	s := strings.Split(noDot, " bags contain ")
	curCol := s[0]
	curBag, ok := bags[curCol]
	if !ok {
		log.Fatal("Invalid bag color.")
	}
	containsSplits := strings.Split(s[1], ",")
	for _, c := range containsSplits {
		t := c
		t = strings.ReplaceAll(t, "bags", "")
		t = strings.ReplaceAll(t, "bag", "")
		t = strings.TrimSpace(t)
		num, _ := strconv.Atoi(string(t[0]))
		color := t[2:]

		otherBag, oOk := bags[color]
		if !oOk {
			log.Fatal("Contains invalid bag.")
		}
		curBag.canContain[otherBag] = num
		otherBag.referencedIn = append(otherBag.referencedIn, curBag)
	}
}

func createBags(input []string) map[string]*Bag {
	bags := make(map[string]*Bag, len(input))
	for _, i := range input {
		bag := parseEmptyBags(i)
		bags[bag.color] = bag
	}
	for _, i := range input {
		linkBags(i, bags)
	}
	return bags
}

func (b *Bag) calculateContainedIn(bags map[string]*Bag, lookup map[string]int) int {
	result := 0
	for _, o := range b.referencedIn {
		_, checked := lookup[o.color]
		if checked {
			continue
		}
		otherBag, ok := bags[o.color]
		if !ok {
			log.Fatal("Bag not found in chain.")
		}
		lookup[otherBag.color] = 0
		result += 1
		result += otherBag.calculateContainedIn(bags, lookup)
	}
	return result
}

func (b *Bag) calculateContained(bags map[string]*Bag) int {
	result := 0
	for oB, c := range b.canContain {
		obc := oB.calculateContained(bags)
		if obc > 0 {
			result += c * obc
		} else {
			result += c
		}
	}
	return result + 1
}

func calculateCanBeFoundIn(bags map[string]*Bag, search string) int {
	bag, ok := bags[search]
	if !ok {
		log.Fatal("Looking for non-existent bag.")
	}

	lookup := make(map[string]int)
	return bag.calculateContainedIn(bags, lookup)
}

func calculateTotalBagsContained(bags map[string]*Bag, search string) int {
	bag, ok := bags[search]
	if !ok {
		log.Fatal("Looking for non-existent bag.")
	}
	return bag.calculateContained(bags)
}

func Day7Part1() {
	lines := helpers.ReadFileLines("./input/07.txt")
	bags := createBags(lines)
	result := calculateCanBeFoundIn(bags, "shiny gold")
	fmt.Printf("Result: %d", result)
}

func Day7Part2() {
	lines := helpers.ReadFileLines("./input/07.txt")
	bags := createBags(lines)
	result := calculateTotalBagsContained(bags, "shiny gold")
	fmt.Printf("Result: %d", result-1)
}
