package leetcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func doPart1(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	l := strings.TrimSpace(line)
	if len(l) == 0 {
		return
	}
	id, games := parse(line)
	for _, r := range games {
		if (r.Blue > 0 && r.Blue > 14) || (r.Green > 0 && r.Green > 13) || (r.Red > 0 && r.Red > 12) {
			return
		}
	}
	ch <- int(id)
}

func doPart2(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	l := strings.TrimSpace(line)
	if len(l) == 0 {
		return
	}
	_, games := parse(line)
	b, r, g := 1, 1, 1
	for _, re := range games {
		if re.Blue > b {
			b = re.Blue
		}

		if re.Red > r {
			r = re.Red
		}

		if re.Green > g {
			g = re.Green
		}
	}
	ch <- b * g * r
}

func parse(line string) (int, []Record) {
	var grid []Record
	re := regexp.MustCompile(`Game (?P<Id>.*): (?P<Games>.*)`)
	idIndex := re.SubexpIndex("Id")
	gamesIndex := re.SubexpIndex("Games")
	matches := re.FindStringSubmatch(line)
	id, err := strconv.Atoi(matches[idIndex])
	if err != nil {
		fmt.Printf("error converting id to num: %v\n", err)
		return 0, grid
	}
	gamesLine := matches[gamesIndex]
	games := parseGameLine(gamesLine)
	for _, game := range games {
		grid = append(grid, parseRecord(game))
	}
	return id, grid
}

func parseGameLine(l string) []string {
	return strings.FieldsFunc(l, func(r rune) bool {
		return r == ';'
	})
}

type Record struct {
	Red   int
	Green int
	Blue  int
}

func parseRecord(l string) Record {
	var record Record
	re := regexp.MustCompile(` ?(?P<num>.*) (?P<name>.*)`)
	numIndex := re.SubexpIndex("num")
	nameIndex := re.SubexpIndex("name")
	cubeLines := strings.FieldsFunc(l, func(r rune) bool {
		return r == ','
	})
	for _, cubeLine := range cubeLines {
		matches := re.FindStringSubmatch(cubeLine)
		num, err := strconv.Atoi(matches[numIndex])
		if err != nil {
			fmt.Printf("error converting to num: %v\n", err)
			return record
		}
		name := matches[nameIndex]
		switch name {
		case "blue":
			record.Blue = num
		case "green":
			record.Green = num
		case "red":
			record.Red = num
		}
	}
	return record
}
