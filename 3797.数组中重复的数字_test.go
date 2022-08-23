package main

import (
	"fmt"
)

func Test_findRepeatNumber() {
	fmt.Sprintln(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	cache := make([]int, 1000)
	for _, num := range nums {
		if cache[num] > 0 {
			return num
		}
		cache[num]++
	}
	return 0
}
