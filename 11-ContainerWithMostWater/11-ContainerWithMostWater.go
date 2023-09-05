package leetcode

import "math"

func maxArea(height []int) int {
    if (len(height) == 0) {
        return 0
    }
    l := 0
    r := len(height) - 1
    max := 0
    for l < r {
        length := r - l
        var min int
        if height[l] > height[r] {
            min = height[r]
            r--
        } else {
            min = height[l]
            l++
        }
        area := length * min
        max = int(math.Max(float64(max), float64(area)))
    }
    return max
}
