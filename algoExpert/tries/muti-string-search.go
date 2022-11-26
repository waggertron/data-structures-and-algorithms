package tries

func MultiStringSearch(bigString string, smallStrings []string) []bool {
	// Write your code here.

	t := NewSuffixTrie()
	t.PopulateSuffixTrieFrom(bigString)

	result := make([]bool, len(smallStrings))
	for i, subStr := range smallStrings {
		if t.Within(subStr) {
			result[i] = true
		}
	}

	return result
}
