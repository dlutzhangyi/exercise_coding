package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	result := []int{}
	for index, num := range nums {
		need := target - num
		if _, ok := sumMap[need]; ok {
			result = append(result, sumMap[need], index)
			return result
		}
		sumMap[num] = index
	}
	return result
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
