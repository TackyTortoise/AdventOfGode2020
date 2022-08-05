package main

import (
	helpers "AoG/helpers"
	"fmt"
	"log"
	"sort"
)

func Day10Part1() {
	nums := helpers.ReadFileLinesAsInt("./input/10.txt")
	sort.Ints(nums)
	if nums[0] > 3 {
		log.Fatalf("Invalid first jolt %d", nums[0])
	}
	oneJumps := 0
	threeJumps := 0
	currJolt := 0
	for _, n := range nums {
		if n-currJolt > 3 {
			log.Fatalf("Encountered invalid jolt %d.", n)
		}
		switch n - currJolt {
		case 1:
			oneJumps++
		case 3:
			threeJumps++
		}
		currJolt = n
	}
	// Final three jolt to device
	threeJumps++
	result := oneJumps * threeJumps
	fmt.Println("Result: ", result)
}

var storage map[int]int

func getPossiblities(input []int) int {
	if len(input) == 1 {
		return 1
	}

	num := input[0]
	v, ok := storage[num]
	if ok {
		return v
	}
	input = input[1:]
	p := 0
	for i, n := range input {
		if n-num <= 3 {
			p += getPossiblities(input[i:])
		}
	}
	storage[num] = p
	return p
}

func Day10Part2() {
	storage = make(map[int]int)
	nums := helpers.ReadFileLinesAsInt("./input/10.txt")
	sort.Ints(nums)
	nums = append([]int{0}, nums...)
	nums = append(nums, nums[len(nums)-1]+3)
	result := getPossiblities(nums)

	fmt.Println("Result: ", result)
}
