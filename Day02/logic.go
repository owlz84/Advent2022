package Day02

import (
	"bufio"
	"os"
	"strings"
)

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var outcomes = map[string]int{
	"AX": 3,
	"AY": 6,
	"AZ": 0,
	"BX": 0,
	"BY": 3,
	"BZ": 6,
	"CX": 6,
	"CY": 0,
	"CZ": 3,
}

var updatedOutcomes = map[string]string{
	"AX": "Z",
	"AY": "X",
	"AZ": "Y",
	"BX": "X",
	"BY": "Y",
	"BZ": "Z",
	"CX": "Y",
	"CY": "Z",
	"CZ": "X",
}

var updatedScores = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

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

func ComputeExpectedScore(filename string) int {

	totalScore := 0
	//matcher, _ := regexp.Compile("([A-C]) ([X-Z])")
	for _, line := range ReadFile(filename) {
		//hands := matcher.FindStringSubmatch(line)
		hands := strings.Split(line, " ")
		totalScore += scores[hands[1]]
		totalScore += outcomes[hands[0]+hands[1]]
	}

	return totalScore
}

func DetermineAction(filename string) int {

	totalScore := 0
	for _, line := range ReadFile(filename) {
		hands := strings.Split(line, " ")
		myMove := updatedOutcomes[hands[0]+hands[1]]
		totalScore += scores[myMove]
		totalScore += updatedScores[hands[1]]
	}
	return totalScore
}
