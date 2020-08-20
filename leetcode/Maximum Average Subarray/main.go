package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/maximum-average-subarray-i/

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := make([]int, len(nums))
	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	average := sum[k-1]
	for i := k; i < len(nums); i++ {
		current := sum[i] - sum[i-k]
		if current > average {
			average = current
		}
	}
	return float64(average) / float64(k)
}

func main() {
	num := []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage(num, 4))
}
