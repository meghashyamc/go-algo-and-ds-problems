package backtracking

import (
	"sort"
)

// Given a set of distinct integers, S, return all possible subsets.

//  Note:
// Elements in a subset must be in non-descending order.
// The solution set must not contain duplicate subsets.
// Also, the subsets should be sorted in ascending ( lexicographic ) order.
// The list is not necessarily sorted.
// Example :

// If S = [1,2,3], a solution is:

// [
//   [],
//   [1],
//   [1, 2],
//   [1, 2, 3],
//   [1, 3],
//   [2],
//   [2, 3],
//   [3],
// ]

func subsets(A []int) [][]int {
	if A == nil {
		return [][]int{}
	}

	result := [][]int{}
	subsetsHelper(A, len(A)-1, &result)
	sort.Slice(result, func(i, j int) bool {
		return compareElements(result, i, j)
	})
	return result

}

func compareElements(result [][]int, i, j int) bool {

	if len(result[i]) == 0 {
		return true
	}
	if len(result[j]) == 0 {
		return false
	}

	for p := 0; ; p++ {
		if p == len(result[i]) {
			return true
		}
		if p == len(result[j]) {
			return false
		}

		if result[i][p] == result[j][p] {
			continue
		}
		if result[i][p] < result[j][p] {
			return true
		}
		return false
	}
}

func subsetsHelper(arr []int, lastindex int, result *[][]int) {

	if lastindex < 0 {
		*result = append(*result, []int{})
		return
	}

	//get all subsets till lastindex-1
	subsetsHelper(arr, lastindex-1, result)
	lenOfResult := len(*result)
	//for each of the gotten subsets, append the element at the last index
	//then add the newly formed subsets to the list of subsets
	for i := 0; i < lenOfResult; i++ {
		currsubset := make([]int, len((*result)[i]))
		copy(currsubset, (*result)[i])
		currsubset = append(currsubset, arr[lastindex])
		sort.Slice(currsubset, func(i, j int) bool {
			return currsubset[i] < currsubset[j]
		})
		*result = append(*result, currsubset)
	}

}
