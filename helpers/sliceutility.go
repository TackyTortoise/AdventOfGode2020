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
