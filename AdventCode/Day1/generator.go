package leetcode

import (
	"sync"
	"unicode"
)

func doPart1(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var cur, l int
	r := len(line) - 1
	var lb, rb bool

	for l < len(line) && (!lb || !rb) {
		if !unicode.IsDigit(rune(line[l])) && !unicode.IsDigit(rune(line[r])) {
			l++
			r--
			continue
		}
		if !lb && unicode.IsDigit(rune(line[l])) {
			cur += int(line[l]-'0') * 10
			lb = true
		}
		if !rb && unicode.IsDigit(rune(line[r])) {
			cur += int(line[r] - '0')
			rb = true
		}
		l++
		r--
	}
	ch <- cur
}

func doPart2(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	tr := initTrieTree()
	var left, right, l int
	r := len(line) - 1
	var lb, rb bool

	for l < len(line) && (!lb || !rb) {
		if !lb {
			if unicode.IsDigit(rune(line[l])) {
				left = int(line[l] - '0')
				lb = true
			} else {
				if v, ok := tr.FindFirst(line[l:]); ok {
					left = v
					lb = true
				}
			}
		}
		if !rb {
			if unicode.IsDigit(rune(line[r])) {
				right = int(line[r] - '0')
				rb = true
			} else {
				if v, ok := tr.FindFirst(line[r:]); ok {
					right = v
					rb = true
				}
			}
		}
		l++
		r--
	}
	ch <- (left * 10) + right
}

func initTrieTree() *TrieTree {
	tr := NewTrieTree()
	tr.AddWord("zero", 0)
	tr.AddWord("one", 1)
	tr.AddWord("two", 2)
	tr.AddWord("three", 3)
	tr.AddWord("four", 4)
	tr.AddWord("five", 5)
	tr.AddWord("six", 6)
	tr.AddWord("seven", 7)
	tr.AddWord("eight", 8)
	tr.AddWord("nine", 9)
	return &tr
}
