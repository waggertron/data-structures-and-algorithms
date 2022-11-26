package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {

	inputs := []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
		{[]int{2, 2}, 2},
		{[]int{2, 3}, 3},
		{[]int{3, 2}, 3},
		{[]int{3, 7}, 28},
	}

	for _, input := range inputs {
		values := input.values
		result := UniquePathsRecursive(values[0], values[1])

		assert.Equal(t, input.expected, result)

		values = input.values
		result = UniquePathsDynamic(values[0], values[1])

		assert.Equal(t, input.expected, result)
	}
}

// assert equality
// assert.Equal(t, 123, 123, "they should be equal")
