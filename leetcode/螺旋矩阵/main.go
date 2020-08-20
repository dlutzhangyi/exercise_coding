package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	result := []int{}

	rows := len(matrix)
	if rows == 0 {
		return result
	}
	cols := len(matrix[0])
	if cols == 0 {
		return result
	}

	top, left := 0, 0
	buttom, right := rows-1, cols-1
	for top <= buttom && left <= right {
		temp := printMatrix(matrix, top, left, buttom, right)
		result = append(result, temp...)
		top = top + 1
		buttom = buttom - 1
		left = left + 1
		right = right - 1
	}

	return result
}

func printMatrix(matrix [][]int, top, left, buttom, right int) []int {
	i, j := top, left
	result := []int{}
	for j <= right {
		result = append(result, matrix[i][j])
		j++
	}
	i = top + 1
	j = right
	for i <= buttom {
		result = append(result, matrix[i][j])
		i++
	}
	j = right - 1
	i = buttom
	for j >= left && i > top {
		result = append(result, matrix[i][j])
		j--
	}
	i = buttom - 1
	j = left
	for i >= (top+1) && j < right {
		result = append(result, matrix[i][j])
		i--
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	result := spiralOrder(matrix)
	fmt.Println(result)
}
