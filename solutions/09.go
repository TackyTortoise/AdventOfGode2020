package main

import (
	helpers "AoG/helpers"
	"fmt"
	"math"
	"strconv"
)

type pair struct {
	first  int
	second int
}

type sumParts struct {
	sum int
	pair
}

func removeFromSumMap(num int, sumMapping []sumParts) []sumParts {
	for i := 0; i < len(sumMapping); i++ {
		p := sumMapping[i]
		if p.first == num || p.second == num {
			sumMapping = append(sumMapping[:i], sumMapping[i+1:]...)
			i--
		}
	}
	return sumMapping
}

func addToSumMap(num int, sumMapping []sumParts, prevNums []int) []sumParts {
	for _, n := range prevNums {
		p := pair{num, n}
		sumMapping = append(sumMapping, sumParts{num + n, p})
	}
	return sumMapping
}

func Day9Part1() int {
	lines := helpers.ReadFileLines("./input/09.txt")
	nums := make([]int, len(lines))
	for i, l := range lines {
		num, _ := strconv.Atoi(l)
		nums[i] = num
	}
	const preamble = 25
	sumMapping := []sumParts{}
	// Parse first preamble
	for i := 0; i < preamble; i++ {
		for j := i + 1; j < preamble; j++ {
			iN := nums[i]
			jN := nums[j]
			p := pair{iN, jN}
			sumMapping = append(sumMapping, sumParts{iN + jN, p})
		}

	}

	for i := preamble; i < len(nums); i++ {
		ok := false
		// Check if we have the sum in sum map
		for _, sp := range sumMapping {
			if sp.sum == nums[i] {
				ok = true
				break
			}
		}

		// If not we have the answer
		if !ok {
			fmt.Printf("Answer: %d\n", nums[i])
			return nums[i]
		}

		// Remove first considered num from sum map
		sumMapping = removeFromSumMap(nums[i-preamble], sumMapping)
		// Add sumParts with new number to sumMapping
		sumMapping = addToSumMap(nums[i], sumMapping, nums[i-preamble+1:i])
	}
	fmt.Println("Failed")
	return -1
}

func Day9Part2() {
	lines := helpers.ReadFileLines("./input/09.txt")
	nums := make([]int, len(lines))
	for i, l := range lines {
		num, _ := strconv.Atoi(l)
		nums[i] = num
	}

	// Searching for result of day 9
	toSearch := Day9Part1()
	for i := 0; i < len(nums); i++ {
		seq := []int{}
		sum := 0
		// Calculate sum starting from i
		for j := i; j < len(nums); j++ {
			seq = append(seq, nums[j])
			sum += nums[j]
			// If we went over the sum, stop
			if sum >= toSearch {
				break
			}
		}
		// If we met exactly the sum, we have winning sequence
		if sum == toSearch {
			// Calculate min and max
			min := math.MaxInt
			max := 0
			for _, s := range seq {
				if s < min {
					min = s
				}
				if s > max {
					max = s
				}
			}
			// Print result
			fmt.Printf("First: %d, Last: %d, Result: %d\n", min, max, min+max)
			return
		}
	}
	fmt.Println("Failed")
}
