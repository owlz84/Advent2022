package Day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ComputeMaxCalories(filename string, scope int) int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err = readFile.Close()
	if err != nil {
		return 0
	}

	var elfFood []int

	totalCalories := 0

	for _, line := range fileLines {
		if line == "" {
			elfFood = append(elfFood, totalCalories)
			totalCalories = 0
		}
		calories := 0
		calories, err = strconv.Atoi(line)
		totalCalories += calories
	}

	sort.Ints(elfFood)

	result := 0
	for i := scope; i > 0; i-- {
		result += elfFood[len(elfFood)-i]
	}

	return result
}
