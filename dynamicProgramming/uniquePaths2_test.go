package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths2(t *testing.T) {

	inputs := []struct {
		values   [][]int
		expected int
	}{
		{
			[][]int{
				{0, 0},
				{0, 0},
			},
			2,
		},
	}

	for _, input := range inputs {
		values := input.values
		result := UniquePaths2(values)

		assert.Equal(t, input.expected, result)
	}
}

// assert equality
// assert.Equal(t, 123, 123, "they should be equal")
