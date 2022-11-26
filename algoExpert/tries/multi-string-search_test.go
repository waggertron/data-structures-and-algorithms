package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type suffixTrieTestInput struct {
	sentence   string
	substrings []string
	expected   []bool
}

func TestNewSuffixTrieContains(t *testing.T) {

	testCases := []suffixTrieTestInput{
		{
			"hello there friend",
			[]string{"hello", "there", "friend", "enemy", "goodbye"},
			[]bool{true, true, true, false, false},
		},
	}

	for _, testCase := range testCases {
		result := MultiStringSearch(testCase.sentence, testCase.substrings)

		assert.ElementsMatch(t, result, testCase.expected)
		// assert.Equal(t, input.expected, result)
	}
}
