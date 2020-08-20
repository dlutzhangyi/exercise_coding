package main

import "fmt"

func sort(array []int, start int, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	key := array[start]
	for i < j {
		if i < j && array[j] >= key {
			j--
		}
		array[i] = array[j]
		if i < j && array[i] <= key {
			i++
		}
		array[j] = array[i]
	}
	array[i] = key
	sort(array, start, i-1)
	sort(array, i+1, end)
}

func sort1(array []int, start int, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	key := array[start]
	for i < j {
		for i < j && array[j] >= key {
			j--
		}
		array[i] = array[j]
		for i < j && array[i] < key {
			i++
		}
		array[j] = array[i]
	}
	array[i] = key
	sort(array, start, i-1)
	sort(array, i+1, end)
}
func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	sort1(array, 0, len(array)-1)
	return array
}

func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	newArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}

		} else {
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}

func main() {
	//array := []int{5, 1, 1, 2, 0, 0, 10, 4, 3}
	array := []int{5, 2, 3, 1}
	fmt.Println(QuickSort(array))
	fmt.Println(MergeSort(array))
}
