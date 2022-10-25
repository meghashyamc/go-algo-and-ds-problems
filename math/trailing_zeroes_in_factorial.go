package math

// Problem Description

// Given an integer A, return the number of trailing zeroes in A!.

// Note: Your solution should be in logarithmic time complexity.

// **Problem Constraints**
// 1 <= A <= 10000

// **Input Format**
// First and only argumment is integer A.

// **Output Format**
// Return an integer, the answer to the problem.

// **Example Input**
// Input 1:

//  A = 4
// Input 2:

//  A = 5

// **Example Output**
// Output 1:

//  0
// Output 2:

//  1

// **Example Explanation**
// Explanation 1:

//  4! = 24
// Explanation 2:

//  5! = 120
func trailingZeroes(A int) int {

	trailingZeroCount := 0

	for A/5 != 0 {
		trailingZeroCount += A / 5
		A /= 5
	}
	return trailingZeroCount
}
