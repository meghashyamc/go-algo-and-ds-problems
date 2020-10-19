package onstrings

// Given a string A consisting only of lowercase characters, we need to check whether it is possible to make this string a palindrome after removing exactly one character from this.

// If it is possible then return 1 else return 0.

// Problem Constraints
// 3 <= |A| <= 105

// A[i] is always a lowercase character.

// Input Format
// First and only argument is an string A.

// Output Format
// Return 1 if it is possible to convert A to palindrome by removing exactly one character else return 0.

// Example Input
// Input 1:

//  A = "abcba"
// Input 2:

//  A = "abecbea"

// Example Output
// Output 1:

//  1
// Output 2:

//  0

// Example Explanation
// Explanation 1:

//  We can remove character ‘c’ to make string palindrome
// Explanation 2:

//  It is not possible to make this string palindrome just by removing one character

func convToPalindrome(s string) int {
	problemChars := 0
	for i, j := 0, len(s)-1; j >= i; {
		if s[i] != s[j] {
			problemChars++
			if i+1 > len(s)-1 {
				break
			}
			if s[i+1] == s[j] {
				i++
				continue
			}
			if j-1 < 0 {
				break
			}
			if s[i] == s[j-1] {
				j--
				continue
			}
			return 0
		}
		i++
		j--

	}
	if problemChars <= 1 {
		return 1
	}
	return 0
}
