package dynamicprogramming

func fibNaive(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibNaive(n - 1)
}

var memo = map[int]int{1: 1, 2: 1}

func fibMemo(n int) int {
	v, ok := memo[n]

	if !ok {
		res := fibMemo(n-1) + fibMemo(n-2)
		
		memo[n] = res
		return res
	}

	return v
}

