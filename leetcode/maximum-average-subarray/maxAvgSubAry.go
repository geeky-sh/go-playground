package maximumaveragesubarray

// [1, 2, 3]

func findMaxAverage(nums []int, k int) float64 {
	var maxAvg float64

	s := 0
	for i := 0; i < k; i++ {
		s += nums[i]
	}
	maxAvg = float64(s) / float64(k)

	for i := k; i < len(nums); i++ {
		s += nums[i] - nums[i-k]

		avg := float64(s) / float64(k)
		if avg > maxAvg {
			maxAvg = avg
		}
	}
	return maxAvg
}
