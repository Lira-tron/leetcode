package leetcode

func longestSubarray(nums []int) int {
	var max, l, r, flip int
	zindex := -1
	for r < len(nums) {
		if nums[r] == 0 {
			flip++
			if zindex < 0 {
				zindex = r
			}
		}
		if flip > 1 {
			if r > 0 && nums[r-1] == 0 && nums[r] == 0 {
				l = r
			} else {
				l = zindex + 1
			}
			zindex = r
			flip--
		}
		r++
		max = maxOf(max, r-l)
	}
	return max - 1
}

func maxOf(i, j int) int {
	if i > j {
		return i
	}
	return j
}
