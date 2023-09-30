package leetcode

func longestOnes(nums []int, k int) int {
	var l, r, flip, max int
	for r < len(nums) {
		if nums[r] == 0 {
			flip++
		}

		if flip > k {
			for nums[l] == 1 {
				l++
			}
			l++
			flip--
		}
		r++
		max = maxOf(max, r-l)
	}
	return max
}

func maxOf(l int, r int) int {
	if l > r {
		return l
	}
	return r
}
