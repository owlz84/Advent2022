package Day04

import (
	"Advent01/Day03"
	"strconv"
	"strings"
)

func checkFullyContains(l1 int, l2 int, h1 int, h2 int) bool {
	if l1 <= l2 && h1 >= h2 {
		return true
	} else if l2 <= l1 && h2 >= h1 {
		return true
	}
	return false
}

func checkPartlyContains(l1 int, l2 int, h1 int, h2 int) bool {
	if l1 <= l2 && h1 >= l2 {
		return true
	} else if l2 <= l1 && h2 >= l1 {
		return true
	}
	return false
}

func FullyContainsCount(filename string) int {
	count := 0
	fileLines := Day03.ReadFile(filename)
	for _, line := range fileLines {
		assignments := strings.Split(line, ",")
		e1 := assignments[0]
		e2 := assignments[1]
		e1Bounds := strings.Split(e1, "-")
		e2Bounds := strings.Split(e2, "-")
		e1Low, _ := strconv.Atoi(e1Bounds[0])
		e1High, _ := strconv.Atoi(e1Bounds[1])
		e2Low, _ := strconv.Atoi(e2Bounds[0])
		e2High, _ := strconv.Atoi(e2Bounds[1])
		if checkFullyContains(e1Low, e2Low, e1High, e2High) {
			count += 1
		}
	}
	return count
}

func PartlyContainsCount(filename string) int {
	count := 0
	fileLines := Day03.ReadFile(filename)
	for _, line := range fileLines {
		assignments := strings.Split(line, ",")
		e1 := assignments[0]
		e2 := assignments[1]
		e1Bounds := strings.Split(e1, "-")
		e2Bounds := strings.Split(e2, "-")
		e1Low, _ := strconv.Atoi(e1Bounds[0])
		e1High, _ := strconv.Atoi(e1Bounds[1])
		e2Low, _ := strconv.Atoi(e2Bounds[0])
		e2High, _ := strconv.Atoi(e2Bounds[1])
		if checkPartlyContains(e1Low, e2Low, e1High, e2High) {
			count += 1
		}
	}
	return count
}
