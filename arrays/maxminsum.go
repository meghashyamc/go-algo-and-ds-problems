package arrays

import "math"

/*
*

	Problem Description

Given an array A of size N. You need to find the sum of Maximum and Minimum element in the given array.

NOTE: You should make minimum number of comparisons.

Problem Constraints
1 <= N <= 10^5

-10^9 <= A[i] <= 10^9

Input Format
First and only argument is an integer array A of size N.

Output Format
Return an integer denoting the sum Maximum and Minimum element in the given array.

Example Input
Input 1:

	A = [-2, 1, -4, 5, 3]

Input 2:

	A = [1, 3, 4, 1]

Example Output
Output 1:

	1

Output 2:

	5

Example Explanation
Explanation 1:

	Maximum Element is 5 and Minimum element is -4. (5 + (-4)) = 1.

Explanation 2:

	Maximum Element is 4 and Minimum element is 1. (4 + 1) = 5.
*/
func getMaxMinElementSum(arr []int) int {

	maxElement := math.MinInt
	minElement := math.MaxInt
	for _, el := range arr {
		if el > maxElement {
			maxElement = el
		}

		if el < minElement {
			minElement = el
		}
	}

	return maxElement + minElement
}
