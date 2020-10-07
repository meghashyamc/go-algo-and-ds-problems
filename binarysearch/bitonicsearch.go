package binarysearch

// Problem Description

// Given a bitonic sequence A of N distinct elements, write a program to find a given element B in the bitonic sequence in O(logN) time.

// NOTE:

// A Bitonic Sequence is a sequence of numbers which is first strictly increasing then after a point strictly decreasing.

// Problem Constraints
// 3 <= N <= 105

// 1 <= A[i], B <= 108

// Given array always contain a bitonic point.

// Array A always contain distinct elements.

// Input Format
// First argument is an integer array A denoting the bitonic sequence.

// Second argument is an integer B.

// Output Format
// Return a single integer denoting the position (0 index based) of the element B in the array A if B doesn't exist in A return -1.

// Example Input
// Input 1:

//  A = [3, 9, 10, 20, 17, 5, 1]
//  B = 20
// Input 2:

//  A = [5, 6, 7, 8, 9, 10, 3, 2, 1]
//  B = 30

// Example Output
// Output 1:

//  3
// Output 2:

//  -1

// Example Explanation
// Explanation 1:

//  B = 20 present in A at index 3
// Explanation 2:

//  B = 30 is not present in A

func solve(A []int, B int) int {

	if len(A) == 0 {
		return -1
	}
	bitonicIndex := getBitonicIndex(A)
	if bitonicIndex == -1 || A[bitonicIndex] < B {
		return -1
	}

	found := binarySearch(A, 0, bitonicIndex, true, B)
	if found != -1 {
		return found
	}

	return binarySearch(A, bitonicIndex, len(A), false, B)
}

func binarySearch(arr []int, lo, hi int, incr bool, elementToFind int) int {

	for {
		if lo >= hi {
			break
		}
		mid := lo + (hi-lo)/2
		// fmt.Println("lo:", lo, "mid:", mid, "hi:", hi)
		if arr[mid] == elementToFind {
			return mid
		}
		if arr[mid] < elementToFind {
			if incr {
				lo = mid + 1
				continue
			}
			hi = mid
			continue
		}

		if incr {
			hi = mid
			continue
		}
		lo = mid + 1

	}
	return -1
}

func getBitonicIndex(arr []int) int {

	lo := 0
	hi := len(arr)

	count := 0
	for {
		count++
		if count > 20 {
			break
		}
		if lo >= hi {
			break
		}

		mid := lo + (hi-lo)/2

		var incrLeft, incrRight int
		if mid > 0 {
			if arr[mid-1] < arr[mid] {
				incrLeft = 1

			} else {
				incrLeft = -1
			}

		}
		if mid < len(arr)-1 {
			if arr[mid] < arr[mid+1] {
				incrRight = 1

			} else {
				incrRight = -1
			}
		}

		if (incrLeft == 1 || incrLeft == 0) && incrRight == 1 {
			lo = mid + 1
			continue
		}

		if incrLeft == 1 && incrRight == -1 {
			return mid
		}
		if incrLeft == -1 && (incrRight == -1 || incrRight == 0) {
			hi = mid
			continue
		}

	}
	return -1
}
