package Day03

import (
	"bufio"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

func ReadFile(filename string) []string {
	readFile, _ := os.Open(filename)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err := readFile.Close()
	if err != nil {
		return nil
	}
	return fileLines
}

func Contains(target []rune, source rune) bool {
	for _, t := range target {
		if t == source {
			return true
		}
	}
	return false
}

func RuneToAscii(r rune) int {
	runeVal := int(r)
	if runeVal > 96 {
		runeVal -= 96
	} else {
		runeVal -= 65 - 27
	}
	return runeVal
}

func PrioritySum(filename string) int {
	fileLines := ReadFile(filename)
	overlap := 0
	for _, line := range fileLines {
		lineLength := len(line)
		r1 := []rune(line[0 : lineLength/2])
		r2 := []rune(line[lineLength/2 : lineLength])
		for _, r := range r1 {
			if Contains(r2, r) {
				overlap += RuneToAscii(r)
				break
			}
		}
	}
	return overlap
}

func BadgePrioritySum(filename string) int {
	fileLines := ReadFile(filename)
	prioritySum := 0
	for i := 0; i < len(fileLines)/3; i++ {
		offset := i * 3
		var (
			e1 = mapset.NewSet[string]()
			e2 = mapset.NewSet[string]()
			e3 = mapset.NewSet[string]()
		)
		for i := 0; i < len(fileLines[offset]); i++ {
			e1.Add(string(fileLines[offset][i]))
		}
		for i := 0; i < len(fileLines[offset+1]); i++ {
			e2.Add(string(fileLines[offset+1][i]))
		}
		for i := 0; i < len(fileLines[offset+2]); i++ {
			e3.Add(string(fileLines[offset+2][i]))
		}
		common, _ := e1.Intersect(e2.Intersect(e3)).Pop()
		prioritySum += RuneToAscii(rune(common[0]))
	}
	return prioritySum
}
