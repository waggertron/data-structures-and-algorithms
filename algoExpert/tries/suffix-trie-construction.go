package tries

// Do not edit the class below except for the
// PopulateSuffixTrieFrom and Contains methods.
// Feel free to add new properties and methods
// to the class.
type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	// Write your code here.
	letters := []byte(str)
	for i := 0; i < len(letters); i++ {
		curTrie := trie
		curLetters := letters[i:]
		for _, char := range curLetters {
			letter := byte(char)
			nextTrie, found := curTrie[letter]
			if !found {
				newTrie := NewSuffixTrie()
				curTrie[letter] = newTrie
				nextTrie = newTrie
			}

			curTrie = nextTrie
		}
		curTrie['*'] = nil
	}

}

func (trie SuffixTrie) Contains(str string) bool {
	// Write your code here.
	curTrie := trie
	for _, char := range str {
		letter := byte(char)
		if t, found := curTrie[letter]; found {
			curTrie = t
			continue
		}
		return false
	}

	if _, found := curTrie['*']; found {
		return true
	}

	return false
}
func (trie SuffixTrie) Within(str string) bool {
	// Write your code here.
	curTrie := trie
	for _, char := range str {
		letter := byte(char)
		if t, found := curTrie[letter]; found {
			curTrie = t
			continue
		}
		return false
	}

	return true
}
