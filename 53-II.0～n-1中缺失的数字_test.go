package main

func Test_missingNumber() {

}

func missingNumber(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] != mid {
			r = mid
		} else if nums[mid] > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
