package helpers

import (
	"log"
	"strconv"
)

func StringSliceToIntSlice(input []string) []int {
	var nums []int
	for _, l := range input {
		num, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func IsValueInSlice[T comparable](s []T, f T) bool {
	for _, v := range s {
		if v == f {
			return true
		}
	}
	return false
}

func IsValidIndex[T string | []any](s T, i int) bool {
	return i >= 0 && i < len(s)
}
