package maxsum

import "testing"

func setupTestCase() ([][]int, []int) {
	testCase := [][]int{
		{1, -2, 4, 5, -1, 1},
		{-1, 1, -1},
		{},
		{-1},
		{1, -10, 12, 4, -5},
	}
	reult := []int{9, 1, 0, -1, 16}
	return testCase, reult
}
func TestMaxSumSubArray(t *testing.T) {
	testCase, result := setupTestCase()
	for i := 0; i < len(testCase); i++ {
		sum := FindMaxSumSubArray(testCase[i])
		if sum != result[i] {
			t.Fatalf("required sum:%d,actual sum:%d", result[i], sum)
		}
	}
}
