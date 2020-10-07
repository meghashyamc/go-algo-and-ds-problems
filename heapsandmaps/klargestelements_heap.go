package heapsandmaps

import (
	"container/heap"
	"errors"
	"fmt"
)

// Problem Description

// Given an 1D integer array A of size N you have to find and return the B largest elements of the array A.

// NOTE:

// Return the largest B elements in any order you like.

// Problem Constraints
// 1 <= N <= 105

// 1 <= B <= N

// 1 <= A[i] <= 103

// Input Format
// First argument is an 1D integer array A

// Second argument is an integer B.

// Output Format
// Return a 1D array of size B denoting the B largest elements.

// Example Input
// Input 1:

//  A = [11, 3, 4]
//  B = 2
// Input 2:

//  A = [11, 3, 4, 6]
//  B = 3

// Example Output
// Output 1:

//  [11, 4]
// Output 2:

//  [4, 6, 11]

// Example Explanation
// Explanation 1:

//  The two largest elements of A are 11 and 4
// Explanation 2:

//  The three largest elements of A are 11, 4 and 6

type intheap []int

func solveHeap(A []int, B int) []int {

	if B > len(A) {
		panic(fmt.Errorf("expected parameter B to be <= length of the array A but found B = %d and length of A = %d", B, len(A)))
	}

	h := intheap(A)
	heap.Init(&h)

	klargestelements := make([]int, 0)
	for i := 0; i < B; i++ {
		largest := heap.Pop(&h)
		klargestelements = append(klargestelements, largest.(int))
	}

	return klargestelements

}

func (h intheap) Len() int {
	return len(h)
}

func (h intheap) Less(i, j int) bool {

	if h == nil || len(h) == 0 {
		panic(errors.New("less can't be called with nil or empty heap"))
	}
	return h[i] > h[j]
}

func (h intheap) Swap(i, j int) {
	if h == nil || len(h) == 0 {
		panic(errors.New("swap can't be called with nil pr empty heap"))
	}
	h[i], h[j] = h[j], h[i]
}

func (h *intheap) Push(x interface{}) {
	if h == nil {
		panic(errors.New("push can't be called with nil heap"))
	}
	*h = append(*h, x.(int))
}

func (h *intheap) Pop() interface{} {
	if h == nil || len(*h) == 0 {
		panic(errors.New("pop can't be called with nil or empty heap"))
	}
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return popped

}
