package backtracking

import (
	"sort"
)

// Given two integers n and k, return all possible combinations of k numbers out of 1 2 3 ... n.

// Make sure the combinations are sorted.

// To elaborate,

// Within every entry, elements should be sorted. [1, 4] is a valid entry while [4, 1] is not.
// Entries should be sorted within themselves.
// Example :
// If n = 4 and k = 2, a solution is:

// [
//   [1,2],
//   [1,3],
//   [1,4],
//   [2,3],
//   [2,4],
//   [3,4],
// ]

func combine(A int, B int) [][]int {

	if B > A {
		return [][]int{}
	}

	result := combineHelper(A, B)
	sort.Slice(result, func(i, j int) bool {
		return compareElements(result, i, j)
	})
	return result
}

func combineHelper(n, k int) [][]int {
	if n < k {
		return [][]int{}
	}

	if k == 1 {
		result := make([][]int, 0)
		for i := 1; i <= n; i++ {
			currCombination := make([]int, 0)
			currCombination = append(currCombination, i)
			result = append(result, currCombination)
		}
		return result
	}

	temp := combineHelper(n-1, k-1)
	result := make([][]int, 0)
	for i := 0; i < len(temp); i++ {
		currCombination := temp[i]
		currCombination = append(currCombination, n)
		result = append(result, currCombination)

	}

	result = append(result, combineHelper(n-1, k)...)
	return result
}
