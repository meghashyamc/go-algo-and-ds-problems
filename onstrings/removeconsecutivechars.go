package onstrings

// Problem Description

// Given a string A and integer B, remove all consecutive same characters that have length exactly B.

// Problem Constraints
// 1 <= |A| <= 100000

// 1 <= B <= |A|

// Input Format
// First Argument is string A.

// Second argument is integer B.

// Output Format
// Return a string after doing the removals.

// Example Input
// Input 1:

// A = "aabcd"
// B = 2
// Input 2:

// A = "aabbccd"
// B = 2

// Example Output
// Output 1:

//  "bcd"
// Output 2:

//  "d"

// Example Explanation
// Explanation 1:

//  "aa" had length 2.
// Explanation 2:

//  "aa", "bb" and "cc" had length of 2.

func solve(A string, B int) string {

	if len(A) == 0 || len(A) <= B {
		return A
	}
	streakStarts := make(map[int]bool)
	streakEnds := make(map[int]bool)

	var streak int
	st := -1
	//add streak starts and ends to maps
	for i := range A {

		if i == len(A)-1 {
			if streak != 0 {
				streak++
			}
			if streak == B {
				streakStarts[st] = true
				streakEnds[st+streak-1] = true

			}
			continue
		}
		if A[i] == A[i+1] {
			streak++
			if st == -1 {
				st = i
			}
		} else {
			if streak != 0 {
				streak++
				if streak == B {
					streakStarts[st] = true
					streakEnds[st+streak-1] = true

				}
				streak = 0
				st = -1
				continue
			}
		}
	}

	var streakGoingOn bool
	outputBytes := make([]byte, 0)
	//form a new string with characters that are not part of any streak
	for i := range A {

		if !streakGoingOn {
			if _, ok := streakStarts[i]; !ok {
				outputBytes = append(outputBytes, A[i])
			} else {
				streakGoingOn = true
			}
		} else {
			if _, ok := streakEnds[i]; ok {
				streakGoingOn = false
			}
		}
	}
	return string(outputBytes)
}
