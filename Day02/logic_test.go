package Day02

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	result := ComputeExpectedScore("example.txt")
	assert.Equal(t, 15, result)
}

func TestFull(t *testing.T) {
	result := ComputeExpectedScore("full.txt")
	fmt.Println("Part 1 result: ", result)
}

func TestDetermineAction(t *testing.T) {
	result := DetermineAction("example.txt")
	assert.Equal(t, 12, result)
}

func TestDetermineActionFull(t *testing.T) {
	result := DetermineAction("full.txt")
	fmt.Println("Part 2 result: ", result)
}
