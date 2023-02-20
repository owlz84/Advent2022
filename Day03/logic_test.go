package Day03

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	result := PrioritySum("example.txt")
	assert.Equal(t, 157, result)
}

func TestFull(t *testing.T) {
	result := PrioritySum("full.txt")
	fmt.Println("Part 1 result: ", result)
}
