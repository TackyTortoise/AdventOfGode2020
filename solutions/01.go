package main

import (
	helpers "AoG/helpers"
	"fmt"
)

func findInSlice[T comparable](s []T, find T, startIndex int) bool {
	for i := startIndex; i < len(s); i++ {
		if s[i] == find {
			return true
		}
	}
	return false
}

func Day_01_part1() {
	lines := helpers.ReadFileLines("./input/01.txt")
	nums := helpers.StringSliceToIntSlice(lines)

	const target = 2020
	for i, num := range nums {
		find := target - num
		if findInSlice[int](nums, find, i+1) {
			fmt.Println(num * find)
			return
		}
	}
}

type intPair struct {
	first  int
	second int
}

func Day_01_part2() {
	lines := helpers.ReadFileLines("./input/01.txt")
	nums := helpers.StringSliceToIntSlice(lines)

	const target = 2020
	sumMap := make(map[int]intPair)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			pair := intPair{nums[i], nums[j]}
			sum := pair.first + pair.second
			if sum > 2020 {
				continue
			}
			//fmt.Printf("%d + %d = %d\n", pair.first, pair.second, sum)
			sumMap[sum] = pair
		}
	}

	for _, num := range nums {
		find := target - num
		value, ok := sumMap[find]
		if ok {
			mult := num * value.first * value.second
			fmt.Println(mult)
			return
		}
	}
}
