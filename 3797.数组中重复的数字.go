package main

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
