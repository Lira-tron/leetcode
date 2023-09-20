package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var max int = math.MinInt
	var sum int
	for i, v := range nums {
		if i >= k {
			sum -= nums[i-k]
		}
		sum += v
		if i >= k-1 && sum > max {
			max = sum
		}
	}
	return roundFloat(float64(max)/float64(k), 5)
}

func roundFloat(val float64, precision int) float64 {
	ratio := math.Pow10(precision)
	return math.Round(val*ratio) / ratio
}
