package slidingWindows

/*
[1,3,-1,-3,5,3,6,7] k=3
[3,3,5,5,6,7]
*/

func GetMaxValueInSlidingWindow(array []int, k int) []int {
	if len(array) == 0 {
		return []int{}
	}
	size := len(array)
	if size < k {
		return []int{}
	}
	result := []int{}
	for i := 0; i <= size-k; i++ {
		maxValue := array[i]
		for j := i; j < i+k; j++ {
			if array[j] > maxValue {
				maxValue = array[j]
			}
		}
		result = append(result, maxValue)
	}
	return result
}
