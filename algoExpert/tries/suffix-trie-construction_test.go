package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiStringSearch(t *testing.T) {
	inputs := []struct {
		values   []string
		expected bool
	}{
		{[]string{"abc", "abc"}, true},
		{[]string{"abb", "abc"}, false},
		{[]string{"dabc", "abc"}, true},
		{[]string{"dabb", "abc"}, false},
		{[]string{"babc", "bab"}, false},
	}

	for _, input := range inputs {
		values := input.values

		tr := NewSuffixTrie()
		tr.PopulateSuffixTrieFrom(values[0])
		result := tr.Contains(values[1])
		assert.Equal(t, input.expected, result)
	}
}
