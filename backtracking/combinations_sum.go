package backtracking

import (
	"sort"
	"strconv"
	"strings"
)

/*Given an array of candidate numbers A and a target number B, find all unique combinations in A where the candidate numbers sums to B.

The same repeated number may be chosen from A unlimited number of times.

Note:

1) All numbers (including target) will be positive integers.

2) Elements in a combination (a1, a2, … , ak) must be in non-descending order. (ie, a1 ≤ a2 ≤ … ≤ ak).

3) The combinations themselves must be sorted in ascending order.

4) CombinationA > CombinationB iff (a1 > b1) OR (a1 = b1 AND a2 > b2) OR ... (a1 = b1 AND a2 = b2 AND ... ai = bi AND ai+1 > bi+1)

5) The solution set must not contain duplicate combinations.





Problem Constraints
1 <= |A| <= 20

1 <= A[i] <= 50

1 <= B <= 500



Input Format
The first argument is an integer array A.

The second argument is integer B.



Output Format
 Return a vector of all combinations that sum up to B.



Example Input
Input 1:

A = [2, 3]
B = 2
Input 2:

A = [2, 3, 6, 7]
B = 7


Example Output
Output 1:

[ [2] ]
Output 2:

[ [2, 2, 3] , [7] ]


Example Explanation
Explanation 1:

All possible combinations are listed.
Explanation 2:

All possible combinations are listed.*/

func getCombinationsWithSum(arr []int, sum int) [][]int {

	sort.Slice(arr, (func(i, j int) bool {
		return arr[i] < arr[j]
	}))
	combinationsWithSum := [][]int{}

	getCombinationsWithSumHelper(arr, []int{}, 0, sum, &combinationsWithSum, map[string]bool{})

	return combinationsWithSum
}

func getCombinationsWithSumHelper(arr, currentArr []int, currentSum, targetSum int, combinationsWithSum *[][]int, combinationUniqueStringsMap map[string]bool) {

	if arr == nil || len(arr) == 0 {
		return
	}
	if currentSum > targetSum {
		return
	}

	if currentSum == targetSum {
		uniqueStringForCombination := formUniqueStringForCombination(currentArr)

		if !combinationUniqueStringsMap[uniqueStringForCombination] {
			combinationUniqueStringsMap[uniqueStringForCombination] = true
			*combinationsWithSum = append(*combinationsWithSum, currentArr)
		}
		return
	}

	for _, el := range arr {
		currentArrCopy := make([]int, len(currentArr))
		copy(currentArrCopy, currentArr)
		currentArrCopy = append(currentArrCopy, el)
		getCombinationsWithSumHelper(arr, currentArrCopy, currentSum+el, targetSum, combinationsWithSum, combinationUniqueStringsMap)
	}

}

func formUniqueStringForCombination(arr []int) string {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	sBuilder := strings.Builder{}
	for _, el := range arr {
		sBuilder.WriteString(strconv.Itoa(el))
	}

	return sBuilder.String()

}
