package stacksandqueues

import "errors"

// Problem Description

// Given a string A consisting only of '(' and ')'.

// You need to find whether parantheses in A is balanced or not ,if it is balanced then return 1 else return 0.

// Problem Constraints
// 1 <= |A| <= 10^5

// Input Format
// First argument is an string A.

// Output Format
// Return 1 if parantheses in string are balanced else return 0.

// Example Input
// Input 1:

//  A = "(()())"
// Input 2:

//  A = "(()"

// Example Output
// Output 1:

//  1
// Output 2:

//  0

// Example Explanation
// Explanation 1:

//  Given string is balanced so we return 1
// Explanation 2:

//  Given string is not balanced so we return 0

const (
	leftParens  = '('
	rightParens = ')'
)

type stack []byte

func (s *stack) push(b byte) {
	if s == nil {
		s = new(stack)
	}
	*s = append(*s, b)
}

func (s *stack) pop() byte {
	if s == nil {
		panic(errors.New("can't pop empty stack"))
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped

}

func solve(A string) int {

	leftParensStack := stack{}
	for i := 0; i < len(A); i++ {
		if A[i] == leftParens {
			leftParensStack.push(A[i])
			continue
		}
		if A[i] == rightParens {
			if len(leftParensStack) == 0 {
				return 0
			}
			leftParensStack.pop()
			continue
		}
		panic(errors.New("found unexpected character in input string: " + string(A[i])))
	}
	if len(leftParensStack) == 0 {
		return 1
	}
	return 0
}
