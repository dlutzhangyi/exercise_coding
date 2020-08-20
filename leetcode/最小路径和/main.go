package main

import "fmt"

//https://leetcode-cn.com/problems/minimum-path-sum/

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i > 0 && j > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[rows-1][cols-1]
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])

	dp := make([]int, cols)
	dp[0] = grid[0][0]
	for j := 1; j < cols; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j > 0 {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			} else {
				dp[j] = dp[j] + grid[i][j]
			}
		}
	}
	return dp[cols-1]
}

func main() {
	var grid = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
