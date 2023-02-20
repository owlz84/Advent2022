package Day03

import (
	"bufio"
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
