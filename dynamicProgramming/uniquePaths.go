package dynamicprogramming

func UniquePathsRecursive(m, n int) int {
	if m == 0 && n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 2 && n == 1 {
		return 1
	}
	if m == 1 && n == 2 {
		return 1
	}

	result := 0

	if m != 1 {
		result += UniquePathsRecursive(m-1, n)
	}
	if n != 1 {
		result += UniquePathsRecursive(m, n-1)
	}

	return result

}

func UniquePathsDynamic(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	table := make([][]int, n+1)
	for i := range table {
		table[i] = make([]int, m+1)
	}
	for i := 0; i < len(table); i++ {
		if len(table[i]) == 0 {
			table[i] = make([]int, m+1)
		}

		row := table[i]
		for j := 0; j < len(row); j++ {
			val := 0
			if i <= 1 {
				if i == 1 {
					val = 1
				}
			} else {
				if j <= 1 {
					if j == 1 {
						val = 1
					}
				} else {
					val = table[i][j-1] + table[i-1][j]
				}
			}
			table[i][j] = val
		}
	}

	return table[n][m]

}
