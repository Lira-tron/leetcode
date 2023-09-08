package leetcode

func maxOperations(nums []int, k int) int {
	var counter int
	mapPos := getMapPositions(nums, k)
	for _, v := range nums {
		if v >= k || mapPos[v] <= 0 {
			continue
		}
		mapPos[v]--
		if val, ok := mapPos[k-v]; ok && val > 0 {
			counter++
			mapPos[k-v]--
		}
	}
	return counter
}

func getMapPositions(nums []int, k int) map[int]int {
	m := make(map[int]int)

	for _, v := range nums {
		if v < k {
			m[v]++
		}
	}
	return m
}
