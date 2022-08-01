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

// Create bags just with color filled in
func parseEmptyBags(in string) *Bag {
	s := strings.Split(in, " bags contain ")
	curCol := s[0]
	return &Bag{color: curCol, referencedIn: []*Bag{}, canContain: map[*Bag]int{}}
}

// Link up what bags are contained in whic + backreferences
func linkBags(in string, bags map[string]*Bag) {
	// No links
	if strings.Contains(in, "no other bags") {
		return
	}

	// Get current bag
	noDot := in[:len(in)-1]
	s := strings.Split(noDot, " bags contain ")
	curCol := s[0]
	curBag, ok := bags[curCol]
	if !ok {
		log.Fatal("Invalid bag color.")
	}

	containsSplits := strings.Split(s[1], ",")
	// Get out other bags
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
		// Add bag to our links
		curBag.canContain[otherBag] = num
		// Add backreference to current bag in other bag
		otherBag.referencedIn = append(otherBag.referencedIn, curBag)
	}
}

// Parse input lines to bags mapped by name
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

// Calculate in how many bags current bag is contained
func (b *Bag) calculateContainedIn(bags map[string]*Bag, lookup map[string]int) int {
	result := 0
	for _, o := range b.referencedIn {
		_, checked := lookup[o.color]
		if checked {
			// If we already encountered outer bag, do not go over it again
			continue
		}

		// Find bag we are referenced in
		otherBag, ok := bags[o.color]
		if !ok {
			log.Fatal("Bag not found in chain.")
		}
		// Add other bag to lookup
		lookup[otherBag.color] = 0

		// Add self to result
		result += 1
		// Add the amount of bags the outer bag is contained in
		result += otherBag.calculateContainedIn(bags, lookup)
	}
	return result
}

// Calculate how many bags are inside this bag
func (b *Bag) calculateContained(bags map[string]*Bag) int {
	result := 0
	// Calculate how many bags current bag contains
	for oB, c := range b.canContain {
		obc := oB.calculateContained(bags)
		result += c * obc
	}
	// Add self to bag count
	return result + 1
}

// Calculate backreference chain count of bag to search
func calculateCanBeFoundIn(bags map[string]*Bag, search string) int {
	bag, ok := bags[search]
	if !ok {
		log.Fatal("Looking for non-existent bag.")
	}

	// Create lookup map
	lookup := make(map[string]int)
	return bag.calculateContainedIn(bags, lookup)
}

// Calculate total amount of bags a given bag contains
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
