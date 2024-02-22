package strings

func isUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if _, got := seen[c]; got {
			return false
		}

		seen[c] = true
	}

	return true
}
