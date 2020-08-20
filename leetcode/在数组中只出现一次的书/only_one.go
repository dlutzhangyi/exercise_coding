package main

import "fmt"

func ValueAppearOnceInArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	num := array[0]
	for i := 1; i < len(array); i++ {
		num = num ^ array[i]
	}
	return num
}

func getLeftShiftValue(len int) int {
	value := 1
	for len > 0 {
		value = value * 2
		len--
	}
	return value
}
func ValueAppearTwiceInArray(array []int) (int, int) {
	if len(array) == 0 {
		return 0, 0
	}
	num := array[0]
	for i := 1; i < len(array); i++ {
		num = num ^ array[i]
	}
	i := 1
	for num > 0 {
		if num&1 != 0 {
			break
		} else {
			num = num >> 1
			fmt.Println(num)
			i++
		}
	}
	digits := getLeftShiftValue(i - 1)
	subArray1 := []int{}
	subArray2 := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]&digits != 0 {
			subArray1 = append(subArray1, array[i])
		} else {
			subArray2 = append(subArray2, array[i])
		}
	}
	return ValueAppearOnceInArray(subArray1), ValueAppearOnceInArray(subArray2)
}

func main() {
	array := []int{2, 2, 1, 3}

	v1, v2 := ValueAppearTwiceInArray(array)
	fmt.Println(v1, v2)
}
