package arrays

// Given a sorted array A containing N integers both positive and negative.

// You need to create another array containing the squares of all the elements in A and return it in non-decreasing order.

// Try to do this in O(N) time.

// Problem Constraints
// 1 <= N <= 10^5.

// -10^3 <= A[i] <= 10^3

// Input Format
// First and only argument is an integer array A.

// Output Format
// Return a integer array as described in the problem above.

// Example Input
// Input 1:

//  A = [-6, -3, -1, 2, 4, 5]
// Input 2:

//  A = [-5, -4, -2, 0, 1]

// Example Output
// Output 1:

//  [1, 4, 9, 16, 25, 36]
// Output 2:

//  [0, 1, 4, 16, 25]

func sortArrayWithSquares(A []int) []int {

	if A == nil || len(A) == 0 {
		return A
	}

	firstNonNegativeIndex := 0
	for i, el := range A {
		if el >= 0 {
			firstNonNegativeIndex = i
			break
		}
	}
	sortedSquares := make([]int, 0)
	for i, j := firstNonNegativeIndex, firstNonNegativeIndex-1; i <= len(A)-1 || j >= 0; {

		var posSquare, negSquare int
		if i > len(A)-1 {
			negSquare = A[j] * A[j]
			sortedSquares = append(sortedSquares, negSquare)
			j--
			continue
		}

		if j < 0 {
			posSquare = A[i] * A[i]
			sortedSquares = append(sortedSquares, posSquare)
			i++
			continue
		}
		posSquare = A[i] * A[i]
		negSquare = A[j] * A[j]

		if posSquare < negSquare {
			sortedSquares = append(sortedSquares, posSquare)
			i++
			continue
		}

		if posSquare > negSquare {
			sortedSquares = append(sortedSquares, negSquare)
			j--
			continue
		}

		if posSquare == negSquare {
			sortedSquares = append(sortedSquares, posSquare)
			sortedSquares = append(sortedSquares, negSquare)
			i++
			j--
			continue
		}
	}

	return sortedSquares
}
