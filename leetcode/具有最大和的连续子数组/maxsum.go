package maxsum

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func FindMaxSumSubArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	dp := make([]int, len(array))
	dp[0] = array[0]
	maxSum := array[0]
	for i := 1; i < len(array); i++ {
		dp[i] = max(dp[i-1]+array[i], array[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}
