package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var max float64 = float64(-1 << 63)
	var sum int
	for i, v := range nums {
		if i >= k {
			sum -= nums[i-k]
		}
		sum += v
		if i >= k-1 {
			max = math.Max(max, float64(sum)/float64(k))
		}
	}
	return roundFloat(max, 5)
}

func roundFloat(val float64, precision int) float64 {
	ratio := math.Pow10(precision)
	return math.Round(val*ratio) / ratio
}
