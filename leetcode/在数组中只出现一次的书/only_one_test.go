package onlyone

import "testing"

func generateTestCase() ([][]int, []int) {
	testCase := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
	}
	result := []int{1, 4}
	return testCase, result
}

func generateTestCase2() ([][]int, []int) {
	testCase := [][]int{
		{2, 2, 1, 3},
		{4, 1, 2, 1, 2, 5},
	}
	result := []int{1, 4}
	return testCase, result
}

func TestValueAppearOnceInArray(t *testing.T) {
	testCase, result := generateTestCase()
	for i := 0; i < len(testCase); i++ {
		num := ValueAppearOnceInArray(testCase[i])
		if num != result[i] {
			t.Fatalf("required:%d actual:%d", result[i], num)
		}
	}
}

func TestValueAppearTwiceInArray(t *testing.T) {
	array := []int{2, 2, 1, 3}
	result := map[int]bool{
		1: true,
		3: true,
	}
	v1, v2 := ValueAppearTwiceInArray(array)
	if _, ok := result[v1]; !ok {
		t.Fatalf("v1:%d", v1)
	}
	if _, ok := result[v2]; !ok {
		t.Fatalf("v2:%d", v2)
	}
}
