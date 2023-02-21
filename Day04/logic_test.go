package Day04

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	result := FullyContainsCount("example.txt")
	assert.Equal(t, 2, result)
}

func TestFull(t *testing.T) {
	result := FullyContainsCount("full.txt")
	fmt.Println("Part 1 result: ", result)
}

func TestPart2(t *testing.T) {
	result := PartlyContainsCount("example.txt")
	assert.Equal(t, 4, result)
}

func TestPart2Full(t *testing.T) {
	result := PartlyContainsCount("full.txt")
	fmt.Println("Part 2 result: ", result)
}
