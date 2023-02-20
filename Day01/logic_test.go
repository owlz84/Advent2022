package Day01

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	result := ComputeMaxCalories("example.txt", 1)
	assert.Equal(t, result, 24000)
}

func TestFull(t *testing.T) {
	result := ComputeMaxCalories("full.txt", 1)
	fmt.Println("Result for part 1: ", result)
}

func TestTop3(t *testing.T) {
	result := ComputeMaxCalories("full.txt", 3)
	fmt.Println("Result for part 2: ", result)
}
