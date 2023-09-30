package leetcode

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func maxVowels(s string, k int) int {
	var cur, max int
	for i, v := range s {
		if i >= k {
			if vowels[rune(s[i-k])] {
				cur--
			}
		}
		if vowels[v] {
			cur++
		}
		if cur > max {
			max = cur
		}
	}

	return max
}
