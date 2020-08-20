package slidingWindows

import "testing"

func generateTestCase() ([][]int, []int, [][]int) {
	testArray := [][]int{
		{1, 3, -1, -3, 5, 3, 6, 7},
	}
	kArray := []int{3}
	required := [][]int{
		{3, 3, 5, 5, 6, 7},
	}
	return testArray, kArray, required
}

func compareTwoSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestGetMaxValueInSlidingWindow(t *testing.T) {
	testArray, kArray, required := generateTestCase()
	for i := 0; i < len(testArray); i++ {
		actual := GetMaxValueInSlidingWindow(testArray[i], kArray[i])
		if !compareTwoSlice(actual, required[i]) {
			t.Fatalf("required:%v actual:%v", required[i], actual)
		}
	}
}
