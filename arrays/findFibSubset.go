package arrays

import "sort"

var failed = []int{}

// What is the largest subset whose elements are Fibonacci numbers
func FindFibSubset(nums []int) []int {
	if len(nums) < 3 {
		return failed
	}

	largestFibSet := make([]int, 3)

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sortedNums = sort.IntSlice(sortedNums)

	highest := sortedNums[0]

	fibs := map[int]int{
		1: 1,
		2: 1,
	}

	for n := 3; ; {
		fibN := fibs[n-1] + fibs[n-2]

		if fibN > highest {
			break
		}

		fibs[n] = fibN
	}

	fibSet := make(map[int]bool)
	for _, v := range fibs {
		fibSet[v] = true
	}

	cur := 0
	for _, n := range sortedNums {
		if cur > 2 {
			break
		}

		_, got := fibSet[n]
		if got {
			largestFibSet = append(largestFibSet, n)
			cur++
		}
	}

	if len(largestFibSet) < 3 {
		return failed
	}

	return largestFibSet
}
