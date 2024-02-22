package arrays

import (
	"fmt"
	"sort"
)

// What is the largest subset whose elements are Fibonacci numbers
func FindFibSubset(nums []int) []int {
	var failed = []int{}

	if len(nums) < 3 {
		return failed
	}

	largestFibSet := make([]int, 3)

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))
	highest := sortedNums[0]

	fibs := map[int]int{
		1: 1,
		2: 1,
	}

	fibSet := map[int]bool{1: true}

	for n := 3; ; {
		fibN := fibs[n-1] + fibs[n-2]

		if fibN > highest {
			break
		}

		fibs[n] = fibN
		fibSet[fibN] = true
	}

	for _, n := range sortedNums {
		if len(largestFibSet) >= 3 {
			break
		}

		if _, got := fibSet[n]; got {
			largestFibSet = append(largestFibSet, n)
		}
	}

	if len(largestFibSet) < 3 {
		return failed
	}

	fmt.Println(fibs)

	return largestFibSet
}
